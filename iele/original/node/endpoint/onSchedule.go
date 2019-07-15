package endpoint

import (
	"errors"

	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
)

// Schedule defines a IELE gas model type
type Schedule int

const (
	// Default is the IELE default gas model
	Default Schedule = iota

	// Albe is the IELE "ALBE" gas model, this was the first version
	Albe

	// Danse is the IELE "DANSE" gas model, this is the latest version
	Danse
)

// ParseSchedule yields the schedule with the given name. It is used in tests.
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

func (vm *OriginalIeleVM) scheduleToK(schedule ielecommon.Schedule) m.KReference {
	switch schedule {
	case Default:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("DEFAULT_IELE-GAS"))
	case Albe:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("ALBE_IELE-CONSTANTS"))
	case Danse:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("DANSE_IELE-CONSTANTS"))
	default:
		panic("unknown IELE schedule")
	}
}
