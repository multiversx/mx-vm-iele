// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestingmodel

// RecycleUnused sends to the recycle bin all objects left without references.
// This goes recursively through the whole sub-tree.
func (ms *ModelState) RecycleUnused(ref KReference) {
	refType, dataRef, value := parseKrefBasic(ref)

	switch refType {
	case boolRef:
	case bottomRef:
	case emptyKseqRef:
	case smallPositiveIntRef:
	case smallNegativeIntRef:
	case stringRef:
	case bytesRef:
	case ktokenRef:
	case bigIntRef:
		obj, _ := ms.getData(dataRef).getBigIntObject(ref)
		if obj.reuseStatus == active && obj.referenceCount < 1 {
			// recycle
			obj.referenceCount = 0
			obj.reuseStatus = inRecycleBin
			ms.getData(dataRef).bigIntRecycleBin = append(ms.getData(dataRef).bigIntRecycleBin, ref)
		}
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		for _, child := range ks {
			ms.RecycleUnused(child)
		}
	case kapplyRef:
		for _, child := range ms.kapplyArgSlice(ref) {
			ms.RecycleUnused(child)
		}
	default:
		// object types
		obj := ms.getData(dataRef).getReferencedObject(value)
		obj.recycleUnused(ms)
	}
}

func (k *InjectedKLabel) recycleUnused(ms *ModelState) {
}

func (k *KVariable) recycleUnused(ms *ModelState) {
}

func (k *Map) recycleUnused(ms *ModelState) {
	for _, v := range k.Data {
		ms.RecycleUnused(v)
	}
}

func (k *List) recycleUnused(ms *ModelState) {
	for _, item := range k.Data {
		ms.RecycleUnused(item)
	}
}

func (k *Set) recycleUnused(ms *ModelState) {
}

func (k *Array) recycleUnused(ms *ModelState) {
	for i := 0; i < len(k.Data.data); i++ {
		if k.Data.data[i] != NullReference {
			ms.RecycleUnused(k.Data.data[i])
		}
	}
}

func (k *MInt) recycleUnused(ms *ModelState) {
}

func (k *Float) recycleUnused(ms *ModelState) {
}

func (k *StringBuffer) recycleUnused(ms *ModelState) {
}
