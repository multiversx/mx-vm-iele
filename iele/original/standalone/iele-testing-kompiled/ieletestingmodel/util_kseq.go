// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:55:22.205

package ieletestingmodel

// EmptyKSequence ... the KSequence with no elements
var EmptyKSequence = KSequence{sequenceIndex: 0, headIndex: 0}

// NewKSequence ... creates new KSequence instance with elements
func (ms *ModelState) NewKSequence(elements []K) KSequence {
	newSequenceIndex := len(ms.allKs)
	ms.allKs = append(ms.allKs, elements)
	return KSequence{sequenceIndex: newSequenceIndex, headIndex: 0}
}

// KSequenceIsEmpty returns true if KSequence has no elements
func (ms *ModelState) KSequenceIsEmpty(k KSequence) bool {
	return ms.KSequenceLength(k) == 0
}

// KSequenceGet yields element at position
// Caution: no checks are performed that the position is valid
func (ms *ModelState) KSequenceGet(k KSequence, position int) K {
	seq := ms.allKs[k.sequenceIndex]
	return seq[k.headIndex+position]
}

// KSequenceLength yields KSequence length
func (ms *ModelState) KSequenceLength(k KSequence) int {
	return len(ms.allKs[k.sequenceIndex]) - k.headIndex
}

// KSequenceToSlice converts KSequence to a slice of K items
func (ms *ModelState) KSequenceToSlice(k KSequence) []K {
	return ms.allKs[k.sequenceIndex][k.headIndex:]
}

// KSequenceSub yields subsequence starting at position
func (ms *ModelState) KSequenceSub(k KSequence, startPosition int) KSequence {
	return KSequence{sequenceIndex: k.sequenceIndex, headIndex: k.headIndex + startPosition}
}

// TrySplitToHeadTail ... extracts first element of a KSequence, extracts the rest, if possible
// will treat non-KSequence as if they were KSequences of length 1
func (ms *ModelState) TrySplitToHeadTail(k K) (ok bool, head K, tail K) {
	if kseq, isKseq := k.(KSequence); isKseq {
		seq := ms.allKs[kseq.sequenceIndex]
		length := len(seq) - kseq.headIndex
		switch length {
		case 0:
			// empty KSequence, no result
			return false, NoResult, EmptyKSequence
		case 1:
			return true, seq[kseq.headIndex], EmptyKSequence
		case 2:
			// the KSequence has length 2
			// this case is special because here the tail is not a KSequence
			return true, seq[kseq.headIndex], seq[kseq.headIndex+1]
		default:
			// advance head
			return true, seq[kseq.headIndex], KSequence{kseq.sequenceIndex, kseq.headIndex + 1}
		}
	}

	// treat non-KSequences as if they were KSequences with 1 element
	return true, k, EmptyKSequence
}

// AssembleKSequence ... appends all elements into a KSequence
// flattens if there are any KSequences among the elements (but only on 1 level, does not handle multiple nesting)
// never returns KSequence of 1 element, it returns the element directly instead
func (ms *ModelState) AssembleKSequence(elements ...K) K {
	var newKs []K
	for _, element := range elements {
		if kseqElem, isKseq := element.(KSequence); isKseq {
			newKs = append(newKs, ms.KSequenceToSlice(kseqElem)...)
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
