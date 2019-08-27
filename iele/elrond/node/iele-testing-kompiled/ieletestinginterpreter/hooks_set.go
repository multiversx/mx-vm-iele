// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

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
	// s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	// s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	// if !isSet1 || !isSet2 {
	// 	return invalidArgsResult()
	// }
	// data := make(map[m.KMapKey]bool)
	// for e1 := range s1.Data {
	// 	_, existsInS2 := s2.Data[e1]
	// 	if !existsInS2 {
	// 		data[e1] = true
	// 	}
	// }
	// return interpreter.Model.NewSet(sort, lbl, data), nil
	return m.NoResult, &hookNotImplementedError{}
}

// tests if kset1 is a subset of kset2
func (setHooksType) inclusion(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	// s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	// if !isSet1 || !isSet2 {
	// 	return invalidArgsResult()
	// }
	// for e1 := range s1.Data {
	// 	_, existsInS2 := s2.Data[e1]
	// 	if !existsInS2 {
	// 		return m.BoolFalse, nil
	// 	}
	// }
	// return m.BoolTrue, nil
	return m.NoResult, &hookNotImplementedError{}
}

func (setHooksType) intersection(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	// s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	// if !isSet1 || !isSet2 {
	// 	return invalidArgsResult()
	// }
	// data := make(map[m.KMapKey]bool)
	// for e1 := range s1.Data {
	// 	data[e1] = true
	// }
	// for e2 := range s2.Data {
	// 	data[e2] = true
	// }
	// return interpreter.Model.NewSet(sort, lbl, data), nil
	return m.NoResult, &hookNotImplementedError{}
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
