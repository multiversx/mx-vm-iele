// File provided by the K Framework Go backend. Timestamp: 2019-05-21 00:58:51.823

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

func trySplitToHeadTail(k m.K) (ok bool, head m.K, tail m.K) {
	if kseq, isKseq := k.(*m.KSequence); isKseq {
		switch len(kseq.Ks) {
		case 0:
			return false, m.NoResult, m.EmptyKSequence
		case 1:
			return true, kseq.Ks[0], m.EmptyKSequence
		case 2:
		    return true, kseq.Ks[0], kseq.Ks[1]
		default:
			return true, kseq.Ks[0], &m.KSequence{Ks: kseq.Ks[1:]}
		}
	}

	// treat non-KSequences as if they were KSequences with 1 element
	return true, k, m.EmptyKSequence
}

// appends all elements to a KSequence
// flattens if there are any KSequences among the elements (but only on 1 level, does not handle multiple nesting)
// never returns KSequence of 1 element, it returns the element directly instead
func assembleKSequence(elements ...m.K) m.K {
	var newKs []m.K
	for _, element := range elements {
		if kseqElem, isKseq := element.(*m.KSequence); isKseq {
			newKs = append(newKs, kseqElem.Ks...)
		} else {
			newKs = append(newKs, element)
		}
	}
	if len(newKs) == 0 {
		return m.EmptyKSequence
	}
	if len(newKs) == 1 {
		return newKs[0]
	}
	return &m.KSequence{Ks: newKs}
}
