// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:46:43.258

package ieletestinginterpreter

import (
	"fmt"
	koreparser "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/koreparser"
	"log"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
	"math"
	"os/exec"
)

func callKast(kdir string, programPath string) []byte {
	cmd := exec.Command("kast", programPath)
	cmd.Dir = kdir
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Kast error: " + err.Error())
	}
	return out
}

// ExecuteSimple interprets the program in the file given at input
func (i *Interpreter) ExecuteSimple(kdir string, execFile string) {
	kast := callKast(kdir, execFile)
	if i.Verbose {
		fmt.Printf("Kast: %s\n\n", kast)
	}

	data := make(map[string][]byte)
	data["PGM"] = kast

	err := i.Execute(data)
	if err != nil {
		panic(err)
	}
	final := i.GetState()
	stepsMade := i.GetNrSteps()

	if i.Verbose {
		fmt.Println("\n\npretty print:")
		fmt.Println(i.Model.PrettyPrint(final))
		fmt.Println("\n\nk print:")
		fmt.Println(i.Model.KPrint(final))
		fmt.Printf("\n\nsteps made: %d\n", stepsMade)
	}
}

// Execute interprets the program with the structure
func (i *Interpreter) Execute(kastMap map[string][]byte) error {
	kConfigMap := make(map[m.KMapKey]m.K)
	for key, kastValue := range kastMap {
		ktoken := m.KToken{Sort: m.SortKConfigVar, Value: "$" + key}
		parsedValue := koreparser.Parse(kastValue)
		kValue := i.convertParserModelToKModel(parsedValue)
		kConfigMap[ktoken] = kValue
	}

	// top cell initialization
	kmap := &m.Map{Sort: m.SortMap, Label: m.KLabelForMap, Data: kConfigMap}
	evalK := &m.KApply{Label: TopCellInitializer, List: []m.K{kmap}}
	kinit, err := i.Eval(evalK, m.InternedBottom)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if i.Verbose {
		fmt.Println("\n\ntop level init:")
		fmt.Println(i.Model.PrettyPrint(kinit))
	}

	// execute
	return i.TakeStepsNoThread(kinit)
}

// TakeStepsNoThread executes as many steps as possible given the starting configuration
func (i *Interpreter) TakeStepsNoThread(k m.K) error {
	i.initializeTrace()
	defer i.closeTrace()

	// start
	i.currentStep = 0
	i.state = k
	i.traceInitialState(k)

	maxSteps := i.MaxSteps
	if maxSteps == 0 {
		// not set, it means we don't limit the number of steps
		maxSteps = math.MaxInt32
	}

	// run - main
	var err error
	err = i.runSteps(maxSteps)
	if err != nil {
		return err
	}

	// try to make stuck, to enable steps dependent on stuck state
	i.currentStep++
	i.traceStepStart()
	i.state, err = i.makeStuck(i.state, i.state)
	if err != nil {
		return err
	}
	i.traceStepEnd()
	i.currentStep++

	// run - stuck steps
	err = i.runSteps(maxSteps)
	if err != nil {
		return err
	}

	return nil
}

func (i *Interpreter) runSteps(maxSteps int) error {
	running := true
	for running {
		if i.currentStep >= maxSteps {
			return errMaxStepsReached
		}
		i.traceStepStart()
		var err error
		i.state, err = i.step(i.state)
		if err == nil {
			i.traceStepEnd()
			i.currentStep++
		} else {
			if _, t := err.(*noStepError); t {
				// no step error, the stop sign
				i.traceNoStep()
				running = false
			} else {
				// unexpected error
				return err
			}
		}
	}

	return nil
}
