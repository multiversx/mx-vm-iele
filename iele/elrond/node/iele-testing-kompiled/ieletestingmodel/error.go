// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestingmodel

import "fmt"

type parseIntError struct {
	parseVal string
}

func (e *parseIntError) Error() string {
	return fmt.Sprintf("Could not parse int from value: %s", e.parseVal)
}

// HookNotImplementedError signals the interpreter that a hook is not implemented.
// Some functions with hooks also provide alternate implementations in K.
// This error signals that the K implementation should be used instead of the hook.a
// The error is declared in the model package so external hooks can also have access to it.
type HookNotImplementedError struct {
}

func (e *HookNotImplementedError) Error() string {
	return "Hook not implemented."
}

var hookNotImplementedErrorInstance = &HookNotImplementedError{}

func GetHookNotImplementedError() *HookNotImplementedError {
	return hookNotImplementedErrorInstance
}
