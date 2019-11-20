package endpoint

import (
	"errors"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
)

var emptyWordStack = m.NewKApplyConstant(m.ParseKLabel(".WordStack_IELE-DATA"))
var wordStackElemLabel = m.ParseKLabel("_:__IELE-DATA")

func (vm *OriginalIeleVM) convertArguments(args [][]byte) m.KReference {
	kargList := vm.kinterpreter.Model.NewKApply(m.LblXdotListXlbracketXquoteoperandListXquoteXrbracket)
	for i := len(args) - 1; i >= 0; i-- {
		newKargList := vm.kinterpreter.Model.NewKApply(m.LblOperandList,
			vm.kinterpreter.Model.FromBigInt(big.NewInt(0).SetBytes(args[i])),
			kargList,
		)
		kargList = newKargList
	}
	return kargList
}

// G0Create yields the initial gas cost of creating a new smart contract
func (vm *OriginalIeleVM) G0Create(input *vmi.ContractCreateInput) (*big.Int, error) {
	kschedule := vm.scheduleToK(vm.schedule)
	kargList := vm.convertArguments(input.Arguments)

	// convert tx data
	txData := []byte(input.ContractCode)
	ws := emptyWordStack
	for i := len(txData) - 1; i >= 0; i-- {
		newWs := vm.kinterpreter.Model.NewKApply(wordStackElemLabel,
			vm.kinterpreter.Model.IntFromByte(txData[i]),
			ws,
		)
		ws = newWs
	}
	kappG0 := vm.kinterpreter.Model.NewKApply(m.LblG0create,
		kschedule,
		ws,
		kargList,
	)

	return vm.evalG0(kappG0)
}

// G0Call yields the initial gas cost of calling an existing smart contract
func (vm *OriginalIeleVM) G0Call(input *vmi.ContractCallInput) (*big.Int, error) {
	kschedule := vm.scheduleToK(vm.schedule)
	kargList := vm.convertArguments(input.Arguments)

	kappG0 := vm.kinterpreter.Model.NewKApply(m.LblG0call,
		kschedule,
		vm.kinterpreter.Model.NewString(input.Function),
		kargList,
	)

	return vm.evalG0(kappG0)
}

func (vm *OriginalIeleVM) evalG0(kappG0 m.KReference) (*big.Int, error) {

	evalRes, evalErr := vm.kinterpreter.Eval(kappG0, m.InternedBottom)
	if evalErr != nil {
		return nil, evalErr
	}
	intRes, isInt := vm.kinterpreter.Model.GetBigInt(evalRes)
	if !isInt {
		return nil, errors.New("g0 function evaluation did not return Int")
	}
	return intRes, nil
}
