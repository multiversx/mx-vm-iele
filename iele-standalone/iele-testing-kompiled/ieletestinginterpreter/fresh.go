package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestingmodel"
)

func freshFunction (s m.Sort, config m.K, counter int) (m.K, error) {
	switch s {
		case m.SortID:
			return evalFreshID(m.NewIntFromInt(counter), config, -1)
		case m.SortInt:
			return evalFreshInt(m.NewIntFromInt(counter), config, -1)
		default:
			panic("Cannot find fresh function for sort " + s.Name())
	}
}

