// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	interpreter := newTestInterpreter()
	var sb m.KReference
	var str m.KReference
	var err error
	sb, err = bufferHooks.empty(m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "", sb, err, interpreter)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", str, err, interpreter)

	_, err = bufferHooks.concat(sb, interpreter.Model.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "abc", sb, err, interpreter)

	_, err = bufferHooks.concat(sb, interpreter.Model.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "abcdef", sb, err, interpreter)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", str, err, interpreter)
}

func assertStringBufferOk(t *testing.T, expectedStr string, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	k, typeOk := interpreter.Model.GetStringBufferObject(actual)
	if !typeOk {
		t.Error("Result is not a StringBuffer.")
		return
	}
	if expectedStr != k.Value.String() {
		t.Errorf("Unexpected StringBuffer.String(). Got: %s Want: %s.",
			k.Value.String(),
			expectedStr)
	}
}
