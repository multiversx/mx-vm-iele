// File provided by the K Framework Go backend. Timestamp: 2019-06-13 13:40:04.133

package ieletestinginterpreter

import (
    "errors"
	"fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type noStepError struct {
}

func (e *noStepError) Error() string {
	return "No step could be performed."
}

var noStep = &noStepError{}

type stuckError struct {
	ms       *m.ModelState
	funcName string
	args     []m.K
}

func (e *stuckError) Error() string {
	if len(e.args) == 0 {
		return "Stuck! Function name: " + e.funcName + ". No args."
	}
	s := "Stuck! Function name: " + e.funcName + ". Args:"
	for i, arg := range e.args {
		s += fmt.Sprintf("\n%d: %s", i, e.ms.PrettyPrint(arg))
	}
	return s
}

type evalArityViolatedError struct {
	funcName      string
	expectedArity int
	actualArity   int
}

func (e *evalArityViolatedError) Error() string {
	return fmt.Sprintf(
		"Eval function arity violated. Function name: %s. Expected arity: %d. Actual arity: %d.",
		e.funcName, e.expectedArity, e.actualArity)
}

type hookNotImplementedError struct {
}

func (e *hookNotImplementedError) Error() string {
	return "Hook not implemented."
}

type hookInvalidArgsError struct {
}

func (e *hookInvalidArgsError) Error() string {
	return "Invalid argument(s) provided to hook."
}

func invalidArgsResult() (m.K, error) {
    return m.NoResult, &hookInvalidArgsError{}
}

type hookDivisionByZeroError struct {
}

func (e *hookDivisionByZeroError) Error() string {
	return "Division by zero."
}

var errInvalidMapKey = errors.New("invalid map key")

var errBadSetElement = errors.New("type cannot be used as a set element")

var errMaxStepsReached = errors.New("Maximum number of steps reached")
