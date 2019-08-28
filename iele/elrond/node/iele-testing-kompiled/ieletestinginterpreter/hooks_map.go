// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type mapHooksType int

const mapHooks mapHooksType = 0

// returns an empty map
func (mapHooksType) unit(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return interpreter.Model.EmptyMap(m.CollectionFor(lbl), sort), nil
}

// returns a map with 1 key to value mapping
func (mapHooksType) element(key m.KReference, val m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	empty := interpreter.Model.EmptyMap(m.CollectionFor(lbl), sort)
	return interpreter.Model.MapUpdate(empty, key, val), nil
}

func (mh mapHooksType) lookup(kmap m.KReference, key m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return mh.lookupOrDefault(kmap, key, m.InternedBottom, lbl, sort, config, interpreter)
}

func (mapHooksType) lookupOrDefault(kmap m.KReference, key m.KReference, defaultValue m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if interpreter.Model.IsMap(kmap) {
		return interpreter.Model.MapGet(kmap, key, defaultValue), nil
	}

	if m.IsBottom(kmap) {
		return m.InternedBottom, nil
	}

	return invalidArgsResult()
}

func (mapHooksType) update(kmap m.KReference, key m.KReference, newValue m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsMap(kmap) {
		return invalidArgsResult()
	}
	return interpreter.Model.MapUpdate(kmap, key, newValue), nil
}

func (mapHooksType) remove(kmap m.KReference, key m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsMap(kmap) {
		return invalidArgsResult()
	}
	return interpreter.Model.MapRemove(kmap, key), nil
}

func (mapHooksType) concat(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapConcatNoUpdate(kmap1, kmap2)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) difference(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapDifference(kmap1, kmap2)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) inKeys(key m.KReference, kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsMap(kmap) {
		return invalidArgsResult()
	}
	keyExists := interpreter.Model.MapContainsKey(kmap, key)
	return m.ToKBool(keyExists), nil
}

func (mapHooksType) keys(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapKeySet(kmap)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) keysList(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapKeyList(kmap)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) values(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapValueList(kmap)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) choice(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapKeyChoice(kmap)
	if !ok {
		return invalidArgsResult()
	}
	return result, nil
}

func (mapHooksType) size(kmap m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if !interpreter.Model.IsMap(kmap) {
		return invalidArgsResult()
	}
	size := interpreter.Model.MapSize(kmap)
	return interpreter.Model.FromInt(size), nil
}

func (mapHooksType) inclusion(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, ok := interpreter.Model.MapInclusion(kmap1, kmap2)
	if !ok {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (mapHooksType) updateAll(kmap1 m.KReference, kmap2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// TODO: test
	// m1, isMap1 := interpreter.Model.GetMapObject(kmap1)
	// m2, isMap2 := interpreter.Model.GetMapObject(kmap2)
	// if !isMap1 || !isMap2 {
	// 	return invalidArgsResult()
	// }
	// if m1.Label != m2.Label {
	// 	return invalidArgsResult()
	// }
	// data := make(map[m.KMapKey]m.KReference)
	// for key, value := range m1.Data {
	// 	data[key] = value
	// }
	// for key, value := range m2.Data {
	// 	data[key] = value
	// }
	// return interpreter.Model.NewMap(m1.Sort, m2.Label, data), nil
	return m.NoResult, m.GetHookNotImplementedError()
}

func (mapHooksType) removeAll(kmap m.KReference, kset m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// TODO: test
	// mp, isMap := interpreter.Model.GetMapObject(kmap)
	// s, isSet := interpreter.Model.GetSetObject(kset)
	// if !isMap || !isSet {
	// 	return invalidArgsResult()
	// }
	// data := make(map[m.KMapKey]m.KReference)
	// for key, value := range mp.Data {
	// 	_, exists := s.Data[key]
	// 	if !exists {
	// 		data[key] = value
	// 	}

	// }
	// return interpreter.Model.NewMap(mp.Sort, mp.Label, data), nil
	return m.NoResult, m.GetHookNotImplementedError()
}
