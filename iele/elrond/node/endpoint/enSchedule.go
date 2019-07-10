package endpoint

import (
	"errors"

	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// Schedule defines a IELE gas model type
type Schedule int

const (
	// VMTests is the IELE default gas model, currently only used in certain VM tests.
	VMTests Schedule = iota

	// Albe is the IELE "ALBE" gas model, this was the first version.
	Albe

	// Danse is the IELE "DANSE" gas model, this is the latest version.
	Danse

	// ElrondDefault is the current gas model for Elrond, it is currently identical to Danse.
	ElrondDefault
)

// ParseSchedule yields the schedule with the given name. It is used in tests.
func ParseSchedule(scheduleName string) (Schedule, error) {
	switch scheduleName {
	case "Default":
		return VMTests, nil
	case "Albe":
		return Albe, nil
	case "Danse":
		return Danse, nil
	case "Elrond":
		return ElrondDefault, nil
	case "ElrondDefault":
		return ElrondDefault, nil
	default:
		return VMTests, errors.New("unknown IELE schedule name")
	}
}

func scheduleToK(schedule Schedule) m.K {
	switch schedule {
	case VMTests:
		return &m.KApply{Label: m.ParseKLabel("DEFAULT_IELE-GAS")}
	case Albe:
		return &m.KApply{Label: m.ParseKLabel("ALBE_IELE-CONSTANTS")}
	case Danse:
		return &m.KApply{Label: m.ParseKLabel("DANSE_IELE-CONSTANTS")}
	case ElrondDefault:
		return &m.KApply{Label: m.ParseKLabel("DANSE_IELE-CONSTANTS")}
	default:
		panic("unknown IELE schedule")
	}
}
