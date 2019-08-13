// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:16:45.638

package ieletestinginterpreter

// The file is duplicated in the interpreter and in the model intentionally,
// for performance reasons.
// This version is for: impmodel

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

type modelDataReference uint64

const (
	mainDataRef  modelDataReference = 3
	memoDataRef  modelDataReference = 2
	constDataRef modelDataReference = 1
	noDataRef    modelDataReference = 0
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
// - next 2 bits: model data identifier
// - the remaining 57 LS bits: type-specific data

const refTypeBits = 5                         // refType gets represented in this many bits
const refTypeMask = (1 << refTypeBits) - 1    // 5 bits of 1
const refTypeShift = 64 - refTypeBits         // shift right this many bits to get the refType
const refModelShift = 2                       // how many bits we use to determine in which model the data resides (main/constants)
const refModelMask = (1 << refModelShift) - 1 // 2 bits of 1
const refBasicDataShift = 64 - refTypeBits - refModelShift
const refBasicDataMask = (1 << refBasicDataShift) - 1

func getRefType(ref KReference) kreferenceType {
	return kreferenceType(uint64(ref) >> refTypeShift)
}

func changeModelDataFlag(ref KReference, dataRef modelDataReference) KReference {
	refRaw := uint64(ref)
	refRaw &^= (refModelMask << refBasicDataShift) // Go has a bit clear operator, cool!
	refRaw |= (uint64(dataRef) << refBasicDataShift)
	return KReference(refRaw)
}

func assertModelDataFlag(ref KReference, expectedModelRef modelDataReference) {
	refType, dataRef, _ := parseKrefBasic(ref)
	if refType == bigIntRef {
		if dataRef != expectedModelRef {
			panic("unexpected model data reference")
		}
	}
}

func parseKrefBasic(ref KReference) (refType kreferenceType, dataRef modelDataReference, rest uint64) {
	refRaw := uint64(ref)
	rest = refRaw & refBasicDataMask
	refRaw >>= refBasicDataShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
	refType = kreferenceType(refRaw)
	return
}

func createKrefBasic(refType kreferenceType, dataRef modelDataReference, rest uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refBasicDataShift
	refRaw |= rest
	return KReference(refRaw)
}

// small int reference structure (from MSB to LSB):
// - 5 bits: reference type
// - 59 bits: absolute value
// note: the model data is unnecessary here and leaving it out saves us 2 bits

const refSmallIntValueBits = 59
const refSmallIntValueMask = (1 << refSmallIntValueBits) - 1

func createKrefSmallInt(i int64) KReference {
	var refRaw, absValue uint64
	if i < 0 {
		refRaw = uint64(smallNegativeIntRef)
		absValue = uint64(-i)
	} else {
		refRaw = uint64(smallPositiveIntRef)
		absValue = uint64(i)
	}

	refRaw <<= refSmallIntValueBits
	if absValue > refSmallIntValueMask {
		panic("small int absolute value too large")
	}
	refRaw |= absValue
	return KReference(refRaw)
}

func parseKrefSmallInt(ref KReference) (int64, bool) {
	refRaw := uint64(ref)
	absValue := refRaw & refSmallIntValueMask
	refRaw >>= refSmallIntValueBits
	if refRaw == uint64(smallPositiveIntRef) {
		return int64(absValue), true
	}
	if refRaw == uint64(smallNegativeIntRef) {
		return -int64(absValue), true
	}
	return 0, false
}

// big int reference structure (from MSB to LSB):
// - 5 bits: reference type
// - 2 bits: model data identifier
// - 25 bits: recycle count
// - 32 bits: index

const refBigIntRecycleCountBits = 25
const refBigIntRecycleCountMask = (1 << refBigIntRecycleCountBits) - 1
const refBigIntIndexBits = 32
const refBigIntIndexMask = (1 << refBigIntIndexBits) - 1

func createKrefBigInt(dataRef modelDataReference, recycleCount uint64, index uint64) KReference {
	refRaw := uint64(bigIntRef) << refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refBigIntRecycleCountBits
	refRaw |= recycleCount
	refRaw <<= refBigIntIndexBits
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefBigInt(ref KReference) (isBigInt bool, dataRef modelDataReference, recycleCount uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refBigIntIndexMask
	refRaw >>= refBigIntIndexBits
	recycleCount = refRaw & refBigIntRecycleCountMask
	refRaw >>= refBigIntRecycleCountBits
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
	isBigInt = refRaw == uint64(bigIntRef)
	return
}

// The collection encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - 2 bits: model data identifier
// - 12 bits: Sort
// - 13 bits: Label
// - 32 bits: object index

const refCollectionSortShift = 12
const refCollectionSortMask = (1 << refCollectionSortShift) - 1
const refCollectionLabelShift = 13
const refCollectionLabelMask = (1 << refCollectionLabelShift) - 1
const refCollectionIndexShift = 32
const refCollectionIndexMask = (1 << refCollectionIndexShift) - 1

func createKrefCollection(refType kreferenceType, dataRef modelDataReference, sortInt uint64, labelInt uint64, index uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refCollectionSortShift
	refRaw |= sortInt
	refRaw <<= refCollectionLabelShift
	refRaw |= labelInt
	refRaw <<= refCollectionIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefCollection(ref KReference) (refType kreferenceType, dataRef modelDataReference, sortInt uint64, labelInt uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refCollectionIndexMask
	refRaw >>= refCollectionIndexShift
	labelInt = refRaw & refCollectionLabelMask
	refRaw >>= refCollectionLabelShift
	sortInt = refRaw & refCollectionSortMask
	refRaw >>= refCollectionSortShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
	refType = kreferenceType(refRaw)
	return
}

// The KApply encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - 2 bits: model data identifier
// - 13 bits: label
// - 12 bits: arity
// - 32 bits: arguments index

const refKApplyLabelShift = 13
const refKApplyLabelMask = (1 << refKApplyLabelShift) - 1
const refKApplyArityShift = 12
const refKApplyArityMask = (1 << refKApplyArityShift) - 1
const refKApplyIndexShift = 32
const refKApplyIndexMask = (1 << refKApplyIndexShift) - 1
const refKApplyTypeAsUint = uint64(kapplyRef)

func createKrefKApply(dataRef modelDataReference, labelInt uint64, arity uint64, index uint64) KReference {
	refRaw := refKApplyTypeAsUint
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refKApplyLabelShift
	refRaw |= labelInt
	refRaw <<= refKApplyArityShift
	refRaw |= arity
	refRaw <<= refKApplyIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefKApply(ref KReference) (isKApply bool, dataRef modelDataReference, labelInt uint64, arity uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refKApplyIndexMask
	refRaw >>= refKApplyIndexShift
	arity = refRaw & refKApplyArityMask
	refRaw >>= refKApplyArityShift
	labelInt = refRaw & refKApplyLabelMask
	refRaw >>= refKApplyLabelShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
	isKApply = refRaw == refKApplyTypeAsUint
	return
}

// The K sequence encoding is as follows (from MSB to LSB):
// - 5 bits: reference type
// - 2 bits: model data identifier
// - 25 bits: length
// - 32 bits: head (element) index

const refNonEmptyKseqLengthShift = 25
const refNonEmptyKseqLengthMask = (1 << refNonEmptyKseqLengthShift) - 1
const refNonEmptyKseqIndexShift = 32
const refNonEmptyKseqIndexMask = (1 << refNonEmptyKseqIndexShift) - 1
const refNonEmptyKseqTypeAsUint = uint64(nonEmptyKseqRef)
const refEmptyKseqTypeAsUint = uint64(emptyKseqRef)

func createKrefNonEmptyKseq(dataRef modelDataReference, elemIndex uint64, length uint64) KReference {
	refRaw := refNonEmptyKseqTypeAsUint
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refNonEmptyKseqLengthShift
	refRaw |= length
	refRaw <<= refNonEmptyKseqIndexShift
	refRaw |= elemIndex
	return KReference(refRaw)
}

func parseKrefKseq(ref KReference) (refType kreferenceType, dataRef modelDataReference, elemIndex uint64, length uint64) {
	refRaw := uint64(ref)
	elemIndex = refRaw & refNonEmptyKseqIndexMask
	refRaw >>= refNonEmptyKseqIndexShift
	length = refRaw & refNonEmptyKseqLengthMask
	refRaw >>= refNonEmptyKseqLengthShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
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
	refRaw >>= refModelShift                     // ignore constant flag

	// refRaw is the reference type at this point
	return refRaw == refNonEmptyKseqTypeAsUint && length >= minimumLength
}

// The K token encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - 2 bits: model data identifier
// - 12 bits: sortInt
// - 13 bits: length
// - 32 bits: value string start index in allBytes

const refKTokenSortShift = 12
const refKTokenSortMask = (1 << refKTokenSortShift) - 1
const refKTokenLengthShift = 13
const refKTokenLengthMask = (1 << refKTokenLengthShift) - 1
const refKTokenIndexShift = 32
const refKTokenIndexMask = (1 << refKTokenIndexShift) - 1
const refKTokenTypeAsUint = uint64(ktokenRef)

func createKrefKToken(dataRef modelDataReference, sortInt uint64, length uint64, index uint64) KReference {
	refRaw := refKTokenTypeAsUint
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refKTokenSortShift
	refRaw |= sortInt
	refRaw <<= refKTokenLengthShift
	refRaw |= length
	refRaw <<= refKTokenIndexShift
	refRaw |= index
	return KReference(refRaw)
}

func parseKrefKToken(ref KReference) (isKToken bool, dataRef modelDataReference, sortInt uint64, length uint64, index uint64) {
	refRaw := uint64(ref)
	index = refRaw & refKTokenIndexMask
	refRaw >>= refKTokenIndexShift
	length = refRaw & refKTokenLengthMask
	refRaw >>= refKTokenLengthShift
	sortInt = refRaw & refKTokenSortMask
	refRaw >>= refKTokenSortShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
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
	refRaw >>= refModelShift              // ignore constant flag
	return refRaw == refKTokenTypeAsUint &&
		sortInt == expectedSort
}

// The byte array and string encoding is as follows (from MSB to LSB):
// - first 5 bits: reference type
// - 2 bits: model data identifier
// - 25 bits: length
// - 32 bits: head (element) index

const refBytesLengthShift = 25
const refBytesLengthMask = (1 << refBytesLengthShift) - 1
const refBytesIndexShift = 32
const refBytesIndexMask = (1 << refBytesIndexShift) - 1

func parseKrefBytes(ref KReference) (refType kreferenceType, dataRef modelDataReference, startIndex uint64, length uint64) {
	refRaw := uint64(ref)
	startIndex = refRaw & refBytesIndexMask
	refRaw >>= refBytesIndexShift
	length = refRaw & refBytesLengthMask
	refRaw >>= refBytesLengthShift
	dataRef = modelDataReference(refRaw & refModelMask)
	refRaw >>= refModelShift
	refType = kreferenceType(refRaw)
	return
}

func createKrefBytes(refType kreferenceType, dataRef modelDataReference, startIndex uint64, length uint64) KReference {
	refRaw := uint64(refType)
	refRaw <<= refModelShift
	refRaw |= uint64(dataRef)
	refRaw <<= refBytesLengthShift
	refRaw |= length
	refRaw <<= refBytesIndexShift
	refRaw |= startIndex
	return KReference(refRaw)
}
