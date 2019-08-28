// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestingmodel

// ChoiceCallback defines a callback to be used in the lookups section.
type ChoiceCallback func(choiceVar KReference) (KReference, error)

// MapKeyChoiceLookup iterates through the keys of a map during a #mapChoice lookup
func (ms *ModelState) MapKeyChoiceLookup(ref KReference, f ChoiceCallback) (KReference, error) {
	return ms.mapKeyChoiceLookup(mapRef, ref, f)
}

// SetChoiceLookup iterates through the elements of a set during a #setChoice lookup
func (ms *ModelState) SetChoiceLookup(ref KReference, f ChoiceCallback) (KReference, error) {
	return ms.mapKeyChoiceLookup(setRef, ref, f)
}

func (ms *ModelState) mapKeyChoiceLookup(collectionType kreferenceType, ref KReference, f ChoiceCallback) (KReference, error) {
	refType, dataRef, _, _, index, length := parseKrefCollection(ref)
	if refType != collectionType {
		panic("argument is not the correct collection type")
	}
	if length > 0 {
		data := ms.getData(dataRef)
		currentIndex := int(index)
		for currentIndex != -1 {
			elem := data.allMapElements[currentIndex]
			choiceResult, err := f(elem.key)
			if choiceResult != InternedBottom || err != nil {
				return choiceResult, err
			}
			currentIndex = elem.next
		}
	}
	return InternedBottom, nil
}
