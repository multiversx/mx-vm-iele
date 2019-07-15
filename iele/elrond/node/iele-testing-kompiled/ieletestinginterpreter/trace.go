// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:08:58.251

// Handles generation of traces
// (what rules were applied, in what order, what were the intermediate states).
// Multiple trace handlers supported.

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

var traceHandlers []traceHandler

// we can have multiple writers to write program execution traces in various formats
// they are all intended for easier debugging
type traceHandler interface {
	initialize()
	closeTrace()
	traceInitialState(state m.KReference)
	traceStepStart(stepNr int, currentState m.KReference)
	traceStepEnd(stepNr int, currentState m.KReference)
	traceNoStep(stepNr int, currentState m.KReference)
	traceRuleApply(ruleType string, stepNr int, ruleInfo string)
}

func (i *Interpreter) initializeTrace() {
	for _, t := range i.traceHandlers {
		t.initialize()
	}
}

func (i *Interpreter) closeTrace() {
	for _, t := range i.traceHandlers {
		t.closeTrace()
	}
}

func (i *Interpreter) traceInitialState(state m.KReference) {
	for _, t := range i.traceHandlers {
		t.traceInitialState(state)
	}
}

func (i *Interpreter) traceStepStart() {
	for _, t := range i.traceHandlers {
		t.traceStepStart(i.currentStep, i.state)
	}
}

func (i *Interpreter) traceStepEnd() {
	for _, t := range i.traceHandlers {
		t.traceStepEnd(i.currentStep, i.state)
	}
}

func (i *Interpreter) traceNoStep() {
	for _, t := range i.traceHandlers {
		t.traceNoStep(i.currentStep, i.state)
	}
}

func (i *Interpreter) traceRuleApply(ruleType string, ruleNr int, ruleInfo string) {
	for _, t := range i.traceHandlers {
		t.traceRuleApply(ruleType, ruleNr, ruleInfo)
	}
}
