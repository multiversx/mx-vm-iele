// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:14:14.526

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestStringConcat(t *testing.T) {
	interpreter := newTestInterpreter()

	result, err := stringHooks.concat(interpreter.Model.NewString("abc"), interpreter.Model.NewString("def"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)
}

func TestStringEq(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error

	result, err = stringHooks.eq(interpreter.Model.NewString("abc"), interpreter.Model.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, true, result, err, interpreter)

	result, err = stringHooks.ne(interpreter.Model.NewString("abc"), interpreter.Model.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, false, result, err, interpreter)

	result, err = stringHooks.eq(interpreter.Model.NewString("yes"), interpreter.Model.NewString("no"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, false, result, err, interpreter)

	result, err = stringHooks.ne(interpreter.Model.NewString(""), interpreter.Model.NewString("s"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertBoolOk(t, true, result, err, interpreter)
}

func TestStringChr(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error

	result, err = stringHooks.chr(interpreter.Model.FromInt(97), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "a", result, err, interpreter)

	result, err = stringHooks.chr(interpreter.Model.FromInt(32), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, " ", result, err, interpreter)

	result, err = stringHooks.chr(interpreter.Model.FromInt(192), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "\xc0", result, err, interpreter)
}

func TestStringFind(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error

	str := interpreter.Model.NewString("abcabcabcd")
	substr := interpreter.Model.NewString("abc")

	result, err = stringHooks.find(str, substr, interpreter.Model.FromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "0", result, err, interpreter)

	result, err = stringHooks.find(str, substr, interpreter.Model.FromInt(1), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.find(str, substr, interpreter.Model.FromInt(3), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.find(str, substr, interpreter.Model.FromInt(7), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, interpreter.Model.FromInt(10), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "6", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, interpreter.Model.FromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, interpreter.Model.FromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)

	result, err = stringHooks.rfind(str, substr, interpreter.Model.FromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", result, err, interpreter)
}

func TestStringLength(t *testing.T) {
	interpreter := newTestInterpreter()
	len, err := stringHooks.length(interpreter.Model.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "3", len, err, interpreter)
}

func TestStringSubstr(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error

	str := interpreter.Model.NewString("abcdef")

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(0), interpreter.Model.FromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "ab", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(0), interpreter.Model.FromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(0), interpreter.Model.FromInt(1000), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "abcdef", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(2), interpreter.Model.FromInt(3), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "c", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(2), interpreter.Model.FromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(0), interpreter.Model.FromInt(0), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)

	result, err = stringHooks.substr(str, interpreter.Model.FromInt(6), interpreter.Model.FromInt(6), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "", result, err, interpreter)
}

func TestString2Base2String(t *testing.T) {
	interpreter := newTestInterpreter()
	var i m.KReference
	var str m.KReference
	var err error
	i, err = stringHooks.string2base(interpreter.Model.NewString("5"), interpreter.Model.FromInt(8), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "5", i, err, interpreter)

	str, err = stringHooks.base2string(i, interpreter.Model.FromInt(8), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "5", str, err, interpreter)

	i, err = stringHooks.string2base(interpreter.Model.NewString("123abcdef123abcdef"), interpreter.Model.FromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	str, err = stringHooks.base2string(i, interpreter.Model.FromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "123abcdef123abcdef", str, err, interpreter)
}

func TestBase2String2Base(t *testing.T) {
	interpreter := newTestInterpreter()
	var i m.KReference
	var str m.KReference
	var err error

	str, err = stringHooks.base2string(interpreter.Model.FromInt(10), interpreter.Model.FromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "1010", str, err, interpreter)

	i, err = stringHooks.string2base(str, interpreter.Model.FromInt(2), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "10", i, err, interpreter)

	str, err = stringHooks.base2string(interpreter.Model.FromInt(123456789123456789), interpreter.Model.FromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	i, err = stringHooks.string2base(str, interpreter.Model.FromInt(16), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertIntOk(t, "123456789123456789", i, err, interpreter)
}

func TestString2Token(t *testing.T) {
	interpreter := newTestInterpreter()
	result, err := stringHooks.string2token(interpreter.Model.NewString("abc"), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	if err != nil {
		t.Error(err, interpreter)
	}
	expected := interpreter.Model.NewKToken(m.SortString, "abc")
	if !interpreter.Model.Equals(result, expected) {
		t.Errorf("Wrong KToken. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(result),
			interpreter.Model.PrettyPrint(expected))
	}
}

func TestToken2String(t *testing.T) {
	interpreter := newTestInterpreter()
	var result m.KReference
	var err error

	ktoken := interpreter.Model.NewKToken(m.SortKResult, "token!")
	result, err = stringHooks.token2string(ktoken, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "token!", result, err, interpreter)

	result, err = stringHooks.token2string(interpreter.Model.FromInt(56), m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "56", result, err, interpreter)

	result, err = stringHooks.token2string(m.BoolTrue, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "true", result, err, interpreter)

	result, err = stringHooks.token2string(m.BoolFalse, m.LblDummy, m.SortString, m.InternedBottom, interpreter)
	assertStringOk(t, "false", result, err, interpreter)

}

func assertStringOk(t *testing.T, expectedStr string, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	if !m.IsString(actual) {
		t.Error("Result is not a String.")
		return
	}
	expected := interpreter.Model.NewString(expectedStr)
	if !interpreter.Model.Equals(expected, actual) {
		t.Errorf("Unexpected String. Got: %s Want: %s.",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
}
