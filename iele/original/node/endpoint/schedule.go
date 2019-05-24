package endpoint

import (
	"errors"
	"math/big"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
	vmi "github.com/ElrondNetwork/elrond-vm/iele/vm-interface"
)

// ParseSchedule ... yields a schedule type based on its string representation
func (OriginalIeleVMType) ParseSchedule(scheduleName string) (vmi.Schedule, error) {
	switch scheduleName {
	case "Default":
		return vmi.Default, nil
	case "Albe":
		return vmi.Albe, nil
	case "Danse":
		return vmi.Danse, nil
	default:
		return vmi.Default, errors.New("unknown IELE schedule name")
	}
}

func scheduleToK(schedule vmi.Schedule) m.K {
	switch schedule {
	case vmi.Default:
		return &m.KApply{Label: m.ParseKLabel("DEFAULT_IELE-GAS")}
	case vmi.Albe:
		return &m.KApply{Label: m.ParseKLabel("ALBE_IELE-CONSTANTS")}
	case vmi.Danse:
		return &m.KApply{Label: m.ParseKLabel("DANSE_IELE-CONSTANTS")}
	default:
		panic("unknown IELE schedule")
	}
}

var emptyWordStack = &m.KApply{Label: m.ParseKLabel(".WordStack_IELE-DATA"), List: nil}
var wordStackElemLabel = m.ParseKLabel("_:__IELE-DATA")

// G0 ... computes the gas cost for starting a smart contract
func (OriginalIeleVMType) G0(schedule vmi.Schedule, txCreate bool, dataForGas string, args []*big.Int) (*big.Int, error) {
	kschedule := scheduleToK(schedule)
	kargList := &m.KApply{Label: m.LblXdotListXlbracketXquoteoperandListXquoteXrbracket, List: []m.K{}}
	for i := len(args) - 1; i >= 0; i-- {
		newKargList := &m.KApply{Label: m.LblOperandList, List: []m.K{
			m.NewInt(args[i]),
			kargList,
		}}
		kargList = newKargList
	}
	var kappG0 m.K

	if txCreate {
		txData := []byte(dataForGas)
		ws := emptyWordStack
		for i := len(txData) - 1; i >= 0; i-- {
			newWs := &m.KApply{Label: wordStackElemLabel, List: []m.K{
				m.NewIntFromByte(txData[i]),
				ws,
			}}
			ws = newWs
		}
		kappG0 = &m.KApply{Label: m.LblG0create, List: []m.K{
			kschedule,
			ws,
			kargList,
		}}
	} else {
		kappG0 = &m.KApply{Label: m.LblG0call, List: []m.K{
			kschedule,
			m.NewString(dataForGas),
			kargList,
		}}
	}

	evalRes, evalErr := interpreter.Eval(kappG0, m.InternedBottom)
	if evalErr != nil {
		return nil, evalErr
	}
	intRes, isInt := evalRes.(*m.Int)
	if !isInt {
		return nil, errors.New("g0 function evaluation did not return Int")
	}
	return intRes.Value, nil
}
