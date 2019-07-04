// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

package ieletestingmodel

import (
	"strings"
)

// String is a KObject that contains a string
type String struct {
	Value string
}

func (*String) referenceType() kreferenceType {
	return stringRef
}

// Bytes is a KObject that contains a slice of bytes
type Bytes struct {
	Value []byte
}

func (*Bytes) referenceType() kreferenceType {
	return bytesRef
}

// StringBuffer is a KObject that contains a string buffer
type StringBuffer struct {
	Value strings.Builder
}

func (*StringBuffer) referenceType() kreferenceType {
	return stringBufferRef
}

// StringEmpty is a reference to an empty string
var StringEmpty = addConstantObject(&String{Value: ""})

// BytesEmpty is a reference to a Bytes item with no bytes (length 0)
var BytesEmpty = addConstantObject(&Bytes{Value: nil})

// IsString returns true if reference points to a string
func IsString(ref KReference) bool {
	return ref.refType == stringRef
}

// IsBytes returns true if reference points to a byte array
func IsBytes(ref KReference) bool {
	return ref.refType == bytesRef
}

// IsStringBuffer returns true if reference points to a string buffer
func IsStringBuffer(ref KReference) bool {
	return ref.refType == stringBufferRef
}

// GetStringObject yields the cast object for a String reference, if possible.
func (ms *ModelState) GetStringObject(ref KReference) (*String, bool) {
	if ref.refType == stringRef {
		obj := ms.getReferencedObject(ref)
		castObj, typeOk := obj.(*String)
		if !typeOk {
			panic("wrong object type for reference")
		}
		return castObj, true
	}

	return nil, false
}

// GetString converts reference to a Go string, if possbile
func (ms *ModelState) GetString(ref KReference) (string, bool) {
	castObj, typeOk := ms.GetStringObject(ref)
	if !typeOk {
		return "", false
	}
	return castObj.Value, true
}

// GetBytesObject yields the cast object for a Bytes reference, if possible.
func (ms *ModelState) GetBytesObject(ref KReference) (*Bytes, bool) {
	if ref.refType != bytesRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*Bytes)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// GetStringBufferObject yields the cast object for a StringBuffer reference, if possible.
func (ms *ModelState) GetStringBufferObject(ref KReference) (*StringBuffer, bool) {
	if ref.refType != stringBufferRef {
		return nil, false
	}
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*StringBuffer)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// NewString creates a new K string object from a Go string
func (ms *ModelState) NewString(str string) KReference {
	return ms.addObject(&String{Value: str})
}

// NewStringConstant creates a new string constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewStringConstant(s string) KReference {
	ref := constantsModel.NewString(s)
	ref.constantObject = true
	return ref
}

// NewBytes creates a new K string object from a Go string
func (ms *ModelState) NewBytes(value []byte) KReference {
	return ms.addObject(&Bytes{Value: value})
}

// NewStringBuffer creates a new object and returns the reference.
func (ms *ModelState) NewStringBuffer() KReference {
	return ms.addObject(&StringBuffer{Value: strings.Builder{}})
}

// IsEmpty returns true if Bytes is the empty byte slice
func (k *Bytes) IsEmpty() bool {
	return len(k.Value) == 0
}

// String yields a Go string representation of the K String
func (k *String) String() string {
	return k.Value
}

// IsEmpty returns true if it is the empty string
func (k *String) IsEmpty() bool {
	return len(k.Value) == 0
}
