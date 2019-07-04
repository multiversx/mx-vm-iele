// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
	"strconv"
	"strings"
)

type stringHooksType int

const stringHooks stringHooksType = 0

func (stringHooksType) concat(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return interpreter.Model.NewString(str1 + str2), nil
}

func (stringHooksType) lt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 < str2), nil
}

func (stringHooksType) le(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 <= str2), nil
}

func (stringHooksType) gt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 > str2), nil
}

func (stringHooksType) ge(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 >= str2), nil
}

func (stringHooksType) eq(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 == str2), nil
}

func (stringHooksType) ne(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str1, ok1 := interpreter.Model.GetString(c1)
	str2, ok2 := interpreter.Model.GetString(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToKBool(str1 != str2), nil
}

func (stringHooksType) chr(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b, ok := interpreter.Model.GetByte(c)
	if !ok {
		return invalidArgsResult()
	}
	bytes := []byte{b}
	return interpreter.Model.NewString(string(bytes)), nil
}

func (stringHooksType) find(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, ok1 := interpreter.Model.GetString(c1)
	substr, ok2 := interpreter.Model.GetString(c2)
	firstIdx, ok3 := interpreter.Model.GetPositiveInt(c3)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if firstIdx > len(str) {
		return invalidArgsResult()
	}

	result := strings.Index(str[firstIdx:], substr)
	if result == -1 {
		return m.IntMinusOne, nil
	}
	return interpreter.Model.FromInt(firstIdx + result), nil
}

func (stringHooksType) rfind(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, ok1 := interpreter.Model.GetString(c1)
	substr, ok2 := interpreter.Model.GetString(c2)
	lastIdx, ok3 := interpreter.Model.GetPositiveInt(c3)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if lastIdx > len(str) {
		return invalidArgsResult()
	}
	result := strings.LastIndex(str[0:lastIdx], substr)
	if result == -1 {
		return m.IntMinusOne, nil
	}
	return interpreter.Model.FromInt(result), nil
}

func (stringHooksType) length(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	k, ok := interpreter.Model.GetString(c)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.FromInt(len(k)), nil
}

func (stringHooksType) substr(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, ok1 := interpreter.Model.GetString(c1)
	from, ok2 := interpreter.Model.GetPositiveInt(c2) // from is inclusive
	to, ok3 := interpreter.Model.GetPositiveInt(c3)   // to is exclusive
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	length := len(str)
	if from > to || from > length {
		return invalidArgsResult()
	}
	if to > length {
		to = length
	}
	return interpreter.Model.NewString(str[from:to]), nil
}

func (stringHooksType) ord(arg m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, ok := interpreter.Model.GetString(arg)
	if !ok {
		return invalidArgsResult()
	}
	asBytes := []byte(str)
	if len(asBytes) == 0 {
		return invalidArgsResult()
	}
	return interpreter.Model.IntFromByte(asBytes[0]), nil
}

func (stringHooksType) int2string(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	iStr, ok := interpreter.Model.GetIntAsDecimalString(c)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewString(iStr), nil
}

func (stringHooksType) string2int(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) string2base(kstr m.KReference, kbase m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, strOk := interpreter.Model.GetString(kstr)
	base, baseOk := interpreter.Model.GetInt(kbase)
	if !strOk || !baseOk {
		return invalidArgsResult()
	}
	if base < 2 || base > 16 {
		return invalidArgsResult()
	}
	resultRef, err := interpreter.Model.ParseIntFromBase(str, base)
	if err != nil {
		return m.NoResult, err
	}
	return resultRef, nil
}

func (stringHooksType) base2string(kint m.KReference, kbase m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	base, baseOk := interpreter.Model.GetInt(kbase)
	if !baseOk {
		return invalidArgsResult()
	}
	if base < 2 || base > 16 {
		return invalidArgsResult()
	}
	str, convOk := interpreter.Model.GetIntToString(kint, base)
	if !convOk {
		return invalidArgsResult()
	}
	return interpreter.Model.NewString(str), nil
}

func (stringHooksType) string2token(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	str, ok := interpreter.Model.GetString(c)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewKToken(sort, str), nil
}

func (stringHooksType) token2string(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if k, typeOk := interpreter.Model.GetKTokenObject(c); typeOk {
		return interpreter.Model.NewString(k.Value), nil
	}
	if k, typeOk := m.CastToBool(c); typeOk {
		return interpreter.Model.NewString(strconv.FormatBool(k)), nil
	}
	if k, typeOk := interpreter.Model.GetString(c); typeOk {
		return interpreter.Model.NewString(k), nil // TODO: should do escaping
	}
	if kIntStr, typeOk := interpreter.Model.GetIntAsDecimalString(c); typeOk {
		return interpreter.Model.NewString(kIntStr), nil
	}
	if _, typeOk := interpreter.Model.GetFloatObject(c); typeOk {
		return m.NoResult, &hookNotImplementedError{}
	}

	return invalidArgsResult()
}

func (stringHooksType) float2string(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) uuid(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) floatFormat(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) string2float(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) replace(argS m.KReference, argToReplace m.KReference, argReplacement m.KReference, argCount m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kS, ok1 := interpreter.Model.GetString(argS)
	kToReplace, ok2 := interpreter.Model.GetString(argToReplace)
	kReplacement, ok3 := interpreter.Model.GetString(argReplacement)
	count, ok4 := interpreter.Model.GetPositiveInt(argCount)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return invalidArgsResult()
	}
	result := strings.Replace(kS, kToReplace, kReplacement, count)
	return interpreter.Model.NewString(result), nil
}

func (stringHooksType) replaceAll(argS m.KReference, argToReplace m.KReference, argReplacement m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kS, ok1 := interpreter.Model.GetString(argS)
	kToReplace, ok2 := interpreter.Model.GetString(argToReplace)
	kReplacement, ok3 := interpreter.Model.GetString(argReplacement)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if len(kS) == 0 {
		return argS, nil // empty
	}
	count := strings.Count(kS, kToReplace)
	if count == 0 {
		return argS, nil
	}
	if count == 1 && strings.HasPrefix(kS, kToReplace) {
		if len(kReplacement) == 0 {
			// just cut off the prefix
			return interpreter.Model.NewString(kS[len(kToReplace):]), nil
		}
		return interpreter.Model.NewString(kReplacement + kS[len(kToReplace):]), nil
	}

	result := strings.ReplaceAll(kS, kToReplace, kReplacement)
	return interpreter.Model.NewString(result), nil
}

func (stringHooksType) replaceFirst(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) countAllOccurrences(argS m.KReference, argToCount m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kS, ok1 := interpreter.Model.GetString(argS)
	kToCount, ok2 := interpreter.Model.GetString(argToCount)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}

	result := strings.Count(kS, kToCount)
	return interpreter.Model.FromInt(result), nil
}

func (stringHooksType) category(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) directionality(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) findChar(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) rfindChar(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}
