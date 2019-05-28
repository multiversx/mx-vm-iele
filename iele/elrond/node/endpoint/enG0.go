package endpoint

import (
	"errors"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

var emptyWordStack = &m.KApply{Label: m.ParseKLabel(".WordStack_IELE-DATA"), List: nil}
var wordStackElemLabel = m.ParseKLabel("_:__IELE-DATA")

func convertArguments(args []*big.Int) m.K {
	kargList := &m.KApply{Label: m.LblXdotListXlbracketXquoteoperandListXquoteXrbracket, List: []m.K{}}
	for i := len(args) - 1; i >= 0; i-- {
		newKargList := &m.KApply{Label: m.LblOperandList, List: []m.K{
			m.NewInt(args[i]),
			kargList,
		}}
		kargList = newKargList
	}
	return kargList
}

// G0Create yields the initial gas cost of creating a new smart contract
func (vm *ElrondIeleVM) G0Create(input *vmi.ContractCreateInput) (*big.Int, error) {
	kschedule := scheduleToK(vm.schedule)
	kargList := convertArguments(input.Arguments)

	// convert tx data
	txData := []byte(input.ContractCode)
	ws := emptyWordStack
	for i := len(txData) - 1; i >= 0; i-- {
		newWs := &m.KApply{Label: wordStackElemLabel, List: []m.K{
			m.NewIntFromByte(txData[i]),
			ws,
		}}
		ws = newWs
	}
	kappG0 := &m.KApply{Label: m.LblG0create, List: []m.K{
		kschedule,
		ws,
		kargList,
	}}

	return vm.evalG0(kappG0)
}

// G0Call yields the initial gas cost of calling an existing smart contract
func (vm *ElrondIeleVM) G0Call(input *vmi.ContractCallInput) (*big.Int, error) {
	kschedule := scheduleToK(vm.schedule)
	kargList := convertArguments(input.Arguments)

	kappG0 := &m.KApply{Label: m.LblG0call, List: []m.K{
		kschedule,
		m.NewString(input.Function),
		kargList,
	}}

	return vm.evalG0(kappG0)
}

func (vm *ElrondIeleVM) evalG0(kappG0 m.K) (*big.Int, error) {

	evalRes, evalErr := vm.kinterpreter.Eval(kappG0, m.InternedBottom)
	if evalErr != nil {
		return nil, evalErr
	}
	intRes, isInt := evalRes.(*m.Int)
	if !isInt {
		return nil, errors.New("g0 function evaluation did not return Int")
	}
	return intRes.Value, nil
}
