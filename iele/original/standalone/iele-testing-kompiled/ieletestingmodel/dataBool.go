// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:19:50.499

package ieletestingmodel

// BoolTrue represents a boolean value with value true
var BoolTrue = createKrefBasic(boolRef, noDataRef, 1)

// BoolFalse represents a boolean value with value false
var BoolFalse = createKrefBasic(boolRef, noDataRef, 0)

// CastToBool converts K Bool to Go bool, if possible.
func CastToBool(ref KReference) (bool, bool) {
	refType, _, value := parseKrefBasic(ref)
	if refType != boolRef {
		return false, false
	}
	return value == 1, true
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
	refType, _, _ := parseKrefBasic(ref)
	return refType == boolRef
}

// IsTrue checks if argument is identical to the K Bool with the value true
func IsTrue(ref KReference) bool {
	refType, _, value := parseKrefBasic(ref)
	if refType != boolRef {
		return false
	}
	return value == 1
}
