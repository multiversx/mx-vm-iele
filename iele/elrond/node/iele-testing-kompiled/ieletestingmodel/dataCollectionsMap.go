// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

package ieletestingmodel

type mapElementData struct {
	key   KReference
	value KReference
	next  int
}

// IsMap returns true if reference points to a map
func (ms *ModelState) IsMap(ref KReference) bool {
	return getRefType(ref) == mapRef
}

// IsMapWithSort returns true if reference points to a map with given sort
func (ms *ModelState) IsMapWithSort(ref KReference, expectedSort Sort) bool {
	refType, _, sort, _, _, _ := parseKrefCollection(ref)
	return refType == mapRef && sort == uint64(expectedSort)
}

// IsMapWithSortAndLabel returns true if reference points to a map with given sort and label
func (ms *ModelState) IsMapWithSortAndLabel(ref KReference, expectedSort Sort, expectedLabel KLabel) bool {
	refType, _, sort, label, _, _ := parseKrefCollection(ref)
	return refType == mapRef && sort == uint64(expectedSort) && label == uint64(expectedLabel)
}

// emptyMap yields an empty map reference.
func emptyMap(sort, label uint64) KReference {
	return createKrefCollection(mapRef, constDataRef, sort, label, 0, 0)
}

// EmptyMap yields an empty map reference.
func (ms *ModelState) EmptyMap(label KLabel, sort Sort) KReference {
	return emptyMap(uint64(sort), uint64(label))
}

// newMapSingleton creates a map with 1 key-value pair.
func (md *ModelData) newMapSingleton(collectionType kreferenceType, label KLabel, sort Sort, key KReference, value KReference) KReference {
	newIndex := len(md.allMapElements)
	md.allMapElements = append(md.allMapElements, mapElementData{
		key:   key,
		value: value,
		next:  -1,
	})
	return createKrefCollection(collectionType, md.selfRef, uint64(sort), uint64(label), uint64(newIndex), 1)
}

// MapSize yields the size of the map.
func (ms *ModelState) MapSize(mp KReference) int {
	refType, _, _, _, _, length := parseKrefCollection(mp)
	if refType != mapRef {
		panic("MapSize argument has wrong collection type")
	}
	return int(length)
}

// MapForEachCallback defines a callback argument for MapForEach.
type MapForEachCallback func(key KReference, value KReference) (shouldBreak bool)

// MapForEach calls argument f for each key-value pair in the map.
func (ms *ModelState) MapForEach(mp KReference, f MapForEachCallback) {
	refType, dataRef, _, _, index, length := parseKrefCollection(mp)
	if refType != mapRef {
		panic("argument has wrong collection type")
	}
	if length > 0 {
		data := ms.getData(dataRef)
		currentIndex := int(index)
		for currentIndex != -1 {
			elem := data.allMapElements[currentIndex]
			if f(elem.key, elem.value) {
				return
			}
			currentIndex = elem.next
		}
	}
}

// MapGet yields the value for a given key.
func (ms *ModelState) MapGet(mp KReference, key KReference, defaultValue KReference) KReference {
	result := defaultValue
	ms.MapForEach(mp, func(k KReference, v KReference) bool {
		if ms.Equals(k, key) {
			result = v
			return true
		}
		return false
	})
	return result
}

// MapContainsKey tells whether a key is present of not in a map.
func (ms *ModelState) MapContainsKey(mp KReference, key KReference) bool {
	found := false
	ms.MapForEach(mp, func(k KReference, v KReference) bool {
		if ms.Equals(k, key) {
			found = true
			return true
		}
		return false
	})
	return found
}

// MapUpdate inserts or updates value for key.
func (ms *ModelState) MapUpdate(mp KReference, key KReference, newValue KReference) KReference {
	return ms.mapUpdate(mapRef, mp, key, newValue)
}

func (ms *ModelState) mapUpdate(collectionType kreferenceType, mp KReference, key KReference, newValue KReference) KReference {
	refType, dataRef, sort, label, index, length := parseKrefCollection(mp)
	if refType != collectionType {
		panic("MapUpdate argument has wrong collection type")
	}
	if length == 0 {
		return ms.mainData.newMapSingleton(collectionType, KLabel(label), Sort(sort), key, newValue)
	}
	if !ms.MapContainsKey(mp, key) {
		// simply prepend new element
		newIndex := len(ms.mainData.allMapElements)
		ms.mainData.allMapElements = append(ms.mainData.allMapElements, mapElementData{
			key:   key,
			value: newValue,
			next:  int(index),
		})
		return createKrefCollection(mapRef, mainDataRef, sort, label, uint64(newIndex), length+1)
	}

	// key exists, copy all elements above element to update
	data := ms.getData(dataRef)
	currentIndex := int(index)
	var newIndex int
	resultIndex := -1 // first index in result
	previousResultIndex := -1
	stillRunning := true
	for stillRunning && currentIndex != -1 {
		elem := data.allMapElements[currentIndex]
		newIndex = len(data.allMapElements)

		// set up chain
		if previousResultIndex != -1 {
			data.allMapElements[previousResultIndex].next = newIndex
		}
		previousResultIndex = newIndex

		// fork structure
		if ms.Equals(key, elem.key) {
			data.allMapElements = append(data.allMapElements, mapElementData{
				key:   elem.key,
				value: newValue,
				next:  elem.next,
			})
			stillRunning = false // forking stops
		} else {
			data.allMapElements = append(data.allMapElements, mapElementData{
				key:   elem.key,
				value: elem.value,
				next:  -1,
			})
			previousResultIndex = newIndex
		}

		currentIndex = elem.next
		if resultIndex == -1 { // first step only
			resultIndex = newIndex
		}
	}
	return createKrefCollection(collectionType, dataRef, sort, label, uint64(resultIndex), length)
}

