// File provided by the K Framework Go backend. Timestamp: 2019-06-24 23:27:10.928

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type mapHooksType int

const mapHooks mapHooksType = 0

// returns a map with 1 key to value mapping
func (mapHooksType) element(key m.K, val m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	kkey, keyOk := m.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	mp := make(map[m.KMapKey]m.K)
	mp[kkey] = val
	return &m.Map{Sort: sort, Label: m.CollectionFor(lbl), Data: mp}, nil
}

// returns an empty map
func (mapHooksType) unit(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	var mp map[m.KMapKey]m.K
	return &m.Map{Sort: sort, Label: m.CollectionFor(lbl), Data: mp}, nil
}

func (mh mapHooksType) lookup(kmap m.K, key m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return mh.lookupOrDefault(kmap, key, m.InternedBottom, lbl, sort, config, interpreter)
}

func (mapHooksType) lookupOrDefault(kmap m.K, key m.K, defaultRes m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	if mp, isMap := kmap.(*m.Map); isMap {
		kkey, keyOk := m.MapKey(key)
		if !keyOk {
			return defaultRes, nil
		}
		elem, found := mp.Data[kkey]
		if found {
			return elem, nil
		}
		return defaultRes, nil
	}

	if _, isBottom := kmap.(*m.Bottom); isBottom {
		return m.InternedBottom, nil
	}

	return invalidArgsResult()
}

func (mapHooksType) update(kmap m.K, key m.K, newValue m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := m.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	// implementing it as an "immutable" map
	// that is, creating a copy for each update (for now)
	// not the most efficient, not sure if necessary, but it is the safest
	newData := make(map[m.KMapKey]m.K)
	for oldKey, oldValue := range mp.Data {
		newData[oldKey] = oldValue
	}
	newData[kkey] = newValue
	return &m.Map{Sort: mp.Sort, Label: mp.Label, Data: newData}, nil
}

func (mapHooksType) remove(kmap m.K, key m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := m.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	// no updating of input map
	newData := make(map[m.KMapKey]m.K)
	for oldKey, oldValue := range mp.Data {
		if oldKey != kkey {
			newData[oldKey] = oldValue
		}
	}
	return &m.Map{Sort: mp.Sort, Label: mp.Label, Data: newData}, nil
}

func (mapHooksType) concat(kmap1 m.K, kmap2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	m1, isMap1 := kmap1.(*m.Map)
	m2, isMap2 := kmap2.(*m.Map)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.K)
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
	return &m.Map{Sort: m1.Sort, Label: m1.Label, Data: data}, nil
}

func (mapHooksType) difference(kmap1 m.K, kmap2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	m1, isMap1 := kmap1.(*m.Map)
	m2, isMap2 := kmap2.(*m.Map)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.K)
	for key, value := range m1.Data {
		_, exists := m2.Data[key]
		if !exists {
			data[key] = value
		}

	}
	return &m.Map{Sort: m1.Sort, Label: m1.Label, Data: data}, nil
}

func (mapHooksType) keys(kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	keySet := make(map[m.KMapKey]bool)
	for key := range mp.Data {
		keySet[key] = true
	}
	return &m.Set{Sort: m.SortSet, Label: m.KLabelForSet, Data: keySet}, nil
}

func (mapHooksType) keysList(kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	var keyList []m.K
	for key := range mp.Data {
		keyAsKItem, err := key.ToKItem()
		if err != nil {
			return m.NoResult, err
		}
		keyList = append(keyList, keyAsKItem)
	}
	return &m.List{Sort: m.SortList, Label: m.KLabelForList, Data: keyList}, nil
}

func (mapHooksType) inKeys(key m.K, kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	kkey, keyOk := m.MapKey(key)
	if !keyOk {
		return m.NoResult, errInvalidMapKey
	}
	_, keyExists := mp.Data[kkey]
	return m.ToBool(keyExists), nil
}

func (mapHooksType) values(kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	var valueList []m.K
	for _, value := range mp.Data {
		valueList = append(valueList, value)
	}
	return &m.List{Sort: m.SortList, Label: m.KLabelForList, Data: valueList}, nil
}

func (mapHooksType) choice(kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	for key := range mp.Data {
		return key.ToKItem()
	}
	return invalidArgsResult()
}

func (mapHooksType) size(kmap m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	if !isMap {
		return invalidArgsResult()
	}
	return m.NewIntFromInt(len(mp.Data)), nil
}

func (mapHooksType) inclusion(kmap1 m.K, kmap2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	m1, isMap1 := kmap1.(*m.Map)
	m2, isMap2 := kmap2.(*m.Map)
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

func (mapHooksType) updateAll(kmap1 m.K, kmap2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	m1, isMap1 := kmap1.(*m.Map)
	m2, isMap2 := kmap2.(*m.Map)
	if !isMap1 || !isMap2 {
		return invalidArgsResult()
	}
	if m1.Label != m2.Label {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.K)
	for key, value := range m1.Data {
		data[key] = value
	}
	for key, value := range m2.Data {
		data[key] = value
	}
	return &m.Map{Sort: m1.Sort, Label: m1.Label, Data: data}, nil
}

func (mapHooksType) removeAll(kmap m.K, kset m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	mp, isMap := kmap.(*m.Map)
	s, isSet := kset.(*m.Set)
	if !isMap || !isSet {
		return invalidArgsResult()
	}
	data := make(map[m.KMapKey]m.K)
	for key, value := range mp.Data {
		_, exists := s.Data[key]
		if !exists {
			data[key] = value
		}

	}
	return &m.Map{Sort: mp.Sort, Label: mp.Label, Data: data}, nil
}
