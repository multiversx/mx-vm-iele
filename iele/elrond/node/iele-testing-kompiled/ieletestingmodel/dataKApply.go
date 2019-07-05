// File provided by the K Framework Go backend. Timestamp: 2019-07-05 04:12:39.818

package ieletestingmodel

// KApply is a KObject representing a KApply item in K.
// Not used internally.
type KApply struct {
	Label KLabel
	List  []KReference
}

// CastKApply returns true if argument is a KApply item.
// Also returns argument, for convenience.
func (ms *ModelState) CastKApply(ref KReference) (KReference, bool) {
	if ref.refType != kapplyRef {
		return NullReference, false
	}
	return ref, true
}

// KApplyLabel returns the label of a KApply item.
func (ms *ModelState) KApplyLabel(ref KReference) KLabel {
	if ref.refType != kapplyRef {
		panic("KApplyLabel called for reference to item other than KApply")
	}
	return KLabel(ref.value1)
}

// KApplyArity returns the arity of a KApply item (nr. of arguments)
func (ms *ModelState) KApplyArity(ref KReference) int {
	if ref.refType != kapplyRef {
		panic("KApplyArity called for reference to item other than KApply")
	}
	return int(ref.value2)
}

// KApplyArg returns the nth argument in a KApply
func (ms *ModelState) KApplyArg(ref KReference, argIndex uint32) KReference {
	if ref.refType != kapplyRef {
		panic("KApplyArg called for reference to item other than KApply")
	}
	if argIndex > ref.value2 {
		panic("KApplyArg called for arg index larger than KApply arity")
	}
	return ms.allKApplyArgs[ref.value3+argIndex]
}

func (ms *ModelState) kapplyArgSlice(ref KReference) []KReference {
	if ref.refType != kapplyRef {
		panic("kapplyArgSlice called for reference to item other than KApply")
	}
	arity := ref.value2
	if arity == 0 {
		return nil
	}
	return ms.allKApplyArgs[ref.value3 : ref.value3+arity]
}

// GetKApplyObject yields the cast object for a KApply reference, if possible.
func (ms *ModelState) GetKApplyObject(ref KReference) (*KApply, bool) {
	if ref.refType != kapplyRef {
		return nil, false
	}
	return &KApply{
		Label: KLabel(ref.value1),
		List:  ms.kapplyArgSlice(ref),
	}, true
}

func newKApplyReference(label KLabel, arity int, argStartIndex uint32) KReference {
	return KReference{
		refType:        kapplyRef,
		constantObject: false,
		value1:         uint32(label),
		value2:         uint32(arity),
		value3:         uint32(argStartIndex),
	}
}

// NewKApply creates a new object and returns the reference.
func (ms *ModelState) NewKApply(label KLabel, arguments ...KReference) KReference {
	argStartIndex := 0
	if len(arguments) > 0 {
		argStartIndex = len(ms.allKApplyArgs)
		ms.allKApplyArgs = append(ms.allKApplyArgs, arguments...)
	}

	return newKApplyReference(label, len(arguments), uint32(argStartIndex))
}

// NewKApplyConstant creates a new integer constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewKApplyConstant(label KLabel, arguments ...KReference) KReference {
	ref := constantsModel.NewKApply(label, arguments...)
	ref.constantObject = true
	return ref
}
