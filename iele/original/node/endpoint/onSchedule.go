package endpoint

import (
	ielecommon "github.com/ElrondNetwork/elrond-vm/iele/common"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
)

func scheduleToK(schedule ielecommon.Schedule) m.K {
	switch schedule {
	case ielecommon.Default:
		return &m.KApply{Label: m.ParseKLabel("DEFAULT_IELE-GAS")}
	case ielecommon.Albe:
		return &m.KApply{Label: m.ParseKLabel("ALBE_IELE-CONSTANTS")}
	case ielecommon.Danse:
		return &m.KApply{Label: m.ParseKLabel("DANSE_IELE-CONSTANTS")}
	default:
		panic("unknown IELE schedule")
	}
}
