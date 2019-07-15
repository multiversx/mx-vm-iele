// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:11:08.386

package ieletestingmodel

// StringEmpty is a reference to an empty string.
// There is no data, so is is irrelevant if we declare it constant or not.
var StringEmpty = createKrefBytes(stringRef, false, 0, 0)

// BytesEmpty is a reference to a Bytes item with no bytes (length 0).
// There is no data, so is is irrelevant if we declare it constant or not.
var BytesEmpty = createKrefBytes(bytesRef, false, 0, 0)

// IsString returns true if reference points to a string
func IsString(ref KReference) bool {
	refType, _, _, _ := parseKrefBytes(ref)
	return refType == stringRef
}

// IsBytes returns true if reference points to a byte array
func IsBytes(ref KReference) bool {
	refType, _, _, _ := parseKrefBytes(ref)
	return refType == bytesRef
}

// GetString converts reference to a Go string, if possbile
func (ms *ModelState) GetString(ref KReference) (string, bool) {
	refType, constant, startIndex, length := parseKrefBytes(ref)
	if refType != stringRef {
		return "", false
	}
	if length == 0 {
		return "", true
	}
	if constant {
		ref = unsetConstantFlag(ref)
		return constantsModel.GetString(ref)
	}
	return string(ms.allBytes[startIndex : startIndex+length]), true
}

// GetBytes yields the cast object for a Bytes reference, if possible.
func (ms *ModelState) GetBytes(ref KReference) ([]byte, bool) {
	refType, constant, startIndex, length := parseKrefBytes(ref)
	if refType != bytesRef {
		return nil, false
	}
	if length == 0 {
		return nil, true
	}
	if constant {
		ref = unsetConstantFlag(ref)
		return constantsModel.GetBytes(ref)
	}
	return ms.allBytes[startIndex : startIndex+length], true
}

// NewString creates a new K string object from a Go string
func (ms *ModelState) NewString(str string) KReference {
	length := len(str)
	if length == 0 {
		return StringEmpty
	}
	startIndex := len(ms.allBytes)
	ms.allBytes = append(ms.allBytes, []byte(str)...)
	return createKrefBytes(stringRef, false, uint64(startIndex), uint64(length))
}

// NewStringConstant creates a new string constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewStringConstant(s string) KReference {
	ref := constantsModel.NewString(s)
	ref = setConstantFlag(ref)
	return ref
}

// NewBytes creates a new K string object from a Go string
func (ms *ModelState) NewBytes(value []byte) KReference {
	length := len(value)
	if length == 0 {
		return BytesEmpty
	}
	startIndex := len(ms.allBytes)
	ms.allBytes = append(ms.allBytes, value...)
	return createKrefBytes(bytesRef, false, uint64(startIndex), uint64(length))
}

// Bytes2String converts a bytes reference to a string reference.
// The neat thing is, because we use the same underlying structure, no data needs to be copied.
func (ms *ModelState) Bytes2String(ref KReference) (KReference, bool) {
	refType, constant, startIndex, length := parseKrefBytes(ref)
	if refType != bytesRef {
		return NullReference, false
	}
	return createKrefBytes(stringRef, constant, startIndex, length), true
}

// String2Bytes converts a string reference to a bytes reference.
// The neat thing is, because we use the same underlying structure, no data needs to be copied.
func (ms *ModelState) String2Bytes(ref KReference) (KReference, bool) {
	refType, constant, startIndex, length := parseKrefBytes(ref)
	if refType != stringRef {
		return NullReference, false
	}
	return createKrefBytes(bytesRef, constant, startIndex, length), true
}

// StringSub yields a reference to a substring of a given string.
// Given the structure of our data, no data needs to be copied or moved in this operation.
func StringSub(ref KReference, fromIndex uint64, toIndex uint64) (KReference, bool) {
	return subString(stringRef, ref, fromIndex, toIndex)
}

// BytesSub yields a reference to a sub-slice of a given byte slice.
// Given the structure of our data, no data needs to be copied or moved in this operation.
func BytesSub(ref KReference, fromIndex uint64, toIndex uint64) (KReference, bool) {
	return subString(bytesRef, ref, fromIndex, toIndex)
}

func subString(expectedRefType kreferenceType, ref KReference, fromIndex uint64, toIndex uint64) (KReference, bool) {
	refType, constant, startIndex, length := parseKrefBytes(ref)
	if refType != expectedRefType {
		return NullReference, false
	}
	if fromIndex > toIndex || fromIndex < 0 || toIndex < 0 || fromIndex > length {
		return NullReference, false
	}
	if toIndex > length {
		toIndex = length
	}
	return createKrefBytes(refType, constant, startIndex+fromIndex, toIndex-fromIndex), true
}

// StringLength yields the length of a string.
func StringLength(ref KReference) (uint64, bool) {
	refType, _, _, length := parseKrefBytes(ref)
	if refType != stringRef {
		return 0, false
	}
	return length, true
}

// BytesLength yields the length of a byte array.
func BytesLength(ref KReference) (uint64, bool) {
	refType, _, _, length := parseKrefBytes(ref)
	if refType != bytesRef {
		return 0, false
	}
	return length, true
}
