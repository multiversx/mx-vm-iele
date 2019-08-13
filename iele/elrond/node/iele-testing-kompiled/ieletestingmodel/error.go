// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestingmodel

import "fmt"

type parseIntError struct {
	parseVal string
}

func (e *parseIntError) Error() string {
	return fmt.Sprintf("Could not parse int from value: %s", e.parseVal)
}
