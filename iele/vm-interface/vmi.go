package vminterface

import "math/big"

// IeleVM ... interface for the IELE VM endpoints
type IeleVM interface {
	RunTransaction(input *VMInput) (*VMOutput, error)
	ParseSchedule(scheduleName string) (Schedule, error)
	G0(schedule Schedule, txCreate bool, dataForGas string, args []*big.Int) (*big.Int, error)
}
