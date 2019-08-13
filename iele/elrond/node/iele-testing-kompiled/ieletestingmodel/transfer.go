// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestingmodel

// transfer copies the structures underlying a reference from one data container to another.
// Will only transfer references that actually point to the "from" model.
// It is a destructive operation, i.e. source sub-tree is no longer usable after transfer.
func transfer(from, to *ModelData, ref KReference) KReference {
	refType, dataRef, value := parseKrefBasic(ref)
	if dataRef != from.selfRef {
		return ref
	}

	// collection types
	if isCollectionType(refType) {
		_, _, sortInt, labelInt, index := parseKrefCollection(ref)
		obj := from.getReferencedObject(index)
		obj.transferContents(from, to)
		return to.addCollectionObject(Sort(sortInt), KLabel(labelInt), obj)
	}

	switch refType {
	case nullRef:
		return ref
	case boolRef:
		return ref
	case bottomRef:
		return ref
	case emptyKseqRef:
		return ref
	case nonEmptyKseqRef:
		ks := from.ksequenceToSlice(ref)
		newKs := make([]KReference, len(ks))
		for i, child := range ks {
			newKs[i] = transfer(from, to, child)
		}
		return to.assembleKSequence(newKs...)
	case smallPositiveIntRef:
		return ref
	case smallNegativeIntRef:
		return ref
	case bigIntRef:
		obj, _ := from.getBigIntObject(ref)
		newRef, newObj := to.newBigIntObject()
		newObj.bigValue.Set(obj.bigValue)
		return newRef
	case kapplyRef:
		_, _, label, arity, index := parseKrefKApply(ref)
		if arity == 0 {
			return createKrefKApply(to.selfRef, label, 0, 0)
		}
		// 1. allocate
		toArgStartIndex := uint64(len(to.allKApplyArgs))
		for i := uint64(0); i < arity; i++ {
			to.allKApplyArgs = append(to.allKApplyArgs, 0)
		}
		// 2. transfer
		for i := uint64(0); i < arity; i++ {
			to.allKApplyArgs[toArgStartIndex+i] = transfer(from, to, from.allKApplyArgs[index+i])
		}
		return createKrefKApply(to.selfRef, label, arity, uint64(toArgStartIndex))
	case stringRef:
		_, _, startIndex, length := parseKrefBytes(ref)
		if length == 0 {
			return ref
		}
		return to.newBytes(stringRef, from.allBytes[startIndex:startIndex+length])
	case bytesRef:
		_, _, startIndex, length := parseKrefBytes(ref)
		if length == 0 {
			return ref
		}
		return to.newBytes(bytesRef, from.allBytes[startIndex:startIndex+length])
	case ktokenRef:
		_, _, sort, length, index := parseKrefKToken(ref)
		return to.newKToken(sort, from.allBytes[index:index+length])
	default:
		// object types
		obj := from.getReferencedObject(value)
		obj.transferContents(from, to)
		from.allObjects[value] = nil
		return to.addObject(obj)
	}
}

func (k *InjectedKLabel) transferContents(from, to *ModelData) {
}

func (k *KVariable) transferContents(from, to *ModelData) {
}

func (k *Map) transferContents(from, to *ModelData) {
	for key, val := range k.Data {
		k.Data[key] = transfer(from, to, val)
	}
}

func (k *List) transferContents(from, to *ModelData) {
	for i, elem := range k.Data {
		k.Data[i] = transfer(from, to, elem)
	}
}

func (k *Set) transferContents(from, to *ModelData) {
}

func (k *Array) transferContents(from, to *ModelData) {
	for i, elem := range k.Data.data {
		k.Data.data[i] = transfer(from, to, elem)
	}
	k.Data.Default = transfer(from, to, k.Data.Default)
}

func (k *MInt) transferContents(from, to *ModelData) {
}

func (k *Float) transferContents(from, to *ModelData) {
}

func (k *StringBuffer) transferContents(from, to *ModelData) {
}
