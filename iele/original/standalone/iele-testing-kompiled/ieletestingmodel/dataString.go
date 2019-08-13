// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:19:50.499

package ieletestingmodel

// StringEmpty is a reference to an empty string.
// There is no data, so is is irrelevant if we declare it constant or not.
var StringEmpty = createKrefBytes(stringRef, noDataRef, 0, 0)

// BytesEmpty is a reference to a Bytes item with no bytes (length 0).
// There is no data, so is is irrelevant if we declare it constant or not.
var BytesEmpty = createKrefBytes(bytesRef, noDataRef, 0, 0)

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
	refType, dataRef, startIndex, length := parseKrefBytes(ref)
	if refType != stringRef {
		return "", false
	}
	if length == 0 {
		return "", true
	}
	return string(ms.getData(dataRef).allBytes[startIndex : startIndex+length]), true
}

// GetBytes yields the cast object for a Bytes reference, if possible.
func (ms *ModelState) GetBytes(ref KReference) ([]byte, bool) {
	refType, dataRef, startIndex, length := parseKrefBytes(ref)
	if refType != bytesRef {
		return nil, false
	}
	if length == 0 {
		return nil, true
	}
	return ms.getData(dataRef).allBytes[startIndex : startIndex+length], true
}

func (md *ModelData) newBytes(refType kreferenceType, bytes []byte) KReference {
	length := len(bytes)
	if length == 0 {
		return createKrefBytes(refType, noDataRef, 0, 0)
	}
	startIndex := len(md.allBytes)
	md.allBytes = append(md.allBytes, bytes...)
	return createKrefBytes(refType, md.selfRef, uint64(startIndex), uint64(length))
}

// NewString creates a new K string object from a Go string
func (ms *ModelState) NewString(str string) KReference {
	return ms.mainData.newBytes(stringRef, []byte(str))
}

// NewStringConstant creates a new string constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewStringConstant(str string) KReference {
	return constantsData.newBytes(stringRef, []byte(str))
}

// NewBytes creates a new K bytes object from a Go byte array
func (ms *ModelState) NewBytes(bytes []byte) KReference {
	return ms.mainData.newBytes(bytesRef, bytes)
}

// Bytes2String converts a bytes reference to a string reference.
// The neat thing is, because we use the same underlying structure, no data needs to be copied.
func (ms *ModelState) Bytes2String(ref KReference) (KReference, bool) {
	refType, dataRef, startIndex, length := parseKrefBytes(ref)
	if refType != bytesRef {
		return NullReference, false
	}
	return createKrefBytes(stringRef, dataRef, startIndex, length), true
}

// String2Bytes converts a string reference to a bytes reference.
// The neat thing is, because we use the same underlying structure, no data needs to be copied.
func (ms *ModelState) String2Bytes(ref KReference) (KReference, bool) {
	refType, dataRef, startIndex, length := parseKrefBytes(ref)
	if refType != stringRef {
		return NullReference, false
	}
	return createKrefBytes(bytesRef, dataRef, startIndex, length), true
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
	refType, dataRef, startIndex, length := parseKrefBytes(ref)
	if refType != expectedRefType {
		return NullReference, false
	}
	if fromIndex > toIndex || fromIndex < 0 || toIndex < 0 || fromIndex > length {
		return NullReference, false
	}
	if toIndex > length {
		toIndex = length
	}
	return createKrefBytes(refType, dataRef, startIndex+fromIndex, toIndex-fromIndex), true
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
