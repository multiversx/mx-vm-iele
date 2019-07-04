// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type mapHooksType int

const mapHooks mapHooksType = 0

// returns a map with 1 key to value mapping
func (mapHooksType) element(key m.KReference, val m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	kkey, keyOk := interpreter.Model.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	mp := make(map[m.KMapKey]m.KReference)
	mp[kkey] = val
	return interpreter.Model.NewMap(sort, m.CollectionFor(lbl), mp), nil
}

// returns an empty map
func (mapHooksType) unit(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	var mp map[m.KMapKey]m.KReference
	return interpreter.Model.NewMap(sort, m.CollectionFor(lbl), mp), nil
}

func (mh mapHooksType) lookup(kmap m.KReference, key m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return mh.lookupOrDefault(kmap, key, m.InternedBottom, lbl, sort, config, interpreter)
}

func (mapHooksType) lookupOrDefault(kmap m.KReference, key m.KReference, defaultRes m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if mp, isMap := interpreter.Model.GetMapObject(kmap); isMap {
		kkey, keyOk := interpreter.Model.MapKey(key)
		if !keyOk {
			return defaultRes, nil
		}
		elem, found := mp.Data[kkey]
		if found {
			return elem, nil
		}
		return defaultRes, nil
	}

	if m.IsBottom(kmap) {
		return m.InternedBottom, nil
	}

	return invalidArgsResult()
}

func (mapHooksType) update(kmap m.KReference, key m.KReference, newValue m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := interpreter.Model.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	// implementing it as an "immutable" map
	// that is, creating a copy for each update (for now)
	// not the most efficient, not sure if necessary, but it is the safest
	newData := make(map[m.KMapKey]m.KReference)
	for oldKey, oldValue := range mp.Data {
		newData[oldKey] = oldValue
	}
	newData[kkey] = newValue
	return interpreter.Model.NewMap(mp.Sort, mp.Label, newData), nil
}

func (mapHooksType) remove(kmap m.KReference, key m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := interpreter.Model.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	// no updating of input map
	newData := make(map[m.KMapKey]m.KReference)
	for oldKey, oldValue := range mp.Data {
		if oldKey != kkey {
			newData[oldKey] = oldValue
		}
	}
	return interpreter.Model.NewMap(mp.Sort, mp.Label, newData), nil
}

func (mapHooksType) concat(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	m1, isMap1 := interpreter.Model.GetMapObject(kmap1)
	m2, isMap2 := interpreter.Model.GetMapObject(kmap2)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.KReference)
	for key, value := range m1.Data {
		data[key] = value
	}
	for key, value := range m2.Data {
		m1Val, exists := m1.Data[key]
		if exists {
			if m1Val != value {
				return invalidArgsResult()
			}
		} else {
			data[key] = value
		}
	}
	return interpreter.Model.NewMap(m1.Sort, m2.Label, data), nil
}

func (mapHooksType) difference(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	m1, isMap1 := interpreter.Model.GetMapObject(kmap1)
	m2, isMap2 := interpreter.Model.GetMapObject(kmap2)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.KReference)
	for key, value := range m1.Data {
		_, exists := m2.Data[key]
		if !exists {
			data[key] = value
		}

	}
	return interpreter.Model.NewMap(m1.Sort, m2.Label, data), nil
}

func (mapHooksType) keys(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	keySet := make(map[m.KMapKey]bool)
	for key := range mp.Data {
		keySet[key] = true
	}
	return interpreter.Model.NewSet(m.SortSet, mp.Label, keySet), nil
}

func (mapHooksType) keysList(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	var keyList []m.KReference
	for key := range mp.Data {
		keyAsKItem, err := interpreter.Model.ToKItem(key)
		if err != nil {
			return m.NoResult, err
		}
		keyList = append(keyList, keyAsKItem)
	}
	return interpreter.Model.NewList(m.SortList, m.KLabelForList, keyList), nil
}

func (mapHooksType) inKeys(key m.KReference, kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := interpreter.Model.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	_, keyExists := mp.Data[kkey]
	return m.ToKBool(keyExists), nil
}

func (mapHooksType) values(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	var valueList []m.KReference
	for _, value := range mp.Data {
		valueList = append(valueList, value)
	}
	return interpreter.Model.NewList(m.SortList, m.KLabelForList, valueList), nil
}

func (mapHooksType) choice(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	for key := range mp.Data {
		return interpreter.Model.ToKItem(key)
	}
	return invalidArgsResult()
}

func (mapHooksType) size(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	if !isMap {
		return invalidArgsResult()
	}
	return interpreter.Model.FromInt(len(mp.Data)), nil
}

func (mapHooksType) inclusion(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	m1, isMap1 := interpreter.Model.GetMapObject(kmap1)
	m2, isMap2 := interpreter.Model.GetMapObject(kmap2)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	for m2Key := range m2.Data {
		_, exists := m1.Data[m2Key]
		if !exists {
			return m.BoolFalse, nil
		}
	}
	return m.BoolTrue, nil
}

func (mapHooksType) updateAll(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	m1, isMap1 := interpreter.Model.GetMapObject(kmap1)
	m2, isMap2 := interpreter.Model.GetMapObject(kmap2)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.KReference)
	for key, value := range m1.Data {
		data[key] = value
	}
	for key, value := range m2.Data {
		data[key] = value
	}
	return interpreter.Model.NewMap(m1.Sort, m2.Label, data), nil
}

func (mapHooksType) removeAll(kmap m.KReference, kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	mp, isMap := interpreter.Model.GetMapObject(kmap)
	s, isSet := interpreter.Model.GetSetObject(kset)
	if !isMap || !isSet {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.KReference)
	for key, value := range mp.Data {
		_, exists := s.Data[key]
		if !exists {
			data[key] = value
		}

	}
	return interpreter.Model.NewMap(mp.Sort, mp.Label, data), nil
}
