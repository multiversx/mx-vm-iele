// File provided by the K Framework Go backend. Timestamp: 2019-07-05 04:12:39.818

package ieletestingmodel

import (
	"bytes"
)

// Equals performs a deep comparison, recursively.
func (ms *ModelState) Equals(ref1 KReference, ref2 KReference) bool {
	if ref1 == ref2 {
		// identical references means the same object
		return true
	}

	// int types
	intEquals, isInt := ms.IntEquals(ref1, ref2)
	if isInt {
		return intEquals
	}

	// for non-int types, refTypes should be equal
	if ref1.refType != ref2.refType {
		return false
	}

	switch ref1.refType {
	case boolRef:
		return false // if they were equal, ref1 == ref2 condition would already have returned true
	case bottomRef:
		panic("there shouldn't be different references of type bottomRef, only one")
	case emptyKseqRef:
		panic("there shouldn't be different references of type emptyKseqRef, only one")
	case nonEmptyKseqRef:
		return ms.ksequenceEquals(ref1, ref2)
	case kapplyRef:
		if ms.KApplyLabel(ref1) != ms.KApplyLabel(ref2) {
			return false
		}
		if ms.KApplyArity(ref1) != ms.KApplyArity(ref2) {
			return false
		}
		argSlice1 := ms.kapplyArgSlice(ref1)
		argSlice2 := ms.kapplyArgSlice(ref2)
		for i := 0; i < len(argSlice1); i++ {
			if !ms.Equals(argSlice1[i], argSlice2[i]) {
				return false
			}
		}
		return true
	default:
		// object types
		obj1 := ms.getReferencedObject(ref1)
		obj2 := ms.getReferencedObject(ref2)
		return obj1.equals(ms, obj2)
	}
}

func (k *InjectedKLabel) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*InjectedKLabel)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	if k.Label != other.Label {
		return false
	}
	return true
}

func (k *KToken) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*KToken)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Value == other.Value
}

func (k *KVariable) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*KVariable)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	if k.Name != other.Name {
		return false
	}
	return true
}

func (k *Map) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*Map)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	if len(k.Data) != len(other.Data) {
		return false
	}
	for key, val := range k.Data {
		otherVal, found := other.Data[key]
		if !found {
			return false
		}
		if !ms.Equals(val, otherVal) {
			return false
		}
	}
	return true
}

func (k *List) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*List)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
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
		if !ms.Equals(k.Data[i], other.Data[i]) {
			return false
		}
	}
	return true
}

func (k *Set) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*Set)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
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

func (k *Array) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*Array)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	if k.Sort != other.Sort {
		return false
	}
	return k.Data.Equals(other.Data)
}

func (k *MInt) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*MInt)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	return k.Value == other.Value
}

func (k *Float) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*Float)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	return k.Value == other.Value
}

func (k *String) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*String)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	return k.Value == other.Value
}

// Equals ... Pointer comparison only for StringBuffer
func (k *StringBuffer) equals(ms *ModelState, arg KObject) bool {
	return k == arg
}

func (k *Bytes) equals(ms *ModelState, arg KObject) bool {
	other, typeOk := arg.(*Bytes)
	if !typeOk {
		panic("equals between different types should have been handled during reference Equals")
	}
	return bytes.Equal(k.Value, other.Value)
}

func (ms *ModelState) ksequenceEquals(ref1 KReference, ref2 KReference) bool {
	s1 := ms.KSequenceToSlice(ref1)
	s2 := ms.KSequenceToSlice(ref2)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if !ms.Equals(s1[i], s2[i]) {
			return false
		}
	}

	return true
}
