package endpoint

import (
	"errors"
	"fmt"
	"math/big"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

// RunSmartContractCreate computes how a smart contract creation should be performed
func (vm *OriginalIeleVM) RunSmartContractCreate(input *vmi.ContractCreateInput) (*vmi.VMOutput, error) {
	// validate input
	if input.Header == nil {
		return nil, errors.New("block header required")
	}
	if len(input.CallerAddr) != AddressLength {
		return nil, fmt.Errorf("caller address is not %d bytes in length", AddressLength)
	}

	// convert input
	kapp := &m.KApply{Label: m.LblRunVM, List: []m.K{
		m.BoolTrue,
		m.IntZero,
		m.NewIntFromBytes(input.CallerAddr),
		m.NewString(string(input.ContractCode)),
		convertArgs(input.Arguments),
		m.NewInt(input.CallValue),
		m.NewInt(input.GasPrice),
		m.NewInt(input.GasProvided),
		m.NewInt(input.Header.Beneficiary),
		m.IntZero, // difficulty
		m.NewInt(input.Header.Number),
		m.NewInt(input.Header.GasLimit),
		m.NewInt(input.Header.Timestamp),
		m.StringEmpty,
	}}

	return vm.runTransaction(kapp)
}

// RunSmartContractCall computes the result of a smart contract call and how the system must change after the execution
func (vm *OriginalIeleVM) RunSmartContractCall(input *vmi.ContractCallInput) (*vmi.VMOutput, error) {
	// validate input
	if input.Header == nil {
		return nil, errors.New("block header required")
	}
	if len(input.CallerAddr) != AddressLength {
		return nil, fmt.Errorf("caller address is not %d bytes in length", AddressLength)
	}
	if len(input.RecipientAddr) != AddressLength {
		return nil, fmt.Errorf("recipient address is not %d bytes in length", AddressLength)
	}

	kapp := &m.KApply{Label: m.LblRunVM, List: []m.K{
		m.BoolFalse,
		m.NewIntFromBytes(input.RecipientAddr),
		m.NewIntFromBytes(input.CallerAddr),
		m.StringEmpty,
		convertArgs(input.Arguments),
		m.NewInt(input.CallValue),
		m.NewInt(input.GasPrice),
		m.NewInt(input.GasProvided),
		m.NewInt(input.Header.Beneficiary),
		m.IntZero, // difficulty
		m.NewInt(input.Header.Number),
		m.NewInt(input.Header.GasLimit),
		m.NewInt(input.Header.Timestamp),
		m.NewString(input.Function),
	}}

	return vm.runTransaction(kapp)
}

func convertArgs(args []*big.Int) m.K {
	kargs := make([]m.K, len(args))
	for i, arg := range args {
		kargs[i] = m.NewInt(arg)
	}
	kargList := &m.List{Sort: m.SortList, Label: m.LblXuListXu, Data: kargs}
	return kargList
}

// RunTransaction ... executes transaction contract code in VM
func (vm *OriginalIeleVM) runTransaction(kinput m.K) (*vmi.VMOutput, error) {

	mode := &m.KApply{Label: m.LblNORMAL}

	kConfigMap := make(map[m.KMapKey]m.K)
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$PGM"}] = kinput
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$MODE"}] = mode
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$SCHEDULE"}] = scheduleToK(vm.schedule)

	// init
	initConfig, initErr := vm.kinterpreter.Eval(
		&m.KApply{
			Label: interpreter.TopCellInitializer,
			List:  []m.K{&m.Map{Sort: m.SortMap, Label: m.LblXuMapXu, Data: kConfigMap}}},
		m.InternedBottom,
	)
	if initErr != nil {
		return nil, initErr
	}

	// execute
	finalState, _, execErr := vm.kinterpreter.TakeStepsNoThread(initConfig)
	if execErr != nil {
		return nil, execErr
	}

	// extract result
	extracted, extractErr := vm.kinterpreter.Eval(
		&m.KApply{
			Label: m.LblExtractConfig,
			List:  []m.K{finalState}},
		m.InternedBottom,
	)
	if extractErr != nil {
		return nil, extractErr
	}

	// parse result
	resultKappArgs, resultKappOk := m.ExtractKApplyArgs(extracted, m.LblVmResult, 8)
	if !resultKappOk {
		return nil, errors.New("unexpected value where vmResult expected")
	}
	//interpreter.DebugPrint(extracted)

	// returns
	kresRets, retsOk := m.ExtractListData(resultKappArgs[0], m.SortList, m.KLabelForList)
	if !retsOk {
		return nil, errors.New("invalid vmResult return list")
	}
	var returnData []*big.Int
	for _, kret := range kresRets {
		kiRet, kiRetOk := kret.(*m.Int)
		if !kiRetOk {
			return nil, errors.New("return value not of type Int")
		}
		returnData = append(returnData, kiRet.Value)
	}

	// gas
	kresGas, gasOk := resultKappArgs[1].(*m.Int)
	if !gasOk {
		return nil, errors.New("invalid vmResult gas")
	}

	// refund
	kresRefund, refundOk := resultKappArgs[2].(*m.Int)
	if !refundOk {
		return nil, errors.New("invalid vmResult refund")
	}

	// status
	kresStatus, statusOk := resultKappArgs[3].(*m.Int)
	if !statusOk {
		return nil, errors.New("invalid vmResult status")
	}

	// self destruct (deleted) account addresses
	kSelfDestruct, kSelfDestructOk := m.ExtractListData(resultKappArgs[4], m.SortList, m.KLabelForList)
	if !kSelfDestructOk {
		return nil, errors.New("invalid vmResult self destruct list")
	}
	var deletedAddr [][]byte
	for _, ksd := range kSelfDestruct {
		isd, isdOk := ksd.(*m.Int)
		if !isdOk {
			return nil, errors.New("self destruct address not of type Int")
		}
		deletedAddr = append(deletedAddr, isd.Value.Bytes())
	}

	// logs
	kLogs, kLogsOk := m.ExtractListData(resultKappArgs[5], m.SortList, m.KLabelForList)
	if !kLogsOk {
		return nil, errors.New("invalid vmResult logs")
	}
	logs := make([]*vmi.LogEntry, len(kLogs))
	for i, klog := range kLogs {
		log, logErr := vm.convertKToLog(klog)
		if logErr != nil {
			return nil, logErr
		}
		logs[i] = log
	}

	// accounts
	kAccountsCellArgs, kAccountsCellArgsOk := m.ExtractKApplyArgs(resultKappArgs[6], m.LblXltaccountsXgt, 1)
	if !kAccountsCellArgsOk {
		return nil, errors.New("invalid vmResult accounts cell")
	}
	kAccountsMapData, kAccountsMapOk := m.ExtractMapData(kAccountsCellArgs[0], m.SortAccountCellMap, m.LblXuAccountCellMapXu)
	if !kAccountsMapOk {
		return nil, errors.New("invalid vmResult account map")
	}
	var modAccounts []*vmi.OutputAccount
	for _, kacc := range kAccountsMapData {
		modAccount, modAccErr := vm.convertKToModifiedAccount(kacc)
		if modAccErr != nil {
			return nil, modAccErr
		}
		modAccounts = append(modAccounts, modAccount)
	}

	kTouched, kTouchedOk := m.ExtractListData(resultKappArgs[7], m.SortList, m.KLabelForList)
	if !kTouchedOk {
		return nil, errors.New("invalid vmResult touched accounts list")
	}
	var touchedAddr [][]byte
	for _, kt := range kTouched {
		it, itOk := kt.(*m.Int)
		if !itOk {
			return nil, errors.New("touched address not of type Int")
		}
		touchedAddr = append(touchedAddr, it.Value.Bytes())
	}

	result := &vmi.VMOutput{
		ReturnData:      returnData,
		GasRemaining:    kresGas.Value,
		GasRefund:       kresRefund.Value,
		ReturnCode:      vmi.ReturnCode(int(kresStatus.Value.Int64())),
		DeletedAccounts: deletedAddr,
		TouchedAccounts: touchedAddr,
		OutputAccounts:  modAccounts,
		Logs:            logs,
	}

	return result, nil
}

func (vm *OriginalIeleVM) convertKToModifiedAccount(kacc m.K) (*vmi.OutputAccount, error) {
	kappAcc, kappAccOk5 := m.ExtractKApplyArgs(kacc, m.LblXltaccountXgt, 5)
	if !kappAccOk5 {
		var kappAccOk6 bool
		kappAcc, kappAccOk6 = m.ExtractKApplyArgs(kacc, m.LblXltaccountXgt, 6)
		if !kappAccOk6 {
			return nil, errors.New("invalid account. Should be KApply with label '<account>' of length 5 or 6")
		}
	}

	// address
	kappAddr, kappAddrOk := m.ExtractKApplyArgs(kappAcc[0], m.LblXltacctIDXgt, 1)
	if !kappAddrOk {
		return nil, errors.New("invalid account address")
	}
	iaddr, iaddrOk := kappAddr[0].(*m.Int)
	if !iaddrOk {
		return nil, errors.New("invalid account address")
	}

	// balance
	kappBalance, kappBalanceOk := m.ExtractKApplyArgs(kappAcc[1], m.LblXltbalanceXgt, 1)
	if !kappBalanceOk {
		return nil, errors.New("invalid account balance")
	}
	ibalance, ibalanceOk := kappBalance[0].(*m.Int)
	if !ibalanceOk {
		return nil, errors.New("invalid account balance")
	}

	// code
	kappCode, kappCodeOk := m.ExtractKApplyArgs(kappAcc[2], m.LblXltcodeXgt, 1)
	if !kappCodeOk {
		return nil, errors.New("invalid account code")
	}
	codeStr, codeErr := vm.getCodeBytes(kappCode[0])
	if codeErr != nil {
		return nil, codeErr
	}

	// storage
	kappStorage, kappStorageOk := m.ExtractKApplyArgs(kappAcc[3], m.LblXltstorageXgt, 1)
	if !kappStorageOk {
		return nil, errors.New("invalid account storage")
	}
	storageData, storageDataOk := m.ExtractMapData(kappStorage[0], m.SortMap, m.LblXuMapXu)
	if !storageDataOk {
		return nil, errors.New("invalid account storage")
	}
	var storageUpdates []*vmi.StorageUpdate
	for kmpkey, kvalue := range storageData {
		kkey, kkeyErr := kmpkey.ToKItem()
		if kkeyErr != nil {
			return nil, kkeyErr
		}
		ikey, ikeyOk := kkey.(*m.Int)
		if !ikeyOk {
			return nil, errors.New("invalid account storage key")
		}
		ivalue, ivalueOk := kvalue.(*m.Int)
		if !ivalueOk {
			return nil, errors.New("invalid account storage value")
		}
		storageUpdates = append(storageUpdates, &vmi.StorageUpdate{
			Offset: ikey.Value.Bytes(),
			Data:   ivalue.Value.Bytes(),
		})
	}

	// kappAcc[4] can be missing or not used

	// nonce
	kappNonce, kappNonceOk := m.ExtractKApplyArgs(kappAcc[len(kappAcc)-1], m.LblXltnonceXgt, 1) // kappAcc[4] or kappAcc[5]
	if !kappNonceOk {
		return nil, errors.New("invalid account nonce")
	}
	inonce, inonceOk := kappNonce[0].(*m.Int)
	if !inonceOk {
		return nil, errors.New("invalid account nonce")
	}

	return &vmi.OutputAccount{
		Address:        iaddr.Value.Bytes(),
		Balance:        ibalance.Value,
		Nonce:          inonce.Value,
		StorageUpdates: storageUpdates,
		Code:           []byte(codeStr),
	}, nil
}

// this one calls the interpreter evaluate function to extract the bytes
func (vm *OriginalIeleVM) getCodeBytes(kcode m.K) (string, error) {
	result, err := vm.kinterpreter.Eval(
		&m.KApply{
			Label: m.LblContractBytes,
			List:  []m.K{kcode},
		},
		m.InternedBottom,
	)
	if err != nil {
		return "", err
	}
	strResult, isStr := result.(*m.String)
	if !isStr {
		return "", errors.New("contract bytes evaluation did not return a K String")
	}
	return strResult.Value, nil
}

func (vm *OriginalIeleVM) convertKToLog(klog m.K) (*vmi.LogEntry, error) {
	logArgs, logKappOk := m.ExtractKApplyArgs(klog, m.LblLogEntry, 3)
	if !logKappOk {
		return nil, errors.New("invalid log entry")
	}
	iAddr, addrOk := logArgs[0].(*m.Int)
	if !addrOk {
		return nil, errors.New("invalid log address")
	}
	ktopics, ktopicsOk := m.ExtractListData(logArgs[1], m.SortList, m.KLabelForList)
	if !ktopicsOk {
		return nil, errors.New("invalid log topics")
	}
	topics := make([]*big.Int, len(ktopics))
	for i, ktopic := range ktopics {
		itopic, itopicOk := ktopic.(*m.Int)
		if !itopicOk {
			return nil, errors.New("invalid log topic")
		}
		topics[i] = itopic.Value
	}

	data := logArgs[2]
	unparseResult, unparseErr := vm.kinterpreter.Eval(
		&m.KApply{Label: m.LblUnparseByteStack, List: []m.K{data}},
		m.InternedBottom,
	)
	if unparseErr != nil {
		return nil, unparseErr
	}
	strResult, isStr := unparseResult.(*m.String)
	if !isStr {
		return nil, errors.New("log data unparse error: result is not String")
	}

	return &vmi.LogEntry{
		Address: iAddr.Value.Bytes(),
		Topics:  topics,
		Data:    []byte(strResult.Value),
	}, nil
}
