// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:14:14.526

package ieletestinginterpreter

// The file is duplicated in the interpreter and in the model intentionally,
// for performance reasons.
// This version is for: ieletestinginterpreter

// kreferenceType identifies the type of K item referenced by a KReference
type kreferenceType uint64

const (
	nullRef kreferenceType = iota
	emptyKseqRef
	nonEmptyKseqRef
	kapplyRef
	injectedKLabelRef
	ktokenRef
	kvariableRef
	mapRef
	setRef
	listRef
	arrayRef
	smallPositiveIntRef
	smallNegativeIntRef
	bigIntRef
	mintRef
	floatRef
	stringRef
	stringBufferRef
	bytesRef
	boolRef
	bottomRef
)

func isCollectionType(refType kreferenceType) bool {
	return refType == mapRef ||
		refType == listRef ||
		refType == setRef ||
		refType == arrayRef
}

// KReference is a reference to a K item.
// For some types, like bool and small int, the entire state can be kept in the reference object.
// For the others, the reference contains enough data to find the object in the model state.
type KReference = uint64

// NullReference is the zero-value of KReference. It doesn't point to anything.
// It has type nullRef.
var NullReference = KReference(0)

// The basic, most general encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: is constant = 1, not constant = 0
// - the remaining 58 LS bits: type-specific data

const refTypeBits = 5                      // refType gets represented in this many bits
const refTypeMask = (1 << refTypeBits) - 1 // 5 bits of 1
const refTypeShift = 59                    // shift right this many bits to get the refType
const refBasicDataShift = 58
const refBasicDataMask = (1 << refBasicDataShift) - 1

func getRefType(ref KReference) kreferenceType {
	return kreferenceType(uint64(ref) >> refTypeShift)
}

func parseKrefBasic(ref KReference) (refType kreferenceType, constant bool, rest uint64) {
	refRaw := uint64(ref)
	rest = refRaw & refBasicDataMask
	refRaw >>= refBasicDataShift
	constant = refRaw&1 == 1
	refRaw >>= 1
	refType = kreferenceType(refRaw)
	return
}

func createKrefBasic(refType kreferenceType, constant bool, rest uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= 1
	if constant {
		refRaw |= 1
	}
	refRaw <<= refBasicDataShift
	refRaw |= rest
	return KReference(refRaw)
}

func setConstantFlag(ref KReference) KReference {
	refRaw := uint64(ref)
	refRaw |= (1 << refBasicDataShift)
	return KReference(refRaw)
}

func unsetConstantFlag(ref KReference) KReference {
	refRaw := uint64(ref)
	refRaw &^= (1 << refBasicDataShift) // bit clear operator
	return KReference(refRaw)
}

// big int reference structure (from MSB to LSB):
// - 5 bits: reference type
// - 1 bit: is constant = 1, not constant = 0
// - 26 bits: recycle count
// - 32 bits: index

const refBigIntRecycleCountBits = 26
const refBigIntRecycleCountMask = (1 << refBigIntRecycleCountBits) - 1
const refBigIntIndexBits = 32
const refBigIntIndexMask = (1 << refBigIntIndexBits) - 1

func createKrefBigInt(constant bool, recycleCount uint64, index uint64) KReference {
	ref := uint64(bigIntRef) << 1
	if constant {
		ref |= 1
	}
	ref <<= refBigIntRecycleCountBits
	ref |= recycleCount
	ref <<= refBigIntIndexBits
	ref |= index
	return KReference(ref)
}

func parseKrefBigInt(ref KReference) (isBigInt bool, constant bool, recycleCount uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refBigIntIndexMask
	refRaw >>= refBigIntIndexBits
	recycleCount = refRaw & refBigIntRecycleCountMask
	refRaw >>= refBigIntRecycleCountBits
	constant = refRaw&1 == 1
	refRaw >>= 1
	isBigInt = refRaw == uint64(bigIntRef)
	return
}

// The collection encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: ignored
// - 13 bits: Sort
// - 13 bits: Label
// - 32 bits: object index

const refCollectionSortShift = 13
const refCollectionSortMask = (1 << refCollectionSortShift) - 1
const refCollectionLabelShift = 13
const refCollectionLabelMask = (1 << refCollectionLabelShift) - 1
const refCollectionIndexShift = 32
const refCollectionIndexMask = (1 << refCollectionIndexShift) - 1

