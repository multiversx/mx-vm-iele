// File provided by the K Framework Go backend. Timestamp: 2019-08-24 18:56:17.501

package ieletestingmodel

// ChoiceCallback defines a callback to be used in the lookups section.
type ChoiceCallback func(choiceVar KReference) (KReference, error)

// MapKeyChoice ...
func (ms *ModelState) MapKeyChoiceLookup(ref KReference, f ChoiceCallback) (KReference, error) {
	refType, dataRef, _, _, index, length := parseKrefCollection(ref)
	if refType != mapRef {
		panic("argument is not a map")
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

// SetChoice ...
func (ms *ModelState) SetChoiceLookup(ref KReference, f ChoiceCallback) (KReference, error) {
	setObj, ok := ms.GetSetObject(ref)
	if !ok {
		panic("argument is not a set")
	}
	for elemAsKey := range setObj.Data {
		setChoiceElem, setChoiceErr := ms.ToKItem(elemAsKey)
		if setChoiceErr != nil {
			panic("couldn't convert key")
		}
		choiceResult, err := f(setChoiceElem)
		if choiceResult != InternedBottom || err != nil {
			return choiceResult, err
		}
	}
	return InternedBottom, nil
}
