// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestingmodel

// kreferenceType identifies the type of K item referenced by a KReference
type kreferenceType byte

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

// KReference is a reference to a K item.
// For some types, like bool and small int, the entire state can be kept in the reference object.
// For most, the reference contains enough data to find the object in the model state.
type KReference struct {
	refType        kreferenceType
	constantObject bool
	value1         uint32
	value2         uint32
}

// NullReference is the zero-value of KReference. It doesn't point to anything.
var NullReference = KReference{refType: nullRef, value1: 0, value2: 0}

// NoResult is the result when a function returns an error
var NoResult = InternedBottom
