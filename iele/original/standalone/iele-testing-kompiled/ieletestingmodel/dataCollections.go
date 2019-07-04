// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestingmodel

// Map is a KObject representing a map in K
type Map struct {
	Sort  Sort
	Label KLabel
	Data  map[KMapKey]KReference
}

func (*Map) referenceType() kreferenceType {
	return mapRef
}

// Set is a KObject representing a set in K
type Set struct {
	Sort  Sort
	Label KLabel
	Data  map[KMapKey]bool
}

func (*Set) referenceType() kreferenceType {
	return setRef
}

// List is a KObject representing a list in K
type List struct {
	Sort  Sort
	Label KLabel
	Data  []KReference
}

func (*List) referenceType() kreferenceType {
	return listRef
}

// Array is a KObject holding an array that can grow
type Array struct {
	Sort Sort
	Data *DynamicArray
}

func (*Array) referenceType() kreferenceType {
	return arrayRef
}

// IsMap returns true if reference points to a map with given sort
func (ms *ModelState) IsMap(ref KReference, expectedSort Sort) bool {
	obj, typeOk := ms.GetMapObject(ref)
	if !typeOk {
		return false
	}
	return obj.Sort == expectedSort
}

// IsSet returns true if reference points to a set with given sort
func (ms *ModelState) IsSet(ref KReference, expectedSort Sort) bool {
	obj, typeOk := ms.GetSetObject(ref)
	if !typeOk {
		return false
	}
	return obj.Sort == expectedSort
}

// IsList returns true if reference points to a list with given sort
func (ms *ModelState) IsList(ref KReference, expectedSort Sort) bool {
	obj, typeOk := ms.GetListObject(ref)
	if !typeOk {
		return false
	}
	return obj.Sort == expectedSort
}

// IsArray returns true if reference points to an array with given sort
func (ms *ModelState) IsArray(ref KReference, expectedSort Sort) bool {
	obj, typeOk := ms.GetArrayObject(ref)
	if !typeOk {
		return false
	}
	return obj.Sort == expectedSort
}

// GetListObject yields the cast object for a List reference, if possible.
func (ms *ModelState) GetListObject(ref KReference) (*List, bool) {
	if ref.refType != listRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*List)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// IsEmptyList returns true only if argument references an empty list, with given sort and label.
func (ms *ModelState) IsEmptyList(ref KReference, expectedSort Sort, expectedLabel KLabel) bool {
	castObj, typeOk := ms.GetListObject(ref)
	if !typeOk {
		return false
	}
    if castObj.Sort != expectedSort {
        return false
    }
    if castObj.Label != expectedLabel {
        return false
    }
	return len(castObj.Data) == 0
}

// ListSplitHeadTail returns true only if argument references an empty list.
// Returns nothing if it is not a list, it is empty, or if sort or label do not match.
func (ms *ModelState) ListSplitHeadTail(ref KReference, expectedSort Sort, expectedLabel KLabel) (ok bool, head KReference, tail KReference) {
	castObj, typeOk := ms.GetListObject(ref)
	if !typeOk {
		return false, NullReference, NullReference
	}
	if castObj.Sort != expectedSort {
	    return false, NullReference, NullReference
	}
	if castObj.Label != expectedLabel {
        return !ok, NullReference, NullReference
    }
    if len(castObj.Data) == 0 {
        return false, NullReference, NullReference
    }
    tailRef := ms.addObject(&List{Sort: castObj.Sort, Label: castObj.Label, Data: castObj.Data[1:]})
	return true, castObj.Data[0], tailRef
}

// GetMapObject yields the cast object for a Map reference, if possible.
func (ms *ModelState) GetMapObject(ref KReference) (*Map, bool) {
	if ref.refType != mapRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*Map)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// GetSetObject yields the cast object for a Set reference, if possible.
func (ms *ModelState) GetSetObject(ref KReference) (*Set, bool) {
	if ref.refType != setRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*Set)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// GetArrayObject yields the cast object for an Array reference, if possible.
func (ms *ModelState) GetArrayObject(ref KReference) (*Array, bool) {
	if ref.refType != arrayRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*Array)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// NewList creates a new object and returns the reference.
func (ms *ModelState) NewList(sort Sort, label KLabel, value []KReference) KReference {
	return ms.addObject(&List{Sort: sort, Label: label, Data: value})
}

// NewMap creates a new object and returns the reference.
func (ms *ModelState) NewMap(sort Sort, label KLabel, value map[KMapKey]KReference) KReference {
	return ms.addObject(&Map{Sort: sort, Label: label, Data: value})
}

// NewSet creates a new object and returns the reference.
func (ms *ModelState) NewSet(sort Sort, label KLabel, value map[KMapKey]bool) KReference {
	return ms.addObject(&Set{Sort: sort, Label: label, Data: value})
}

// NewArray creates a new object and returns the reference.
func (ms *ModelState) NewArray(sort Sort, value *DynamicArray) KReference {
	return ms.addObject(&Array{Sort: sort, Data: value})
}