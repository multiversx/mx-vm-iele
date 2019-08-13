// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:53:01.019

package ieletestingmodel

// EmptyKSequence is the KSequence with no elements.
// To simplify things, it is a separate reference type.
var EmptyKSequence = createKrefBasic(emptyKseqRef, noDataRef, 0)

type ksequenceElem struct {
	head KReference
	tail KReference
}

// IsNonEmptyKSequenceMinimumLength returns true for any K sequence with length greater of equal than given argument.
// Returns false for EmptyKSequence.
// Especially used for pattern matching.
func (ms *ModelState) IsNonEmptyKSequenceMinimumLength(ref KReference, minimumLength uint64) bool {
	refType, _, _, length := parseKrefKseq(ref)
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
		refType, dataRef, elemIndex, _ := parseKrefKseq(ref)
		if refType != nonEmptyKseqRef {
			panic("bad argument to KSequenceGet: position exceeds K sequence length")
		}
		elem := ms.getData(dataRef).allKsElements[elemIndex]
		ref = elem.tail
	}

	refType, dataRef, elemIndex, _ := parseKrefKseq(ref)
	if refType != nonEmptyKseqRef {
		return ref
	}
	return ms.getData(dataRef).allKsElements[elemIndex].head
}

// KSequenceLength yields KSequence length
func (ms *ModelState) KSequenceLength(ref KReference) uint64 {
	refType, _, _, length := parseKrefKseq(ref)
	if refType != nonEmptyKseqRef {
		panic("bad argument to KSequenceLength: ref is not a reference to a K sequence")
	}
	return length
}

// KSequenceToSlice converts KSequence to a slice of K items
func (ms *ModelState) KSequenceToSlice(ref KReference) []KReference {
	_, dataRef, _, _ := parseKrefKseq(ref)
	return ms.getData(dataRef).ksequenceToSlice(ref)
}

func (md *ModelData) ksequenceToSlice(ref KReference) []KReference {
	refType, dataRef, elemIndex, length := parseKrefKseq(ref)
	if refType == emptyKseqRef {
		return nil
	}
	if dataRef != md.selfRef {
		panic("trying to retrieve K sequence elements from wrong data container")
	}
	if refType != nonEmptyKseqRef {
		panic("bad argument to KSequenceToSlice: ref is not a reference to a K sequence")
	}

	var result []KReference
	for refType == nonEmptyKseqRef {
		elem := md.allKsElements[elemIndex]
		result = append(result, elem.head)
		ref = elem.tail
		var dataRef modelDataReference
		refType, dataRef, elemIndex, _ = parseKrefKseq(ref)
		if refType == nonEmptyKseqRef && dataRef != md.selfRef {
			panic("chaining K sequence elements from differnt data containers is not supported")
		}
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
		refType, dataRef, elemIndex, _ := parseKrefKseq(ref)
		if refType != nonEmptyKseqRef {
			if i == startPosition-1 {
				return EmptyKSequence
			}
			panic("bad argument to KSequenceSub: startPosition exceeds original K sequence")
		} else {
			elem := ms.getData(dataRef).allKsElements[elemIndex]
			ref = elem.tail
		}
	}

	return ref
}

// KSequenceSplitHeadTail  extracts first element of a KSequence, extracts the rest, if possible
// will treat non-KSequence as if they were KSequences of length 1
func (ms *ModelState) KSequenceSplitHeadTail(ref KReference) (ok bool, head KReference, tail KReference) {
	refType, dataRef, elemIndex, _ := parseKrefKseq(ref)
	if refType == emptyKseqRef {
		return false, NoResult, EmptyKSequence
	}

	if refType == nonEmptyKseqRef {
		elem := ms.getData(dataRef).allKsElements[elemIndex]
		return true, elem.head, elem.tail
	}

	// treat non-KSequences as if they were KSequences with 1 element
	return true, ref, EmptyKSequence
}

// AssembleKSequence appends all given arguments into a KSequence.
// It flattens any KSequences among the arguments.
// Never returns KSequence of 1 element, it returns the element directly instead
func (ms *ModelState) AssembleKSequence(refs ...KReference) KReference {
	return ms.mainData.assembleKSequence(refs...)
}

func (md *ModelData) assembleKSequence(refs ...KReference) KReference {
	head := EmptyKSequence
	var resultLength uint64

	for i := len(refs) - 1; i >= 0; i-- {
		ref := refs[i]
		refType, _, _, refLength := parseKrefKseq(ref)
		headType, _, _, _ := parseKrefKseq(head)
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
					slice := md.ksequenceToSlice(ref)
					slice = append(slice, head)
					head = md.assembleKSequence(slice...)
					_, _, _, resultLength = parseKrefKseq(head)
				} else {
					// add 1 element to beginning of list
					newHead := ksequenceElem{
						head: ref,
						tail: head,
					}
					resultLength++
					newIndex := uint64(len(md.allKsElements))
					md.allKsElements = append(md.allKsElements, newHead)
					head = createKrefNonEmptyKseq(md.selfRef, newIndex, resultLength)
				}
			}
		}
	}

	return head
}
