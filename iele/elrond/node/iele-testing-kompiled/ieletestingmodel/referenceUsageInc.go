// File provided by the K Framework Go backend. Timestamp: 2019-07-04 01:26:11.488

package ieletestingmodel

// IncreaseUsage increments all reference counters in tree below given root.
// It goes recursively through the whole tree.
func (ms *ModelState) IncreaseUsage(ref KReference) {
	if ref.constantObject {
		return
	}

	switch ref.refType {
	case boolRef:
	case bottomRef:
	case emptyKseqRef:
	case smallPositiveIntRef:
	case smallNegativeIntRef:
	case bigIntRef:
		obj, _ := ms.getBigIntObject(ref)
		if obj.reuseStatus == active {
		    if obj.referenceCount < 1 {
		        obj.referenceCount = 1
		    } else {
		        obj.referenceCount++
		    }
		}
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		for _, child := range ks {
			ms.IncreaseUsage(child)
		}
	default:
		// object types
		obj := ms.getReferencedObject(ref)
		obj.increaseUsage(ms)
	}
}

func (k *KApply) increaseUsage(ms *ModelState) {
	for _, child := range k.List {
		ms.IncreaseUsage(child)
	}
}

func (k *InjectedKLabel) increaseUsage(ms *ModelState) {
}

func (k *KToken) increaseUsage(ms *ModelState) {
}

func (k *KVariable) increaseUsage(ms *ModelState) {
}

func (k *Map) increaseUsage(ms *ModelState) {
	for _, v := range k.Data {
		ms.IncreaseUsage(v)
	}
}

func (k *List) increaseUsage(ms *ModelState) {
	for _, item := range k.Data {
		ms.IncreaseUsage(item)
	}
}

func (k *Set) increaseUsage(ms *ModelState) {
}

func (k *Array) increaseUsage(ms *ModelState) {
	for i := 0; i < len(k.Data.data); i++ {
		if k.Data.data[i] != NullReference {
			ms.IncreaseUsage(k.Data.data[i])
		}
	}
}

func (k *MInt) increaseUsage(ms *ModelState) {
}

func (k *Float) increaseUsage(ms *ModelState) {
}

func (k *String) increaseUsage(ms *ModelState) {
}

func (k *StringBuffer) increaseUsage(ms *ModelState) {
}

func (k *Bytes) increaseUsage(ms *ModelState) {
}