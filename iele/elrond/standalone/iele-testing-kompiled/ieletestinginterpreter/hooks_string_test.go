// File provided by the K Framework Go backend. Timestamp: 2019-06-25 00:00:28.701

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringConcat(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := stringHooks.concat(m.NewString("abc"), m.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)
}

func TestStringEq(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.K
	var err error

	result, err = stringHooks.eq(m.NewString("abc"), m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, true, result, err, interpreter)

	result, err = stringHooks.ne(m.NewString("abc"), m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, false, result, err, interpreter)

	result, err = stringHooks.eq(m.NewString("yes"), m.NewString("no"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, false, result, err, interpreter)

	result, err = stringHooks.ne(m.NewString(""), m.NewString("s"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, true, result, err, interpreter)
}

func TestStringChr(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.K
	var err error

	result, err = stringHooks.chr(m.NewIntFromInt(97), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "a", result, err, interpreter)

	result, err = stringHooks.chr(m.NewIntFromInt(32), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, " ", result, err, interpreter)

	result, err = stringHooks.chr(m.NewIntFromInt(192), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "\xc0", result, err, interpreter)
}

func TestStringFind(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.K
	var err error

	str := m.NewString("abcabcabcd")
	substr := m.NewString("abc")

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "0", result, err, interpreter)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(1), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(3), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.find(str, substr, m.NewIntFromInt(7), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(10), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "6", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)
}

func TestStringLength(t *testing.T) {
	interpreter := newTestInterpreter()
	len, err := stringHooks.length(m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", len, err, interpreter)
}

func TestStringSubstr(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.K
	var err error

	str := m.NewString("abcdef")

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "ab", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(1000), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(2), m.NewIntFromInt(3), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "c", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(2), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(0), m.NewIntFromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)

	result, err = stringHooks.substr(str, m.NewIntFromInt(6), m.NewIntFromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)
}

func TestString2Base2String(t *testing.T) {
	interpreter := newTestInterpreter()
	var i m.K
	var str m.K
	var err error
	i, err = stringHooks.string2base(m.NewString("5"), m.NewIntFromInt(8), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "5", i, err, interpreter)

	str, err = stringHooks.base2string(i, m.NewIntFromInt(8), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "5", str, err, interpreter)

	i, err = stringHooks.string2base(m.NewString("123abcdef123abcdef"), m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	str, err = stringHooks.base2string(i, m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "123abcdef123abcdef", str, err, interpreter)
}

func TestBase2String2Base(t *testing.T) {
	interpreter := newTestInterpreter()
	var i m.K
	var str m.K
	var err error

	str, err = stringHooks.base2string(m.NewIntFromInt(10), m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "1010", str, err, interpreter)

	i, err = stringHooks.string2base(str, m.NewIntFromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "10", i, err, interpreter)

	str, err = stringHooks.base2string(m.NewIntFromInt(123456789123456789), m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	i, err = stringHooks.string2base(str, m.NewIntFromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "123456789123456789", i, err, interpreter)
}

func TestString2Token(t *testing.T) {
	interpreter := newTestInterpreter()
	result, err := stringHooks.string2token(m.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	if err != nil {
		t.Error(err, interpreter)
	}
	expected := &m.KToken{Sort: m.SortString, Value: "abc"}
	if !interpreter.Model.Equals(result, expected) {
		t.Errorf("Wrong KToken. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(result),
			interpreter.Model.PrettyPrint(expected))
	}
}

func TestToken2String(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.K
	var err error

	ktoken := &m.KToken{Sort: m.SortKResult, Value: "token!"}
	result, err = stringHooks.token2string(ktoken, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "token!", result, err, interpreter)

	result, err = stringHooks.token2string(m.NewIntFromInt(56), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "56", result, err, interpreter)

	result, err = stringHooks.token2string(m.BoolTrue, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "true", result, err, interpreter)

	result, err = stringHooks.token2string(m.BoolFalse, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "false", result, err, interpreter)

}

func assertStringOk(t *testing.T, expectedStr string, actual m.K, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	k, isString := actual.(*m.String)
	if !isString {
		t.Error("Result is not a String.")
		return
	}
	expected := m.NewString(expectedStr)
	if !interpreter.Model.Equals(expected, actual) {
		t.Errorf("Unexpected String. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(k),
			interpreter.Model.PrettyPrint(expected))
	}
}
