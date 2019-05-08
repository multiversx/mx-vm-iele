package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
)

// TopCellInitializer ... label passed to Eval to initialize the top cell
const TopCellInitializer m.KLabel = m.LblInitGeneratedTopCell

// Eval ... evaluates a KApply item based on its label and arguments
func Eval(c m.K, config m.K) (m.K, error) {
	kapp, isKapply := c.(*m.KApply)
	if !isKapply {
		return c, nil
	}
	switch kapp.Label {
		case m.LblIsSStoreInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSStoreInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSStoreInst(kapp.List[0], config, -1)
		case m.LblXhashargv:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashargv", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashargv(config, -1)
		case m.LblIsCallValueCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallValueCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallValueCell(kapp.List[0], config, -1)
		case m.LblMapXcolonlookup:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMapXcolonlookup", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMapXcolonlookup(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashrlpEncodeIntsAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeIntsAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeIntsAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXlsqbXuXrsqbXuARRAYXhyphenSYNTAX:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXrsqbXuARRAYXhyphenSYNTAX", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXrsqbXuARRAYXhyphenSYNTAX(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashpadToWidth:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashpadToWidth", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashpadToWidth(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsNregsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNregsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNregsCellOpt(kapp.List[0], config, -1)
		case m.LblIsFuncIDCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncIDCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncIDCellOpt(kapp.List[0], config, -1)
		case m.LblIsAssignInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAssignInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAssignInst(kapp.List[0], config, -1)
		case m.LblIntSize:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntSize", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIntSize(kapp.List[0], config, -1)
		case m.LblIsArray:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsArray", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsArray(kapp.List[0], config, -1)
		case m.LblXuXltXeqSetXuXuSET:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXeqSetXuXuSET", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXeqSetXuXuSET(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTxGasPriceCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxGasPriceCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxGasPriceCellOpt(kapp.List[0], config, -1)
		case m.LblIsIOError:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIOError", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIOError(kapp.List[0], config, -1)
		case m.LblIsDataCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDataCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDataCellOpt(kapp.List[0], config, -1)
		case m.LblIsOrInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOrInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOrInst(kapp.List[0], config, -1)
		case m.LblMakeList:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMakeList", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMakeList(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashisValidLoad:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidLoad", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashisValidLoad(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsLabeledBlocks:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLabeledBlocks", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLabeledBlocks(kapp.List[0], config, -1)
		case m.LblIsMulModInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMulModInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMulModInst(kapp.List[0], config, -1)
		case m.LblIsCurrentInstructionsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentInstructionsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentInstructionsCellOpt(kapp.List[0], config, -1)
		case m.LblXhashdasmContract:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmContract", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashdasmContract(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashunlockXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashunlockXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashunlockXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsActiveAccountsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsActiveAccountsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsActiveAccountsCellOpt(kapp.List[0], config, -1)
		case m.LblIsNonceCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNonceCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNonceCellOpt(kapp.List[0], config, -1)
		case m.LblXhashparseMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseMap(kapp.List[0], config, -1)
		case m.LblIsCurrentMemoryCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentMemoryCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentMemoryCellOpt(kapp.List[0], config, -1)
		case m.LblIsLocalCallInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalCallInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalCallInst(kapp.List[0], config, -1)
		case m.LblXhashecrec:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashecrec", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashecrec(kapp.List[0], config, -1)
		case m.LblIsContractDeclaration:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractDeclaration", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractDeclaration(kapp.List[0], config, -1)
		case m.LblIsNumberCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNumberCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNumberCell(kapp.List[0], config, -1)
		case m.LblIsLabelsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLabelsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLabelsCellOpt(kapp.List[0], config, -1)
		case m.LblIsInstructionsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInstructionsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInstructionsCellOpt(kapp.List[0], config, -1)
		case m.LblXhashloadLen:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadLen", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashloadLen(kapp.List[0], config, -1)
		case m.LblIsOutputCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOutputCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOutputCellOpt(kapp.List[0], config, -1)
		case m.LblCexpmod:
			if len(kapp.List) != 5 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCexpmod", expectedArity: 5, actualArity: len(kapp.List)}
			}
			return evalCexpmod(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], config, -1)
		case m.LblProjectXcolonMode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalProjectXcolonMode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalProjectXcolonMode(kapp.List[0], config, -1)
		case m.LblIsFunctionParameters:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionParameters", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionParameters(kapp.List[0], config, -1)
		case m.LblXhashcomputeNRegs:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcomputeNRegs", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashcomputeNRegs(kapp.List[0], config, -1)
		case m.LblTopLevelAppend:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalTopLevelAppend", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalTopLevelAppend(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsCurrentContractCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentContractCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentContractCellFragment(kapp.List[0], config, -1)
		case m.LblIsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCell(kapp.List[0], config, -1)
		case m.LblValues:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalValues", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalValues(kapp.List[0], config, -1)
		case m.LblIsRefundCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsRefundCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsRefundCell(kapp.List[0], config, -1)
		case m.LblIsUnOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsUnOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsUnOp(kapp.List[0], config, -1)
		case m.LblXhashloadDeclarations:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadDeclarations", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashloadDeclarations(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsFunctionNameCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionNameCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionNameCellOpt(kapp.List[0], config, -1)
		case m.LblIsExpModInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExpModInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExpModInst(kapp.List[0], config, -1)
		case m.LblIsTopLevelDefinition:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTopLevelDefinition", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTopLevelDefinition(kapp.List[0], config, -1)
		case m.LblIsFidCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFidCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFidCell(kapp.List[0], config, -1)
		case m.LblKeccak:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalKeccak", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalKeccak(kapp.List[0], config, -1)
		case m.LblXhashconfigurationXuKXhyphenREFLECTION:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashconfigurationXuKXhyphenREFLECTION", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashconfigurationXuKXhyphenREFLECTION(config, -1)
		case m.LblXhashdecodeLengthPrefixLength:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdecodeLengthPrefixLength", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashdecodeLengthPrefixLength(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsFloat:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFloat", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFloat(kapp.List[0], config, -1)
		case m.LblInitPeakMemoryCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitPeakMemoryCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitPeakMemoryCell(config, -1)
		case m.LblIsLocalNames:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalNames", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalNames(kapp.List[0], config, -1)
		case m.LblIsBlockhashCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBlockhashCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBlockhashCellOpt(kapp.List[0], config, -1)
		case m.LblChop:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalChop", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalChop(kapp.List[0], config, -1)
		case m.LblIsRegsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsRegsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsRegsCell(kapp.List[0], config, -1)
		case m.LblIsArgsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsArgsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsArgsCellOpt(kapp.List[0], config, -1)
		case m.LblXuXplusIntXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusIntXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusIntXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsSendtoCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSendtoCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSendtoCell(kapp.List[0], config, -1)
		case m.LblIsTimestampCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTimestampCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTimestampCellOpt(kapp.List[0], config, -1)
		case m.LblXhashrev:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrev", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrev(kapp.List[0], kapp.List[1], config, -1)
		case m.LblReplaceAtBytes:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalReplaceAtBytes", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalReplaceAtBytes(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsTypeCheckingCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTypeCheckingCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTypeCheckingCell(kapp.List[0], config, -1)
		case m.LblIsLValues:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLValues", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLValues(kapp.List[0], config, -1)
		case m.LblAccountCellMapItem:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalAccountCellMapItem", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalAccountCellMapItem(kapp.List[0], kapp.List[1], config, -1)
		case m.LblArrayCtor:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalArrayCtor", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalArrayCtor(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFuncIDsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncIDsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncIDsCell(kapp.List[0], config, -1)
		case m.LblXhashparseByteStackRawAux:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseByteStackRawAux", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashparseByteStackRawAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsArgsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsArgsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsArgsCell(kapp.List[0], config, -1)
		case m.LblGetInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalGetInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalGetInt(kapp.List[0], config, -1)
		case m.LblIsProgramCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsProgramCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsProgramCellOpt(kapp.List[0], config, -1)
		case m.LblBool2Word:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBool2Word", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalBool2Word(kapp.List[0], config, -1)
		case m.LblXhashparseByteStackRaw:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseByteStackRaw", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseByteStackRaw(kapp.List[0], config, -1)
		case m.LblInitDataCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitDataCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitDataCell(config, -1)
		case m.LblXuListXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuListXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuListXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsNonEmptyInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNonEmptyInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNonEmptyInts(kapp.List[0], config, -1)
		case m.LblIsFunctionBodiesCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionBodiesCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionBodiesCellOpt(kapp.List[0], config, -1)
		case m.LblXuXhyphenMapXuXuMAP:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXhyphenMapXuXuMAP", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXhyphenMapXuXuMAP(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashloadCodeAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadCodeAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashloadCodeAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashsort:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsort", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsort(kapp.List[0], config, -1)
		case m.LblXuXeqXeqKXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXeqKXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXeqKXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsInstructions:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInstructions", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInstructions(kapp.List[0], config, -1)
		case m.LblIsBinOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBinOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBinOp(kapp.List[0], config, -1)
		case m.LblIsCallDepthCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallDepthCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallDepthCell(kapp.List[0], config, -1)
		case m.LblReplaceFirstXlparenXuXcommaXuXcommaXuXrparenXuSTRING:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalReplaceFirstXlparenXuXcommaXuXcommaXuXrparenXuSTRING", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalReplaceFirstXlparenXuXcommaXuXcommaXuXrparenXuSTRING(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsAccountsCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountsCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountsCellFragment(kapp.List[0], config, -1)
		case m.LblIsTwosInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTwosInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTwosInst(kapp.List[0], config, -1)
		case m.LblXdotMap:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotMap", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotMap(config, -1)
		case m.LblIsFuncIDsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncIDsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncIDsCellOpt(kapp.List[0], config, -1)
		case m.LblIsCallFrameCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallFrameCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallFrameCell(kapp.List[0], config, -1)
		case m.LblInitCallFrameCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallFrameCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallFrameCell(config, -1)
		case m.LblInitNonceCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitNonceCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitNonceCell(config, -1)
		case m.LblIsJumpTableCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJumpTableCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJumpTableCell(kapp.List[0], config, -1)
		case m.LblXuXeqXslashXeqStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXslashXeqStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXslashXeqStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsAccountCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountCellFragment(kapp.List[0], config, -1)
		case m.LblIsBswapInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBswapInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBswapInst(kapp.List[0], config, -1)
		case m.LblIsReturnOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsReturnOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsReturnOp(kapp.List[0], config, -1)
		case m.LblIsAccounts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccounts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccounts(kapp.List[0], config, -1)
		case m.LblIsNetworkCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNetworkCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNetworkCellOpt(kapp.List[0], config, -1)
		case m.LblIsCallerCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallerCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallerCell(kapp.List[0], config, -1)
		case m.LblIsType:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsType", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsType(kapp.List[0], config, -1)
		case m.LblSignextend:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSignextend", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalSignextend(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashtoBlocks:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtoBlocks", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashtoBlocks(kapp.List[0], config, -1)
		case m.LblIsJSONKey:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJSONKey", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJSONKey(kapp.List[0], config, -1)
		case m.LblIsCopyCreateOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCopyCreateOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCopyCreateOp(kapp.List[0], config, -1)
		case m.LblXhashtake:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtake", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashtake(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTxNonceCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxNonceCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxNonceCellOpt(kapp.List[0], config, -1)
		case m.LblIsSubInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubInst(kapp.List[0], config, -1)
		case m.LblXhashfresh:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashfresh", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashfresh(kapp.List[0], config, -1)
		case m.LblXhashregRangeAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashregRangeAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashregRangeAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTypesCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTypesCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTypesCellOpt(kapp.List[0], config, -1)
		case m.LblXhashallBut64th:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashallBut64th", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashallBut64th(kapp.List[0], config, -1)
		case m.LblXuXstarIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXstarIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXstarIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInts(kapp.List[0], config, -1)
		case m.LblIsNparamsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNparamsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNparamsCellOpt(kapp.List[0], config, -1)
		case m.LblInitContractsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitContractsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitContractsCell(config, -1)
		case m.LblXdotMessageCellMap:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotMessageCellMap", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotMessageCellMap(config, -1)
		case m.LblXuXltXeqStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXeqStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXeqStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitCallerCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallerCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallerCell(config, -1)
		case m.LblXhashcontractBytesAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcontractBytesAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashcontractBytesAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashnumArgs:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashnumArgs", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashnumArgs(kapp.List[0], config, -1)
		case m.LblXhashdecodeLengthPrefixLengthAux:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdecodeLengthPrefixLengthAux", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashdecodeLengthPrefixLengthAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblListToInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalListToInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalListToInts(kapp.List[0], config, -1)
		case m.LblXhashchangesState:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashchangesState", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashchangesState(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsLocalCall:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalCall", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalCall(kapp.List[0], config, -1)
		case m.LblIsGlobalName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGlobalName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGlobalName(kapp.List[0], config, -1)
		case m.LblWord2Bool:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalWord2Bool", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalWord2Bool(kapp.List[0], config, -1)
		case m.LblIsAccountCellMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountCellMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountCellMap(kapp.List[0], config, -1)
		case m.LblInitLabelsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitLabelsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitLabelsCell(config, -1)
		case m.LblRegistersOperands:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRegistersOperands", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRegistersOperands(kapp.List[0], config, -1)
		case m.LblXhashdasmInstructionAux:
			if len(kapp.List) != 6 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmInstructionAux", expectedArity: 6, actualArity: len(kapp.List)}
			}
			return evalXhashdasmInstructionAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], config, -1)
		case m.LblInitCheckGasCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCheckGasCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCheckGasCell(config, -1)
		case m.LblIsAddInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAddInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAddInst(kapp.List[0], config, -1)
		case m.LblIsValidContractAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsValidContractAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalIsValidContractAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblUpdateArray:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUpdateArray", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalUpdateArray(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXlsqbXuXltXhyphenundefXrsqbXuARRAYXhyphenSYNTAX:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXltXhyphenundefXrsqbXuARRAYXhyphenSYNTAX", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXltXhyphenundefXrsqbXuARRAYXhyphenSYNTAX(kapp.List[0], kapp.List[1], config, -1)
		case m.LblSizeList:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSizeList", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSizeList(kapp.List[0], config, -1)
		case m.LblIsRefundCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsRefundCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsRefundCellOpt(kapp.List[0], config, -1)
		case m.LblIsNumberCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNumberCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNumberCellOpt(kapp.List[0], config, -1)
		case m.LblString2ID:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2ID", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalString2ID(kapp.List[0], config, -1)
		case m.LblInitDeclaredContractsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitDeclaredContractsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitDeclaredContractsCell(config, -1)
		case m.LblXuXeqXslashXeqBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXslashXeqBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXslashXeqBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashisValidStringTable:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidStringTable", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashisValidStringTable(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashcomputeNRegsAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcomputeNRegsAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashcomputeNRegsAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashsenderAux:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsenderAux", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashsenderAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXhashisCodeEmpty:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisCodeEmpty", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashisCodeEmpty(kapp.List[0], config, -1)
		case m.LblIsValidPoint:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsValidPoint", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsValidPoint(kapp.List[0], config, -1)
		case m.LblInitGasCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitGasCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitGasCell(config, -1)
		case m.LblIsJumpTableCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJumpTableCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJumpTableCellOpt(kapp.List[0], config, -1)
		case m.LblXuXltXltXuXgtXgtXuIELEXhyphenGAS:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXltXuXgtXgtXuIELEXhyphenGAS", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXltXuXgtXgtXuIELEXhyphenGAS(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXpercentlXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY:
			if len(kapp.List) != 5 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXpercentlXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY", expectedArity: 5, actualArity: len(kapp.List)}
			}
			return evalXpercentlXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], config, -1)
		case m.LblIsStorageCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStorageCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStorageCell(kapp.List[0], config, -1)
		case m.LblRlpDecode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRlpDecode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRlpDecode(kapp.List[0], config, -1)
		case m.LblRfindString:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRfindString", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalRfindString(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblUpdateList:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUpdateList", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalUpdateList(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblStringBuffer2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalStringBuffer2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalStringBuffer2String(kapp.List[0], config, -1)
		case m.LblPowmod:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPowmod", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalPowmod(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblCategoryChar:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCategoryChar", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalCategoryChar(kapp.List[0], config, -1)
		case m.LblIsJumpInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJumpInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJumpInst(kapp.List[0], config, -1)
		case m.LblIsCallerCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallerCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallerCellOpt(kapp.List[0], config, -1)
		case m.LblIsKCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsKCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsKCellOpt(kapp.List[0], config, -1)
		case m.LblInitSendtoCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitSendtoCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitSendtoCell(config, -1)
		case m.LblXhashdrop:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdrop", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashdrop(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFromCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFromCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFromCellOpt(kapp.List[0], config, -1)
		case m.LblIsCallStackCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallStackCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallStackCellOpt(kapp.List[0], config, -1)
		case m.LblInitCallValueCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallValueCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallValueCell(config, -1)
		case m.LblString2Float:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2Float", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalString2Float(kapp.List[0], config, -1)
		case m.LblMapXcolonlookupOrDefault:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMapXcolonlookupOrDefault", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalMapXcolonlookupOrDefault(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXampsIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXampsIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXampsIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsGasPriceCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasPriceCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasPriceCell(kapp.List[0], config, -1)
		case m.LblXuXplusXplusXuXuIELEXhyphenDATA:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusXplusXuXuIELEXhyphenDATA", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusXplusXuXuIELEXhyphenDATA(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsAcctIDCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAcctIDCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAcctIDCellOpt(kapp.List[0], config, -1)
		case m.LblIsInterimStatesCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInterimStatesCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInterimStatesCell(kapp.List[0], config, -1)
		case m.LblPadLeftBytes:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPadLeftBytes", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalPadLeftBytes(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXeqXslashXeqIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXslashXeqIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXslashXeqIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblLog2Int:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalLog2Int", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalLog2Int(kapp.List[0], config, -1)
		case m.LblIsJSON:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJSON", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJSON(kapp.List[0], config, -1)
		case m.LblIsMessageCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessageCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessageCellFragment(kapp.List[0], config, -1)
		case m.LblXhashstdinXuKXhyphenIO:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashstdinXuKXhyphenIO", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashstdinXuKXhyphenIO(config, -1)
		case m.LblIsPrecompiledOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPrecompiledOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPrecompiledOp(kapp.List[0], config, -1)
		case m.LblInitNetworkCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitNetworkCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitNetworkCell(config, -1)
		case m.LblIsID:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsID", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsID(kapp.List[0], config, -1)
		case m.LblIsFuncCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncCell(kapp.List[0], config, -1)
		case m.LblInitDifficultyCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitDifficultyCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitDifficultyCell(config, -1)
		case m.LblUnparseByteStack:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUnparseByteStack", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalUnparseByteStack(kapp.List[0], config, -1)
		case m.LblInitNparamsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitNparamsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitNparamsCell(config, -1)
		case m.LblInitProgramSizeCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitProgramSizeCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitProgramSizeCell(config, -1)
		case m.LblXhashmainContract:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashmainContract", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashmainContract(kapp.List[0], config, -1)
		case m.LblInitProgramCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitProgramCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitProgramCell(config, -1)
		case m.LblIsOriginCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOriginCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOriginCell(kapp.List[0], config, -1)
		case m.LblIsRevertInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsRevertInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsRevertInst(kapp.List[0], config, -1)
		case m.LblIsFuncIDCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncIDCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncIDCell(kapp.List[0], config, -1)
		case m.LblIsLoadInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLoadInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLoadInst(kapp.List[0], config, -1)
		case m.LblXhashparseByteStack:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseByteStack", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseByteStack(kapp.List[0], config, -1)
		case m.LblSrandInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSrandInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSrandInt(kapp.List[0], config, -1)
		case m.LblIntSizesAux:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntSizesAux", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalIntSizesAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXhashaddrXquesXlparenXuXrparenXuIELEXhyphenINFRASTRUCTURE:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashaddrXquesXlparenXuXrparenXuIELEXhyphenINFRASTRUCTURE", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashaddrXquesXlparenXuXrparenXuIELEXhyphenINFRASTRUCTURE(kapp.List[0], config, -1)
		case m.LblIsLocalMemCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalMemCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalMemCellOpt(kapp.List[0], config, -1)
		case m.LblIsLogDataCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLogDataCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLogDataCellOpt(kapp.List[0], config, -1)
		case m.LblIsGasUsedCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasUsedCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasUsedCellOpt(kapp.List[0], config, -1)
		case m.LblXhashcomputeJumpTableAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcomputeJumpTableAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashcomputeJumpTableAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsNumericIeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNumericIeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNumericIeleName(kapp.List[0], config, -1)
		case m.LblString2Base:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2Base", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalString2Base(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsGasLimitCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasLimitCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasLimitCellOpt(kapp.List[0], config, -1)
		case m.LblIsContractsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractsCell(kapp.List[0], config, -1)
		case m.LblStringIeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalStringIeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalStringIeleName(kapp.List[0], config, -1)
		case m.LblIsAccount:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccount", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccount(kapp.List[0], config, -1)
		case m.LblBN128AtePairing:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBN128AtePairing", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBN128AtePairing(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitFuncLabelsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFuncLabelsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFuncLabelsCell(config, -1)
		case m.LblIeleNameToken2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIeleNameToken2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIeleNameToken2String(kapp.List[0], config, -1)
		case m.LblIsFunctionBodiesCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionBodiesCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionBodiesCell(kapp.List[0], config, -1)
		case m.LblIsBalanceCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBalanceCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBalanceCellOpt(kapp.List[0], config, -1)
		case m.LblIsProgramCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsProgramCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsProgramCellFragment(kapp.List[0], config, -1)
		case m.LblNotBoolXu:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalNotBoolXu", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalNotBoolXu(kapp.List[0], config, -1)
		case m.LblXuXltXeqIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXeqIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXeqIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsLogDataCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLogDataCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLogDataCell(kapp.List[0], config, -1)
		case m.LblIsIELESimulation:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIELESimulation", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIELESimulation(kapp.List[0], config, -1)
		case m.LblXhashgetenv:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetenv", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetenv(kapp.List[0], config, -1)
		case m.LblIsEndianness:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsEndianness", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsEndianness(kapp.List[0], config, -1)
		case m.LblInitScheduleCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitScheduleCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitScheduleCell(kapp.List[0], config, -1)
		case m.LblIntersectSet:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntersectSet", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalIntersectSet(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFunctionCellMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionCellMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionCellMap(kapp.List[0], config, -1)
		case m.LblInitFunctionsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFunctionsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFunctionsCell(config, -1)
		case m.LblStringIeleName2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalStringIeleName2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalStringIeleName2String(kapp.List[0], config, -1)
		case m.LblXhashasUnsigned:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashasUnsigned", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashasUnsigned(kapp.List[0], config, -1)
		case m.LblXhashisValidInstruction:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidInstruction", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashisValidInstruction(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsFidCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFidCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFidCellOpt(kapp.List[0], config, -1)
		case m.LblIsOutputCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOutputCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOutputCell(kapp.List[0], config, -1)
		case m.LblInitFunctionNameCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFunctionNameCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFunctionNameCell(config, -1)
		case m.LblXuXlsqbXuXltXhyphenundefXrsqb:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXltXhyphenundefXrsqb", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXltXhyphenundefXrsqb(kapp.List[0], kapp.List[1], config, -1)
		case m.LblUnescape:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUnescape", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalUnescape(kapp.List[0], config, -1)
		case m.LblIsCallFrameCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallFrameCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallFrameCellFragment(kapp.List[0], config, -1)
		case m.LblXhashloadCode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadCode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashloadCode(kapp.List[0], config, -1)
		case m.LblXuinXukeysXlparenXuXrparenXuARRAYXhyphenSYNTAX:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuinXukeysXlparenXuXrparenXuARRAYXhyphenSYNTAX", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuinXukeysXlparenXuXrparenXuARRAYXhyphenSYNTAX(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXeqXeqIntXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXeqIntXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXeqIntXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuandThenBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuandThenBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuandThenBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashparseInModule:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseInModule", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashparseInModule(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsOriginCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOriginCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOriginCellOpt(kapp.List[0], config, -1)
		case m.LblIsCodeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCodeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCodeCell(kapp.List[0], config, -1)
		case m.LblInitLocalCallsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitLocalCallsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitLocalCallsCell(config, -1)
		case m.LblIsLogInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLogInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLogInst(kapp.List[0], config, -1)
		case m.LblExtractConfig:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalExtractConfig", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalExtractConfig(kapp.List[0], config, -1)
		case m.LblXuXpercentIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXpercentIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXpercentIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblCmem:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCmem", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCmem(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXgtXgtIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXgtXgtIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXgtXgtIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashtoBlockAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtoBlockAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashtoBlockAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblReplaceAllXlparenXuXcommaXuXcommaXuXrparenXuSTRING:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalReplaceAllXlparenXuXcommaXuXcommaXuXrparenXuSTRING", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalReplaceAllXlparenXuXcommaXuXcommaXuXrparenXuSTRING(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashrlpDecodeList:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpDecodeList", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrlpDecodeList(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsValueCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsValueCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsValueCellOpt(kapp.List[0], config, -1)
		case m.LblIsLabeledBlock:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLabeledBlock", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLabeledBlock(kapp.List[0], config, -1)
		case m.LblInitFunctionCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFunctionCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFunctionCell(config, -1)
		case m.LblIsGasLimitCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasLimitCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasLimitCell(kapp.List[0], config, -1)
		case m.LblXuXxorIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXxorIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXxorIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblFindString:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFindString", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalFindString(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblAbsInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalAbsInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalAbsInt(kapp.List[0], config, -1)
		case m.LblIsCallDataCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallDataCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallDataCellOpt(kapp.List[0], config, -1)
		case m.LblIsTxPendingCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxPendingCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxPendingCellOpt(kapp.List[0], config, -1)
		case m.LblXuXeqXeqStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXeqStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXeqStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsRegsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsRegsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsRegsCellOpt(kapp.List[0], config, -1)
		case m.LblCmul:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCmul", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalCmul(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsDeclaredContractsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDeclaredContractsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDeclaredContractsCell(kapp.List[0], config, -1)
		case m.LblInitTimestampCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTimestampCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTimestampCell(config, -1)
		case m.LblPow30XuIELEXhyphenDATA:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPow30XuIELEXhyphenDATA", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalPow30XuIELEXhyphenDATA(config, -1)
		case m.LblInitMessageCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitMessageCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitMessageCell(config, -1)
		case m.LblListXcolonget:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalListXcolonget", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalListXcolonget(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsReturnInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsReturnInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsReturnInst(kapp.List[0], config, -1)
		case m.LblSet2List:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSet2List", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSet2List(kapp.List[0], config, -1)
		case m.LblIsQuadOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsQuadOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsQuadOp(kapp.List[0], config, -1)
		case m.LblIsGeneratedTopCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGeneratedTopCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGeneratedTopCell(kapp.List[0], config, -1)
		case m.LblInitAccountsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitAccountsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitAccountsCell(config, -1)
		case m.LblIsSubstateLogEntry:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateLogEntry", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateLogEntry(kapp.List[0], config, -1)
		case m.LblInitTxGasLimitCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTxGasLimitCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTxGasLimitCell(config, -1)
		case m.LblXdotList:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotList", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotList(config, -1)
		case m.LblIsLValue:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLValue", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLValue(kapp.List[0], config, -1)
		case m.LblMakeEmptyArray:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMakeEmptyArray", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalMakeEmptyArray(kapp.List[0], config, -1)
		case m.LblIsSelfDestructCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSelfDestructCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSelfDestructCellOpt(kapp.List[0], config, -1)
		case m.LblIsGasCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasCellOpt(kapp.List[0], config, -1)
		case m.LblIsXorInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsXorInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsXorInst(kapp.List[0], config, -1)
		case m.LblIsJSONList:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsJSONList", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsJSONList(kapp.List[0], config, -1)
		case m.LblInitWellFormednessCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitWellFormednessCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitWellFormednessCell(kapp.List[0], config, -1)
		case m.LblListXcolonrange:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalListXcolonrange", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalListXcolonrange(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsMessageCellMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessageCellMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessageCellMap(kapp.List[0], config, -1)
		case m.LblIsPseudoInstruction:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPseudoInstruction", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPseudoInstruction(kapp.List[0], config, -1)
		case m.LblIsModeCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsModeCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsModeCellOpt(kapp.List[0], config, -1)
		case m.LblIsKCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsKCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsKCell(kapp.List[0], config, -1)
		case m.LblUnescapeAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUnescapeAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalUnescapeAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblInitMessagesCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitMessagesCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitMessagesCell(config, -1)
		case m.LblXdotFunctionCellMap:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotFunctionCellMap", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotFunctionCellMap(config, -1)
		case m.LblXuXgtXeqIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXgtXeqIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXgtXeqIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashtakeAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtakeAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashtakeAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsTxGasPriceCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxGasPriceCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxGasPriceCell(kapp.List[0], config, -1)
		case m.LblSizeWordStackAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSizeWordStackAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalSizeWordStackAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitGasUsedCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitGasUsedCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitGasUsedCell(config, -1)
		case m.LblFillList:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFillList", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalFillList(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblInitAcctIDCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitAcctIDCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitAcctIDCell(config, -1)
		case m.LblJSONListToInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalJSONListToInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalJSONListToInts(kapp.List[0], config, -1)
		case m.LblXuXltStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsNotInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNotInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNotInst(kapp.List[0], config, -1)
		case m.LblIsSelfDestructCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSelfDestructCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSelfDestructCell(kapp.List[0], config, -1)
		case m.LblXhashrlpEncodeWord:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeWord", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeWord(kapp.List[0], config, -1)
		case m.LblIsXhashLowerID:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsXhashLowerID", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsXhashLowerID(kapp.List[0], config, -1)
		case m.LblIsCurrentFunctionCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentFunctionCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentFunctionCellFragment(kapp.List[0], config, -1)
		case m.LblIsModInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsModInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsModInst(kapp.List[0], config, -1)
		case m.LblXhashlogToFile:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlogToFile", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashlogToFile(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitCallDepthCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallDepthCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallDepthCell(config, -1)
		case m.LblXhashreadXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashreadXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashreadXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashrlpEncodeLengthAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeLengthAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeLengthAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsContractCodeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractCodeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractCodeCell(kapp.List[0], config, -1)
		case m.LblXhashcallAddressAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcallAddressAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashcallAddressAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsHexConstant:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsHexConstant", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsHexConstant(kapp.List[0], config, -1)
		case m.LblID2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalID2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalID2String(kapp.List[0], config, -1)
		case m.LblInitJumpTableCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitJumpTableCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitJumpTableCell(config, -1)
		case m.LblIsCurrentContractCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentContractCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentContractCellOpt(kapp.List[0], config, -1)
		case m.LblMapXcolonchoice:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMapXcolonchoice", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalMapXcolonchoice(kapp.List[0], config, -1)
		case m.LblIsMessageCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessageCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessageCell(kapp.List[0], config, -1)
		case m.LblIsPreviousGasCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPreviousGasCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPreviousGasCell(kapp.List[0], config, -1)
		case m.LblIsTypesCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTypesCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTypesCell(kapp.List[0], config, -1)
		case m.LblXhashsizeNames:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeNames", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsizeNames(kapp.List[0], config, -1)
		case m.LblCxfer:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCxfer", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCxfer(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXplusBytesXuXuBYTESXhyphenHOOKED:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusBytesXuXuBYTESXhyphenHOOKED", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusBytesXuXuBYTESXhyphenHOOKED(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFuncLabelsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncLabelsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncLabelsCell(kapp.List[0], config, -1)
		case m.LblIsBool:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBool", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBool(kapp.List[0], config, -1)
		case m.LblInitValueCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitValueCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitValueCell(config, -1)
		case m.LblXtildeIntXuXuINT:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXtildeIntXuXuINT", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXtildeIntXuXuINT(kapp.List[0], config, -1)
		case m.LblInitCurrentMemoryCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCurrentMemoryCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCurrentMemoryCell(config, -1)
		case m.LblOrdChar:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalOrdChar", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalOrdChar(kapp.List[0], config, -1)
		case m.LblInitIDCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitIDCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitIDCell(config, -1)
		case m.LblInitMsgIDCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitMsgIDCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitMsgIDCell(config, -1)
		case m.LblIntSizes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntSizes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIntSizes(kapp.List[0], config, -1)
		case m.LblInitKCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitKCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitKCell(kapp.List[0], config, -1)
		case m.LblIsDataCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDataCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDataCell(kapp.List[0], config, -1)
		case m.LblIsNetworkCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNetworkCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNetworkCellFragment(kapp.List[0], config, -1)
		case m.LblBytesRange:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBytesRange", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalBytesRange(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsStringBuffer:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStringBuffer", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStringBuffer(kapp.List[0], config, -1)
		case m.LblRemoveAll:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRemoveAll", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalRemoveAll(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuandBoolXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuandBoolXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuandBoolXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsCallDepthCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallDepthCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallDepthCellOpt(kapp.List[0], config, -1)
		case m.LblRlpEncodeInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRlpEncodeInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRlpEncodeInts(kapp.List[0], config, -1)
		case m.LblIsProgramCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsProgramCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsProgramCell(kapp.List[0], config, -1)
		case m.LblIsUnlabeledBlock:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsUnlabeledBlock", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsUnlabeledBlock(kapp.List[0], config, -1)
		case m.LblIsShiftInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsShiftInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsShiftInst(kapp.List[0], config, -1)
		case m.LblCextra:
			if len(kapp.List) != 5 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCextra", expectedArity: 5, actualArity: len(kapp.List)}
			}
			return evalCextra(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], config, -1)
		case m.LblIsNonEmptyOperands:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNonEmptyOperands", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNonEmptyOperands(kapp.List[0], config, -1)
		case m.LblXhashdecodeLengthPrefix:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdecodeLengthPrefix", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashdecodeLengthPrefix(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsCurrentContractCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentContractCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentContractCell(kapp.List[0], config, -1)
		case m.LblCheckInit:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCheckInit", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCheckInit(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXudividesIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXudividesIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXudividesIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsException:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsException", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsException(kapp.List[0], config, -1)
		case m.LblIsCurrentMemoryCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentMemoryCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentMemoryCell(kapp.List[0], config, -1)
		case m.LblSetXcolonchoice:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSetXcolonchoice", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSetXcolonchoice(kapp.List[0], config, -1)
		case m.LblInitPreviousGasCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitPreviousGasCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitPreviousGasCell(config, -1)
		case m.LblXhashinvalidXquesXlsqbXuXrsqbXuIELE:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashinvalidXquesXlsqbXuXrsqbXuIELE", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashinvalidXquesXlsqbXuXrsqbXuIELE(kapp.List[0], config, -1)
		case m.LblXhashparseHexWord:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseHexWord", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseHexWord(kapp.List[0], config, -1)
		case m.LblXhashaddr:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashaddr", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashaddr(kapp.List[0], config, -1)
		case m.LblIsCondJumpInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCondJumpInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCondJumpInst(kapp.List[0], config, -1)
		case m.LblXhashgetStorageData:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetStorageData", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashgetStorageData(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsStaticCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStaticCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStaticCellOpt(kapp.List[0], config, -1)
		case m.LblIsCallAddressInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallAddressInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallAddressInst(kapp.List[0], config, -1)
		case m.LblIsExitCodeCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExitCodeCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExitCodeCellOpt(kapp.List[0], config, -1)
		case m.LblIsSCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSCellOpt(kapp.List[0], config, -1)
		case m.LblXhashrlpDecodeAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpDecodeAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrlpDecodeAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitTxPendingCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTxPendingCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTxPendingCell(config, -1)
		case m.LblInitCurrentInstructionsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCurrentInstructionsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCurrentInstructionsCell(config, -1)
		case m.LblIsLocalCallOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalCallOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalCallOp(kapp.List[0], config, -1)
		case m.LblXhashlambdaXuXu:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlambdaXuXu", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashlambdaXuXu(kapp.List[0], config, -1)
		case m.LblFreshInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFreshInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalFreshInt(kapp.List[0], config, -1)
		case m.LblXhashwriteXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashwriteXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashwriteXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsDifficultyCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDifficultyCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDifficultyCell(kapp.List[0], config, -1)
		case m.LblIsNregsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNregsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNregsCell(kapp.List[0], config, -1)
		case m.LblXhashloadOffset:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadOffset", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashloadOffset(kapp.List[0], config, -1)
		case m.LblIsSExtInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSExtInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSExtInst(kapp.List[0], config, -1)
		case m.LblXhashsenderAux2:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsenderAux2", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsenderAux2(kapp.List[0], config, -1)
		case m.LblXpercentoXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY:
			if len(kapp.List) != 5 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXpercentoXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY", expectedArity: 5, actualArity: len(kapp.List)}
			}
			return evalXpercentoXlparenXuXcommaXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], config, -1)
		case m.LblIsNonceCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNonceCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNonceCell(kapp.List[0], config, -1)
		case m.LblIsNullOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNullOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNullOp(kapp.List[0], config, -1)
		case m.LblList2Set:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalList2Set", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalList2Set(kapp.List[0], config, -1)
		case m.LblXuXltXltIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXltIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXltIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitCallDataCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallDataCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallDataCell(config, -1)
		case m.LblIsMulInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMulInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMulInst(kapp.List[0], config, -1)
		case m.LblIsAcctIDCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAcctIDCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAcctIDCell(kapp.List[0], config, -1)
		case m.LblBitsInWords:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBitsInWords", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBitsInWords(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsAndInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAndInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAndInst(kapp.List[0], config, -1)
		case m.LblIsFuncLabelsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncLabelsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncLabelsCellOpt(kapp.List[0], config, -1)
		case m.LblFunctionCellMapItem:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFunctionCellMapItem", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalFunctionCellMapItem(kapp.List[0], kapp.List[1], config, -1)
		case m.LblCcallarg:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCcallarg", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCcallarg(kapp.List[0], kapp.List[1], config, -1)
		case m.LblSha256:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSha256", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSha256(kapp.List[0], config, -1)
		case m.LblIsNparamsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNparamsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNparamsCell(kapp.List[0], config, -1)
		case m.LblIsFiveOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFiveOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFiveOp(kapp.List[0], config, -1)
		case m.LblXhashloadFunction:
			if len(kapp.List) != 6 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashloadFunction", expectedArity: 6, actualArity: len(kapp.List)}
			}
			return evalXhashloadFunction(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], config, -1)
		case m.LblBytes2Int:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBytes2Int", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalBytes2Int(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblString2Bytes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2Bytes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalString2Bytes(kapp.List[0], config, -1)
		case m.LblXuFunctionCellMapXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuFunctionCellMapXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuFunctionCellMapXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsWellFormednessCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWellFormednessCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWellFormednessCellOpt(kapp.List[0], config, -1)
		case m.LblMinIntXlparenXuXcommaXuXrparenXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMinIntXlparenXuXcommaXuXrparenXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMinIntXlparenXuXcommaXuXrparenXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMap(kapp.List[0], config, -1)
		case m.LblInitRegsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitRegsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitRegsCell(config, -1)
		case m.LblXuXltXuXgtXuIELEXhyphenGAS:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXuXgtXuIELEXhyphenGAS", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXuXgtXuIELEXhyphenGAS(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsXhashRuleTag:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsXhashRuleTag", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsXhashRuleTag(kapp.List[0], config, -1)
		case m.LblXuXltXltByteXuXuIELEXhyphenDATA:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXltByteXuXuIELEXhyphenDATA", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXltByteXuXuIELEXhyphenDATA(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsExportedCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExportedCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExportedCell(kapp.List[0], config, -1)
		case m.LblReplaceXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuSTRING:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalReplaceXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuSTRING", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalReplaceXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuSTRING(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXhashtellXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtellXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashtellXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblLengthBytes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalLengthBytes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalLengthBytes(kapp.List[0], config, -1)
		case m.LblInitBlockhashCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitBlockhashCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitBlockhashCell(config, -1)
		case m.LblXhashunparseByteStack:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashunparseByteStack", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashunparseByteStack(kapp.List[0], config, -1)
		case m.LblIsIeleCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIeleCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIeleCell(kapp.List[0], config, -1)
		case m.LblInitOutputCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitOutputCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitOutputCell(config, -1)
		case m.LblXhashcomputeJumpTable:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcomputeJumpTable", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashcomputeJumpTable(kapp.List[0], config, -1)
		case m.LblXhashisValidContract:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidContract", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashisValidContract(kapp.List[0], config, -1)
		case m.LblXhashgetNonce:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetNonce", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetNonce(kapp.List[0], config, -1)
		case m.LblXhashseekXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashseekXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashseekXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsLengthPrefix:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLengthPrefix", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLengthPrefix(kapp.List[0], config, -1)
		case m.LblSignExtendBitRangeInt:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSignExtendBitRangeInt", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalSignExtendBitRangeInt(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsWellFormednessCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWellFormednessCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWellFormednessCell(kapp.List[0], config, -1)
		case m.LblXuXeqXeqBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXeqBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXeqBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsSet:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSet", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSet(kapp.List[0], config, -1)
		case m.LblInitExportedCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitExportedCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitExportedCell(config, -1)
		case m.LblIntrinsicTypesXuIELEXhyphenWELLXhyphenFORMEDNESS:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntrinsicTypesXuIELEXhyphenWELLXhyphenFORMEDNESS", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalIntrinsicTypesXuIELEXhyphenWELLXhyphenFORMEDNESS(config, -1)
		case m.LblInitContractNameCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitContractNameCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitContractNameCell(config, -1)
		case m.LblIsAccountCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountCell(kapp.List[0], config, -1)
		case m.LblXhashparse:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparse", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashparse(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsSendtoCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSendtoCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSendtoCellOpt(kapp.List[0], config, -1)
		case m.LblIsFunctionsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionsCellOpt(kapp.List[0], config, -1)
		case m.LblIsLengthPrefixType:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLengthPrefixType", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLengthPrefixType(kapp.List[0], config, -1)
		case m.LblRegistersLValues:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRegistersLValues", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRegistersLValues(kapp.List[0], config, -1)
		case m.LblByte:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalByte", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalByte(kapp.List[0], kapp.List[1], config, -1)
		case m.LblBN128Add:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBN128Add", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBN128Add(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitCurrentContractCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCurrentContractCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCurrentContractCell(config, -1)
		case m.LblXhashisValidFunctions:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidFunctions", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashisValidFunctions(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXcolonXslashXeqKXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXcolonXslashXeqKXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXcolonXslashXeqKXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblRipEmd160:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRipEmd160", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRipEmd160(kapp.List[0], config, -1)
		case m.LblIsExportedCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExportedCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExportedCellOpt(kapp.List[0], config, -1)
		case m.LblCnew:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCnew", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalCnew(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsStaticCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStaticCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStaticCell(kapp.List[0], config, -1)
		case m.LblInitBalanceCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitBalanceCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitBalanceCell(config, -1)
		case m.LblPadRightBytes:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPadRightBytes", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalPadRightBytes(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblReverseBytes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalReverseBytes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalReverseBytes(kapp.List[0], config, -1)
		case m.LblCgascap:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCgascap", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalCgascap(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblInitInterimStatesCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitInterimStatesCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitInterimStatesCell(config, -1)
		case m.LblIsStringIeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStringIeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStringIeleName(kapp.List[0], config, -1)
		case m.LblIsSubstateStackCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateStackCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateStackCellOpt(kapp.List[0], config, -1)
		case m.LblGcdInt:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalGcdInt", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalGcdInt(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsOperand:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOperand", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOperand(kapp.List[0], config, -1)
		case m.LblIsKConfigVar:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsKConfigVar", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsKConfigVar(kapp.List[0], config, -1)
		case m.LblIsGasCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasCell(kapp.List[0], config, -1)
		case m.LblIsSubstateCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateCellOpt(kapp.List[0], config, -1)
		case m.LblIsGasUsedCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasUsedCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasUsedCell(kapp.List[0], config, -1)
		case m.LblXhashparseByteStackAux:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseByteStackAux", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashparseByteStackAux(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblInitFunctionBodiesCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFunctionBodiesCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFunctionBodiesCell(config, -1)
		case m.LblSubstrBytes:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSubstrBytes", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalSubstrBytes(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXltIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitGasPriceCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitGasPriceCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitGasPriceCell(config, -1)
		case m.LblIsCallDataCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallDataCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallDataCell(kapp.List[0], config, -1)
		case m.LblIsMessagesCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessagesCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessagesCellFragment(kapp.List[0], config, -1)
		case m.LblChrChar:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalChrChar", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalChrChar(kapp.List[0], config, -1)
		case m.LblXudivIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXudivIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXudivIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsSLoadInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSLoadInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSLoadInst(kapp.List[0], config, -1)
		case m.LblInitTypeCheckingCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTypeCheckingCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTypeCheckingCell(config, -1)
		case m.LblIsSelfdestructInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSelfdestructInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSelfdestructInst(kapp.List[0], config, -1)
		case m.LblInitFromCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFromCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFromCell(config, -1)
		case m.LblIsIsZeroInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIsZeroInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIsZeroInst(kapp.List[0], config, -1)
		case m.LblIsCallFrameCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallFrameCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallFrameCellOpt(kapp.List[0], config, -1)
		case m.LblXuorBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuorBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuorBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblUpdateMap:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalUpdateMap", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalUpdateMap(kapp.List[0], kapp.List[1], config, -1)
		case m.LblCeilDiv:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCeilDiv", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCeilDiv(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInt2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInt2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInt2String(kapp.List[0], config, -1)
		case m.LblXuXeqXslashXeqKXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXeqXslashXeqKXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXeqXslashXeqKXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsInstructionsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInstructionsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInstructionsCell(kapp.List[0], config, -1)
		case m.LblIsScheduleCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsScheduleCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsScheduleCellOpt(kapp.List[0], config, -1)
		case m.LblIsLocalName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalName(kapp.List[0], config, -1)
		case m.LblIsInstruction:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInstruction", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInstruction(kapp.List[0], config, -1)
		case m.LblXhashopenXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashopenXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashopenXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitOriginCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitOriginCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitOriginCell(config, -1)
		case m.LblGXstarXlparenXuXcommaXuXcommaXuXrparenXuIELEXhyphenGAS:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalGXstarXlparenXuXcommaXuXcommaXuXrparenXuIELEXhyphenGAS", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalGXstarXlparenXuXcommaXuXcommaXuXrparenXuIELEXhyphenGAS(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXpipeXhyphenXgtXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXpipeXhyphenXgtXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXpipeXhyphenXgtXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitFuncCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFuncCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFuncCell(config, -1)
		case m.LblString2IeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2IeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalString2IeleName(kapp.List[0], config, -1)
		case m.LblXhashopWidth:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashopWidth", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashopWidth(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashparseWord:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseWord", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseWord(kapp.List[0], config, -1)
		case m.LblXhashprecompiledXuIELEXhyphenPRECOMPILED:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashprecompiledXuIELEXhyphenPRECOMPILED", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashprecompiledXuIELEXhyphenPRECOMPILED(config, -1)
		case m.LblXhashdasmFunctions:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmFunctions", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashdasmFunctions(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsContract:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContract", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContract(kapp.List[0], config, -1)
		case m.LblXhashoverApproxKara:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashoverApproxKara", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashoverApproxKara(kapp.List[0], config, -1)
		case m.LblXhashputcXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashputcXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashputcXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTxGasLimitCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxGasLimitCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxGasLimitCellOpt(kapp.List[0], config, -1)
		case m.LblAssignBytesRange:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalAssignBytesRange", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalAssignBytesRange(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsIELECommand:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIELECommand", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIELECommand(kapp.List[0], config, -1)
		case m.LblXuXlsqbXuXltXhyphenXuXrsqb:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXltXhyphenXuXrsqb", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXltXhyphenXuXrsqb(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashremoveZerosAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashremoveZerosAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashremoveZerosAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInt(kapp.List[0], config, -1)
		case m.LblIsPredicate:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPredicate", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPredicate(kapp.List[0], config, -1)
		case m.LblIsPreviousGasCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPreviousGasCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPreviousGasCellOpt(kapp.List[0], config, -1)
		case m.LblXdotBytesXuBYTESXhyphenHOOKED:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotBytesXuBYTESXhyphenHOOKED", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotBytesXuBYTESXhyphenHOOKED(config, -1)
		case m.LblIsAccountsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountsCellOpt(kapp.List[0], config, -1)
		case m.LblXuimpliesBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuimpliesBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuimpliesBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsIeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIeleName(kapp.List[0], config, -1)
		case m.LblInitTxNonceCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTxNonceCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTxNonceCell(config, -1)
		case m.LblIsExitCodeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExitCodeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExitCodeCell(kapp.List[0], config, -1)
		case m.LblMaxIntXlparenXuXcommaXuXrparenXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMaxIntXlparenXuXcommaXuXrparenXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMaxIntXlparenXuXcommaXuXrparenXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblFillArray:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFillArray", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalFillArray(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsActiveAccountsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsActiveAccountsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsActiveAccountsCell(kapp.List[0], config, -1)
		case m.LblXhashsizeRegs:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeRegs", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsizeRegs(kapp.List[0], config, -1)
		case m.LblXuMapXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuMapXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuMapXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXhyphenIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXhyphenIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXhyphenIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsBalanceCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBalanceCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBalanceCell(kapp.List[0], config, -1)
		case m.LblFloat2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFloat2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalFloat2String(kapp.List[0], config, -1)
		case m.LblBN128Mul:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBN128Mul", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBN128Mul(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsReturnType:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsReturnType", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsReturnType(kapp.List[0], config, -1)
		case m.LblInt2Bytes:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInt2Bytes", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalInt2Bytes(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXcolonXeqKXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXcolonXeqKXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXcolonXeqKXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsBeneficiaryCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBeneficiaryCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBeneficiaryCellOpt(kapp.List[0], config, -1)
		case m.LblIeleName2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIeleName2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIeleName2String(kapp.List[0], config, -1)
		case m.LblInitGasLimitCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitGasLimitCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitGasLimitCell(config, -1)
		case m.LblXhashadjustedBitLength:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashadjustedBitLength", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashadjustedBitLength(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitTypesCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTypesCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTypesCell(config, -1)
		case m.LblIsSubstateStackCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateStackCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateStackCell(kapp.List[0], config, -1)
		case m.LblInitWellFormednessScheduleCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitWellFormednessScheduleCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitWellFormednessScheduleCell(kapp.List[0], config, -1)
		case m.LblIsContractCodeCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractCodeCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractCodeCellOpt(kapp.List[0], config, -1)
		case m.LblIsG1Point:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsG1Point", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsG1Point(kapp.List[0], config, -1)
		case m.LblIsCurrentInstructionsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentInstructionsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentInstructionsCell(kapp.List[0], config, -1)
		case m.LblIsLocalMemCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalMemCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalMemCell(kapp.List[0], config, -1)
		case m.LblIsSCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSCell(kapp.List[0], config, -1)
		case m.LblInitSCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitSCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitSCell(kapp.List[0], config, -1)
		case m.LblXhashsubcontract:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsubcontract", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashsubcontract(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXxorXpercentIntXuXuXuINT:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXxorXpercentIntXuXuXuINT", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXuXxorXpercentIntXuXuXuINT(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsDifficultyCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDifficultyCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDifficultyCellOpt(kapp.List[0], config, -1)
		case m.LblIsWellFormednessScheduleCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWellFormednessScheduleCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWellFormednessScheduleCell(kapp.List[0], config, -1)
		case m.LblXhashstatXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashstatXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashstatXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblIsFunctionsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionsCell(kapp.List[0], config, -1)
		case m.LblIsTopLevelDefinitions:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTopLevelDefinitions", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTopLevelDefinitions(kapp.List[0], config, -1)
		case m.LblSetXcolondifference:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSetXcolondifference", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalSetXcolondifference(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsCallOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallOp(kapp.List[0], config, -1)
		case m.LblIsProgramSizeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsProgramSizeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsProgramSizeCell(kapp.List[0], config, -1)
		case m.LblIsStrategy:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStrategy", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStrategy(kapp.List[0], config, -1)
		case m.LblInitCodeCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCodeCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCodeCell(config, -1)
		case m.LblXhashifXuXhashthenXuXhashelseXuXhashfiXuKXhyphenEQUAL:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashifXuXhashthenXuXhashelseXuXhashfiXuKXhyphenEQUAL", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashifXuXhashthenXuXhashelseXuXhashfiXuKXhyphenEQUAL(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXplusXdotXplusIeleNameXuXuIELEXhyphenBINARY:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusXdotXplusIeleNameXuXuIELEXhyphenBINARY", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusXdotXplusIeleNameXuXuIELEXhyphenBINARY(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashdasmOpCode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmOpCode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashdasmOpCode(kapp.List[0], config, -1)
		case m.LblXhashstdoutXuKXhyphenIO:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashstdoutXuKXhyphenIO", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashstdoutXuKXhyphenIO(config, -1)
		case m.LblInitActiveAccountsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitActiveAccountsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitActiveAccountsCell(config, -1)
		case m.LblXhashrlpDecodeListAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpDecodeListAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashrlpDecodeListAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsFunctionNameCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionNameCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionNameCell(kapp.List[0], config, -1)
		case m.LblIsCheckGasCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCheckGasCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCheckGasCellOpt(kapp.List[0], config, -1)
		case m.LblXuXgtXeqStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXgtXeqStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXgtXeqStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashcallAddress:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcallAddress", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashcallAddress(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblAssignWordStackRange:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalAssignWordStackRange", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalAssignWordStackRange(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblSizeMap:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSizeMap", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSizeMap(kapp.List[0], config, -1)
		case m.LblIsSubstateCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateCell(kapp.List[0], config, -1)
		case m.LblXhashsizeLVals:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeLVals", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsizeLVals(kapp.List[0], config, -1)
		case m.LblG0create:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalG0create", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalG0create(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblSubstrString:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSubstrString", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalSubstrString(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsCurrentFunctionCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentFunctionCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentFunctionCellOpt(kapp.List[0], config, -1)
		case m.LblIsGeneratedTopCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGeneratedTopCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGeneratedTopCellFragment(kapp.List[0], config, -1)
		case m.LblSize:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSize", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSize(kapp.List[0], config, -1)
		case m.LblIsCallStackCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallStackCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallStackCell(kapp.List[0], config, -1)
		case m.LblIsDivInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDivInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDivInst(kapp.List[0], config, -1)
		case m.LblCpricedmem:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCpricedmem", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCpricedmem(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFunctionsCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionsCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionsCellFragment(kapp.List[0], config, -1)
		case m.LblInitFidCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFidCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFidCell(config, -1)
		case m.LblInitCallStackCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCallStackCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCallStackCell(config, -1)
		case m.LblXhashrlpEncodeBytes:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeBytes", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeBytes(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsIeleCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIeleCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIeleCellOpt(kapp.List[0], config, -1)
		case m.LblIsWellFormednessCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWellFormednessCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWellFormednessCellFragment(kapp.List[0], config, -1)
		case m.LblIsAccountCallInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountCallInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountCallInst(kapp.List[0], config, -1)
		case m.LblIsCreateOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCreateOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCreateOp(kapp.List[0], config, -1)
		case m.LblIsG2Point:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsG2Point", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsG2Point(kapp.List[0], config, -1)
		case m.LblIsIeleCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIeleCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIeleCellFragment(kapp.List[0], config, -1)
		case m.LblIsXhashUpperID:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsXhashUpperID", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsXhashUpperID(kapp.List[0], config, -1)
		case m.LblInitTxOrderCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTxOrderCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTxOrderCell(config, -1)
		case m.LblIsKItem:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsKItem", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsKItem(kapp.List[0], config, -1)
		case m.LblIsStoreInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStoreInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStoreInst(kapp.List[0], config, -1)
		case m.LblListXcolonset:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalListXcolonset", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalListXcolonset(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblKeys:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalKeys", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalKeys(kapp.List[0], config, -1)
		case m.LblIsMessagesCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessagesCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessagesCellOpt(kapp.List[0], config, -1)
		case m.LblIsTxOrderCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxOrderCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxOrderCellOpt(kapp.List[0], config, -1)
		case m.LblXhashprecompiledAccountXuIELEXhyphenPRECOMPILED:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashprecompiledAccountXuIELEXhyphenPRECOMPILED", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashprecompiledAccountXuIELEXhyphenPRECOMPILED(config, -1)
		case m.LblIsCreateInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCreateInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCreateInst(kapp.List[0], config, -1)
		case m.LblXhashdasmLoad:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmLoad", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashdasmLoad(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsFunctionDefinition:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionDefinition", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionDefinition(kapp.List[0], config, -1)
		case m.LblBswap:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBswap", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBswap(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsCheckGasCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCheckGasCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCheckGasCell(kapp.List[0], config, -1)
		case m.LblIsBytes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBytes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBytes(kapp.List[0], config, -1)
		case m.LblIsValidG2Point:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsValidG2Point", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsValidG2Point(kapp.List[0], config, -1)
		case m.LblXhashlambdaXuXu2:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlambdaXuXu2", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXhashlambdaXuXu2(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXhashstderrXuKXhyphenIO:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashstderrXuKXhyphenIO", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXhashstderrXuKXhyphenIO(config, -1)
		case m.LblInitExitCodeCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitExitCodeCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitExitCodeCell(config, -1)
		case m.LblXuinXukeysXlparenXuXrparenXuMAP:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuinXukeysXlparenXuXrparenXuMAP", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuinXukeysXlparenXuXrparenXuMAP(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitFuncIDsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFuncIDsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFuncIDsCell(config, -1)
		case m.LblFindChar:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFindChar", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalFindChar(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblSetXcolonin:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSetXcolonin", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalSetXcolonin(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsK:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsK", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsK(kapp.List[0], config, -1)
		case m.LblIsScheduleFlag:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsScheduleFlag", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsScheduleFlag(kapp.List[0], config, -1)
		case m.LblString2Int:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalString2Int", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalString2Int(kapp.List[0], config, -1)
		case m.LblInitStorageCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitStorageCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitStorageCell(config, -1)
		case m.LblBytesInWords:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBytesInWords", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalBytesInWords(kapp.List[0], config, -1)
		case m.LblCexp:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCexp", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalCexp(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblIsCurrentFunctionCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCurrentFunctionCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCurrentFunctionCell(kapp.List[0], config, -1)
		case m.LblIsLocalCallsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalCallsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalCallsCellOpt(kapp.List[0], config, -1)
		case m.LblXhashsizeLValuesAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeLValuesAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashsizeLValuesAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashrlpEncodeLength:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeLength", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeLength(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashnewAddr:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashnewAddr", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashnewAddr(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsScheduleConst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsScheduleConst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsScheduleConst(kapp.List[0], config, -1)
		case m.LblXhashsystem:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsystem", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsystem(kapp.List[0], config, -1)
		case m.LblIsString:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsString", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsString(kapp.List[0], config, -1)
		case m.LblIsGasPriceCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGasPriceCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGasPriceCellOpt(kapp.List[0], config, -1)
		case m.LblIsList:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsList", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsList(kapp.List[0], config, -1)
		case m.LblInitRefundCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitRefundCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitRefundCell(config, -1)
		case m.LblIsFunctionCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionCellFragment(kapp.List[0], config, -1)
		case m.LblXuXlsqbXuXdotXdotXuXrsqbXuIELEXhyphenDATA:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXdotXdotXuXrsqbXuIELEXhyphenDATA", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXdotXdotXuXrsqbXuIELEXhyphenDATA(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashlambdaXuXu4:
			if len(kapp.List) != 11 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlambdaXuXu4", expectedArity: 11, actualArity: len(kapp.List)}
			}
			return evalXhashlambdaXuXu4(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], kapp.List[6], kapp.List[7], kapp.List[8], kapp.List[9], kapp.List[10], config, -1)
		case m.LblInitNumberCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitNumberCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitNumberCell(config, -1)
		case m.LblIsContractsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractsCellOpt(kapp.List[0], config, -1)
		case m.LblXuXgtStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXgtStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXgtStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsModeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsModeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsModeCell(kapp.List[0], config, -1)
		case m.LblXhashlambdaXuXu3:
			if len(kapp.List) != 10 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlambdaXuXu3", expectedArity: 10, actualArity: len(kapp.List)}
			}
			return evalXhashlambdaXuXu3(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], kapp.List[6], kapp.List[7], kapp.List[8], kapp.List[9], config, -1)
		case m.LblInitStaticCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitStaticCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitStaticCell(config, -1)
		case m.LblInitLocalMemCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitLocalMemCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitLocalMemCell(config, -1)
		case m.LblXhashgetBlockhash:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetBlockhash", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetBlockhash(kapp.List[0], config, -1)
		case m.LblXhashregisters:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashregisters", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashregisters(kapp.List[0], config, -1)
		case m.LblIsKResult:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsKResult", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsKResult(kapp.List[0], config, -1)
		case m.LblIsOpCode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOpCode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOpCode(kapp.List[0], config, -1)
		case m.LblIsOperands:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsOperands", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsOperands(kapp.List[0], config, -1)
		case m.LblKeccak256:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalKeccak256", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalKeccak256(kapp.List[0], config, -1)
		case m.LblIsByteInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsByteInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsByteInst(kapp.List[0], config, -1)
		case m.LblIsInterimStatesCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInterimStatesCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInterimStatesCellOpt(kapp.List[0], config, -1)
		case m.LblXhashlstatXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlstatXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashlstatXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblSetItem:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalSetItem", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalSetItem(kapp.List[0], config, -1)
		case m.LblIsIeleBuiltin:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIeleBuiltin", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIeleBuiltin(kapp.List[0], config, -1)
		case m.LblIsTernOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTernOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTernOp(kapp.List[0], config, -1)
		case m.LblXhashdasmFunction:
			if len(kapp.List) != 9 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmFunction", expectedArity: 9, actualArity: len(kapp.List)}
			}
			return evalXhashdasmFunction(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], kapp.List[6], kapp.List[7], kapp.List[8], config, -1)
		case m.LblIsAddModInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAddModInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAddModInst(kapp.List[0], config, -1)
		case m.LblRandInt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRandInt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalRandInt(kapp.List[0], config, -1)
		case m.LblIntSizesArr:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIntSizesArr", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalIntSizesArr(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashcloseXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcloseXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashcloseXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblIsTimestampCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTimestampCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTimestampCell(kapp.List[0], config, -1)
		case m.LblKeysXulistXlparenXuXrparenXuMAP:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalKeysXulistXlparenXuXrparenXuMAP", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalKeysXulistXlparenXuXrparenXuMAP(kapp.List[0], config, -1)
		case m.LblFreshID:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFreshID", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalFreshID(kapp.List[0], config, -1)
		case m.LblXuorElseBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuorElseBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuorElseBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsPeakMemoryCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPeakMemoryCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPeakMemoryCellOpt(kapp.List[0], config, -1)
		case m.LblXhashregRange:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashregRange", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashregRange(kapp.List[0], config, -1)
		case m.LblXhashisValidFunction:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashisValidFunction", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashisValidFunction(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXdotAccountCellMap:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotAccountCellMap", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotAccountCellMap(config, -1)
		case m.LblIsSignedness:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSignedness", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSignedness(kapp.List[0], config, -1)
		case m.LblInitNregsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitNregsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitNregsCell(config, -1)
		case m.LblXhashopCodeWidth:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashopCodeWidth", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashopCodeWidth(kapp.List[0], config, -1)
		case m.LblPow256XuIELEXhyphenDATA:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPow256XuIELEXhyphenDATA", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalPow256XuIELEXhyphenDATA(config, -1)
		case m.LblXhashlockXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashlockXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashlockXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsIntConstant:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIntConstant", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIntConstant(kapp.List[0], config, -1)
		case m.LblCountAllOccurrencesXlparenXuXcommaXuXrparenXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCountAllOccurrencesXlparenXuXcommaXuXrparenXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCountAllOccurrencesXlparenXuXcommaXuXrparenXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXgtIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXgtIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXgtIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsBeneficiaryCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBeneficiaryCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBeneficiaryCell(kapp.List[0], config, -1)
		case m.LblCkara:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCkara", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalCkara(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitSubstateCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitSubstateCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitSubstateCell(config, -1)
		case m.LblIsBlocks:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBlocks", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBlocks(kapp.List[0], config, -1)
		case m.LblXhashgcdInt:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgcdInt", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashgcdInt(kapp.List[0], kapp.List[1], config, -1)
		case m.LblBitRangeInt:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBitRangeInt", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalBitRangeInt(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsProgramSizeCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsProgramSizeCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsProgramSizeCellOpt(kapp.List[0], config, -1)
		case m.LblXhashunparseByteStackAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashunparseByteStackAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashunparseByteStackAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuxorBoolXuXuBOOL:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuxorBoolXuXuBOOL", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuxorBoolXuXuBOOL(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashrlpEncodeString:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashrlpEncodeString", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashrlpEncodeString(kapp.List[0], config, -1)
		case m.LblInitGeneratedTopCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitGeneratedTopCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitGeneratedTopCell(kapp.List[0], config, -1)
		case m.LblXdotStringBufferXuSTRINGXhyphenBUFFERXhyphenHOOKED:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotStringBufferXuSTRINGXhyphenBUFFERXhyphenHOOKED", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotStringBufferXuSTRINGXhyphenBUFFERXhyphenHOOKED(config, -1)
		case m.LblXhashdasmInstruction:
			if len(kapp.List) != 5 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmInstruction", expectedArity: 5, actualArity: len(kapp.List)}
			}
			return evalXhashdasmInstruction(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], config, -1)
		case m.LblLookupRegisters:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalLookupRegisters", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalLookupRegisters(kapp.List[0], kapp.List[1], config, -1)
		case m.LblAccountEmpty:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalAccountEmpty", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalAccountEmpty(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashtoList:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashtoList", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashtoList(kapp.List[0], config, -1)
		case m.LblIsSubstateCellFragment:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSubstateCellFragment", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSubstateCellFragment(kapp.List[0], config, -1)
		case m.LblXhashopenXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashopenXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashopenXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblIsCodeCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCodeCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCodeCellOpt(kapp.List[0], config, -1)
		case m.LblIsTypeCheckingCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTypeCheckingCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTypeCheckingCellOpt(kapp.List[0], config, -1)
		case m.LblCsstore:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCsstore", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalCsstore(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblContractAppend:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalContractAppend", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalContractAppend(kapp.List[0], kapp.List[1], config, -1)
		case m.LblG0aux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalG0aux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalG0aux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsInts:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInts", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInts(kapp.List[0], config, -1)
		case m.LblXuAccountCellMapXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuAccountCellMapXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuAccountCellMapXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTxPendingCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxPendingCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxPendingCell(kapp.List[0], config, -1)
		case m.LblXuSetXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuSetXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuSetXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashgetcXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetcXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetcXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblInt2BytesNoLen:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInt2BytesNoLen", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalInt2BytesNoLen(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblInitBeneficiaryCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitBeneficiaryCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitBeneficiaryCell(config, -1)
		case m.LblIsScheduleCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsScheduleCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsScheduleCell(kapp.List[0], config, -1)
		case m.LblXhashdasmContractAux2:
			if len(kapp.List) != 8 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmContractAux2", expectedArity: 8, actualArity: len(kapp.List)}
			}
			return evalXhashdasmContractAux2(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], kapp.List[4], kapp.List[5], kapp.List[6], kapp.List[7], config, -1)
		case m.LblXhashcontractSize:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashcontractSize", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashcontractSize(kapp.List[0], kapp.List[1], config, -1)
		case m.LblRfindChar:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalRfindChar", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalRfindChar(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXumodIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXumodIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXumodIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTypes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTypes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTypes(kapp.List[0], config, -1)
		case m.LblIsPeakMemoryCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsPeakMemoryCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsPeakMemoryCell(kapp.List[0], config, -1)
		case m.LblCdiv:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalCdiv", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalCdiv(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblDirectionalityChar:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalDirectionalityChar", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalDirectionalityChar(kapp.List[0], config, -1)
		case m.LblIsIDCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIDCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIDCell(kapp.List[0], config, -1)
		case m.LblXhashopendirXlparenXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashopendirXlparenXuXrparenXuKXhyphenIO", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashopendirXlparenXuXrparenXuKXhyphenIO(kapp.List[0], config, -1)
		case m.LblTwos:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalTwos", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalTwos(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsGlobalDefinition:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsGlobalDefinition", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsGlobalDefinition(kapp.List[0], config, -1)
		case m.LblIsContractNameCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractNameCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractNameCell(kapp.List[0], config, -1)
		case m.LblIsBExp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBExp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBExp(kapp.List[0], config, -1)
		case m.LblIsTxGasLimitCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxGasLimitCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxGasLimitCell(kapp.List[0], config, -1)
		case m.LblIsFunctionSignature:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionSignature", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionSignature(kapp.List[0], config, -1)
		case m.LblIsContractDefinition:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractDefinition", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractDefinition(kapp.List[0], config, -1)
		case m.LblIsLocalCallsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLocalCallsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLocalCallsCell(kapp.List[0], config, -1)
		case m.LblIsMode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMode(kapp.List[0], config, -1)
		case m.LblXdotSet:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotSet", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotSet(config, -1)
		case m.LblInitTxGasPriceCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitTxGasPriceCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitTxGasPriceCell(config, -1)
		case m.LblIsFunctionCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFunctionCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFunctionCell(kapp.List[0], config, -1)
		case m.LblInitFuncIDCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitFuncIDCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitFuncIDCell(config, -1)
		case m.LblInitArgsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitArgsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitArgsCell(config, -1)
		case m.LblIsCallValueCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCallValueCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCallValueCellOpt(kapp.List[0], config, -1)
		case m.LblIsSchedule:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSchedule", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSchedule(kapp.List[0], config, -1)
		case m.LblIsNetworkCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsNetworkCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsNetworkCell(kapp.List[0], config, -1)
		case m.LblIsFuncCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFuncCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFuncCellOpt(kapp.List[0], config, -1)
		case m.LblLengthString:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalLengthString", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalLengthString(kapp.List[0], config, -1)
		case m.LblXuMessageCellMapXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuMessageCellMapXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuMessageCellMapXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashremoveZeros:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashremoveZeros", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashremoveZeros(kapp.List[0], config, -1)
		case m.LblFloatFormat:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalFloatFormat", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalFloatFormat(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsInternalOp:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsInternalOp", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsInternalOp(kapp.List[0], config, -1)
		case m.LblInitContractCodeCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitContractCodeCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitContractCodeCell(config, -1)
		case m.LblXuXplusStringXuXuSTRINGXhyphenBUFFERXhyphenHOOKED:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusStringXuXuSTRINGXhyphenBUFFERXhyphenHOOKED", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusStringXuXuSTRINGXhyphenBUFFERXhyphenHOOKED(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsFromCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsFromCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsFromCell(kapp.List[0], config, -1)
		case m.LblXuXplusStringXuXuSTRING:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXplusStringXuXuSTRING", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXplusStringXuXuSTRING(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXpipeIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXpipeIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXpipeIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitSubstateStackCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitSubstateStackCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitSubstateStackCell(config, -1)
		case m.LblIsStorageCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStorageCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStorageCellOpt(kapp.List[0], config, -1)
		case m.LblIsExpInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsExpInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsExpInst(kapp.List[0], config, -1)
		case m.LblXhashdasmContractAux1:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdasmContractAux1", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashdasmContractAux1(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashsizeWordStack:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeWordStack", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashsizeWordStack(kapp.List[0], config, -1)
		case m.LblIsDeclaredContractsCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsDeclaredContractsCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsDeclaredContractsCellOpt(kapp.List[0], config, -1)
		case m.LblProjectXcolonSchedule:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalProjectXcolonSchedule", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalProjectXcolonSchedule(kapp.List[0], config, -1)
		case m.LblXhashsizeRegsAux:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashsizeRegsAux", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashsizeRegsAux(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitIeleCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitIeleCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitIeleCell(config, -1)
		case m.LblInitLogDataCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitLogDataCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitLogDataCell(config, -1)
		case m.LblInitAccountCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitAccountCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitAccountCell(config, -1)
		case m.LblIsMsgIDCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMsgIDCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMsgIDCellOpt(kapp.List[0], config, -1)
		case m.LblIsTxNonceCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxNonceCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxNonceCell(kapp.List[0], config, -1)
		case m.LblIsIDCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsIDCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsIDCellOpt(kapp.List[0], config, -1)
		case m.LblXpercentXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXpercentXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalXpercentXlparenXuXcommaXuXcommaXuXcommaXuXrparenXuIELEXhyphenBINARY(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXuxorIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuxorIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuxorIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblInitInstructionsCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitInstructionsCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitInstructionsCell(config, -1)
		case m.LblXdotArrayXuIELEXhyphenDATA:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXdotArrayXuIELEXhyphenDATA", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalXdotArrayXuIELEXhyphenDATA(config, -1)
		case m.LblXhashgetBalance:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetBalance", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetBalance(kapp.List[0], config, -1)
		case m.LblIsCmpInst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsCmpInst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsCmpInst(kapp.List[0], config, -1)
		case m.LblBytes2String:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBytes2String", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalBytes2String(kapp.List[0], config, -1)
		case m.LblMessageCellMapItem:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMessageCellMapItem", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMessageCellMapItem(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashgetCode:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashgetCode", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashgetCode(kapp.List[0], config, -1)
		case m.LblContractBytes:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalContractBytes", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalContractBytes(kapp.List[0], config, -1)
		case m.LblBase2String:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalBase2String", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalBase2String(kapp.List[0], kapp.List[1], config, -1)
		case m.LblListItem:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalListItem", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalListItem(kapp.List[0], config, -1)
		case m.LblIsStream:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsStream", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsStream(kapp.List[0], config, -1)
		case m.LblInitCurrentFunctionCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitCurrentFunctionCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitCurrentFunctionCell(config, -1)
		case m.LblIsWordStack:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWordStack", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWordStack(kapp.List[0], config, -1)
		case m.LblIsAccountsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsAccountsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsAccountsCell(kapp.List[0], config, -1)
		case m.LblXuXltXeqMapXuXuMAP:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXltXeqMapXuXuMAP", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXltXeqMapXuXuMAP(kapp.List[0], kapp.List[1], config, -1)
		case m.LblNewUUIDXuSTRING:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalNewUUIDXuSTRING", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalNewUUIDXuSTRING(config, -1)
		case m.LblInitSelfDestructCell:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitSelfDestructCell", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalInitSelfDestructCell(config, -1)
		case m.LblIsMessagesCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMessagesCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMessagesCell(kapp.List[0], config, -1)
		case m.LblECDSARecover:
			if len(kapp.List) != 4 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalECDSARecover", expectedArity: 4, actualArity: len(kapp.List)}
			}
			return evalECDSARecover(kapp.List[0], kapp.List[1], kapp.List[2], kapp.List[3], config, -1)
		case m.LblXhashpoint:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashpoint", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashpoint(kapp.List[0], config, -1)
		case m.LblXhashasAccount:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashasAccount", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashasAccount(kapp.List[0], config, -1)
		case m.LblMakeArrayOcaml:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalMakeArrayOcaml", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalMakeArrayOcaml(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXhashparseAddr:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashparseAddr", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashparseAddr(kapp.List[0], config, -1)
		case m.LblIsWellFormednessScheduleCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsWellFormednessScheduleCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsWellFormednessScheduleCellOpt(kapp.List[0], config, -1)
		case m.LblXuinListXu:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuinListXu", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuinListXu(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsContractNameCellOpt:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsContractNameCellOpt", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsContractNameCellOpt(kapp.List[0], config, -1)
		case m.LblInitModeCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalInitModeCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalInitModeCell(kapp.List[0], config, -1)
		case m.LblIsValueCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsValueCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsValueCell(kapp.List[0], config, -1)
		case m.LblPow160XuIELEXhyphenDATA:
			if len(kapp.List) != 0 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalPow160XuIELEXhyphenDATA", expectedArity: 0, actualArity: len(kapp.List)}
			}
			return evalPow160XuIELEXhyphenDATA(config, -1)
		case m.LblIsMsgIDCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsMsgIDCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsMsgIDCell(kapp.List[0], config, -1)
		case m.LblIsConstant:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsConstant", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsConstant(kapp.List[0], config, -1)
		case m.LblIsBlockhashCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsBlockhashCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsBlockhashCell(kapp.List[0], config, -1)
		case m.LblIsSHA3Inst:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsSHA3Inst", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsSHA3Inst(kapp.List[0], config, -1)
		case m.LblXhashdecodeLengthPrefixAux:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashdecodeLengthPrefixAux", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXhashdecodeLengthPrefixAux(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXuXslashIntXuXuINT:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXslashIntXuXuINT", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXuXslashIntXuXuINT(kapp.List[0], kapp.List[1], config, -1)
		case m.LblXuXlsqbXuXltXhyphenXuXrsqbXuMAP:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXuXlsqbXuXltXhyphenXuXrsqbXuMAP", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalXuXlsqbXuXltXhyphenXuXrsqbXuMAP(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblXhashadjustedBitLengthAux:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashadjustedBitLengthAux", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashadjustedBitLengthAux(kapp.List[0], config, -1)
		case m.LblGetKLabel:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalGetKLabel", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalGetKLabel(kapp.List[0], config, -1)
		case m.LblG0call:
			if len(kapp.List) != 3 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalG0call", expectedArity: 3, actualArity: len(kapp.List)}
			}
			return evalG0call(kapp.List[0], kapp.List[1], kapp.List[2], config, -1)
		case m.LblIsLabelsCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsLabelsCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsLabelsCell(kapp.List[0], config, -1)
		case m.LblXhashseekEndXlparenXuXcommaXuXrparenXuKXhyphenIO:
			if len(kapp.List) != 2 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashseekEndXlparenXuXcommaXuXrparenXuKXhyphenIO", expectedArity: 2, actualArity: len(kapp.List)}
			}
			return evalXhashseekEndXlparenXuXcommaXuXrparenXuKXhyphenIO(kapp.List[0], kapp.List[1], config, -1)
		case m.LblIsTxOrderCell:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalIsTxOrderCell", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalIsTxOrderCell(kapp.List[0], config, -1)
		case m.LblGetIeleName:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalGetIeleName", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalGetIeleName(kapp.List[0], config, -1)
		case m.LblXhashappliedRule:
			if len(kapp.List) != 1 {
				return m.NoResult, &evalArityViolatedError{funcName:"evalXhashappliedRule", expectedArity: 1, actualArity: len(kapp.List)}
			}
			return evalXhashappliedRule(kapp.List[0], config, -1)
		default:
			return c, nil
	}
}

