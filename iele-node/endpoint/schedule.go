package endpoint

import (
	"errors"
	"math/big"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
)

// Schedule ... IELE gas model type
type Schedule int

const (
	// Default ... IELE default gas model
	Default Schedule = iota

	// Albe ... IELE "ALBE" gas model, this was their first version
	Albe

	// Danse ... IELE "DANSE" gas model, this is the latest version
	Danse
)

// ParseSchedule ... yields a schedule type based on its string representation
func ParseSchedule(scheduleName string) (Schedule, error) {
	switch scheduleName {
	case "Default":
		return Default, nil
	case "Albe":
		return Albe, nil
	case "Danse":
		return Danse, nil
	default:
		return Default, errors.New("unknown IELE schedule name")
	}
}

func scheduleToK(schedule Schedule) m.K {
	switch schedule {
	case Default:
		return &m.KApply{Label: m.ParseKLabel("DEFAULT_IELE-GAS")}
	case Albe:
		return &m.KApply{Label: m.ParseKLabel("ALBE_IELE-CONSTANTS")}
	case Danse:
		return &m.KApply{Label: m.ParseKLabel("DANSE_IELE-CONSTANTS")}
	default:
		panic("unknown IELE schedule")
	}
}

var emptyWordStack = &m.KApply{Label: m.ParseKLabel(".WordStack_IELE-DATA"), List: nil}
var wordStackElemLabel = m.ParseKLabel("_:__IELE-DATA")

// G0 ... computes the gas cost for starting a smart contract
func G0(schedule Schedule, txCreate bool, dataForGas string, args []*big.Int) (*big.Int, error) {
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
