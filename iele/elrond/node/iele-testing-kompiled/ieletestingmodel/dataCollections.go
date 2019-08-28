// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestingmodel

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

// CollectionSortName yields the sort name of a collection, if argument is a collection.
func (ms *ModelState) CollectionSortName(ref KReference) (string, bool) {
	refType, _, sort, _, _, _ := parseKrefCollection(ref)
	if refType == mapRef ||
		refType == listRef ||
		refType == setRef ||
		refType == arrayRef {
		return Sort(sort).Name(), true
	}
	return "", false
}

// IsList returns true if reference points to a list with given sort
func (ms *ModelState) IsList(ref KReference, expectedSort Sort) bool {
	refType, _, sort, _, _, _ := parseKrefCollection(ref)
	return refType == listRef && sort == uint64(expectedSort)
}

// IsArray returns true if reference points to an array with given sort
func (ms *ModelState) IsArray(ref KReference, expectedSort Sort) bool {
	refType, _, sort, _, _, _ := parseKrefCollection(ref)
	return refType == arrayRef && sort == uint64(expectedSort)
}

// GetListObject yields the cast object for a List reference, if possible.
func (ms *ModelState) GetListObject(ref KReference) (*List, bool) {
	refType, dataRef, _, _, index, _ := parseKrefCollection(ref)
	if refType != listRef {
		return nil, false
	}
	obj := ms.getData(dataRef).getReferencedObject(index)
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
	tailRef := ms.NewList(castObj.Sort, castObj.Label, castObj.Data[1:])
	return true, castObj.Data[0], tailRef
}

// GetSetObject yields the cast object for a Set reference, if possible.
// func (ms *ModelState) GetSetObject(ref KReference) (*Set, bool) {
// 	refType, dataRef, _, _, index, _ := parseKrefCollection(ref)
// 	if refType != setRef {
// 		return nil, false
// 	}
// 	obj := ms.getData(dataRef).getReferencedObject(index)
// 	castObj, typeOk := obj.(*Set)
// 	if !typeOk {
// 		panic("wrong object type for reference")
// 	}
// 	return castObj, true
// }

// GetArrayObject yields the cast object for an Array reference, if possible.
func (ms *ModelState) GetArrayObject(ref KReference) (*Array, bool) {
	refType, dataRef, _, _, index, _ := parseKrefCollection(ref)
	if refType != arrayRef {
		return nil, false
	}
	obj := ms.getData(dataRef).getReferencedObject(index)
	castObj, typeOk := obj.(*Array)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

func (md *ModelData) addCollectionObject(sort Sort, label KLabel, obj KObject) KReference {
	newIndex := len(md.allObjects)
	md.allObjects = append(md.allObjects, obj)
	return createKrefCollection(obj.referenceType(), md.selfRef, uint64(sort), uint64(label), uint64(newIndex), 0)
}

func (ms *ModelState) addCollectionObject(sort Sort, label KLabel, obj KObject) KReference {
	return ms.mainData.addCollectionObject(sort, label, obj)
}

// NewList creates a new object and returns the reference.
func (ms *ModelState) NewList(sort Sort, label KLabel, value []KReference) KReference {
	return ms.addCollectionObject(sort, label, &List{Sort: sort, Label: label, Data: value})
}

// NewArray creates a new object and returns the reference.
func (ms *ModelState) NewArray(sort Sort, value *DynamicArray) KReference {
	return ms.addCollectionObject(sort, KLabel(0), &Array{Sort: sort, Data: value})
}