// MapRemove removes a key from a map
func (ms *ModelState) MapRemove(mp KReference, elem KReference) KReference {
	return ms.mapRemove(mapRef, mp, elem)
}

func (ms *ModelState) mapRemove(collectionType kreferenceType, mp KReference, key KReference) KReference {
	refType, dataRef, sort, label, index, length := parseKrefCollection(mp)
	if refType != collectionType {
		panic("MapRemove argument has wrong collection type")
	}
	if !ms.MapContainsKey(mp, key) {
		return mp // nothing to remove
	}

	// key exists, copy all elements above element to remove
	data := ms.getData(dataRef)
	currentIndex := int(index)
	var newIndex int
	resultIndex := -1 // first index in result
	previousResultIndex := -1
	stillRunning := true
	for stillRunning && currentIndex != -1 {
		elem := data.allMapElements[currentIndex]

		// fork structure
		if ms.Equals(key, elem.key) {
			if previousResultIndex == -1 {
				// first element is the one to remove
				// we simply return the map starting with the second element, no copying required
				resultIndex = elem.next
				if resultIndex == -1 {
					resultIndex = 0 // input has 1 item, result is empty
				}
				return createKrefCollection(mapRef, dataRef, sort, label, uint64(resultIndex), length-1)
			}

			// connect copied part of the chain with the original elements below
			data.allMapElements[previousResultIndex].next = elem.next
			stillRunning = false // forking stops
		} else {
			// copy element
			newIndex = len(data.allMapElements)
			data.allMapElements = append(data.allMapElements, mapElementData{
				key:   elem.key,
				value: elem.value,
				next:  -1,
			})
			if previousResultIndex != -1 {
				data.allMapElements[previousResultIndex].next = newIndex
			}
			previousResultIndex = newIndex
		}

		currentIndex = elem.next
		if resultIndex == -1 { // first step only
			resultIndex = newIndex
		}
	}
	return createKrefCollection(collectionType, dataRef, sort, label, uint64(resultIndex), length-1)
}

// MapConcatNoUpdate concatenates 2 maps. The maps cannot have different values for the same key.
func (ms *ModelState) MapConcatNoUpdate(mp1, mp2 KReference) (KReference, bool) {
	return ms.mapConcatNoUpdate(mapRef, mp1, mp2)
}

func (ms *ModelState) mapConcatNoUpdate(collectionType kreferenceType, mp1, mp2 KReference) (KReference, bool) {
	refType1, dataRef1, sort1, label1, index1, length1 := parseKrefCollection(mp1)
	refType2, dataRef2, _, label2, index2, length2 := parseKrefCollection(mp2)
	if refType1 != collectionType || refType2 != collectionType {
		return NoResult, false
	}
	if label1 != label2 {
		return NoResult, false
	}
	if length1 == 0 {
		return mp2, true
	}
	if length2 == 0 {
		return mp1, true
	}
	if dataRef1 != dataRef2 {
		// only check this after getting the empty aps out of the way, they don't have the same dataRef
		return NoResult, false
	}
	if length2 > length1 {
		// in order to only copy the shorter one, invert the arguments
		tempMp := mp1
		mp1 = mp2
		mp2 = tempMp
		refType1, dataRef1, sort1, label1, index1, length1 = parseKrefCollection(mp1)
		refType2, dataRef2, _, label2, index2, length2 = parseKrefCollection(mp2)
	}

	// copy
	md := ms.getData(dataRef1)
	fromIndex := int(index1)
	firstIndex := -1
	lastIndex := -1
	copyLength := uint64(0)
	for fromIndex != -1 {
		elem := md.allMapElements[fromIndex]
		var keyFound bool
		if collectionType == mapRef {
			mp2Value := ms.MapGet(mp2, elem.key, NullReference)
			if mp2Value != NullReference && !ms.Equals(mp2Value, elem.value) {
				// if key appears in both, values should also match
				return NoResult, false
			}
			keyFound = mp2Value != NullReference
		} else if collectionType == setRef {
			keyFound = ms.SetContains(mp2, elem.key)
		}

		if keyFound {
			newIndex := len(md.allMapElements)
			md.allMapElements = append(md.allMapElements, mapElementData{
				key:   elem.key,
				value: elem.value,
				next:  -1,
			})
			if lastIndex != -1 {
				md.allMapElements[lastIndex].next = newIndex
			}
			lastIndex = newIndex
			if firstIndex == -1 {
				firstIndex = newIndex
			}
			copyLength++
		}
		fromIndex = elem.next
	}

	if firstIndex == -1 {
		// copy is empty, all elements from mp1 were found in mp2
		return mp2, true
	}

	// chain mp1 copy -> mp2
	md.allMapElements[lastIndex].next = int(index2)

	return createKrefCollection(collectionType, dataRef1, sort1, label1, uint64(firstIndex), copyLength+length2), true
}

