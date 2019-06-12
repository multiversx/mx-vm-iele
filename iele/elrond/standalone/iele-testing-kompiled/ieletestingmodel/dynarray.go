// File provided by the K Framework Go backend. Timestamp: 2019-06-12 17:39:27.917

package ieletestingmodel

import (
	"errors"
)

// DynamicArray ... an array that resizes automatically
type DynamicArray struct {
	MaxSize uint64
	data    []K
	Default K
	ms      *ModelState
}

// ErrIndexOutOfBounds ... returned when the index exceeds DynamicArray max size
var ErrIndexOutOfBounds = errors.New("DynamicArray index out of bounds")

// MakeDynamicArray ... create new DynamicArray instance
func (ms *ModelState) MakeDynamicArray(maxSize uint64, defaultVal K) *DynamicArray {
	initialSize := maxSize
	if initialSize > 10 {
		initialSize = 10
	}
	data := make([]K, initialSize)
	return &DynamicArray{
		MaxSize: maxSize,
		data:    data,
		Default: defaultVal,
		ms:      ms,
	}
}

// Get ... get element at index
func (da *DynamicArray) Get(index uint64) (K, error) {
	if index >= da.MaxSize {
		return NoResult, ErrIndexOutOfBounds
	}
	if index >= uint64(len(da.data)) {
		return da.Default, nil
	}
	val := da.data[index]
	if val == nil {
		return da.Default, nil
	}
	return val, nil
}

// UpgradeSize ... increases the size of the underlying slice if necessary
func (da *DynamicArray) UpgradeSize(newSize uint64) {
	if newSize > da.MaxSize {
		newSize = da.MaxSize
	}
	currentLen := uint64(len(da.data))
	if newSize > currentLen {
		sizeInc := newSize - currentLen
		da.data = append(da.data, make([]K, sizeInc)...)
	}
}

// Set ... get element at index
func (da *DynamicArray) Set(index uint64, value K) error {
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
		da.data[index] = nil
	} else {
		da.data[index] = value
	}
	return nil
}

// Equals ... deep equals
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

// ToSlice ... convert DynamicArray to a slice of K items
func (da *DynamicArray) ToSlice() []K {
	slice := make([]K, len(da.data))
	for i := 0; i < len(da.data); i++ {
		if da.data[i] == nil {
			slice[i] = da.Default
		} else {
			slice[i] = da.data[i]
		}
	}
	return slice
}
