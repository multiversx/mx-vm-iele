package ielecommon

import (
	"errors"
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

// ParseSchedule ... get schedule with name
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
