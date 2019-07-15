// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:26:24.140

package ieletestingmodel

// KApply is a KObject representing a KApply item in K.
// Not used internally.
type KApply struct {
	Label KLabel
	List  []KReference
}

// KApplyMatch returns true if reference is a KApply with correct label and arity
func KApplyMatch(ref KReference, expectedLabel KLabel, expectedArity uint32) bool {
	isKApply, label, arity, _ := parseKrefKApply(ref)
	if !isKApply {
		return false
	}
	if label != uint64(expectedLabel) {
		return false
	}
	if arity != uint64(expectedArity) {
		return false
	}
	return true
}

// KApplyLabel returns the label of a KApply item.
func (ms *ModelState) KApplyLabel(ref KReference) KLabel {
	isKApply, label, _, _ := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyLabel called for reference to item other than KApply")
	}
	return KLabel(label)
}

// KApplyArity returns the arity of a KApply item (nr. of arguments)
func (ms *ModelState) KApplyArity(ref KReference) int {
	isKApply, _, arity, _ := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyArity called for reference to item other than KApply")
	}
	return int(arity)
}

// KApplyArg returns the nth argument in a KApply
func (ms *ModelState) KApplyArg(ref KReference, argIndex uint64) KReference {
	isKApply, _, arity, index := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyArg called for reference to item other than KApply")
	}
	if argIndex > arity {
		panic("KApplyArg called for arg index larger than KApply arity")
	}
	return ms.allKApplyArgs[index+argIndex]
}

func (ms *ModelState) kapplyArgSlice(ref KReference) []KReference {
	isKApply, _, arity, index := parseKrefKApply(ref)
	if !isKApply {
		panic("kapplyArgSlice called for reference to item other than KApply")
	}
	if arity == 0 {
		return nil
	}
	return ms.allKApplyArgs[index : index+arity]
}

// GetKApplyObject yields the cast object for a KApply reference, if possible.
func (ms *ModelState) GetKApplyObject(ref KReference) (*KApply, bool) {
	isKApply, label, _, _ := parseKrefKApply(ref)
	if !isKApply {
		return nil, false
	}
	return &KApply{
		Label: KLabel(label),
		List:  ms.kapplyArgSlice(ref),
	}, true
}

// NewKApply creates a new object and returns the reference.
func (ms *ModelState) NewKApply(label KLabel, arguments ...KReference) KReference {
	argStartIndex := 0
	if len(arguments) > 0 {
		argStartIndex = len(ms.allKApplyArgs)
		ms.allKApplyArgs = append(ms.allKApplyArgs, arguments...)
	}

	return createKrefKApply(uint64(label), uint64(len(arguments)), uint64(argStartIndex))
}

// NewKApplyConstant creates a new integer constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewKApplyConstant(label KLabel, arguments ...KReference) KReference {
	ref := constantsModel.NewKApply(label, arguments...)
	ref = setConstantFlag(ref)
	return ref
}
