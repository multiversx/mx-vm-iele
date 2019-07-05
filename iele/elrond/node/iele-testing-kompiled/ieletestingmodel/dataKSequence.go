// File provided by the K Framework Go backend. Timestamp: 2019-07-05 04:12:39.818

package ieletestingmodel

// EmptyKSequence is the KSequence with no elements.
// To simplify things, it is a separate reference type.
var EmptyKSequence = KReference{refType: emptyKseqRef}

type ksequenceSlice struct {
	firstKsHead int
	data        []KReference
}

type ksequenceSliceContainer struct {
	allSlices []*ksequenceSlice
}

func (c *ksequenceSliceContainer) getSlice(index int) *ksequenceSlice {
	return c.allSlices[index]
}

func (c *ksequenceSliceContainer) addSlice(s *ksequenceSlice) int {
	newSliceIndex := len(c.allSlices)
	c.allSlices = append(c.allSlices, s)
	return newSliceIndex
}

func createNonEmptyKseqRef(sequenceIndex int, headIndex int) KReference {
	return KReference{refType: nonEmptyKseqRef, value1: uint32(sequenceIndex), value2: uint32(headIndex)}
}

func nonEmptyKseqRefParse(ref KReference) (ok bool, sliceIndex int, headIndex int) {
	if ref.refType != nonEmptyKseqRef {
		return false, 0, 0
	}
	return true, int(ref.value1), int(ref.value2)
}

// IsNonEmptyKSequenceMinimumLength returns true for any K sequence with length greater of equal than given argument.
// Returns false for EmptyKSequence.
// Especially used for pattern matching.
func (ms *ModelState) IsNonEmptyKSequenceMinimumLength(ref KReference, minimumLength int) bool {
	if ref.refType == emptyKseqRef {
		return false
	}
	if ref.refType != nonEmptyKseqRef {
		return minimumLength == 1
	}
	_, sliceIndex, headIndex := nonEmptyKseqRefParse(ref)
	slice := ms.allKs.getSlice(sliceIndex)
	length := len(slice.data) - headIndex
	return length >= minimumLength
}

// NewKSequence creates new KSequence instance with given references
func (ms *ModelState) NewKSequence(elements []KReference) KReference {
	slice := &ksequenceSlice{firstKsHead: 0, data: elements}
	newSliceIndex := ms.allKs.addSlice(slice)
	return createNonEmptyKseqRef(newSliceIndex, 0)
}

// KSequenceIsEmpty returns true if KSequence has no elements
func (ms *ModelState) KSequenceIsEmpty(ref KReference) bool {
	if ref.refType == emptyKseqRef {
		return true
	}
	if ref.refType == nonEmptyKseqRef {
		if ms.KSequenceLength(ref) == 0 {
			panic("empty K sequence should have type emptyKseqRef, not nonEmptyKseqRef")
		}
	}
	return ms.KSequenceLength(ref) == 0
}

// KSequenceGet yields element at position
// Caution: no checks are performed that the position is valid
func (ms *ModelState) KSequenceGet(ref KReference, position int) KReference {
	ok, sliceIndex, headIndex := nonEmptyKseqRefParse(ref)
	if ok {
		slice := ms.allKs.getSlice(sliceIndex)
		return slice.data[headIndex+position]
	} else if ref.refType == emptyKseqRef {
		panic("bad argument to KSequenceGet: empty K sequence not allowed")
	}

	panic("bad argument to KSequenceGet: ref is not a reference to a K sequence")
}

// KSequenceLength yields KSequence length
func (ms *ModelState) KSequenceLength(ref KReference) int {
	ok, sliceIndex, headIndex := nonEmptyKseqRefParse(ref)
	if !ok {
		panic("bad argument to KSequenceLength: ref is not a reference to a K sequence")
	}
	slice := ms.allKs.getSlice(sliceIndex)
	return len(slice.data) - headIndex
}

// KSequenceToSlice converts KSequence to a slice of K items
func (ms *ModelState) KSequenceToSlice(ref KReference) []KReference {
	ok, sliceIndex, headIndex := nonEmptyKseqRefParse(ref)
	if !ok {
		panic("bad argument to KSequenceToSlice: ref is not a reference to a K sequence")
	}
	slice := ms.allKs.getSlice(sliceIndex)
	return slice.data[headIndex:]
}

// KSequenceSub yields subsequence starting at position
func (ms *ModelState) KSequenceSub(ref KReference, startPosition int) KReference {
	isNonEmptyKs, sliceIndex, headIndex := nonEmptyKseqRefParse(ref)
	if isNonEmptyKs {
		slice := ms.allKs.getSlice(sliceIndex)
		subLength := len(slice.data) - headIndex - startPosition
		if subLength < 0 {
			panic("bad argument to KSequenceSub: startPosition exceeds original K sequence")
		} else if subLength == 0 {
			return EmptyKSequence
		}
		return createNonEmptyKseqRef(sliceIndex, headIndex+startPosition)
	} else if ref.refType == emptyKseqRef {
		if startPosition == 0 {
			return EmptyKSequence
		}
		panic("bad argument to KSequenceSub: startPosition exceeds original K sequence (empty K sequence, startPosition > 0)")
	}

	panic("bad argument to KSequenceSub: ref is not a reference to a K sequence")
}

// KSequenceSplitHeadTail  extracts first element of a KSequence, extracts the rest, if possible
// will treat non-KSequence as if they were KSequences of length 1
func (ms *ModelState) KSequenceSplitHeadTail(ref KReference) (ok bool, head KReference, tail KReference) {
	if ref.refType == emptyKseqRef {
		return false, NoResult, EmptyKSequence
	}
	if isNonEmptyKseq, sliceIndex, headIndex := nonEmptyKseqRefParse(ref); isNonEmptyKseq {
		slice := ms.allKs.getSlice(sliceIndex)
		length := len(slice.data) - headIndex
		switch length {
		case 0:
			panic("empty K sequence should have type emptyKseqRef, not nonEmptyKseqRef")
		case 1:
			return true, slice.data[headIndex], EmptyKSequence
		case 2:
			// the KSequence has length 2
			// this case is special because here the tail is not a KSequence
			return true, slice.data[headIndex], slice.data[headIndex+1]
		default:
			// advance head
			return true, slice.data[headIndex], createNonEmptyKseqRef(sliceIndex, headIndex+1)
		}
	}

	// treat non-KSequences as if they were KSequences with 1 element
	return true, ref, EmptyKSequence
}

// AssembleKSequence appends all elements into a KSequence.
// It flattens any KSequences among the elements (but only on 1 level, does not handle multiple nesting).
// Never returns KSequence of 1 element, it returns the element directly instead
func (ms *ModelState) AssembleKSequence(elements ...KReference) KReference {
	var newKs []KReference
	for _, element := range elements {
		if element.refType == emptyKseqRef {
			// nothing, ignore
		} else if element.refType == nonEmptyKseqRef {
			newKs = append(newKs, ms.KSequenceToSlice(element)...)
		} else {
			newKs = append(newKs, element)
		}
	}
	if len(newKs) == 0 {
		return EmptyKSequence
	}
	if len(newKs) == 1 {
		return newKs[0]
	}
	return ms.NewKSequence(newKs)
}
