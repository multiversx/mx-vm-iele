package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestBoolHooks(t *testing.T) {
	var z m.K
	var err error

	z, err = boolHooks.and(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.and(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.and(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.and(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.andThen(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.andThen(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.andThen(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.andThen(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.or(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.or(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.or(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.or(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.orElse(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.orElse(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.orElse(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.orElse(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.not(m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.not(m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.implies(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.implies(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.implies(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.implies(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.ne(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.ne(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.ne(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.ne(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.eq(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.eq(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.eq(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.eq(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.xor(m.BoolTrue, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

	z, err = boolHooks.xor(m.BoolTrue, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.xor(m.BoolFalse, m.BoolTrue, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, true, z, err)

	z, err = boolHooks.xor(m.BoolFalse, m.BoolFalse, m.LblDummy, m.SortBool, m.InternedBottom)
	assertBoolOk(t, false, z, err)

}

func assertBoolOk(t *testing.T, expected bool, actual m.K, err error) {
	if err != nil {
		t.Error(err)
	}
    expectedK := m.ToBool(expected)
    if !actual.Equals(expectedK) {
        t.Errorf("Unexpected result. Got:%s Want:%s", m.PrettyPrint(actual), m.PrettyPrint(expectedK))
    }
}
