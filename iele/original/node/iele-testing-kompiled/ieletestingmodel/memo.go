// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:03:30.337

package ieletestingmodel

// MemoTable is a reference to a memoization table
type MemoTable int

// GetMemoizedValue searches for a value in the memo tables structure of the model.
func (ms *ModelState) GetMemoizedValue(memoTable MemoTable, keys ...KMapKey) (KReference, bool) {
	if ms.memoTables == nil {
		return NullReference, false
	}
	currentObj, tableFound := ms.memoTables[memoTable]
	if !tableFound {
		return NullReference, false
	}
	for _, key := range keys {
		currentMap, isMap := currentObj.(map[KMapKey]interface{})
		if !isMap {
			panic("wrong object found: memo tables need a level of map[KMapKey]interface{} for each key")
		}
		objectForKey, isKeyPresent := currentMap[key]
		if !isKeyPresent {
			return NullReference, false
		}
		currentObj = objectForKey
	}
	kref, isKref := currentObj.(KReference)
	if !isKref {
		panic("wrong object found: memo tables need to have a KReference on the last level")
	}
	return kref, true
}

// SetMemoizedValue inserts a value into the memo table structure, for a variable number of keys.
// It extends the tree up to where it is required.
func (ms *ModelState) SetMemoizedValue(memoized KReference, memoTable MemoTable, keys ...KMapKey) {
    // it is very important to prevent memoized values from being recycled
    ms.Preserve(memoized)

    // create necessary structures and insert
	if ms.memoTables == nil {
		ms.memoTables = make(map[MemoTable]interface{})
	}
	if len(keys) == 0 {
		// no keys, memo table is not really a table, it just contains one value
		ms.memoTables[memoTable] = memoized
		return
	}

	currentMapObj, tableFound := ms.memoTables[memoTable]
	if !tableFound {
		currentMapObj = make(map[KMapKey]interface{})
		ms.memoTables[memoTable] = currentMapObj
	}
	for i, key := range keys {
		currentMap, isMap := currentMapObj.(map[KMapKey]interface{})
		if !isMap {
			panic("wrong object found: memo tables need a level of map[KMapKey]interface{} for each key")
		}
		if i < len(keys)-1 {
			nextMap, nextMapExists := currentMap[key]
			if !nextMapExists {
				nextMap = make(map[KMapKey]interface{})
				currentMap[key] = nextMap
				currentMapObj = nextMap
			}
		} else {
			// last key
			currentMap[key] = memoized
		}
	}
}
