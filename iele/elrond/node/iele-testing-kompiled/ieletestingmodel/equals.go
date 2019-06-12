// File provided by the K Framework Go backend. Timestamp: 2019-06-12 11:57:09.485

package ieletestingmodel

import (
	"bytes"
)

// Equals ... Deep comparison
func (ms *ModelState) Equals(arg1 K, arg2 K) bool {
	return arg1.equals(ms, arg2)
}

func (k *KApply) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*KApply)
	if !typeOk {
		return false
	}
	if k.Label != other.Label {
		return false
	}
	if len(k.List) != len(other.List) {
		return false
	}
	for i := 0; i < len(k.List); i++ {
		if !k.List[i].equals(ms, other.List[i]) {
			return false
		}
	}
	return true
}

func (k *InjectedKLabel) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*InjectedKLabel)
	if !typeOk {
		return false
	}
	if k.Label != other.Label {
		return false
	}
	return true
}

func (k *KToken) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*KToken)
	if !typeOk {
		return false
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Value == other.Value
}

func (k *KVariable) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*KVariable)
	if !typeOk {
		return false
	}
	if k.Name != other.Name {
		return false
	}
	return true
}

func (k *Map) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Map)
	if !typeOk {
		return false
	}
	if len(k.Data) != len(other.Data) {
		return false
	}
	for key, val := range k.Data {
		otherVal, found := other.Data[key]
		if !found {
			return false
		}
		if !val.equals(ms, otherVal) {
			return false
		}
	}
	return true
}

func (k *List) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*List)
	if !typeOk {
		return false
	}
	if k.Sort != other.Sort {
		return false
	}
	if k.Label != other.Label {
		return false
	}
	if len(k.Data) != len(other.Data) {
		return false
	}
	for i := 0; i < len(k.Data); i++ {
		if !k.Data[i].equals(ms, other.Data[i]) {
			return false
		}
	}
	return true
}

func (k *Set) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Set)
	if !typeOk {
		return false
	}
	if len(k.Data) != len(other.Data) {
		return false
	}
	for key := range k.Data {
		_, found := other.Data[key]
		if !found {
			return false
		}
	}
	return true
}

func (k *Array) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Array)
	if !typeOk {
		return false
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Data.Equals(other.Data)
}

func (k *Int) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Int)
	if !typeOk {
		return false
	}
	return k.Value.Cmp(other.Value) == 0
}

func (k *MInt) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*MInt)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

func (k *Float) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Float)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

func (k *String) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*String)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Pointer comparison only for StringBuffer
func (k *StringBuffer) equals(ms *ModelState, arg K) bool {
	return k == arg
}

func (k *Bytes) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Bytes)
	if !typeOk {
		return false
	}
	return bytes.Equal(k.Value, other.Value)
}

func (k *Bool) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(*Bool)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

func (k *Bottom) equals(ms *ModelState, arg K) bool {
	_, typeOk := arg.(*Bottom)
	if !typeOk {
		return false
	}
	return true
}

func (k KSequence) equals(ms *ModelState, arg K) bool {
	other, typeOk := arg.(KSequence)
	if !typeOk {
		return false
	}

	length := ms.KSequenceLength(k)
	if length != ms.KSequenceLength(other) {
		return false
	}

	seq1 := ms.allKs[k.sequenceIndex]
	seq2 := ms.allKs[other.sequenceIndex]

	for i := 0; i < length; i++ {
		if !seq1[k.headIndex+i].equals(ms, seq2[other.headIndex+i]) {
			return false // element mismatch
		}
	}

	return true
}
