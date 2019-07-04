// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestingmodel

// BoolTrue represents a boolean value with value true
var BoolTrue = KReference{refType: boolRef, value1: 1, value2: 0}

// BoolFalse represents a boolean value with value false
var BoolFalse = KReference{refType: boolRef, value1: 0, value2: 0}

// CastToBool converts K Bool to Go bool, if possible.
func CastToBool(ref KReference) (bool, bool) {
	if ref.refType != boolRef {
		return false, false
	}
	return ref.value1 == 1, true
}

// ToKBool converts Go bool to K Bool.
func ToKBool(b bool) KReference {
	if b {
		return BoolTrue
	}
	return BoolFalse
}

// IsBool checks if the argument is a bool reference
func IsBool(ref KReference) bool {
	return ref.refType == boolRef
}

// IsTrue checks if argument is identical to the K Bool with the value true
func IsTrue(ref KReference) bool {
	if ref.refType != boolRef {
		return false
	}
	return ref.value1 == 1
}
