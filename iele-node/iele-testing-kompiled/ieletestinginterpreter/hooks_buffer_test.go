package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	var sb m.K
	var str m.K
	var err error
	sb, err = bufferHooks.empty(m.LblDummy, m.SortString, m.InternedBottom)
	assertStringBufferOk(t, "", sb, err)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "", str, err)

	_, err = bufferHooks.concat(sb, m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringBufferOk(t, "abc", sb, err)

	_, err = bufferHooks.concat(sb, m.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringBufferOk(t, "abcdef", sb, err)

	str, err = bufferHooks.toString(sb, m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "abcdef", str, err)
}

func assertStringBufferOk(t *testing.T, expectedStr string, actual m.K, err error) {
	if err != nil {
		t.Error(err)
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
