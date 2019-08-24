// File provided by the K Framework Go backend. Timestamp: 2019-08-24 18:56:17.501

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestArrayMake(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var int3 = interpreter.Model.FromInt(3)
	var int5 = interpreter.Model.FromInt(5)
	var bottom = m.InternedBottom

	arr, err = arrayHooks.makeEmpty(int3, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, bottom, bottom}, arr, err, interpreter)

	arr, err = arrayHooks.make(int3, int5, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, int5, []m.KReference{int5, int5, int5}, arr, err, interpreter)

	arr, err = arrayHooks.ctor(bottom, int5, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, bottom, bottom, bottom, bottom}, arr, err, interpreter)

}

// Without default (default = bottom)
func TestArrayMakeUpdateRemoveLookup1(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var elem m.KReference
	var int3 = interpreter.Model.FromInt(3)
	var int5 = interpreter.Model.FromInt(5)
	var int7 = interpreter.Model.FromInt(7)
	var bottom = m.InternedBottom

	// create
	arr, err = arrayHooks.makeEmpty(int3, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, bottom, bottom}, arr, err, interpreter)

	// updates
	arrayHooks.update(arr, interpreter.Model.FromInt(1), int5, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, int5, bottom}, arr, err, interpreter)

	arrayHooks.update(arr, interpreter.Model.FromInt(2), int7, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, int5, int7}, arr, err, interpreter)

	// test some lookups
	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(0), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBottomOk(t, elem, err, interpreter)

	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "5", elem, err, interpreter)

	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(2), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "7", elem, err, interpreter)

	// remove
	arrayHooks.remove(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, bottom, int7}, arr, err, interpreter)

	// lookup again
	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBottomOk(t, elem, err, interpreter)
}

// With default
func TestArrayMakeUpdateRemoveLookup2(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var elem m.KReference
	var int3 = interpreter.Model.FromInt(3)
	var int5 = interpreter.Model.FromInt(5)
	var int7 = interpreter.Model.FromInt(7)
	var defElem = interpreter.Model.FromInt(20)

	// create
	arr, err = arrayHooks.make(int3, defElem, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem, defElem}, arr, err, interpreter)

	// updates
	arrayHooks.update(arr, interpreter.Model.FromInt(1), int5, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, int5, defElem}, arr, err, interpreter)

	arrayHooks.update(arr, interpreter.Model.FromInt(2), int7, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, int5, int7}, arr, err, interpreter)

	// test some lookups
	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(0), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "20", elem, err, interpreter)

	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "5", elem, err, interpreter)

	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(2), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "7", elem, err, interpreter)

	// remove
	arrayHooks.remove(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem, int7}, arr, err, interpreter)

	// lookup again
	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(1), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "20", elem, err, interpreter)
}

