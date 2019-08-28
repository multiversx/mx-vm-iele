// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestingmodel

// IsSet returns true if reference points to a set
func (ms *ModelState) IsSet(ref KReference) bool {
	return getRefType(ref) == setRef
}

// IsSetWithSort returns true if reference points to a set with given sort
func (ms *ModelState) IsSetWithSort(ref KReference, expectedSort Sort) bool {
	refType, _, sort, _, _, _ := parseKrefCollection(ref)
	return refType == setRef && sort == uint64(expectedSort)
}

// IsSetWithSortAndLabel returns true if reference points to a map with given sort and label
func (ms *ModelState) IsSetWithSortAndLabel(ref KReference, expectedSort Sort, expectedLabel KLabel) bool {
	refType, _, sort, label, _, _ := parseKrefCollection(ref)
	return refType == setRef && sort == uint64(expectedSort) && label == uint64(expectedLabel)
}

// emptySet yields an empty set reference.
func emptySet(sort, label uint64) KReference {
	return createKrefCollection(setRef, constDataRef, sort, label, 0, 0)
}

// EmptySet yields an empty set reference.
func (ms *ModelState) EmptySet(label KLabel, sort Sort) KReference {
	return emptySet(uint64(sort), uint64(label))
}

// SetSize yields the size of the set.
func (ms *ModelState) SetSize(mp KReference) int {
	refType, _, _, _, _, length := parseKrefCollection(mp)
	if refType != setRef {
		panic("SetSize argument is not a set")
	}
	return int(length)
}

// SetForEachCallback defines a callback argument for MapForEach.
type SetForEachCallback func(elem KReference) (shouldBreak bool)

// SetForEach calls argument f for each element in a set.
func (ms *ModelState) SetForEach(mp KReference, f SetForEachCallback) {
	refType, dataRef, _, _, index, length := parseKrefCollection(mp)
	if refType != setRef {
		panic("argument is not a map")
	}
	if length > 0 {
		data := ms.getData(dataRef)
		currentIndex := int(index)
		for currentIndex != -1 {
			elem := data.allMapElements[currentIndex]
			if f(elem.key) {
				return
			}
			currentIndex = elem.next
		}
	}
}

// SetContains tells whether an object is present or not in a set.
func (ms *ModelState) SetContains(mp KReference, element KReference) bool {
	found := false
	ms.SetForEach(mp, func(elem KReference) bool {
		if ms.Equals(elem, element) {
			found = true
			return true
		}
		return false
	})
	return found
}

// SetAdd inserts an element in a set.
func (ms *ModelState) SetAdd(mp KReference, elem KReference) KReference {
	return ms.mapUpdate(setRef, mp, elem, BoolTrue)
}

// SetConcat concatenates 2 sets.
func (ms *ModelState) SetConcat(mp1, mp2 KReference) (KReference, bool) {
	return ms.mapConcatNoUpdate(setRef, mp1, mp2)
}

// SetChoice yields an element (any element) from the set.
func (ms *ModelState) SetChoice(mp KReference) (KReference, bool) {
	return ms.mapKeyChoice(setRef, mp)
}
