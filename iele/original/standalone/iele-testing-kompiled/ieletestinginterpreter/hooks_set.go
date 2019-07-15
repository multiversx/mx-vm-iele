// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:05:41.660

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type setHooksType int

const setHooks setHooksType = 0

func (setHooksType) in(e m.KReference, kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s, isSet := interpreter.Model.GetSetObject(kset)
	if !isSet {
		return invalidArgsResult()
	}
	setElem, setElemOk := interpreter.Model.MapKey(e)
	if !setElemOk {
		return m.NoResult, errBadSetElement
	}
	_, exists := s.Data[setElem]
	return m.ToKBool(exists), nil
}

func (setHooksType) unit(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	var data map[m.KMapKey]bool
	return interpreter.Model.NewSet(sort, m.CollectionFor(lbl), data), nil
}

// returns a set with 1 element
func (setHooksType) element(e m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	setElem, setElemOk := interpreter.Model.MapKey(e)
	if !setElemOk {
		return m.NoResult, errBadSetElement
	}
	data := make(map[m.KMapKey]bool)
	data[setElem] = true
	return interpreter.Model.NewSet(sort, m.CollectionFor(lbl), data), nil
}

func (setHooksType) concat(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	if !isSet1 || !isSet2 {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]bool)
	for e1 := range s1.Data {
		data[e1] = true
	}
	for e2 := range s2.Data {
		data[e2] = true
	}
	return interpreter.Model.NewSet(sort, lbl, data), nil
}

func (setHooksType) difference(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	if !isSet1 || !isSet2 {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]bool)
	for e1 := range s1.Data {
		_, existsInS2 := s2.Data[e1]
		if !existsInS2 {
			data[e1] = true
		}
	}
	return interpreter.Model.NewSet(sort, lbl, data), nil
}

// tests if kset1 is a subset of kset2
func (setHooksType) inclusion(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	if !isSet1 || !isSet2 {
		return invalidArgsResult()
	}
	for e1 := range s1.Data {
		_, existsInS2 := s2.Data[e1]
		if !existsInS2 {
			return m.BoolFalse, nil
		}
	}
	return m.BoolTrue, nil
}

func (setHooksType) intersection(kset1 m.KReference, kset2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s1, isSet1 := interpreter.Model.GetSetObject(kset1)
	s2, isSet2 := interpreter.Model.GetSetObject(kset2)
	if !isSet1 || !isSet2 {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]bool)
	for e1 := range s1.Data {
		data[e1] = true
	}
	for e2 := range s2.Data {
		data[e2] = true
	}
	return interpreter.Model.NewSet(sort, lbl, data), nil
}

func (setHooksType) choice(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s, isSet := interpreter.Model.GetSetObject(kset)
	if !isSet {
		return invalidArgsResult()
	}
	for e := range s.Data {
		return interpreter.Model.ToKItem(e)
	}
	return invalidArgsResult()
}

func (setHooksType) size(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s, isSet := interpreter.Model.GetSetObject(kset)
	if !isSet {
		return invalidArgsResult()
	}
	return interpreter.Model.FromInt(len(s.Data)), nil
}

func (setHooksType) set2list(kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	s, isSet := interpreter.Model.GetSetObject(kset)
	if !isSet {
		return invalidArgsResult()
	}
	var list []m.KReference
	for e := range s.Data {
		elemAsKItem, err := interpreter.Model.ToKItem(e)
		if err != nil {
			return m.NoResult, err
		}
		list = append(list, elemAsKItem)
	}
	return interpreter.Model.NewList(m.SortList, m.KLabelForList, list), nil
}

func (setHooksType) list2set(klist m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	l, isList := interpreter.Model.GetListObject(klist)
	if !isList {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]bool)
	for _, e := range l.Data {
		setElem, setElemOk := interpreter.Model.MapKey(e)
		if !setElemOk {
			return m.NoResult, errBadSetElement
		}
		data[setElem] = true
	}
	return interpreter.Model.NewSet(m.SortSet, m.KLabelForSet, data), nil
}
