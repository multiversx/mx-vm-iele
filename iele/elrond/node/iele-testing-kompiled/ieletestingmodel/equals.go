// File provided by the K Framework Go backend. Timestamp: 2019-05-21 00:58:51.823

package ieletestingmodel

import (
	"bytes"
)

// Equals ... Deep comparison
func (k *KApply) Equals(arg K) bool {
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
		if !k.List[i].Equals(other.List[i]) {
			return false
		}
	}
	return true
}

// Equals ... Deep comparison
func (k *InjectedKLabel) Equals(arg K) bool {
	other, typeOk := arg.(*InjectedKLabel)
	if !typeOk {
		return false
	}
	if k.Label != other.Label {
		return false
	}
	return true
}

// Equals ... Deep comparison
func (k *KToken) Equals(arg K) bool {
	other, typeOk := arg.(*KToken)
	if !typeOk {
		return false
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Deep comparison
func (k *KVariable) Equals(arg K) bool {
	other, typeOk := arg.(*KVariable)
	if !typeOk {
		return false
	}
	if k.Name != other.Name {
		return false
	}
	return true
}

// Equals ... Deep comparison
func (k *Map) Equals(arg K) bool {
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
		if !val.Equals(otherVal) {
			return false
		}
	}
	return true
}

// Equals ... Deep comparison
func (k *List) Equals(arg K) bool {
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
		if !k.Data[i].Equals(other.Data[i]) {
			return false
		}
	}
	return true
}

// Equals ... Deep comparison
func (k *Set) Equals(arg K) bool {
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

// Equals ... Deep comparison
func (k *Array) Equals(arg K) bool {
	other, typeOk := arg.(*Array)
	if !typeOk {
		return false
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Data.Equals(other.Data)
}

// Equals ... Deep comparison
func (k *Int) Equals(arg K) bool {
	other, typeOk := arg.(*Int)
	if !typeOk {
		return false
	}
	return k.Value.Cmp(other.Value) == 0
}

// Equals ... Deep comparison
func (k *MInt) Equals(arg K) bool {
	other, typeOk := arg.(*MInt)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Deep comparison
func (k *Float) Equals(arg K) bool {
	other, typeOk := arg.(*Float)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Deep comparison
func (k *String) Equals(arg K) bool {
	other, typeOk := arg.(*String)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Pointer comparison only for StringBuffer
func (k *StringBuffer) Equals(arg K) bool {
	return k == arg
}

// Equals ... Deep comparison
func (k *Bytes) Equals(arg K) bool {
	other, typeOk := arg.(*Bytes)
	if !typeOk {
		return false
	}
	return bytes.Equal(k.Value, other.Value)
}

// Equals ... Deep comparison
func (k *Bool) Equals(arg K) bool {
	other, typeOk := arg.(*Bool)
	if !typeOk {
		return false
	}
	return k.Value == other.Value
}

// Equals ... Deep comparison
func (k *Bottom) Equals(arg K) bool {
	_, typeOk := arg.(*Bottom)
	if !typeOk {
		return false
	}
	return true
}

// Equals ... Deep comparison
func (k *KSequence) Equals(arg K) bool {
	other, typeOk := arg.(*KSequence)
	if !typeOk {
		return false
	}
	if len(k.Ks) != len(other.Ks) {
		return false
	}
	for i := 0; i < len(k.Ks); i++ {
		if !k.Ks[i].Equals(other.Ks[i]) {
			return false
		}
	}
	return true
}
