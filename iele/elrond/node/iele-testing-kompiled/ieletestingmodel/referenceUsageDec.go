// File provided by the K Framework Go backend. Timestamp: 2019-08-24 18:56:17.501

package ieletestingmodel

// DecreaseUsage decrements all reference counters in tree below given root
// and sends to the recycle bin all objects left without references.
// This goes recursively through the whole sub-tree.
func (ms *ModelState) DecreaseUsage(ref KReference) {
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
		if obj.reuseStatus == active {
			obj.referenceCount--
		}
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		for _, child := range ks {
			ms.DecreaseUsage(child)
		}
	case kapplyRef:
		for _, child := range ms.kapplyArgSlice(ref) {
			ms.DecreaseUsage(child)
		}
	default:
		// object types
		obj := ms.getData(dataRef).getReferencedObject(value)
		obj.decreaseUsage(ms)
	}
}

func (k *InjectedKLabel) decreaseUsage(ms *ModelState) {
}

func (k *KVariable) decreaseUsage(ms *ModelState) {
}

// func (k *Map) decreaseUsage(ms *ModelState) {
// 	for _, v := range k.Data {
// 		ms.DecreaseUsage(v)
// 	}
// }

func (k *List) decreaseUsage(ms *ModelState) {
	for _, item := range k.Data {
		ms.DecreaseUsage(item)
	}
}

func (k *Set) decreaseUsage(ms *ModelState) {
}

func (k *Array) decreaseUsage(ms *ModelState) {
	for i := 0; i < len(k.Data.data); i++ {
		if k.Data.data[i] != NullReference {
			ms.DecreaseUsage(k.Data.data[i])
		}
	}
}

func (k *MInt) decreaseUsage(ms *ModelState) {
}

func (k *Float) decreaseUsage(ms *ModelState) {
}

func (k *StringBuffer) decreaseUsage(ms *ModelState) {
}
