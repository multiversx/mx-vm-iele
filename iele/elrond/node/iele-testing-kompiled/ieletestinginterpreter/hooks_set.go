// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type setHooksType int

const setHooks setHooksType = 0

func (setHooksType) unit(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return interpreter.Model.EmptySet(m.CollectionFor(lbl), sort), nil
}

// returns a set with 1 element
func (setHooksType) element(e m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	empty := interpreter.Model.EmptySet(m.CollectionFor(lbl), sort)
	return interpreter.Model.SetAdd(empty, e), nil
}

func (setHooksType) in(e m.KReference, kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset) {
		return invalidArgsResult()
	}
	elemExists := interpreter.Model.SetContains(kset, e)
	return m.ToKBool(elemExists), nil
}

func (setHooksType) concat(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.SetConcat(kset1, kset2)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (setHooksType) difference(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset1) {
		return invalidArgsResult()
	}
	if !interpreter.Model.IsSet(kset2) {
		return invalidArgsResult()
	}
	result := interpreter.Model.EmptySet(lbl, sort)
	interpreter.Model.SetForEach(kset1, func(elem1 KReference) bool {
		if !interpreter.Model.SetContains(kset2, elem1) {
			result = interpreter.Model.SetAdd(result, elem1)
		}
		return false
	})
	return result, nil
}

// tests if kset1 is a subset of kset2
func (setHooksType) inclusion(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset1) {
		return invalidArgsResult()
	}
	if !interpreter.Model.IsSet(kset2) {
		return invalidArgsResult()
	}
	result := true
	interpreter.Model.SetForEach(kset1, func(elem1 KReference) bool {
		if !interpreter.Model.SetContains(kset2, elem1) {
			result = false
		}
		return true
	})
	return m.ToKBool(result), nil
}

func (setHooksType) intersection(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset1) {
		return invalidArgsResult()
	}
	if !interpreter.Model.IsSet(kset2) {
		return invalidArgsResult()
	}
	result := interpreter.Model.EmptySet(lbl, sort)
	interpreter.Model.SetForEach(kset1, func(elem1 KReference) bool {
		if interpreter.Model.SetContains(kset2, elem1) {
			result = interpreter.Model.SetAdd(result, elem1)
		}
		return false
	})
	return result, nil
}

func (setHooksType) choice(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.SetChoice(kset)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (setHooksType) size(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset) {
		return invalidArgsResult()
	}
	size := interpreter.Model.SetSize(kset)
	return interpreter.Model.FromInt(size), nil
}

func (setHooksType) set2list(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsSet(kset) {
		return invalidArgsResult()
	}
	var list []KReference
	interpreter.Model.SetForEach(kset, func(elem KReference) bool {
		list = append(list, elem)
		return false
	})
	return interpreter.Model.NewList(m.SortList, m.KLabelForList, list), nil
}

func (setHooksType) list2set(klist m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	l, isList := interpreter.Model.GetListObject(klist)
	if !isList {
		return invalidArgsResult()
	}
	result := interpreter.Model.EmptySet(m.KLabelForSet, m.SortSet)
	for _, e := range l.Data {
		result = interpreter.Model.SetAdd(result, e)
	}
	return result, nil
}
