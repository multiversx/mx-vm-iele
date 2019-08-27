// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

type concatCommonKeyTestKVP struct {
	key   int
	value int
}

func TestMapEmpty(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{}, result, err, interpreter)
}

func TestMapSingleton(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := mapHooks.element(
		interpreter.Model.FromInt(1),
		interpreter.Model.FromInt(3),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 3}, result, err, interpreter)
}

func TestMapUpdateLookupRemove(t *testing.T) {
	interpreter := newTestInterpreter()
	var result, lookup, keyFound m.KReference
	var err error

	result, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)

	result, err = mapHooks.update(result,
		interpreter.Model.FromInt(1),
		interpreter.Model.FromInt(3),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 3}, result, err, interpreter)

	lookup, err = mapHooks.lookup(result,
		interpreter.Model.FromInt(1),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertIntOk(t, "3", lookup, err, interpreter)
	keyFound, err = mapHooks.inKeys(
		interpreter.Model.FromInt(1),
		result,
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertBoolOk(t, true, keyFound, err, interpreter)

	lookup, err = mapHooks.lookup(result,
		interpreter.Model.FromInt(1234),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertBottomOk(t, lookup, err, interpreter)
	lookup, err = mapHooks.lookupOrDefault(result,
		interpreter.Model.FromInt(1234),
		interpreter.Model.FromInt(555),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertIntOk(t, "555", lookup, err, interpreter)
	keyFound, err = mapHooks.inKeys(
		interpreter.Model.FromInt(1234),
		result,
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertBoolOk(t, false, keyFound, err, interpreter)

	result, err = mapHooks.remove(result,
		interpreter.Model.FromInt(12345),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 3}, result, err, interpreter)

	result, err = mapHooks.remove(result,
		interpreter.Model.FromInt(1),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{}, result, err, interpreter)

	lookup, err = mapHooks.lookup(result,
		interpreter.Model.FromInt(1),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertBottomOk(t, lookup, err, interpreter)
	keyFound, err = mapHooks.inKeys(
		interpreter.Model.FromInt(1),
		result,
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertBoolOk(t, false, keyFound, err, interpreter)
}

func testRemove(t *testing.T, contents []concatCommonKeyTestKVP, keysToRemove []int) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error
	expected := make(map[int]int)

	result, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	for _, c := range contents {
		result, err = mapHooks.update(result,
			interpreter.Model.FromInt(c.key),
			interpreter.Model.FromInt(c.value),
			m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
		expected[c.key] = c.value
	}
	assertMapOk(t, expected, result, err, interpreter)

	for _, rem := range keysToRemove {
		result, err = mapHooks.remove(result,
			interpreter.Model.FromInt(rem),
			m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
		delete(expected, rem)
		assertMapOk(t, expected, result, err, interpreter)
	}
}

func TestMapRemove1(t *testing.T) {
	testRemove(t,
		[]concatCommonKeyTestKVP{},
		[]int{1})
}

func TestMapRemove2(t *testing.T) {
	testRemove(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]int{1})
}

func TestMapRemove3(t *testing.T) {
	testRemove(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		},
		[]int{2, 1})
}

func TestMapRemove4(t *testing.T) {
	testRemove(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		},
		[]int{1, 2, 1})
}

func TestMapRemove5(t *testing.T) {
	testRemove(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 3, value: 30},
		},
		[]int{2, 3, 1})
}

func TestMapConcatSelf1(t *testing.T) {
	interpreter := newTestInterpreter()
	var m1, result m.KReference
	var err error

	m1, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	m1, err = mapHooks.update(m1,
		interpreter.Model.FromInt(1),
		interpreter.Model.FromInt(10),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 10}, m1, err, interpreter)

	result, err = mapHooks.concat(m1, m1, m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 10}, result, err, interpreter)
	assertMapOk(t, map[int]int{1: 10}, m1, err, interpreter) // m1 the same
}

func TestMapConcatSelf2(t *testing.T) {
	interpreter := newTestInterpreter()
	var m1, result m.KReference
	var err error

	m1, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	m1, err = mapHooks.update(m1,
		interpreter.Model.FromInt(1),
		interpreter.Model.FromInt(10),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	m1, err = mapHooks.update(m1,
		interpreter.Model.FromInt(2),
		interpreter.Model.FromInt(20),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)

	assertMapOk(t, map[int]int{1: 10, 2: 20}, m1, err, interpreter)

	result, err = mapHooks.concat(m1, m1, m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 10, 2: 20}, result, err, interpreter)
	assertMapOk(t, map[int]int{1: 10, 2: 20}, m1, err, interpreter) // m1 the same
}

