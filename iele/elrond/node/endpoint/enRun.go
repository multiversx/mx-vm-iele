package endpoint

import (
	"errors"
	"fmt"
	"math/big"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

// RunSmartContractCreate computes how a smart contract creation should be performed
func (vm *ElrondIeleVM) RunSmartContractCreate(input *vmi.ContractCreateInput) (*vmi.VMOutput, error) {
	// validate input
	if input.Header == nil {
		return nil, errors.New("block header required")
	}
	if len(input.CallerAddr) != AddressLength {
		return nil, fmt.Errorf("caller address is not %d bytes in length", AddressLength)
	}

	vm.logCreateInput(input)

	// reset the VM state without freeing up the memory,
	// so the same memory can be reused on the next execution
	vm.kinterpreter.Model.Clear()
	vm.blockchainAdapter.InitAdapter()
	vm.blockchainAdapter.InitNewAddress(input.NewContractAddress, input.CallerAddr)

	// subtract initial gas (G0)
	g0, g0Err := vm.G0Create(input)
	if g0Err != nil {
		return nil, g0Err
	}
	gasProvided := big.NewInt(0).Sub(input.GasProvided, g0)

	// convert input
	kapp := vm.kinterpreter.Model.NewKApply(m.LblRunVM,
		m.BoolTrue,
		m.IntZero,
		vm.kinterpreter.Model.IntFromBytes(input.CallerAddr),
		vm.kinterpreter.Model.NewString(string(input.ContractCode)),
		vm.convertArgs(input.Arguments),
		vm.kinterpreter.Model.FromBigInt(input.CallValue),
		vm.kinterpreter.Model.FromBigInt(input.GasPrice),
		vm.kinterpreter.Model.FromBigInt(gasProvided),
		vm.kinterpreter.Model.FromBigInt(input.Header.Beneficiary),

		m.IntZero, // difficulty
		vm.kinterpreter.Model.FromBigInt(input.Header.Number),
		vm.kinterpreter.Model.FromBigInt(input.Header.GasLimit),
		vm.kinterpreter.Model.FromBigInt(input.Header.Timestamp),
		m.StringEmpty,
	)

	return vm.runTransaction(kapp)
}

// RunSmartContractCall computes the result of a smart contract call and how the system must change after the execution
func (vm *ElrondIeleVM) RunSmartContractCall(input *vmi.ContractCallInput) (*vmi.VMOutput, error) {
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

	vm.logCallInput(input)

	// reset the VM state without freeing up the memory,
	// so the same memory can be reused on the next execution
	vm.kinterpreter.Model.Clear()
	vm.blockchainAdapter.InitAdapter()

	// subtract initial gas (G0)
	g0, g0Err := vm.G0Call(input)
	if g0Err != nil {
		return nil, g0Err
	}
	gasProvided := big.NewInt(0).Sub(input.GasProvided, g0)

	kapp := vm.kinterpreter.Model.NewKApply(m.LblRunVM,
		m.BoolFalse,
		vm.kinterpreter.Model.IntFromBytes(input.RecipientAddr),
		vm.kinterpreter.Model.IntFromBytes(input.CallerAddr),
		m.StringEmpty,
		vm.convertArgs(input.Arguments),
		vm.kinterpreter.Model.FromBigInt(input.CallValue),
		vm.kinterpreter.Model.FromBigInt(input.GasPrice),
		vm.kinterpreter.Model.FromBigInt(gasProvided),
		vm.kinterpreter.Model.FromBigInt(input.Header.Beneficiary),
		m.IntZero, // difficulty
		vm.kinterpreter.Model.FromBigInt(input.Header.Number),
		vm.kinterpreter.Model.FromBigInt(input.Header.GasLimit),
		vm.kinterpreter.Model.FromBigInt(input.Header.Timestamp),
		vm.kinterpreter.Model.NewString(input.Function),
	)

	return vm.runTransaction(kapp)
}

func (vm *ElrondIeleVM) convertArgs(args []*big.Int) m.KReference {
	kargs := make([]m.KReference, len(args))
	for i, arg := range args {
		kargs[i] = vm.kinterpreter.Model.FromBigInt(arg)
	}
	kargList := vm.kinterpreter.Model.NewList(m.SortList, m.LblXuListXu, kargs)
	return kargList
}

// RunTransaction executes transaction contract code in VM
func (vm *ElrondIeleVM) runTransaction(kinput m.KReference) (*vmi.VMOutput, error) {

	mode := vm.kinterpreter.Model.NewKApply(m.LblNORMAL)

	kConfigMap := vm.kinterpreter.Model.EmptyMap(m.LblXuMapXu, m.SortMap)
	kConfigMap = vm.kinterpreter.Model.MapUpdate(kConfigMap,
		vm.kinterpreter.Model.NewKToken(m.SortKConfigVar, "$PGM"),
		kinput)
	kConfigMap = vm.kinterpreter.Model.MapUpdate(kConfigMap,
		vm.kinterpreter.Model.NewKToken(m.SortKConfigVar, "$MODE"),
		mode)
	kConfigMap = vm.kinterpreter.Model.MapUpdate(kConfigMap,
		vm.kinterpreter.Model.NewKToken(m.SortKConfigVar, "$SCHEDULE"),
		vm.scheduleToK(vm.schedule))

	// init
	initConfig, initErr := vm.kinterpreter.Eval(
		vm.kinterpreter.Model.NewKApply(interpreter.TopCellInitializer, kConfigMap),
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
	var returnData []*big.Int
	for _, kret := range kresRets {
		kiRet, kiRetOk := vm.kinterpreter.Model.GetBigInt(kret)
		if !kiRetOk {
			return nil, errors.New("return value not of type Int")
		}
		returnData = append(returnData, kiRet)
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
	vmResultMap := kAccountsCellArgs[0]
	if !vm.kinterpreter.Model.IsMapWithSortAndLabel(
		vmResultMap, m.SortAccountCellMap, m.LblXuAccountCellMapXu) {
		return nil, errors.New("invalid vmResult account map")
	}
	var modAccounts []*vmi.OutputAccount
	var modAccErr error
	vm.kinterpreter.Model.MapForEach(vmResultMap, func(_, kacc m.KReference) bool {
		var modAccount *vmi.OutputAccount
		modAccount, modAccErr = vm.convertKToModifiedAccount(kacc)
		if modAccErr != nil {
			return true
		}
		modAccounts = append(modAccounts, modAccount)
		return false
	})
	if modAccErr != nil {
		return nil, modAccErr
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

	vm.logInputAccounts()
	vm.logOutput(result)

	return result, nil
}

func (vm *ElrondIeleVM) convertKToModifiedAccount(kacc m.KReference) (*vmi.OutputAccount, error) {
	kappAcc, kappAccOk6 := vm.kinterpreter.Model.ExtractKApplyArgs(kacc, m.LblXltaccountXgt, 6)
	if !kappAccOk6 {
		return nil, errors.New("invalid account. Should be KApply with label '<account>' of length 6")
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

	// balance delta
	kappBalance, kappBalanceOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[1], m.LblXltbalanceXgt, 1)
	if !kappBalanceOk {
		return nil, errors.New("invalid account balance")
	}
	ibalance, ibalanceOk := vm.kinterpreter.Model.GetBigInt(kappBalance[0])
	if !ibalanceOk {
		return nil, errors.New("invalid account balance")
	}
	initialBalance, initialBalanceExists := vm.blockchainAdapter.InitialBalances[string(iaddr)]
	if !initialBalanceExists {
		return nil, errors.New("output account balance does not have a corresponding input balance")
	}
	balanceDelta := big.NewInt(0).Sub(ibalance, initialBalance)

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
	storageMap := kappStorage[0]
	if !vm.kinterpreter.Model.IsMapWithSortAndLabel(
		storageMap, m.SortMap, m.LblXuMapXu) {
		return nil, errors.New("invalid account storage")
	}
	var storageUpdates []*vmi.StorageUpdate
	var suErr error
	vm.kinterpreter.Model.MapForEach(storageMap, func(kkey, kvalue m.KReference) bool {
		ikey, ikeyOk := vm.kinterpreter.Model.GetBigInt(kkey)
		if !ikeyOk {
			suErr = errors.New("invalid account storage key")
			return true
		}
		ivalue, ivalueOk := vm.kinterpreter.Model.GetBigInt(kvalue)
		if !ivalueOk {
			suErr = errors.New("invalid account storage value")
			return true
		}
		storageUpdates = append(storageUpdates, &vmi.StorageUpdate{
			Offset: ikey.Bytes(),
			Data:   ivalue.Bytes(),
		})
		return false
	})
	if suErr != nil {
		return nil, suErr
	}

	// nonce
	kappNonce, kappNonceOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[4], m.LblXltnonceXgt, 1)
	if !kappNonceOk {
		return nil, errors.New("invalid account nonce")
	}
	inonce, inonceOk := vm.kinterpreter.Model.GetUint64(kappNonce[0])
	if !inonceOk {
		return nil, errors.New("invalid account nonce")
	}

	// exists, only checking that it is trues
	kappExists, kappExistsOk := vm.kinterpreter.Model.ExtractKApplyArgs(kappAcc[5], m.LblXltexistsXgt, 1)
	if !kappExistsOk {
		return nil, errors.New("invalid account exists tag")
	}
	if !m.IsTrue(kappExists[0]) {
		return nil, errors.New("VM should only output existing accounts")
	}

	return &vmi.OutputAccount{
		Address:        iaddr,
		Balance:        ibalance,
		BalanceDelta:   balanceDelta,
		Nonce:          inonce,
		StorageUpdates: storageUpdates,
		Code:           []byte(codeStr),
	}, nil
}

// this one calls the interpreter evaluate function to extract the bytes
func (vm *ElrondIeleVM) getCodeBytes(kcode m.KReference) (string, error) {
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

func (vm *ElrondIeleVM) convertKToLog(klog m.KReference) (*vmi.LogEntry, error) {
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
	strResult, isStr := vm.kinterpreter.Model.GetString(unparseResult) // Warning! comes as small endian
	if !isStr {
		return nil, errors.New("log data unparse error: result is not String")
	}
	littleEndianData := []byte(strResult)
	bigEndianData := make([]byte, len(littleEndianData))
	for i, ch := range littleEndianData {
		bigEndianData[len(littleEndianData)-1-i] = ch
	}

	return &vmi.LogEntry{
		Address: iAddr.Bytes(),
		Topics:  topics,
		Data:    bigEndianData,
	}, nil
}
