// File provided by the K Framework Go backend. Timestamp: 2019-07-15 13:11:08.386

package ieletestingmodel

// EmptyKSequence is the KSequence with no elements.
// To simplify things, it is a separate reference type.
var EmptyKSequence = createKrefBasic(emptyKseqRef, true, 0)

type ksequenceElem struct {
	head KReference
	tail KReference
}

// IsNonEmptyKSequenceMinimumLength returns true for any K sequence with length greater of equal than given argument.
// Returns false for EmptyKSequence.
// Especially used for pattern matching.
func (ms *ModelState) IsNonEmptyKSequenceMinimumLength(ref KReference, minimumLength uint64) bool {
	refType, _, length := parseKrefKseq(ref)
	if refType == emptyKseqRef {
		return false
	}
	if refType != nonEmptyKseqRef {
		return minimumLength == 1
	}
	return length >= minimumLength
}

// NewKSequence creates new KSequence instance with given references
func (ms *ModelState) NewKSequence(elements []KReference) KReference {
	return ms.AssembleKSequence(elements...)
}

// KSequenceIsEmpty returns true if KSequence has no elements
func (ms *ModelState) KSequenceIsEmpty(ref KReference) bool {
	refType, _, _ := parseKrefBasic(ref)
	return refType == emptyKseqRef
}

// KSequenceGet yields element at position.
func (ms *ModelState) KSequenceGet(ref KReference, position int) KReference {
	for i := 0; i < position; i++ {
		refType, elemIndex, _ := parseKrefKseq(ref)
		if refType != nonEmptyKseqRef {
			panic("bad argument to KSequenceGet: position exceeds K sequence length")
		}
		elem := ms.allKsElements[elemIndex]
		ref = elem.tail
	}

	refType, elemIndex, _ := parseKrefKseq(ref)
	if refType != nonEmptyKseqRef {
		return ref
	}
	return ms.allKsElements[elemIndex].head
}

// KSequenceLength yields KSequence length
func (ms *ModelState) KSequenceLength(ref KReference) uint64 {
	refType, _, length := parseKrefKseq(ref)
	if refType != nonEmptyKseqRef {
		panic("bad argument to KSequenceLength: ref is not a reference to a K sequence")
	}
	return length
}

// KSequenceToSlice converts KSequence to a slice of K items
func (ms *ModelState) KSequenceToSlice(ref KReference) []KReference {
	refType, elemIndex, length := parseKrefKseq(ref)
	if refType == emptyKseqRef {
		return nil
	}
	if refType != nonEmptyKseqRef {
		panic("bad argument to KSequenceToSlice: ref is not a reference to a K sequence")
	}

	var result []KReference
	for refType == nonEmptyKseqRef {
		elem := ms.allKsElements[elemIndex]
		result = append(result, elem.head)
		ref = elem.tail
		refType, elemIndex, _ = parseKrefKseq(ref)
	}

	// last element is not a K sequence
	result = append(result, ref)

	if uint64(len(result)) != length {
		panic("K sequence reference length does not match actual length of K sequence")
	}

	return result
}

// KSequenceSub yields subsequence starting at position
func (ms *ModelState) KSequenceSub(ref KReference, startPosition int) KReference {
	for i := 0; i < startPosition; i++ {
		refType, elemIndex, _ := parseKrefKseq(ref)
		if refType != nonEmptyKseqRef {
			if i == startPosition-1 {
				return EmptyKSequence
			}
			panic("bad argument to KSequenceSub: startPosition exceeds original K sequence")
		} else {
			elem := ms.allKsElements[elemIndex]
			ref = elem.tail
		}
	}

	return ref
}

// KSequenceSplitHeadTail  extracts first element of a KSequence, extracts the rest, if possible
// will treat non-KSequence as if they were KSequences of length 1
func (ms *ModelState) KSequenceSplitHeadTail(ref KReference) (ok bool, head KReference, tail KReference) {
	refType, elemIndex, _ := parseKrefKseq(ref)
	if refType == emptyKseqRef {
		return false, NoResult, EmptyKSequence
	}

	if refType == nonEmptyKseqRef {
		elem := ms.allKsElements[elemIndex]
		return true, elem.head, elem.tail
	}

	// treat non-KSequences as if they were KSequences with 1 element
	return true, ref, EmptyKSequence
}

// AssembleKSequence appends all given arguments into a KSequence.
// It flattens any KSequences among the arguments.
// Never returns KSequence of 1 element, it returns the element directly instead
func (ms *ModelState) AssembleKSequence(refs ...KReference) KReference {
	head := EmptyKSequence
	var resultLength uint64

	for i := len(refs) - 1; i >= 0; i-- {
		ref := refs[i]
		refType, _, refLength := parseKrefKseq(ref)
		headType, _, _ := parseKrefKseq(head)
		if refType == emptyKseqRef {
			// nothing, ignore
		} else {
			if headType == emptyKseqRef {
				// first to be added
				head = ref
				if refType == nonEmptyKseqRef {
					// last ref is a K sequence that we
					resultLength = refLength
				} else {
					resultLength = 1
				}
			} else {
				// append to the simple linked list
				// like the cons in cons lists
				if refType == nonEmptyKseqRef {
					// flatten K sequence given as argument
					// concatenate entire sub-sequence to beginning of result sequence
					slice := ms.KSequenceToSlice(ref)
					slice = append(slice, head)
					head = ms.AssembleKSequence(slice...)
					_, _, resultLength = parseKrefKseq(head)
				} else {
					// add 1 element to beginning of list
					newHead := ksequenceElem{
						head: ref,
						tail: head,
					}
					resultLength++
					newIndex := uint64(len(ms.allKsElements))
					ms.allKsElements = append(ms.allKsElements, newHead)
					head = createKrefNonEmptyKseq(newIndex, resultLength)
				}
			}
		}
	}

	return head
}
