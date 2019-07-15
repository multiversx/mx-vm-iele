package endpoint

import (
	ielecommon "github.com/ElrondNetwork/elrond-vm/iele/common"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
)

func (vm *OriginalIeleVM) scheduleToK(schedule ielecommon.Schedule) m.KReference {
	switch schedule {
	case ielecommon.Default:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("DEFAULT_IELE-GAS"))
	case ielecommon.Albe:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("ALBE_IELE-CONSTANTS"))
	case ielecommon.Danse:
		return vm.kinterpreter.Model.NewKApply(m.ParseKLabel("DANSE_IELE-CONSTANTS"))
	default:
		panic("unknown IELE schedule")
	}
}
