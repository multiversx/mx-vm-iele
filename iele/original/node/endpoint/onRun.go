package endpoint

import (
	"errors"
	"fmt"
	"math/big"

	twos "github.com/ElrondNetwork/big-int-util/twos-complement"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

// RunSmartContractCreate computes how a smart contract creation should be performed
func (vm *OriginalIeleVM) RunSmartContractCreate(input *vmi.ContractCreateInput) (*vmi.VMOutput, error) {
	// validate input
	if len(input.CallerAddr) != AddressLength {
		return nil, fmt.Errorf("caller address is not %d bytes in length", AddressLength)
	}

	// reset the VM state without freeing up the memory,
	// so the same memory can be reused on the next execution
	vm.kinterpreter.Model.Clear()

	// subtract initial gas (G0)
	g0, g0Err := vm.G0Create(input)
	if g0Err != nil {
		return nil, g0Err
	}
	gasProvided := big.NewInt(0).Sub(big.NewInt(0).SetUint64(input.GasProvided), g0)

	// convert input
	kapp := vm.kinterpreter.Model.NewKApply(m.LblRunVM,
		m.BoolTrue,
		m.IntZero,
		vm.kinterpreter.Model.IntFromBytes(input.CallerAddr),
		vm.kinterpreter.Model.NewString(string(input.ContractCode)),
		vm.convertArgs(input.Arguments),
		vm.kinterpreter.Model.FromBigInt(input.CallValue),
		vm.kinterpreter.Model.FromUint64(input.GasPrice),
		vm.kinterpreter.Model.FromBigInt(gasProvided),
		m.IntZero, // beneficiary
		m.IntZero, // difficulty
		m.IntZero, // number
		m.IntZero, // block gas limit
		vm.kinterpreter.Model.FromUint64(vm.blockchainAdapter.Upstream.CurrentTimeStamp()),
		m.StringEmpty,
	)

	return vm.runTransaction(kapp)
}

// RunSmartContractCall computes the result of a smart contract call and how the system must change after the execution
func (vm *OriginalIeleVM) RunSmartContractCall(input *vmi.ContractCallInput) (*vmi.VMOutput, error) {
	// validate input
	if len(input.CallerAddr) != AddressLength {
		return nil, fmt.Errorf("caller address is not %d bytes in length", AddressLength)
	}
	if len(input.RecipientAddr) != AddressLength {
		return nil, fmt.Errorf("recipient address is not %d bytes in length", AddressLength)
	}

	// subtract initial gas (G0)
	g0, g0Err := vm.G0Call(input)
	if g0Err != nil {
		return nil, g0Err
	}
	gasProvided := big.NewInt(0).Sub(big.NewInt(0).SetUint64(input.GasProvided), g0)

	kapp := vm.kinterpreter.Model.NewKApply(m.LblRunVM,
		m.BoolFalse,
		vm.kinterpreter.Model.IntFromBytes(input.RecipientAddr),
		vm.kinterpreter.Model.IntFromBytes(input.CallerAddr),
		m.StringEmpty,
		vm.convertArgs(input.Arguments),
		vm.kinterpreter.Model.FromBigInt(input.CallValue),
		vm.kinterpreter.Model.FromUint64(input.GasPrice),
		vm.kinterpreter.Model.FromBigInt(gasProvided),
		m.IntZero, // beneficiary
		m.IntZero, // difficulty
		m.IntZero, // number
		m.IntZero, // block gas limit
		vm.kinterpreter.Model.FromUint64(vm.blockchainAdapter.Upstream.CurrentTimeStamp()),
		vm.kinterpreter.Model.NewString(input.Function),
	)

	return vm.runTransaction(kapp)
}

func (vm *OriginalIeleVM) convertArgs(args [][]byte) m.KReference {
	kargs := make([]m.KReference, len(args))
	for i, arg := range args {
		kargs[i] = vm.kinterpreter.Model.FromBigInt(twos.FromBytes(arg))
	}
	kargList := vm.kinterpreter.Model.NewList(m.SortList, m.LblXuListXu, kargs)
	return kargList
}

// RunTransaction ... executes transaction contract code in VM
func (vm *OriginalIeleVM) runTransaction(kinput m.KReference) (*vmi.VMOutput, error) {

	mode := vm.kinterpreter.Model.NewKApply(m.LblNORMAL)

	kConfigMap := make(map[m.KMapKey]m.KReference)
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$PGM"}] = kinput
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$MODE"}] = mode
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$SCHEDULE"}] = vm.scheduleToK(vm.schedule)

	// init
	initConfig, initErr := vm.kinterpreter.Eval(
		vm.kinterpreter.Model.NewKApply(interpreter.TopCellInitializer,
			vm.kinterpreter.Model.NewMap(m.SortMap, m.LblXuMapXu, kConfigMap),
		),
		m.InternedBottom,
	)
	if initErr != nil {
		return nil, initErr
	}

	// execute
	execErr := vm.kinterpreter.TakeStepsNoThread(initConfig)
	if execErr != nil {
		return nil, execErr
	}
	finalState := vm.kinterpreter.GetState()

	// extract result
	extracted, extractErr := vm.kinterpreter.Eval(
		vm.kinterpreter.Model.NewKApply(m.LblExtractConfig, finalState),
		m.InternedBottom,
	)
	if extractErr != nil {
		return nil, extractErr
	}

	// parse result
	resultKappArgs, resultKappOk := vm.kinterpreter.Model.ExtractKApplyArgs(extracted, m.LblVmResult, 8)
	if !resultKappOk {
		return nil, errors.New("unexpected value where vmResult expected")
	}
	//interpreter.DebugPrint(extracted)

	// returns
	kresRets, retsOk := vm.kinterpreter.Model.ExtractListData(resultKappArgs[0], m.SortList, m.KLabelForList)
	if !retsOk {
		return nil, errors.New("invalid vmResult return list")
	}
	var returnData [][]byte
	for _, kret := range kresRets {
		kiRet, kiRetOk := vm.kinterpreter.Model.GetBigInt(kret)
		if !kiRetOk {
			return nil, errors.New("return value not of type Int")
		}
		returnData = append(returnData, twos.ToBytes(kiRet))
	}

	// gas
	kresGas, gasOk := vm.kinterpreter.Model.GetBigInt(resultKappArgs[1])
	if !gasOk {
		return nil, errors.New("invalid vmResult gas")
	}

	// refund
	kresRefund, refundOk := vm.kinterpreter.Model.GetBigInt(resultKappArgs[2])
	if !refundOk {
		return nil, errors.New("invalid vmResult refund")
	}

	// status
	kresStatus, statusOk := vm.kinterpreter.Model.GetBigInt(resultKappArgs[3])
	if !statusOk {
		return nil, errors.New("invalid vmResult status")
	}

	// self destruct (deleted) account addresses
	kSelfDestruct, kSelfDestructOk := vm.kinterpreter.Model.ExtractListData(resultKappArgs[4], m.SortList, m.KLabelForList)
	if !kSelfDestructOk {
		return nil, errors.New("invalid vmResult self destruct list")
	}
	var deletedAddr [][]byte
	for _, ksd := range kSelfDestruct {
		addrSD, sdOk := vm.blockchainAdapter.ConvertKIntToAddress(ksd, vm.kinterpreter.Model)
		if !sdOk {
			return nil, errors.New("self destruct address not of type Int")
		}
		deletedAddr = append(deletedAddr, addrSD)
	}

	// logs
	kLogs, kLogsOk := vm.kinterpreter.Model.ExtractListData(resultKappArgs[5], m.SortList, m.KLabelForList)
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
	kAccountsCellArgs, kAccountsCellArgsOk := vm.kinterpreter.Model.ExtractKApplyArgs(
		resultKappArgs[6], m.LblXltaccountsXgt, 1)
	if !kAccountsCellArgsOk {
		return nil, errors.New("invalid vmResult accounts cell")
	}
	kAccountsMapData, kAccountsMapOk := vm.kinterpreter.Model.ExtractMapData(
		kAccountsCellArgs[0], m.SortAccountCellMap, m.LblXuAccountCellMapXu)
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

	kTouched, kTouchedOk := vm.kinterpreter.Model.ExtractListData(resultKappArgs[7], m.SortList, m.KLabelForList)
	if !kTouchedOk {
		return nil, errors.New("invalid vmResult touched accounts list")
	}
	var touchedAddr [][]byte
	for _, kt := range kTouched {
		it, itOk := vm.blockchainAdapter.ConvertKIntToAddress(kt, vm.kinterpreter.Model)
		if !itOk {
			return nil, errors.New("touched address not a valid address")
		}
		touchedAddr = append(touchedAddr, it)
	}

	result := &vmi.VMOutput{
		ReturnData:      returnData,
		GasRemaining:    kresGas,
		GasRefund:       kresRefund,
		ReturnCode:      vmi.ReturnCode(int(kresStatus.Int64())),
		DeletedAccounts: deletedAddr,
		TouchedAccounts: touchedAddr,
		OutputAccounts:  modAccounts,
		Logs:            logs,
	}

	//vm.kinterpreter.Model.PrintStats()

	return result, nil
}

func (vm *OriginalIeleVM) convertKToModifiedAccount(kacc m.KReference) (*vmi.OutputAccount, error) {
	kappAcc, kappAccOk5 := vm.kinterpreter.Model.ExtractKApplyArgs(kacc, m.LblXltaccountXgt, 5)
	if !kappAccOk5 {
		var kappAccOk6 bool
		kappAcc, kappAccOk6 = vm.kinterpreter.Model.ExtractKApplyArgs(kacc, m.LblXltaccountXgt, 6)
		if !kappAccOk6 {
			return nil, errors.New("invalid account. Should be KApply with label '<account>' of length 5 or 6")
		}
	}

	// address
	kappAddr, kappAddrOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[0], m.LblXltacctIDXgt, 1)
	if !kappAddrOk {
		return nil, errors.New("invalid account address")
	}
	iaddr, iaddrOk := vm.blockchainAdapter.ConvertKIntToAddress(kappAddr[0], vm.kinterpreter.Model)
	if !iaddrOk {
		return nil, errors.New("invalid account address")
	}

	// balance
	kappBalance, kappBalanceOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[1], m.LblXltbalanceXgt, 1)
	if !kappBalanceOk {
		return nil, errors.New("invalid account balance")
	}
	ibalance, ibalanceOk := vm.kinterpreter.Model.GetBigInt(kappBalance[0])
	if !ibalanceOk {
		return nil, errors.New("invalid account balance")
	}

	// code
	kappCode, kappCodeOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[2], m.LblXltcodeXgt, 1)
	if !kappCodeOk {
		return nil, errors.New("invalid account code")
	}
	codeStr, codeErr := vm.getCodeBytes(kappCode[0])
	if codeErr != nil {
		return nil, codeErr
	}

	// storage
	kappStorage, kappStorageOk := vm.kinterpreter.Model.ExtractKApplyArgs(
		kappAcc[3], m.LblXltstorageXgt, 1)
	if !kappStorageOk {
		return nil, errors.New("invalid account storage")
	}
	storageData, storageDataOk := vm.kinterpreter.Model.ExtractMapData(
		kappStorage[0], m.SortMap, m.LblXuMapXu)
	if !storageDataOk {
		return nil, errors.New("invalid account storage")
	}
	var storageUpdates []*vmi.StorageUpdate
	for kmpkey, kvalue := range storageData {
		kkey, kkeyErr := vm.kinterpreter.Model.ToKItem(kmpkey)
		if kkeyErr != nil {
			return nil, kkeyErr
		}
		ikey, ikeyOk := vm.kinterpreter.Model.GetBigInt(kkey)
		if !ikeyOk {
			return nil, errors.New("invalid account storage key")
		}
		ivalue, ivalueOk := vm.kinterpreter.Model.GetBigInt(kvalue)
		if !ivalueOk {
			return nil, errors.New("invalid account storage value")
		}
		storageUpdates = append(storageUpdates, &vmi.StorageUpdate{
			Offset: ikey.Bytes(),
			Data:   ivalue.Bytes(),
		})
	}

	// kappAcc[4] can be missing or not used

	// nonce
	kappNonce, kappNonceOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[len(kappAcc)-1], m.LblXltnonceXgt, 1) // kappAcc[4] or kappAcc[5]
	if !kappNonceOk {
		return nil, errors.New("invalid account nonce")
	}
	inonce, inonceOk := vm.kinterpreter.Model.GetUint64(kappNonce[0])
	if !inonceOk {
		return nil, errors.New("invalid account nonce")
	}

	return &vmi.OutputAccount{
		Address:        iaddr,
		Balance:        ibalance,
		Nonce:          inonce,
		StorageUpdates: storageUpdates,
		Code:           []byte(codeStr),
	}, nil
}

