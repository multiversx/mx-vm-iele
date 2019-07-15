// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:14:14.526

package ieletestingmodel

import (
	"strings"
)

// InternedBottom is usually used as a dummy object
var InternedBottom = createKrefBasic(bottomRef, true, 0)

// NoResult is the result when a function returns an error
var NoResult = InternedBottom

// IsBottom returns true if reference points to bottom
func IsBottom(ref KReference) bool {
	refType, _, _ := parseKrefBasic(ref)
	return refType == bottomRef
}

// Float is a KObject representing a float in K
type Float struct {
	Value float32
}

func (*Float) referenceType() kreferenceType {
	return floatRef
}

// GetFloatObject yields the cast object for a KApply reference, if possible.
func (ms *ModelState) GetFloatObject(ref KReference) (*Float, bool) {
	refType, _, index := parseKrefBasic(ref)
	if refType != floatRef {
		return nil, false
	}
	obj := ms.getReferencedObject(index, false)
	castObj, typeOk := obj.(*Float)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// IsFloat returns true if reference points to a float
func IsFloat(ref KReference) bool {
	refType, _, _ := parseKrefBasic(ref)
	return refType == floatRef
}

// MInt is a KObject representing a machine integer in K
type MInt struct {
	Value int32
}

func (*MInt) referenceType() kreferenceType {
	return mintRef
}

// IsMInt returns true if reference points to a string buffer
func IsMInt(ref KReference) bool {
	refType, _, _ := parseKrefBasic(ref)
	return refType == mintRef
}

// InjectedKLabel is a KObject representing an InjectedKLabel item in K
type InjectedKLabel struct {
	Label KLabel
}

func (*InjectedKLabel) referenceType() kreferenceType {
	return injectedKLabelRef
}

// NewInjectedKLabel creates a new InjectedKLabel object and returns the reference.
func (ms *ModelState) NewInjectedKLabel(label KLabel) KReference {
	return ms.addObject(&InjectedKLabel{Label: label})
}

// StringBuffer is a KObject that contains a string buffer
type StringBuffer struct {
	Value strings.Builder
}

func (*StringBuffer) referenceType() kreferenceType {
	return stringBufferRef
}

// IsStringBuffer returns true if reference points to a string buffer
func IsStringBuffer(ref KReference) bool {
	refType, _, _ := parseKrefBasic(ref)
	return refType == stringBufferRef
}

// GetStringBufferObject yields the cast object for a StringBuffer reference, if possible.
func (ms *ModelState) GetStringBufferObject(ref KReference) (*StringBuffer, bool) {
	refType, _, index := parseKrefBasic(ref)
	if refType != stringBufferRef {
		return nil, false
	}
	obj := ms.getReferencedObject(index, false)
	castObj, typeOk := obj.(*StringBuffer)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// NewStringBuffer creates a new object and returns the reference.
func (ms *ModelState) NewStringBuffer() KReference {
	return ms.addObject(&StringBuffer{Value: strings.Builder{}})
}
