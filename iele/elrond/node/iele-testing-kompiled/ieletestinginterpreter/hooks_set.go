// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:43:22.780

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type setHooksType int

const setHooks setHooksType = 0

func (setHooksType) in(e m.K, kset m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s, isSet := kset.(*m.Set)
	if !isSet {
		return invalidArgsResult()
	}
	setElem, setElemOk := m.MapKey(e)
	if !setElemOk {
		return m.NoResult, errBadSetElement
	}
	_, exists := s.Data[setElem]
	return m.ToBool(exists), nil
}

func (setHooksType) unit(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	var data map[m.KMapKey]bool
	return &m.Set{Sort: sort, Label: m.CollectionFor(lbl), Data: data}, nil
}

// returns a set with 1 element
func (setHooksType) element(e m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	setElem, setElemOk := m.MapKey(e)
	if !setElemOk {
		return m.NoResult, errBadSetElement
	}
	data := make(map[m.KMapKey]bool)
	data[setElem] = true
	return &m.Set{Sort: sort, Label: m.CollectionFor(lbl), Data: data}, nil
}

func (setHooksType) concat(kset1 m.K, kset2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s1, isSet1 := kset1.(*m.Set)
	s2, isSet2 := kset2.(*m.Set)
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
	return &m.Set{Sort: sort, Label: lbl, Data: data}, nil
}

func (setHooksType) difference(kset1 m.K, kset2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s1, isSet1 := kset1.(*m.Set)
	s2, isSet2 := kset2.(*m.Set)
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
	return &m.Set{Sort: sort, Label: lbl, Data: data}, nil
}

// tests if kset1 is a subset of kset2
func (setHooksType) inclusion(kset1 m.K, kset2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s1, isSet1 := kset1.(*m.Set)
	s2, isSet2 := kset2.(*m.Set)
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

func (setHooksType) intersection(kset1 m.K, kset2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s1, isSet1 := kset1.(*m.Set)
	s2, isSet2 := kset2.(*m.Set)
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
	return &m.Set{Sort: sort, Label: lbl, Data: data}, nil
}

func (setHooksType) choice(kset m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s, isSet := kset.(*m.Set)
	if !isSet {
		return invalidArgsResult()
	}
	for e := range s.Data {
		return e.ToKItem()
	}
	return invalidArgsResult()
}

func (setHooksType) size(kset m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s, isSet := kset.(*m.Set)
	if !isSet {
		return invalidArgsResult()
	}
	return m.NewIntFromInt(len(s.Data)), nil
}

func (setHooksType) set2list(kset m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	s, isSet := kset.(*m.Set)
	if !isSet {
		return invalidArgsResult()
	}
	var list []m.K
	for e := range s.Data {
		elemAsKItem, err := e.ToKItem()
		if err != nil {
			return m.NoResult, err
		}
		list = append(list, elemAsKItem)
	}
	return &m.List{Sort: m.SortList, Label: m.KLabelForList, Data: list}, nil
}

func (setHooksType) list2set(klist m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l, isList := klist.(*m.List)
	if !isList {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]bool)
	for _, e := range l.Data {
		setElem, setElemOk := m.MapKey(e)
		if !setElemOk {
			return m.NoResult, errBadSetElement
		}
		data[setElem] = true
	}
	return &m.Set{Sort: m.SortSet, Label: m.KLabelForSet, Data: data}, nil
}