func testMapConcat(t *testing.T, contents1, contents2 []concatCommonKeyTestKVP) {
	interpreter := newTestInterpreter()
	var m1, m2, result m.KReference
	var err error
	expected1 := make(map[int]int)
	expected2 := make(map[int]int)
	expectedConcat := make(map[int]int)

	m1, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	for _, c1 := range contents1 {
		m1, err = mapHooks.update(m1,
			interpreter.Model.FromInt(c1.key),
			interpreter.Model.FromInt(c1.value),
			m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
		expected1[c1.key] = c1.value
		expectedConcat[c1.key] = c1.value
	}
	assertMapOk(t, expected1, m1, err, interpreter)

	m2, err = mapHooks.unit(m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	for _, c2 := range contents2 {
		m2, err = mapHooks.update(m2,
			interpreter.Model.FromInt(c2.key),
			interpreter.Model.FromInt(c2.value),
			m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
		expected2[c2.key] = c2.value
		expectedConcat[c2.key] = c2.value
	}
	assertMapOk(t, expected2, m2, err, interpreter)

	result, err = mapHooks.concat(m1, m2, m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, expectedConcat, result, err, interpreter)
	assertMapOk(t, expected1, m1, err, interpreter) // m1 the same
	assertMapOk(t, expected2, m2, err, interpreter) // m2 the same

	result, err = mapHooks.concat(m2, m1, m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, expectedConcat, result, err, interpreter)
	assertMapOk(t, expected1, m1, err, interpreter) // m1 the same
	assertMapOk(t, expected2, m2, err, interpreter) // m2 the same
}

func TestMapConcatEmpty1(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{},
		[]concatCommonKeyTestKVP{})
}

func TestMapConcatEmpty2(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{})
}

func TestMapConcat1(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
		})
}

func TestMapConcat2(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 3, value: 30},
		})
}

func TestMapConcat3(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		})
}

func TestMapConcat4(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 5, value: 50},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		})
}

func TestMapConcat5(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 5, value: 50},
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
		})
}

func TestMapConcat6(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 3, value: 30},
		})
}

func TestMapConcat7(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 3, value: 30},
			concatCommonKeyTestKVP{key: 1, value: 10},
		})
}

func TestMapConcat8(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 3, value: 30},
			concatCommonKeyTestKVP{key: 1, value: 10},
		})
}

func TestMapConcat9(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		})
}

func TestMapConcat10(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 1, value: 10},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
		})
}

func TestMapConcat11(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 3, value: 30},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
		})
}

func TestMapConcat12(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 3, value: 30},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 4, value: 40},
			concatCommonKeyTestKVP{key: 3, value: 30},
		})
}

func TestMapConcat13(t *testing.T) {
	testMapConcat(t,
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 3, value: 30},
			concatCommonKeyTestKVP{key: 4, value: 40},
		},
		[]concatCommonKeyTestKVP{
			concatCommonKeyTestKVP{key: 4, value: 40},
			concatCommonKeyTestKVP{key: 2, value: 20},
			concatCommonKeyTestKVP{key: 1, value: 10},
			concatCommonKeyTestKVP{key: 3, value: 30},
		})
}

func TestMapChoice(t *testing.T) {
	interpreter := newTestInterpreter()
	var result, choice m.KReference
	var err error

	result, err = mapHooks.element(
		interpreter.Model.FromInt(1),
		interpreter.Model.FromInt(3),
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertMapOk(t, map[int]int{1: 3}, result, err, interpreter)

	choice, err = mapHooks.choice(
		result,
		m.KLabelForMap, m.SortMap, m.InternedBottom, interpreter)
	assertIntOk(t, "1", choice, err, interpreter)
}

func assertMapOk(t *testing.T, expectedValues map[int]int, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	if !interpreter.Model.IsMap(actual) {
		t.Error("result is not a map")
	}
	if len(expectedValues) != interpreter.Model.MapSize(actual) {
		t.Errorf("map length doesn't match, have: %s",
			interpreter.Model.PrettyPrint(actual))
	}
	for eKey, eVal := range expectedValues {
		kKey := interpreter.Model.FromInt(eKey)
		if !interpreter.Model.MapContainsKey(actual, kKey) {
			t.Errorf("key expected but not found: %s",
				interpreter.Model.PrettyPrint(kKey))
		}
		kVal := interpreter.Model.FromInt(eVal)
		actualValue := interpreter.Model.MapGet(actual, kKey, m.NullReference)
		if !interpreter.Model.Equals(actualValue, kVal) {
			t.Errorf("map value mismatch for key %s. Have: %s Want: %s",
				interpreter.Model.PrettyPrint(kKey),
				interpreter.Model.PrettyPrint(actualValue),
				interpreter.Model.PrettyPrint(kVal))
		}
	}

	interpreter.Model.MapForEach(actual, func(k KReference, v KReference) bool {
		ik, _ := interpreter.Model.GetInt(k)
		iv, _ := interpreter.Model.GetInt(v)
		expectedVal, found := expectedValues[ik]
		if !found {
			t.Errorf("key not found: %d", ik)
		} else {
			if expectedVal != iv {
				t.Errorf("wrong value for key %d. Have: %d Want %d", ik, iv, expectedVal)
			}
		}
		return false
	})
}
