// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:16:45.638

package ieletestingmodel

import (
	"errors"
)

// DynamicArray is an array that resizes automatically.
type DynamicArray struct {
	MaxSize uint64
	data    []KReference
	Default KReference
	ms      *ModelState
}

// ErrIndexOutOfBounds is returned when the index exceeds DynamicArray max size.
var ErrIndexOutOfBounds = errors.New("DynamicArray index out of bounds")

// MakeDynamicArray creates new DynamicArray instance.
func (ms *ModelState) MakeDynamicArray(maxSize uint64, defaultVal KReference) *DynamicArray {
	initialSize := maxSize
	if initialSize > 10 {
		initialSize = 10
	}
	data := make([]KReference, initialSize)
	return &DynamicArray{
		MaxSize: maxSize,
		data:    data,
		Default: defaultVal,
		ms:      ms,
	}
}

// Get retrieves element at index
func (da *DynamicArray) Get(index uint64) (KReference, error) {
	if index >= da.MaxSize {
		return NoResult, ErrIndexOutOfBounds
	}
	if index >= uint64(len(da.data)) {
		return da.Default, nil
	}
	val := da.data[index]
	if val == NullReference {
		return da.Default, nil
	}
	return val, nil
}

// UpgradeSize increases the size of the underlying slice if necessary
func (da *DynamicArray) UpgradeSize(newSize uint64) {
	if newSize > da.MaxSize {
		newSize = da.MaxSize
	}
	currentLen := uint64(len(da.data))
	if newSize > currentLen {
		sizeInc := newSize - currentLen
		da.data = append(da.data, make([]KReference, sizeInc)...)
	}
}

// Set updates a position in the array with a new value. It extends the array if necessary.
func (da *DynamicArray) Set(index uint64, value KReference) error {
	if index >= da.MaxSize {
		return ErrIndexOutOfBounds
	}
	currentLen := uint64(len(da.data))
	if index >= currentLen {
		if !da.ms.Equals(value, da.Default) {
			da.UpgradeSize(index + 1) // extend if necessary
			da.data[index] = value
		}
		return nil
	}

	if da.ms.Equals(value, da.Default) {
		da.data[index] = NullReference
	} else {
		da.data[index] = value
	}
	return nil
}

// Equals is a deep comparison.
func (da *DynamicArray) Equals(other *DynamicArray) bool {
	if da == other {
		return true
	}
	if da.MaxSize != other.MaxSize {
		return false
	}
	if !da.ms.Equals(da.Default, other.Default) {
		return false
	}
	maxLen := len(da.data)
	if len(other.data) > maxLen {
		maxLen = len(other.data)
	}
	for i := uint64(0); i < uint64(maxLen); i++ {
		val1, _ := da.Get(i)
		val2, _ := other.Get(i)
		if !da.ms.Equals(val1, val2) {
			return false
		}
	}
	return true
}

// ToSlice converts the DynamicArray to a slice of K references
func (da *DynamicArray) ToSlice() []KReference {
	slice := make([]KReference, len(da.data))
	for i := 0; i < len(da.data); i++ {
		if da.data[i] == NullReference {
			slice[i] = da.Default
		} else {
			slice[i] = da.data[i]
		}
	}
	return slice
}
