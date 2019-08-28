// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestSetEmpty(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := setHooks.unit(m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertSetOk(t, map[int]bool{}, result, err, interpreter)
}

func TestSetSingleton(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := setHooks.element(
		interpreter.Model.FromInt(1),
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertSetOk(t, map[int]bool{1: true}, result, err, interpreter)
}

func TestSetAddContains(t *testing.T) {
	interpreter := newTestInterpreter()
	var result, found m.KReference
	var err error

	result, err = setHooks.unit(m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)

	result, err = setAdd(result, 1, interpreter)
	assertSetOk(t, map[int]bool{1: true}, result, err, interpreter)

	found, err = setHooks.in(
		interpreter.Model.FromInt(1),
		result,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertBoolOk(t, true, found, err, interpreter)

	result, err = setAdd(result, 2, interpreter)
	assertSetOk(t, map[int]bool{1: true, 2: true}, result, err, interpreter)

	found, err = setHooks.in(
		interpreter.Model.FromInt(1),
		result,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertBoolOk(t, true, found, err, interpreter)
	found, err = setHooks.in(
		interpreter.Model.FromInt(2),
		result,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertBoolOk(t, true, found, err, interpreter)
}

type setOperation func(set1, set2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error)

func testSetOperation(t *testing.T, op setOperation, set1, set2, expectedResult []int) {
	interpreter := newTestInterpreter()
	var kset1, kset2, result m.KReference
	var err error
	var expectedMap map[int]bool

	kset1, err = setOf(interpreter, set1...)
	expectedMap = make(map[int]bool)
	for _, i := range set1 {
		expectedMap[i] = true
	}
	assertSetOk(t, expectedMap, kset1, err, interpreter)

	kset2, err = setOf(interpreter, set2...)
	expectedMap = make(map[int]bool)
	for _, i := range set2 {
		expectedMap[i] = true
	}
	assertSetOk(t, expectedMap, kset2, err, interpreter)

	result, err = op(
		kset1, kset2,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	expectedMap = make(map[int]bool)
	for _, i := range expectedResult {
		expectedMap[i] = true
	}
	assertSetOk(t, expectedMap, result, err, interpreter)
}

func TestSetConcatEmpty(t *testing.T) {
	testSetOperation(t, setHooks.concat,
		[]int{},
		[]int{},
		[]int{})

	testSetOperation(t, setHooks.concat,
		[]int{},
		[]int{1},
		[]int{1})

	testSetOperation(t, setHooks.concat,
		[]int{1},
		[]int{},
		[]int{1})

	testSetOperation(t, setHooks.concat,
		[]int{},
		[]int{1, 2, 3},
		[]int{1, 2, 3})

	testSetOperation(t, setHooks.concat,
		[]int{1, 2, 3},
		[]int{},
		[]int{1, 2, 3})
}

func TestSetConcat(t *testing.T) {
	testSetOperation(t, setHooks.concat,
		[]int{4},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4})

	testSetOperation(t, setHooks.concat,
		[]int{1, 2, 3},
		[]int{4},
		[]int{1, 2, 3, 4})
}

func TestSetIntersectionEmpty(t *testing.T) {
	testSetOperation(t, setHooks.intersection,
		[]int{},
		[]int{},
		[]int{})

	testSetOperation(t, setHooks.intersection,
		[]int{},
		[]int{1},
		[]int{})

	testSetOperation(t, setHooks.intersection,
		[]int{1},
		[]int{},
		[]int{})

	testSetOperation(t, setHooks.intersection,
		[]int{},
		[]int{1, 2, 3},
		[]int{})

	testSetOperation(t, setHooks.intersection,
		[]int{1, 2, 3},
		[]int{},
		[]int{})
}

func TestSetIntersection(t *testing.T) {
	testSetOperation(t, setHooks.intersection,
		[]int{1},
		[]int{1, 2, 3},
		[]int{1})

	testSetOperation(t, setHooks.intersection,
		[]int{1, 2, 3},
		[]int{1},
		[]int{1})

	testSetOperation(t, setHooks.intersection,
		[]int{1, 3},
		[]int{1, 2, 3},
		[]int{1, 3})

	testSetOperation(t, setHooks.intersection,
		[]int{1, 2, 3},
		[]int{1, 3},
		[]int{1, 3})
}

func TestSetDifferenceEmpty(t *testing.T) {
	testSetOperation(t, setHooks.difference,
		[]int{},
		[]int{},
		[]int{})

	testSetOperation(t, setHooks.difference,
		[]int{1},
		[]int{},
		[]int{1})

	testSetOperation(t, setHooks.difference,
		[]int{1, 2, 3},
		[]int{},
		[]int{1, 2, 3})

	testSetOperation(t, setHooks.difference,
		[]int{},
		[]int{1, 2, 3},
		[]int{})
}

func TestSetDifference(t *testing.T) {
	testSetOperation(t, setHooks.difference,
		[]int{1},
		[]int{1},
		[]int{})

	testSetOperation(t, setHooks.difference,
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{})

	testSetOperation(t, setHooks.difference,
		[]int{2},
		[]int{1, 2, 3},
		[]int{})

	testSetOperation(t, setHooks.difference,
		[]int{2, 4},
		[]int{1, 2, 3},
		[]int{4})

	testSetOperation(t, setHooks.difference,
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3},
		[]int{4})
}

func TestSetChoice(t *testing.T) {
	interpreter := newTestInterpreter()
	var kset, choiceElem m.KReference
	var err error

	kset, err = setOf(interpreter, 1)
	assertSetOk(t, map[int]bool{1: true}, kset, err, interpreter)

	choiceElem, err = setHooks.choice(
		kset,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	assertIntOk(t, "1", choiceElem, err, interpreter)
}

func setAdd(set m.KReference, i int, interpreter *Interpreter) (m.KReference, error) {
	elemSingleton, err := setHooks.element(
		interpreter.Model.FromInt(i),
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	if err != nil {
		return NullReference, err
	}

	set, err = setHooks.concat(
		set, elemSingleton,
		m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	return set, err
}

func setOf(interpreter *Interpreter, values ...int) (m.KReference, error) {
	result, err := setHooks.unit(m.KLabelForSet, m.SortSet, m.InternedBottom, interpreter)
	if err != nil {
		return m.NoResult, err
	}

	for _, i := range values {
		result, err = setAdd(result, i, interpreter)
		if err != nil {
			return m.NoResult, err
		}
	}

	return result, nil
}

func assertSetOk(t *testing.T, expectedValues map[int]bool, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	if !interpreter.Model.IsSet(actual) {
		t.Error("result is not a set")
	}
	if len(expectedValues) != interpreter.Model.SetSize(actual) {
		t.Errorf("set length doesn't match, have: %s",
			interpreter.Model.PrettyPrint(actual))
	}
	for eKey := range expectedValues {
		kKey := interpreter.Model.FromInt(eKey)
		if !interpreter.Model.SetContains(actual, kKey) {
			t.Errorf("key expected but not found: %s",
				interpreter.Model.PrettyPrint(kKey))
		}
	}

	realLength := 0
	interpreter.Model.SetForEach(actual, func(elem KReference) bool {
		ielem, _ := interpreter.Model.GetInt(elem)
		_, found := expectedValues[ielem]
		if !found {
			t.Errorf("key not found: %d", ielem)
		}
		realLength++
		return false
	})
	if len(expectedValues) != realLength {
		t.Errorf("real set length doesn't match, have: %d\n%s",
			realLength,
			interpreter.Model.PrettyPrint(actual))
	}
}
