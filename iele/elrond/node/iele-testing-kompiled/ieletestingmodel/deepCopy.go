// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestingmodel

// DeepCopy yields a fresh copy of the K item given as argument.
// The copies end up in the main data container, even if the original objects don't reside there.
func (ms *ModelState) DeepCopy(ref KReference) KReference {
	refType, dataRef, value := parseKrefBasic(ref)

	// collection types
	if isCollectionType(refType) {
		_, _, _, _, index, _ := parseKrefCollection(ref)
		md := ms.getData(dataRef)
		obj := md.getReferencedObject(index)
		copiedObj := obj.deepCopy(ms)
		return ms.mainData.addObject(copiedObj)
	}

	switch refType {
	case boolRef:
		return ref
	case bottomRef:
		return ref
	case emptyKseqRef:
		return ref
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		newKs := make([]KReference, len(ks))
		for i, child := range ks {
			newKs[i] = ms.DeepCopy(child)
		}
		return ms.NewKSequence(newKs)
	case smallPositiveIntRef:
		return ref
	case smallNegativeIntRef:
		return ref
	case bigIntRef:
		md := ms.getData(dataRef)
		obj, _ := md.getBigIntObject(ref)
		newRef, newObj := md.newBigIntObjectNoRecycle()
		newObj.bigValue.Set(obj.bigValue)
		return newRef
	case kapplyRef:
		argSlice := ms.kapplyArgSlice(ref)
		argCopy := make([]KReference, len(argSlice))
		for i, child := range argSlice {
			argCopy[i] = ms.DeepCopy(child)
		}
		return ms.NewKApply(ms.KApplyLabel(ref), argCopy...)
	case stringRef:
		str, _ := ms.GetString(ref)
		return ms.NewString(str)
	case bytesRef:
		bytes, _ := ms.GetBytes(ref)
		return ms.NewBytes(bytes)
	case ktokenRef:
		ktoken, _ := ms.GetKTokenObject(ref)
		return ms.NewKToken(ktoken.Sort, ktoken.Value)
	case setRef:
		fallthrough
	case mapRef:
		_, _, sort, label, index, length := parseKrefCollection(ref)
		if length == 0 {
			return ref
		}
		md := ms.getData(dataRef)
		fromIndex := int(index)
		var toIndex int
		previousToIndex := -1
		for fromIndex != -1 {
			elem := md.allMapElements[fromIndex]
			toIndex = len(md.allMapElements)
			md.allMapElements = append(md.allMapElements, mapElementData{
				key:   ms.DeepCopy(elem.key),
				value: ms.DeepCopy(elem.value),
				next:  -1,
			})
			if previousToIndex != -1 {
				md.allMapElements[previousToIndex].next = toIndex
			}

			previousToIndex = toIndex
			fromIndex = elem.next
		}
		return createKrefCollection(refType, dataRef, sort, label, uint64(toIndex), length)
	default:
		// object types
		md := ms.getData(dataRef)
		obj := md.getReferencedObject(value)
		copiedObj := obj.deepCopy(ms)
		if copiedObj == obj {
			// if no new instance was created,
			// it means that the object does not need to be deep copied
			return ref
		}
		return ms.mainData.addObject(copiedObj)
	}
}

func (k *InjectedKLabel) deepCopy(ms *ModelState) KObject {
	return &InjectedKLabel{Label: k.Label}
}

func (k *KVariable) deepCopy(ms *ModelState) KObject {
	return &KVariable{Name: k.Name}
}

func (k *List) deepCopy(ms *ModelState) KObject {
	listCopy := make([]KReference, len(k.Data))
	for i, elem := range k.Data {
		listCopy[i] = ms.DeepCopy(elem)
	}
	return &List{
		Sort:  k.Sort,
		Label: k.Label,
		Data:  listCopy,
	}
}

func (k *Array) deepCopy(ms *ModelState) KObject {
	return k // TODO: not implemented
}

func (k *MInt) deepCopy(ms *ModelState) KObject {
	return k // not implemented
}

func (k *Float) deepCopy(ms *ModelState) KObject {
	return k // not implemented
}

func (k *StringBuffer) deepCopy(ms *ModelState) KObject {
	return k // no deep copy needed here
}