func TestArrayIncreaseSize(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var elem m.KReference
	var defElem = interpreter.Model.FromInt(120)

	// create
	arr, err = arrayHooks.make(interpreter.Model.FromInt(20), defElem, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem}, arr, err, interpreter)

	// grow
	_, err = arrayHooks.update(arr, interpreter.Model.FromInt(11), interpreter.Model.FromInt(500), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, interpreter.Model.FromInt(500)}, arr, err, interpreter)

	// remove
	_, err = arrayHooks.remove(arr, interpreter.Model.FromInt(11), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem, defElem}, arr, err, interpreter)

	// test below the limit
	_, err = arrayHooks.update(arr, interpreter.Model.FromInt(19), interpreter.Model.FromInt(700), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	if err != nil {
		t.Error(err, interpreter)
	}
	elem, err = arrayHooks.lookup(arr, interpreter.Model.FromInt(19), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "700", elem, err, interpreter)

	// test above the limit
	_, err = arrayHooks.update(arr, interpreter.Model.FromInt(20), interpreter.Model.FromInt(700), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	if err == nil {
		t.Error("ErrIndexOutOfBounds expected")
	}
}

func TestArrayUpdateAll1(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var bottom = m.InternedBottom

	arr, _ = arrayHooks.makeEmpty(interpreter.Model.FromInt(4), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	list1 := interpreter.Model.NewList(m.SortInt, m.KLabelForList,
		[]m.KReference{
			interpreter.Model.FromInt(1),
			interpreter.Model.FromInt(2),
		})
	arrayHooks.updateAll(arr, interpreter.Model.FromInt(1), list1, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, interpreter.Model.FromInt(1), interpreter.Model.FromInt(2), bottom}, arr, err, interpreter)
}

func TestArrayUpdateAll2(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var bottom = m.InternedBottom

	arr, _ = arrayHooks.makeEmpty(interpreter.Model.FromInt(4), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	list2 := interpreter.Model.NewList(m.SortInt, m.KLabelForList,
		[]m.KReference{
			interpreter.Model.FromInt(1),
			interpreter.Model.FromInt(2),
			interpreter.Model.FromInt(3),
			interpreter.Model.FromInt(4),
		})
	arrayHooks.updateAll(arr, interpreter.Model.FromInt(1), list2, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, interpreter.Model.FromInt(1), interpreter.Model.FromInt(2), interpreter.Model.FromInt(3)}, arr, err, interpreter)
}

func TestArrayFill1(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var bottom = m.InternedBottom
	var fill = interpreter.Model.FromInt(123)

	arr, _ = arrayHooks.makeEmpty(interpreter.Model.FromInt(4), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	arrayHooks.fill(arr, interpreter.Model.FromInt(1), interpreter.Model.FromInt(3), fill, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, fill, fill, bottom}, arr, err, interpreter)
}

func TestArrayFill2(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var bottom = m.InternedBottom
	var fill = interpreter.Model.FromInt(123)

	arr, _ = arrayHooks.makeEmpty(interpreter.Model.FromInt(4), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	arrayHooks.fill(arr, interpreter.Model.FromInt(1), interpreter.Model.FromInt(10), fill, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, bottom, []m.KReference{bottom, fill, fill, fill}, arr, err, interpreter)
}

func TestArrayInKeys(t *testing.T) {
	interpreter := newTestInterpreter()
	var arr m.KReference
	var err error
	var result m.KReference
	var defElem = interpreter.Model.FromInt(20)

	// create
	arr, err = arrayHooks.make(interpreter.Model.FromInt(2), defElem, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertArrayOk(t, defElem, []m.KReference{defElem, defElem}, arr, err, interpreter)

	// updates
	arrayHooks.update(arr, m.IntZero, interpreter.Model.FromInt(5), m.LblDummy, m.SortInt, m.InternedBottom, interpreter)

	// test
	result, err = arrayHooks.inKeys(m.IntZero, arr, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, result, err, interpreter)

	result, err = arrayHooks.inKeys(interpreter.Model.FromInt(1), arr, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, result, err, interpreter)
}

func assertArrayOk(t *testing.T, expectedDefault m.KReference, expectedElems []m.KReference, a m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	/*expectedData := m.MakeDynamicArray(expectedMaxSize, expectedDefault)
	for i := 0; i < len(expectedElems); i++ {
		expectedData.Set(uint64(i), expectedElems[i])
	}
	expectedArray := &m.Array{Sort: m.SortInt, Data: expectedData}

	if !a.Equals(expectedArray) {

	}*/

	arr, isArray := interpreter.Model.GetArrayObject(a)
	if !isArray {
		t.Error("Result is not an Array.")
		return
	}
	if !interpreter.Model.Equals(expectedDefault, arr.Data.Default) {
		t.Errorf("Unexpected Array default. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(arr.Data.Default),
			interpreter.Model.PrettyPrint(expectedDefault))
	}
	sliceCopy := arr.Data.ToSlice()
	if len(expectedElems) != len(sliceCopy) {
		t.Errorf("Unexpected Array length. Got: %d Want: %d.",
			len(sliceCopy),
			len(expectedElems))
		return
	}
	for i := 0; i < len(expectedElems); i++ {
		if !interpreter.Model.Equals(sliceCopy[i], expectedElems[i]) {
			t.Errorf("Unexpected element at position %d. Got: %s Want: %s.",
				i,
				interpreter.Model.PrettyPrint(sliceCopy[i]),
				interpreter.Model.PrettyPrint(expectedElems[i]))
		}
	}
}

func assertBottomOk(t *testing.T, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	if !m.IsBottom(actual) {
		t.Errorf("Bottom expected. Got: %s", interpreter.Model.PrettyPrint(actual))
		return
	}
}
