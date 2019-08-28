// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestingmodel

// CollectionsToK converts all collections to standard K items, like KApply, KToken, etc.
func (ms *ModelState) CollectionsToK(ref KReference) KReference {
	refType := getRefType(ref)
	if refType == nonEmptyKseqRef {
		ks := ms.KSequenceToSlice(ref)
		newKs := make([]KReference, len(ks))
		for i, child := range ks {
			newKs[i] = ms.CollectionsToK(child)
		}
		return ms.NewKSequence(newKs)
	} else if refType == kapplyRef {
		argSlice := ms.kapplyArgSlice(ref)
		newArgs := make([]KReference, len(argSlice))
		for i, child := range argSlice {
			newArgs[i] = ms.CollectionsToK(child)
		}
		return ms.NewKApply(ms.KApplyLabel(ref), newArgs...)
	} else if refType == mapRef {
		_, _, _, lbl, _, _ := parseKrefCollection(ref)
		orderedKVPairs := ms.mapOrderedKeyValuePairs(ref)
		return ms.mapElementsToK(KLabel(lbl), orderedKVPairs)
	} else if refType == setRef {
		_, _, _, lbl, _, _ := parseKrefCollection(ref)
		orderedElems := ms.setOrderedElements(ref)
		return ms.setElementsToK(KLabel(lbl), orderedElems)
	} else if isCollectionType(refType) {
		_, dataRef, _, _, index, _ := parseKrefCollection(ref)
		obj := ms.getData(dataRef).getReferencedObject(index)
		return obj.collectionsToK(ms)
	}

	// no processing required for the others
	return ref
}

func (ms *ModelState) mapElementsToK(label KLabel, orderedKVPairs []MapKeyValuePair) KReference {
	elemLabel := ElementFor(label)
	var result KReference
	for i, pair := range orderedKVPairs {
		elemK := ms.NewKApply(elemLabel, pair.Key, ms.CollectionsToK(pair.Value))
		if i == 0 {
			result = elemK
		} else {
			newResult := ms.NewKApply(label, result, elemK)
			result = newResult
		}
	}

	return result
}

func (ms *ModelState) setElementsToK(label KLabel, orderedElems []KReference) KReference {
	// process
	elemLabel := ElementFor(label)
	var result KReference
	for i, key := range orderedElems {
		elemK := ms.NewKApply(elemLabel, key)
		if i == 0 {
			result = elemK
		} else {
			newResult := ms.NewKApply(label, result, elemK)
			result = newResult
		}
	}

	return result
}

func (k *List) collectionsToK(ms *ModelState) KReference {
	if len(k.Data) == 0 {
		return ms.NewKApply(UnitFor(k.Label))
	}

	elemLabel := ElementFor(k.Label)
	var result KReference
	for i, elem := range k.Data {
		elemK := ms.NewKApply(elemLabel, elem)
		if i == 0 {
			result = elemK
		} else {
			newResult := ms.NewKApply(k.Label, result, elemK)
			result = newResult
		}
	}

	return result
}

func (k *Array) collectionsToK(ms *ModelState) KReference {
	result := ms.NewKApply(UnitForArray(k.Sort),
		ms.NewString("uid"),
		ms.FromUint64(k.Data.MaxSize))

	slice := k.Data.ToSlice()
	for i, item := range slice {
		newResult := ms.NewKApply(ElementForArray(k.Sort),
			result,
			ms.FromInt(i),
			item)
		result = newResult
	}

	return result

}

func (k *InjectedKLabel) collectionsToK(ms *ModelState) KReference {
	panic("collectionsToK shouldn't be called for this object type")
}

func (k *KVariable) collectionsToK(ms *ModelState) KReference {
	panic("collectionsToK shouldn't be called for this object type")
}

func (k *MInt) collectionsToK(ms *ModelState) KReference {
	panic("collectionsToK shouldn't be called for this object type")
}

func (k *Float) collectionsToK(ms *ModelState) KReference {
	panic("collectionsToK shouldn't be called for this object type")
}

func (k *StringBuffer) collectionsToK(ms *ModelState) KReference {
	panic("collectionsToK shouldn't be called for this object type")
}
