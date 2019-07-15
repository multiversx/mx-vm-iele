// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:11:08.386

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type kreflectionHooksType int

const kreflectionHooks kreflectionHooksType = 0

var constKReflectionSortInt = m.NewStringConstant("Int")
var constKReflectionSortString = m.NewStringConstant("String")
var constKReflectionSortBytes = m.NewStringConstant("Bytes")
var constKReflectionSortBool = m.NewStringConstant("Bool")

func (kreflectionHooksType) sort(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if obj, t := interpreter.Model.GetKTokenObject(c); t {
		return interpreter.Model.NewString(obj.Sort.Name()), nil
	}
	if m.IsInt(c) {
		return constKReflectionSortInt, nil
	}
	if m.IsString(c) {
		return constKReflectionSortString, nil
	}
	if m.IsBytes(c) {
		return constKReflectionSortBytes, nil
	}
	if m.IsBool(c) {
		return constKReflectionSortBool, nil
	}
	if obj, t := interpreter.Model.GetMapObject(c); t {
		return interpreter.Model.NewString(obj.Sort.Name()), nil
	}
	if obj, t := interpreter.Model.GetListObject(c); t {
		return interpreter.Model.NewString(obj.Sort.Name()), nil
	}
	if obj, t := interpreter.Model.GetSetObject(c); t {
		return interpreter.Model.NewString(obj.Sort.Name()), nil
	}
	if obj, t := interpreter.Model.GetArrayObject(c); t {
		return interpreter.Model.NewString(obj.Sort.Name()), nil
	}

	return m.NoResult, &hookNotImplementedError{}

}

func (kreflectionHooksType) getKLabel(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if k, t := interpreter.Model.GetKApplyObject(c); t {
		return interpreter.Model.NewInjectedKLabel(k.Label), nil
	}
	return m.InternedBottom, nil
}

func (kreflectionHooksType) configuration(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return config, nil
}

var freshCounter int

func (kreflectionHooksType) fresh(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if str, t := interpreter.Model.GetString(c); t {
		sort := m.ParseSort(str)
		result, err := interpreter.freshFunction(sort, config, freshCounter)
		if err != nil {
			return m.NoResult, err
		}
		freshCounter++
		return result, nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (kreflectionHooksType) isConcrete(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.BoolTrue, nil
}

func (kreflectionHooksType) getenv(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (kreflectionHooksType) argv(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}
