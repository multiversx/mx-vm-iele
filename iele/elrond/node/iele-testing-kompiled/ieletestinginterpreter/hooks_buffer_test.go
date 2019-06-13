// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:38:24.453

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	interpreter := newTestInterpreter()
	var sb m.K
	var str m.K
	var err error
	sb, err = bufferHooks.empty(m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "", sb, err, interpreter)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", str, err, interpreter)

	_, err = bufferHooks.concat(sb, m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "abc", sb, err, interpreter)

	_, err = bufferHooks.concat(sb, m.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringBufferOk(t, "abcdef", sb, err, interpreter)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", str, err, interpreter)
}

func assertStringBufferOk(t *testing.T, expectedStr string, actual m.K, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	k, typeOk := actual.(*m.StringBuffer)
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