// this one calls the interpreter evaluate function to extract the bytes
func (vm *OriginalIeleVM) getCodeBytes(kcode m.KReference) (string, error) {
	result, err := vm.kinterpreter.Eval(
		vm.kinterpreter.Model.NewKApply(m.LblContractBytes, kcode),
		m.InternedBottom,
	)
	if err != nil {
		return "", err
	}
	strResult, isStr := vm.kinterpreter.Model.GetString(result)
	if !isStr {
		return "", errors.New("contract bytes evaluation did not return a K String")
	}
	return strResult, nil
}

func (vm *OriginalIeleVM) convertKToLog(klog m.KReference) (*vmi.LogEntry, error) {
	logArgs, logKappOk := vm.kinterpreter.Model.ExtractKApplyArgs(klog, m.LblLogEntry, 3)
	if !logKappOk {
		return nil, errors.New("invalid log entry")
	}
	iAddr, addrOk := vm.kinterpreter.Model.GetBigInt(logArgs[0])
	if !addrOk {
		return nil, errors.New("invalid log address")
	}
	ktopics, ktopicsOk := vm.kinterpreter.Model.ExtractListData(logArgs[1], m.SortList, m.KLabelForList)
	if !ktopicsOk {
		return nil, errors.New("invalid log topics")
	}
	topics := make([]*big.Int, len(ktopics))
	for i, ktopic := range ktopics {
		itopic, itopicOk := vm.kinterpreter.Model.GetBigInt(ktopic)
		if !itopicOk {
			return nil, errors.New("invalid log topic")
		}
		topics[i] = itopic
	}

	data := logArgs[2]
	unparseResult, unparseErr := vm.kinterpreter.Eval(
		vm.kinterpreter.Model.NewKApply(m.LblUnparseByteStack, data),
		m.InternedBottom,
	)
	if unparseErr != nil {
		return nil, unparseErr
	}
	strResult, isStr := vm.kinterpreter.Model.GetString(unparseResult)
	if !isStr {
		return nil, errors.New("log data unparse error: result is not String")
	}

	return &vmi.LogEntry{
		Address: iAddr.Bytes(),
		Topics:  topics,
		Data:    []byte(strResult),
	}, nil
}
