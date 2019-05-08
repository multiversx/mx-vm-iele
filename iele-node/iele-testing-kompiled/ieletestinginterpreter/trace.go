// Handles generation of traces
// (what rules were applied, in what order, what were the intermediate states).
// Multiple trace handlers supported.

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
)

var traceHandlers []traceHandler

// we can have multiple writers to write program execution traces in various formats
// they are all intended for easier debugging
type traceHandler interface {
	initialize()
	closeTrace()
	traceInitialState(state m.K)
	traceStepStart(stepNr int, currentState m.K)
	traceStepEnd(stepNr int, currentState m.K)
	traceNoStep(stepNr int, currentState m.K)
	traceRuleApply(ruleType string, stepNr int, ruleInfo string)
}

func initializeTrace() {
	for _, t := range traceHandlers {
		t.initialize()
	}
}

func closeTrace() {
	for _, t := range traceHandlers {
		t.closeTrace()
	}
}

func traceInitialState(state m.K) {
	for _, t := range traceHandlers {
		t.traceInitialState(state)
	}
}

func traceStepStart(stepNr int, currentState m.K) {
	for _, t := range traceHandlers {
		t.traceStepStart(stepNr, currentState)
	}
}

func traceStepEnd(stepNr int, currentState m.K) {
	for _, t := range traceHandlers {
		t.traceStepEnd(stepNr, currentState)
	}
}

func traceNoStep(stepNr int, currentState m.K) {
	for _, t := range traceHandlers {
		t.traceNoStep(stepNr, currentState)
	}
}

func traceRuleApply(ruleType string, stepNr int, ruleInfo string) {
	for _, t := range traceHandlers {
		t.traceRuleApply(ruleType, stepNr, ruleInfo)
	}
}
