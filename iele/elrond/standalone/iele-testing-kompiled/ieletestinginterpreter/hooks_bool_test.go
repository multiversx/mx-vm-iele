// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:53:22.151

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestBoolHooks(t *testing.T) {
	interpreter := newTestInterpreter()
	var z m.K
	var err error

	z, err = boolHooks.and(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.and(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.and(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.and(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.andThen(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.andThen(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.andThen(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.andThen(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.or(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.or(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.or(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.or(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.orElse(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.orElse(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.orElse(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.orElse(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.not(m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.not(m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.implies(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.implies(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.implies(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.implies(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.ne(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.ne(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.ne(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.ne(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.eq(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.eq(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.eq(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.eq(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.xor(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = boolHooks.xor(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.xor(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = boolHooks.xor(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

}

func assertBoolOk(t *testing.T, expected bool, actual m.K, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	expectedK := m.ToBool(expected)
	if !interpreter.Model.Equals(actual, expectedK) {
		t.Errorf("Unexpected result. Got:%s Want:%s",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expectedK))
	}
}
