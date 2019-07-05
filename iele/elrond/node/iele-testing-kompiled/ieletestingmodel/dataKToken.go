// File provided by the K Framework Go backend. Timestamp: 2019-07-05 04:12:39.818

package ieletestingmodel

// KToken is a KObject representing a KToken item in K
type KToken struct {
	Value string
	Sort  Sort
}

func (*KToken) referenceType() kreferenceType {
	return ktokenRef
}

// GetKTokenObject yields the cast object for a KApply reference, if possible.
func (ms *ModelState) GetKTokenObject(ref KReference) (*KToken, bool) {
	if ref.refType != ktokenRef {
		return nil, false
	}
	ms.getReferencedObject(ref)
	obj := ms.getReferencedObject(ref)
	castObj, typeOk := obj.(*KToken)
	if !typeOk {
		panic("wrong object type for reference")
	}
	return castObj, true
}

// NewKToken creates a new object and returns the reference.
func (ms *ModelState) NewKToken(sort Sort, value string) KReference {
	return ms.addObject(&KToken{Sort: sort, Value: value})
}

// NewKTokenConstant creates a new KToken constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewKTokenConstant(sort Sort, value string) KReference {
	ref := constantsModel.NewKToken(sort, value)
	ref.constantObject = true
	return ref
}