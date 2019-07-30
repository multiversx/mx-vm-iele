// File provided by the K Framework Go backend. Timestamp: 2019-07-30 16:33:19.058

package ieletestingmodel

// KApply is a KObject representing a KApply item in K.
// Not used internally.
type KApply struct {
	Label KLabel
	List  []KReference
}

// KApplyMatch returns true if reference is a KApply with correct label and arity
func KApplyMatch(ref KReference, expectedLabel KLabel, expectedArity uint32) bool {
	isKApply, _, label, arity, _ := parseKrefKApply(ref)
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
	isKApply, _, label, _, _ := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyLabel called for reference to item other than KApply")
	}
	return KLabel(label)
}

// KApplyArity returns the arity of a KApply item (nr. of arguments)
func (ms *ModelState) KApplyArity(ref KReference) int {
	isKApply, _, _, arity, _ := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyArity called for reference to item other than KApply")
	}
	return int(arity)
}

// KApplyArg returns the nth argument in a KApply
func (ms *ModelState) KApplyArg(ref KReference, argIndex uint64) KReference {
	isKApply, dataRef, _, arity, index := parseKrefKApply(ref)
	if !isKApply {
		panic("KApplyArg called for reference to item other than KApply")
	}
	if argIndex > arity {
		panic("KApplyArg called for arg index larger than KApply arity")
	}
	return ms.getData(dataRef).allKApplyArgs[index+argIndex]
}

func (ms *ModelState) kapplyArgSlice(ref KReference) []KReference {
	isKApply, dataRef, _, arity, index := parseKrefKApply(ref)
	if !isKApply {
		panic("kapplyArgSlice called for reference to item other than KApply")
	}
	if arity == 0 {
		return nil
	}
	return ms.getData(dataRef).allKApplyArgs[index : index+arity]
}

// GetKApplyObject yields the cast object for a KApply reference, if possible.
func (ms *ModelState) GetKApplyObject(ref KReference) (*KApply, bool) {
	isKApply, _, label, _, _ := parseKrefKApply(ref)
	if !isKApply {
		return nil, false
	}
	return &KApply{
		Label: KLabel(label),
		List:  ms.kapplyArgSlice(ref),
	}, true
}

func (md *ModelData) newKApply(labelInt uint64, arguments ...KReference) KReference {
	argStartIndex := 0
	if len(arguments) > 0 {
		argStartIndex = len(md.allKApplyArgs)
		md.allKApplyArgs = append(md.allKApplyArgs, arguments...)
	}

	return createKrefKApply(md.selfRef, labelInt, uint64(len(arguments)), uint64(argStartIndex))
}

// NewKApply creates a new KApply object and returns the reference.
func (ms *ModelState) NewKApply(label KLabel, arguments ...KReference) KReference {
	return ms.mainData.newKApply(uint64(label), arguments...)
}

// NewKApplyConstant creates a new integer constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewKApplyConstant(label KLabel, arguments ...KReference) KReference {
	return constantsData.newKApply(uint64(label), arguments...)
}
