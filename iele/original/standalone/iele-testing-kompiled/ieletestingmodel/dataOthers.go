// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestingmodel

// InternedBottom is usually used as a dummy object
var InternedBottom = KReference{refType: bottomRef, value1: 0, value2: 0}

// IsBottom returns true if reference points to bottom
func IsBottom(ref KReference) bool {
	return ref.refType == bottomRef
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
	if ref.refType != floatRef {
		return nil, false
	}
	ms.getReferencedObject(ref)
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*Float)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// IsFloat returns true if reference points to a float
func IsFloat(ref KReference) bool {
	return ref.refType == floatRef
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
	return ref.refType == mintRef
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

