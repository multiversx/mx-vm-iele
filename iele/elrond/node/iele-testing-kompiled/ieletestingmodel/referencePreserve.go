// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

package ieletestingmodel

// Preserve prevents argument and any objects contained by it from being recycled ever again.
func (ms *ModelState) Preserve(ref KReference) {
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
		obj, _ := ms.getBigIntObject(ref)
		obj.reuseStatus = preserved
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		for _, child := range ks {
			ms.Preserve(child)
		}
	case kapplyRef:
		for _, child := range ms.kapplyArgSlice(ref) {
			ms.Preserve(child)
		}
	default:
		// object types
		obj := ms.getData(dataRef).getReferencedObject(value)
		obj.preserve(ms)
	}
}

func (k *InjectedKLabel) preserve(ms *ModelState) {
}

func (k *KVariable) preserve(ms *ModelState) {
}

// func (k *Map) preserve(ms *ModelState) {
// 	for _, v := range k.Data {
// 		ms.Preserve(v)
// 	}
// }

func (k *List) preserve(ms *ModelState) {
	for _, item := range k.Data {
		ms.Preserve(item)
	}
}

func (k *Array) preserve(ms *ModelState) {
	for i := 0; i < len(k.Data.data); i++ {
		if k.Data.data[i] != NullReference {
			ms.Preserve(k.Data.data[i])
		}
	}
}

func (k *MInt) preserve(ms *ModelState) {
}

func (k *Float) preserve(ms *ModelState) {
}

func (k *StringBuffer) preserve(ms *ModelState) {
}
