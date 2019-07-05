// File provided by the K Framework Go backend. Timestamp: 2019-07-05 04:12:39.818

package ieletestingmodel

// DeepCopy yields a fresh copy of the K item given as argument.
func (ms *ModelState) DeepCopy(ref KReference) KReference {
	switch ref.refType {
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
		obj, _ := ms.getBigIntObject(ref)
		newRef, newObj := ms.newBigIntObjectNoRecycle()
		newObj.bigValue.Set(obj.bigValue)
		return newRef
	case kapplyRef:
		argSlice := ms.kapplyArgSlice(ref)
		argCopy := make([]KReference, len(argSlice))
		for i, child := range argSlice {
			argCopy[i] = ms.DeepCopy(child)
		}
		return ms.NewKApply(ms.KApplyLabel(ref), argCopy...)
	default:
		// object types
		obj := ms.getReferencedObject(ref)
		copiedObj := obj.deepCopy(ms)
		if copiedObj == obj {
			// if no new instance was created,
			// it means that the object does not need to be deep copied
			return ref
		}
		return ms.addObject(obj)
	}
}

func (k *InjectedKLabel) deepCopy(ms *ModelState) KObject {
	return &InjectedKLabel{Label: k.Label}
}

func (k *KToken) deepCopy(ms *ModelState) KObject {
	return &KToken{Sort: k.Sort, Value: k.Value}
}

func (k *KVariable) deepCopy(ms *ModelState) KObject {
	return &KVariable{Name: k.Name}
}

func (k *Map) deepCopy(ms *ModelState) KObject {
	mapCopy := make(map[KMapKey]KReference)
	for key, val := range k.Data {
		mapCopy[key] = ms.DeepCopy(val)
	}
	return &Map{Data: mapCopy}
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

func (k *Set) deepCopy(ms *ModelState) KObject {
	mapCopy := make(map[KMapKey]bool)
	for key := range k.Data {
		mapCopy[key] = true
	}
	return &Set{Data: mapCopy}
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

func (k *String) deepCopy(ms *ModelState) KObject {
	return &String{Value: k.Value}
}

func (k *StringBuffer) deepCopy(ms *ModelState) KObject {
	return k // no deep copy needed here
}

func (k *Bytes) deepCopy(ms *ModelState) KObject {
	bytesCopy := make([]byte, len(k.Value))
	copy(bytesCopy, k.Value)
	return &Bytes{Value: bytesCopy}
}