// MapDifference yields a map that has all keys from the first that are not present in the second.
func (ms *ModelState) MapDifference(mp1, mp2 KReference) (KReference, bool) {
	refType1, dataRef1, sort1, label1, index1, length1 := parseKrefCollection(mp1)
	refType2, dataRef2, _, label2, _, length2 := parseKrefCollection(mp2)
	if refType1 != mapRef || refType2 != mapRef {
		return NoResult, false
	}
	if dataRef1 != dataRef2 {
		return NoResult, false
	}
	if label1 != label2 {
		return NoResult, false
	}
	if length1 == 0 || length2 == 0 {
		return mp1, true
	}

	// copy
	md := ms.getData(dataRef1)
	fromIndex := int(index1)
	firstIndex := -1
	lastIndex := -1
	resultLength := 0
	for fromIndex != -1 {
		elem := md.allMapElements[fromIndex]
		if !ms.MapContainsKey(mp2, elem.key) {
			newIndex := len(md.allMapElements)
			md.allMapElements = append(md.allMapElements, mapElementData{
				key:   elem.key,
				value: elem.value,
				next:  -1,
			})
			resultLength++
			if lastIndex != -1 {
				md.allMapElements[lastIndex].next = newIndex
			}
			lastIndex = newIndex
			if firstIndex == -1 {
				firstIndex = newIndex
			}
		}
		fromIndex = elem.next
	}

	if firstIndex == -1 {
		// copy is empty, all keys from mp1 were found in mp2
		return emptyMap(sort1, label1), true
	}

	return createKrefCollection(mapRef, dataRef1, sort1, label1, uint64(firstIndex), uint64(resultLength)), true
}

// MapKeySet yields a set containing all keys.
func (ms *ModelState) MapKeySet(mp KReference) (KReference, bool) {
	refType, _, _, label, _, _ := parseKrefCollection(mp)
	if refType != mapRef {
		return NoResult, false
	}
	result := ms.EmptySet(KLabel(label), SortSet)
	ms.MapForEach(mp, func(k, _ KReference) bool {
		result = ms.SetAdd(result, k)
		return false
	})
	return result, true
}

// MapKeyList yields a list containing all keys.
func (ms *ModelState) MapKeyList(mp KReference) (KReference, bool) {
	refType, _, _, _, _, _ := parseKrefCollection(mp)
	if refType != mapRef {
		return NoResult, false
	}
	var keyList []KReference
	ms.MapForEach(mp, func(k, _ KReference) bool {
		keyList = append(keyList, k)
		return false
	})
	return ms.NewList(SortList, KLabelForList, keyList), true
}

// MapValueList yields a list containing all values.
func (ms *ModelState) MapValueList(mp KReference) (KReference, bool) {
	refType, _, _, _, _, _ := parseKrefCollection(mp)
	if refType != mapRef {
		return NoResult, false
	}
	var valueList []KReference
	ms.MapForEach(mp, func(_, v KReference) bool {
		valueList = append(valueList, v)
		return false
	})
	return ms.NewList(SortList, KLabelForList, valueList), true
}

// MapKeyChoice yields a key (any key) from the map.
func (ms *ModelState) MapKeyChoice(mp KReference) (KReference, bool) {
	return ms.mapKeyChoice(mapRef, mp)
}

func (ms *ModelState) mapKeyChoice(collectionType kreferenceType, mp KReference) (KReference, bool) {
	refType, dataRef, _, _, index, length := parseKrefCollection(mp)
	if refType != collectionType {
		return NoResult, false
	}
	if refType != mapRef {
		return NoResult, false
	}
	if length == 0 {
		return NoResult, false
	}
	elem := ms.getData(dataRef).allMapElements[int(index)]
	return elem.key, true
}

// MapInclusion returns true if map2 is included in map1.
func (ms *ModelState) MapInclusion(mp1, mp2 KReference) (bool, bool) {
	refType1, dataRef1, _, label1, _, _ := parseKrefCollection(mp1)
	refType2, dataRef2, _, label2, index2, length2 := parseKrefCollection(mp2)
	if refType1 != mapRef || refType2 != mapRef {
		return false, false
	}
	if dataRef1 != dataRef2 {
		return false, false
	}
	if label1 != label2 {
		return false, false
	}
	if length2 == 0 {
		return true, true
	}

	data := ms.getData(dataRef2)
	currentIndex := int(index2)
	for currentIndex != -1 {
		elem := data.allMapElements[currentIndex]
		if !ms.MapContainsKey(mp1, elem.key) {
			return false, true
		}
		currentIndex = elem.next
	}
	return true, true
}
