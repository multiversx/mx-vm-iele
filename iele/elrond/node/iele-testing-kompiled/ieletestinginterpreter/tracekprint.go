// File provided by the K Framework Go backend. Timestamp: 2019-07-04 01:26:11.488

package ieletestinginterpreter

import (
	"bufio"
	"fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"os"
	"time"
)

// this one writes the state at each step, all in one file
type traceKPrint struct {
	file       *os.File
	fileWriter *bufio.Writer
	interpreter *Interpreter
}

func (t *traceKPrint) initialize() {
	formattedNow := time.Now().Format("20060102150405")
	fileName := "traceKPrint_" + formattedNow + ".log"
	var err error
	t.file, err = os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while creating trace file:" + err.Error())
	}

	t.fileWriter = bufio.NewWriter(t.file)
}

func (t *traceKPrint) closeTrace() {
	t.fileWriter.Flush()
	t.file.Close()
}

func (t *traceKPrint) traceInitialState(state m.KReference) {
}

func (t *traceKPrint) traceStepStart(stepNr int, currentState m.KReference) {
	kast := t.interpreter.Model.KPrint(currentState)
	t.fileWriter.WriteString(fmt.Sprintf("\nstep %d %s", stepNr, kast))
}

func (t *traceKPrint) traceStepEnd(stepNr int, currentState m.KReference) {
}

func (t *traceKPrint) traceNoStep(stepNr int, currentState m.KReference) {
}

func (t *traceKPrint) traceRuleApply(ruleType string, stepNr int, ruleInfo string) {
}