func createKrefCollection(refType kreferenceType, sortInt uint64, labelInt uint64, index uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= 1
	refRaw <<= refCollectionSortShift
	refRaw |= sortInt
	refRaw <<= refCollectionLabelShift
	refRaw |= labelInt
	refRaw <<= refCollectionIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefCollection(ref KReference) (refType kreferenceType, sortInt uint64, labelInt uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refCollectionIndexMask
	refRaw >>= refCollectionIndexShift
	labelInt = refRaw & refCollectionLabelMask
	refRaw >>= refCollectionLabelShift
	sortInt = refRaw & refCollectionSortMask
	refRaw >>= refCollectionSortShift
	refRaw >>= 1 // ignore constant flag
	refType = kreferenceType(refRaw)
	return
}

// The KApply encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: ignored
// - 13 bits: label
// - 13 bits: arity
// - 32 bits: arguments index

const refKApplyLabelShift = 13
const refKApplyLabelMask = (1 << refKApplyLabelShift) - 1
const refKApplyArityShift = 13
const refKApplyArityMask = (1 << refKApplyArityShift) - 1
const refKApplyIndexShift = 32
const refKApplyIndexMask = (1 << refKApplyIndexShift) - 1
const refKApplyTypeAsUint = uint64(kapplyRef)

func createKrefKApply(labelInt uint64, arity uint64, index uint64) KReference {
	refRaw := refKApplyTypeAsUint
	refRaw <<= 1
	refRaw <<= refKApplyLabelShift
	refRaw |= labelInt
	refRaw <<= refKApplyArityShift
	refRaw |= arity
	refRaw <<= refKApplyIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefKApply(ref KReference) (isKApply bool, labelInt uint64, arity uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refKApplyIndexMask
	refRaw >>= refKApplyIndexShift
	arity = refRaw & refKApplyArityMask
	refRaw >>= refKApplyArityShift
	labelInt = refRaw & refKApplyLabelMask
	refRaw >>= refKApplyLabelShift
	refRaw >>= 1 // ignore constant flag
	isKApply = refRaw == refKApplyTypeAsUint
	return
}

// MatchKApply returns true if reference is a KApply with correct label and arity.
// Function should be inlined, for performance reasons.
func MatchKApply(ref KReference, expectedLabel uint64, expectedArity uint64) bool {
	refRaw := uint64(ref)
	refRaw >>= refKApplyIndexShift // ignore index here
	arity := refRaw & refKApplyArityMask
	refRaw >>= refKApplyArityShift
	labelInt := refRaw & refKApplyLabelMask
	refRaw >>= refKApplyLabelShift
	refRaw >>= 1 // ignore constant flag
	return refRaw == refKApplyTypeAsUint &&
		labelInt == expectedLabel &&
		arity == expectedArity
}

// The K sequence encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: ignored
// - 26 bits: length
// - 32 bits: head (element) index

const refNonEmptyKseqLengthShift = 26
const refNonEmptyKseqLengthMask = (1 << refNonEmptyKseqLengthShift) - 1
const refNonEmptyKseqIndexShift = 32
const refNonEmptyKseqIndexMask = (1 << refNonEmptyKseqIndexShift) - 1
const refNonEmptyKseqTypeAsUint = uint64(nonEmptyKseqRef)
const refEmptyKseqTypeAsUint = uint64(emptyKseqRef)

func createKrefNonEmptyKseq(elemIndex uint64, length uint64) KReference {
	refRaw := refNonEmptyKseqTypeAsUint
	refRaw <<= 1
	refRaw <<= refNonEmptyKseqLengthShift
	refRaw |= length
	refRaw <<= refNonEmptyKseqIndexShift
	refRaw |= elemIndex
	return KReference(refRaw)
}

func parseKrefKseq(ref KReference) (refType kreferenceType, elemIndex uint64, length uint64) {
	refRaw := uint64(ref)
	elemIndex = refRaw & refNonEmptyKseqIndexMask
	refRaw >>= refNonEmptyKseqIndexShift
	length = refRaw & refNonEmptyKseqLengthMask
	refRaw >>= refNonEmptyKseqLengthShift
	refRaw >>= 1 // ignore constant flag
	refType = kreferenceType(refRaw)
	return
}

// MatchNonEmptyKSequence returns true if reference is a K sequence with at least this many items,
// OR another any item type other than empty K sequence.
// Function should be inlined, for performance reasons.
func MatchNonEmptyKSequence(ref KReference) bool {
	refType := uint64(ref) >> refTypeShift
	return refType != refEmptyKseqTypeAsUint
}

// MatchNonEmptyKSequenceMinLength returns true if reference is a K sequence with at least this many items.
// Argument minimumLength must be minimum 2.
// Function should be inlined, for performance reasons.
func MatchNonEmptyKSequenceMinLength(ref KReference, minimumLength uint64) bool {
	refRaw := uint64(ref)
	refRaw >>= refNonEmptyKseqIndexShift         // ignore element index
	length := refRaw & refNonEmptyKseqLengthMask // length for matching
	refRaw >>= refNonEmptyKseqLengthShift        //
	refRaw >>= 1                                 // ignore constant flag

	// refRaw is the reference type at this point
	return refRaw == refNonEmptyKseqTypeAsUint && length >= minimumLength
}

// The K token encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: is constant = 1, not constant = 0
// - 13 bits: sortInt
// - 13 bits: length
// - 32 bits: value string start index in allBytes

const refKTokenSortShift = 13
const refKTokenSortMask = (1 << refKTokenSortShift) - 1
const refKTokenLengthShift = 13
const refKTokenLengthMask = (1 << refKTokenLengthShift) - 1
const refKTokenIndexShift = 32
const refKTokenIndexMask = (1 << refKTokenIndexShift) - 1
const refKTokenTypeAsUint = uint64(ktokenRef)

func createKrefKToken(constant bool, sortInt uint64, length uint64, index uint64) KReference {
	refRaw := refKTokenTypeAsUint
	refRaw <<= 1
	if constant {
		refRaw |= 1
	}
	refRaw <<= refKTokenSortShift
	refRaw |= sortInt
	refRaw <<= refKTokenLengthShift
	refRaw |= length
	refRaw <<= refKTokenIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefKToken(ref KReference) (isKToken bool, constant bool, sortInt uint64, length uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refKTokenIndexMask
	refRaw >>= refKTokenIndexShift
	length = refRaw & refKTokenLengthMask
	refRaw >>= refKTokenLengthShift
	sortInt = refRaw & refKTokenSortMask
	refRaw >>= refKTokenSortShift
	constant = refRaw&1 == 1
	refRaw >>= 1
	isKToken = refRaw == refKTokenTypeAsUint
	return
}

// MatchKToken returns true if reference is a KToken with correct sort.
// Function should be inlined, for performance reasons.
func MatchKToken(ref KReference, expectedSort uint64) bool {
	refRaw := uint64(ref)
	refRaw >>= refKTokenIndexShift        // ignore index
	refRaw >>= refKTokenLengthShift       // ignore length
	sortInt := refRaw & refKTokenSortMask // get sort
	refRaw >>= refKTokenSortShift         // for matching
	refRaw >>= 1                          // ignore constant flag
	return refRaw == refKTokenTypeAsUint &&
		sortInt == expectedSort
}

// The byte array and string encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - next 1 bit: is constant = 1, not constant = 0
// - 26 bits: length
// - 32 bits: head (element) index

const refBytesLengthShift = 26
const refBytesLengthMask = (1 << refBytesLengthShift) - 1
const refBytesIndexShift = 32
const refBytesIndexMask = (1 << refBytesIndexShift) - 1

func parseKrefBytes(ref KReference) (refType kreferenceType, constant bool, startIndex uint64, length uint64) {
	refRaw := uint64(ref)
	startIndex = refRaw & refBytesIndexMask
	refRaw >>= refBytesIndexShift
	length = refRaw & refBytesLengthMask
	refRaw >>= refBytesLengthShift
	constant = refRaw&1 == 1
	refRaw >>= 1
	refType = kreferenceType(refRaw)
	return
}

func createKrefBytes(refType kreferenceType, constant bool, startIndex uint64, length uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= 1
	if constant {
		refRaw |= 1
	}
	refRaw <<= refBytesLengthShift
	refRaw |= length
	refRaw <<= refBytesIndexShift
	refRaw |= startIndex
	return KReference(refRaw)
}
