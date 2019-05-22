package endpoint

import (
	"errors"
	"math/big"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// InterpreterOptions ... options used by the interpreter, shouldn't need to change other than for debugging
var InterpreterOptions = &interpreter.ExecuteOptions{
	TracePretty: false,
	TraceKPrint: false,
	Verbose:     false,
	MaxSteps:    0,
}

// RunTransaction ... executes transaction contract code in VM
func RunTransaction(input *VMInput) (*VMOutput, error) {
	if input.BlockHeader == nil {
		return nil, errors.New("block header required")
	}

	kargs := make([]m.K, len(input.Arguments))
	for i, arg := range input.Arguments {
		kargs[i] = m.NewInt(arg)
	}
	kargList := &m.List{Sort: m.SortList, Label: m.LblXuListXu, Data: kargs}

	kapp := &m.KApply{Label: m.LblRunVM, List: []m.K{
		m.ToBool(input.IsCreate),
		m.NewInt(input.RecipientAddr),
		m.NewInt(input.CallerAddr),
		m.NewString(input.InputData),
		kargList,
		m.NewInt(input.CallValue),
		m.NewInt(input.GasPrice),
		m.NewInt(input.GasProvided),
		m.NewInt(input.BlockHeader.Beneficiary),
		m.NewInt(input.BlockHeader.Difficulty),
		m.NewInt(input.BlockHeader.Number),
		m.NewInt(input.BlockHeader.GasLimit),
		m.NewInt(input.BlockHeader.UnixTimestamp),
		m.NewString(input.Function),
	}}

	mode := &m.KApply{Label: m.LblNORMAL}

	kConfigMap := make(map[m.KMapKey]m.K)
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$PGM"}] = kapp
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$MODE"}] = mode
	kConfigMap[m.KToken{Sort: m.SortKConfigVar, Value: "$SCHEDULE"}] = scheduleToK(input.Schedule)

	// init
	initConfig, initErr := interpreter.Eval(
		&m.KApply{
			Label: interpreter.TopCellInitializer,
			List:  []m.K{&m.Map{Sort: m.SortMap, Label: m.LblXuMapXu, Data: kConfigMap}}},
		m.InternedBottom,
	)
	if initErr != nil {
		return nil, initErr
	}

	// execute
	finalState, _, execErr := interpreter.TakeStepsNoThread(initConfig, InterpreterOptions)
	if execErr != nil {
		return nil, execErr
	}

	// extract result
	extracted, extractErr := interpreter.Eval(
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
		deletedAddr = append(deletedAddr, world.AccountAddress(isd.Value))
	}

	// logs
	kLogs, kLogsOk := m.ExtractListData(resultKappArgs[5], m.SortList, m.KLabelForList)
	if !kLogsOk {
		return nil, errors.New("invalid vmResult logs")
	}
	logs := make([]*LogEntry, len(kLogs))
	for i, klog := range kLogs {
		log, logErr := convertKToLog(klog)
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
	var modAccounts []*world.ModifiedAccount
	for _, kacc := range kAccountsMapData {
		modAccount, modAccErr := convertKToModifiedAccount(kacc)
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
		touchedAddr = append(touchedAddr, world.AccountAddress(it.Value))
	}

	result := &VMOutput{
		ReturnData:       returnData,
		GasRemaining:     kresGas.Value,
		GasRefund:        kresRefund.Value,
		ReturnCode:       kresStatus.Value,
		DeletedAccounts:  deletedAddr,
		TouchedAccounts:  touchedAddr,
		ModifiedAccounts: modAccounts,
		Logs:             logs,
	}

	return result, nil
}

func convertKToModifiedAccount(kacc m.K) (*world.ModifiedAccount, error) {
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
	codeStr, codeErr := getCodeBytes(kappCode[0])
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
	var storageUpdates []*world.StorageUpdate
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
		storageUpdates = append(storageUpdates, &world.StorageUpdate{
			Offset: ikey.Value,
			Data:   ivalue.Value,
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

	return &world.ModifiedAccount{
		Address:        world.AccountAddress(iaddr.Value),
		Balance:        ibalance.Value,
		Nonce:          inonce.Value,
		StorageUpdates: storageUpdates,
		Code:           codeStr,
	}, nil
}

// this one calls the interpreter evaluate function to extract the bytes
func getCodeBytes(kcode m.K) (string, error) {
	result, err := interpreter.Eval(
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

func convertKToLog(klog m.K) (*LogEntry, error) {
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
	unparseResult, unparseErr := interpreter.Eval(
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

	return &LogEntry{
		Address: iAddr.Value,
		Topics:  topics,
		Data:    []byte(strResult.Value),
	}, nil
}
