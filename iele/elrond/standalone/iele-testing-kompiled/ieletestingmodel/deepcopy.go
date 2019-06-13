// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:44:33.979

package ieletestingmodel

import (
	"math/big"
)


// DeepCopy ... complete copy of the object
func (ms *ModelState) DeepCopy(k K) K {
	return k.deepCopy(ms)
}

func (k *KApply) deepCopy(ms *ModelState) K {
	listCopy := make([]K, len(k.List))
	for i, child := range k.List {
		listCopy[i] = child.deepCopy(ms)
	}
	return &KApply{Label: k.Label, List: listCopy}
}

func (k *InjectedKLabel) deepCopy(ms *ModelState) K {
	return &InjectedKLabel{Label: k.Label}
}

func (k *KToken) deepCopy(ms *ModelState) K {
	return &KToken{Sort: k.Sort, Value: k.Value}
}

func (k *KVariable) deepCopy(ms *ModelState) K {
	return &KVariable{Name: k.Name}
}

func (k *Map) deepCopy(ms *ModelState) K {
	mapCopy := make(map[KMapKey]K)
	for key, val := range k.Data {
		mapCopy[key] = val.deepCopy(ms)
	}
	return &Map{Data: mapCopy}
}

func (k *List) deepCopy(ms *ModelState) K {
	listCopy := make([]K, len(k.Data))
	for i, elem := range k.Data {
		listCopy[i] = elem.deepCopy(ms)
	}
	return &List{Sort: k.Sort, Label: k.Label, Data: listCopy}
}

func (k *Set) deepCopy(ms *ModelState) K {
	mapCopy := make(map[KMapKey]bool)
	for key := range k.Data {
		mapCopy[key] = true
	}
	return &Set{Data: mapCopy}
}

func (k *Array) deepCopy(ms *ModelState) K {
	return k // TODO: not implemented
}

func (k *Int) deepCopy(ms *ModelState) K {
	intCopy := new(big.Int)
	intCopy.Set(k.Value)
	return &Int{Value: intCopy}
}

func (k *MInt) deepCopy(ms *ModelState) K {
	return k // not implemented
}

func (k *Float) deepCopy(ms *ModelState) K {
	return k // not implemented
}

func (k *String) deepCopy(ms *ModelState) K {
	return &String{Value: k.Value}
}

func (k *StringBuffer) deepCopy(ms *ModelState) K {
	return k // no deep copy needed here
}

func (k *Bytes) deepCopy(ms *ModelState) K {
	bytesCopy := make([]byte, len(k.Value))
	copy(bytesCopy, k.Value)
	return &Bytes{Value: bytesCopy}
}

func (k *Bool) deepCopy(ms *ModelState) K {
	return &Bool{Value: k.Value}
}

func (k *Bottom) deepCopy(ms *ModelState) K {
	return &Bottom{}
}

func (k KSequence) deepCopy(ms *ModelState) K {
	ks := ms.KSequenceToSlice(k)
	ksCopy := make([]K, len(ks))
	for i, elem := range ks {
		ksCopy[i] = elem.deepCopy(ms)
	}
	return ms.NewKSequence(ksCopy)
}
