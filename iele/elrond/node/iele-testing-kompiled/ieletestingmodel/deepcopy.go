// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:38:10.632

package ieletestingmodel

import (
	"math/big"
)

// DeepCopy ... complete copy of the object
func (k *KApply) DeepCopy() K {
	listCopy := make([]K, len(k.List))
	for i, child := range k.List {
		listCopy[i] = child.DeepCopy()
	}
	return &KApply{Label: k.Label, List: listCopy}
}

// DeepCopy ... complete copy of the object
func (k *InjectedKLabel) DeepCopy() K {
	return &InjectedKLabel{Label: k.Label}
}

// DeepCopy ... complete copy of the object
func (k *KToken) DeepCopy() K {
	return &KToken{Sort: k.Sort, Value: k.Value}
}

// DeepCopy ... complete copy of the object
func (k *KVariable) DeepCopy() K {
	return &KVariable{Name: k.Name}
}

// DeepCopy ... complete copy of the object
func (k *Map) DeepCopy() K {
	mapCopy := make(map[KMapKey]K)
	for key, val := range k.Data {
		mapCopy[key] = val.DeepCopy()
	}
	return &Map{Data: mapCopy}
}

// DeepCopy ... complete copy of the object
func (k *List) DeepCopy() K {
	listCopy := make([]K, len(k.Data))
	for i, elem := range k.Data {
		listCopy[i] = elem.DeepCopy()
	}
	return &List{Sort: k.Sort, Label: k.Label, Data: listCopy}
}

// DeepCopy ... complete copy of the object
func (k *Set) DeepCopy() K {
	mapCopy := make(map[KMapKey]bool)
	for key := range k.Data {
		mapCopy[key] = true
	}
	return &Set{Data: mapCopy}
}

// DeepCopy ... complete copy of the object
func (k *Array) DeepCopy() K {
	return k // TODO: not implemented
}

// DeepCopy ... complete copy of the object
func (k *Int) DeepCopy() K {
	intCopy := new(big.Int)
	intCopy.Set(k.Value)
	return &Int{Value: intCopy}
}

// DeepCopy ... complete copy of the object
func (k *MInt) DeepCopy() K {
	return k // not implemented
}

// DeepCopy ... complete copy of the object
func (k *Float) DeepCopy() K {
	return k // not implemented
}

// DeepCopy ... complete copy of the object
func (k *String) DeepCopy() K {
	return &String{Value: k.Value}
}

// DeepCopy ... complete copy of the object
func (k *StringBuffer) DeepCopy() K {
	return k // no deep copy needed here
}

// DeepCopy ... complete copy of the object
func (k *Bytes) DeepCopy() K {
	bytesCopy := make([]byte, len(k.Value))
	copy(bytesCopy, k.Value)
	return &Bytes{Value: bytesCopy}
}

// DeepCopy ... complete copy of the object
func (k *Bool) DeepCopy() K {
	return &Bool{Value: k.Value}
}

// DeepCopy ... complete copy of the object
func (k *Bottom) DeepCopy() K {
	return &Bottom{}
}

// DeepCopy ... complete copy of the object
func (k *KSequence) DeepCopy() K {
	ksCopy := make([]K, len(k.Ks))
	for i, elem := range k.Ks {
		ksCopy[i] = elem.DeepCopy()
	}
	return &KSequence{Ks: ksCopy}
}
