// File provided by the K Framework Go backend. Timestamp: 2019-05-21 00:58:51.823

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringConcat(t *testing.T) {
	result, err := stringHooks.concat(m.NewString("abc"), m.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "abcdef", result, err)
}

func TestStringEq(t *testing.T) {
	var result m.K
	var err error

	result, err = stringHooks.eq(m.NewString("abc"), m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom)
	assertBoolOk(t, true, result, err)

	result, err = stringHooks.ne(m.NewString("abc"), m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom)
	assertBoolOk(t, false, result, err)

	result, err = stringHooks.eq(m.NewString("yes"), m.NewString("no"), m.LblDummy, m.SortString, m.InternedBottom)
	assertBoolOk(t, false, result, err)

	result, err = stringHooks.ne(m.NewString(""), m.NewString("s"), m.LblDummy, m.SortString, m.InternedBottom)
	assertBoolOk(t, true, result, err)
}

func TestStringChr(t *testing.T) {
	var result m.K
	var err error

	result, err = stringHooks.chr(m.NewIntFromInt(97), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "a", result, err)

	result, err = stringHooks.chr(m.NewIntFromInt(32), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, " ", result, err)

	result, err = stringHooks.chr(m.NewIntFromInt(192), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "\xc0", result, err)
}

func TestStringFind(t *testing.T) {
	var result m.K
	var err error

	str := m.NewString("abcabcabcd")
	substr := m.NewString("abc")

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "0", result, err)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(1), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "3", result, err)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(3), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "3", result, err)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(7), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "-1", result, err)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(10), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "6", result, err)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "3", result, err)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "-1", result, err)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "-1", result, err)
}

func TestStringLength(t *testing.T) {
	len, err := stringHooks.length(m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "3", len, err)
}

func TestStringSubstr(t *testing.T) {
	var result m.K
	var err error

	str := m.NewString("abcdef")

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "ab", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "abcdef", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(1000), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "abcdef", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(2), m.NewIntFromInt(3), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "c", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(2), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "", result, err)

	result, err = stringHooks.substr(str, m.NewIntFromInt(6), m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "", result, err)
}

func TestString2Base2String(t *testing.T) {
	var i m.K
	var str m.K
	var err error
	i, err = stringHooks.string2base(m.NewString("5"), m.NewIntFromInt(8), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "5", i, err)

	str, err = stringHooks.base2string(i, m.NewIntFromInt(8), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "5", str, err)

	i, err = stringHooks.string2base(m.NewString("123abcdef123abcdef"), m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom)
	str, err = stringHooks.base2string(i, m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "123abcdef123abcdef", str, err)
}

func TestBase2String2Base(t *testing.T) {
	var i m.K
	var str m.K
	var err error

	str, err = stringHooks.base2string(m.NewIntFromInt(10), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "1010", str, err)

	i, err = stringHooks.string2base(str, m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "10", i, err)

	str, err = stringHooks.base2string(m.NewIntFromInt(123456789123456789), m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom)
	i, err = stringHooks.string2base(str, m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom)
	assertIntOk(t, "123456789123456789", i, err)
}

func TestString2Token(t *testing.T) {
	result, err := stringHooks.string2token(m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom)
	if err != nil {
		t.Error(err)
	}
	expected := &m.KToken{Sort: m.SortString, Value: "abc"}
	if !result.Equals(expected) {
		t.Errorf("Wrong KToken. Got: %s Want: %s.",
			m.PrettyPrint(result),
			m.PrettyPrint(expected))
	}
}

func TestToken2String(t *testing.T) {
	var result m.K
	var err error

	ktoken := &m.KToken{Sort: m.SortKResult, Value: "token!"}
	result, err = stringHooks.token2string(ktoken, m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "token!", result, err)

	result, err = stringHooks.token2string(m.NewIntFromInt(56), m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "56", result, err)

	result, err = stringHooks.token2string(m.BoolTrue, m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "true", result, err)

	result, err = stringHooks.token2string(m.BoolFalse, m.LblDummy, m.SortString, m.InternedBottom)
	assertStringOk(t, "false", result, err)

}

func assertStringOk(t *testing.T, expectedStr string, actual m.K, err error) {
	if err != nil {
		t.Error(err)
	}
	k, isString := actual.(*m.String)
	if !isString {
		t.Error("Result is not a String.")
		return
	}
	expected := m.NewString(expectedStr)
	if !expected.Equals(actual) {
		t.Errorf("Unexpected String. Got: %s Want: %s.",
			m.PrettyPrint(k),
			m.PrettyPrint(expected))
	}
}
