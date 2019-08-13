// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:25:08.138

package ieletestinginterpreter

import (
	"fmt"
	koreparser "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/koreparser"
	"log"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
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
	kConfigMap := make(map[m.KMapKey]m.KReference)
	for key, kastValue := range kastMap {
		ktokenRef := i.Model.NewKToken(m.SortKConfigVar, "$"+key)
		ktokenKey, _ := i.Model.MapKey(ktokenRef)
		parsedValue := koreparser.Parse(kastValue)
		kValue := i.convertParserModelToKModel(parsedValue)
		kConfigMap[ktokenKey] = kValue
	}

	// top cell initialization
	kmap := i.Model.NewMap(m.SortMap, m.KLabelForMap, kConfigMap)
	evalK := i.Model.NewKApply(TopCellInitializer, kmap)
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
func (i *Interpreter) TakeStepsNoThread(k m.KReference) error {
	i.initializeTrace()
	defer i.closeTrace()

	// start
	i.currentStep = 0
	i.checksSinceLastGc = 0
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

	completelyStuck := false
	for !completelyStuck {
		// try to make stuck, to allow execution of steps that depend on stuck state
		// it is possible to set stuck multiple times
		i.currentStep++
		i.traceStepStart()
		i.state, err = i.makeStuck(i.state, i.state)
		if err != nil {
			return err
		}
		i.traceStepEnd()
		i.currentStep++

		// run - stuck steps
		stepBeforeStuck := i.currentStep
		err = i.runSteps(maxSteps)
		if err != nil {
			return err
		}
		if stepBeforeStuck == i.currentStep {
			// will only stop when no other step can be performed after stuck
			completelyStuck = true
		}
	}

	return nil
}

// gcFrequencyMask indicates how often we check if garbage collection is needed.
// every 1024 steps in this case.
const gcFrequencyMask = (1 << 11) - 1

// minChecksBetweenGc makes sure we don't run the Gc too often
const minChecksBetweenGc = 14

func (i *Interpreter) runSteps(maxSteps int) error {
	running := true
	for running {
		if i.currentStep >= maxSteps {
			return errMaxStepsReached
		}
		i.traceStepStart()

		var err error
		i.state, err = i.step(i.state)

		// periodically clean up model of old data
		if i.currentStep&gcFrequencyMask == 0 {
			if i.checksSinceLastGc > minChecksBetweenGc && i.Model.ShouldRunGc() {
				//fmt.Printf("clean: %d | %d | [%d of %d] \n", i.checksSinceLastGc, i.currentStep, i.Model.SizeUsed(), i.Model.SizeAllocated())
				i.state = i.Model.Gc(i.state)
				i.checksSinceLastGc = 0
			} else {
				i.checksSinceLastGc++
			}
		}

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
