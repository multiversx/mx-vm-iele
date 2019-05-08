package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestAssembleKSequenceEmpty1(t *testing.T) {
	kseq := assembleKSequence()
	assertKSequenceEmpty(t, kseq)
}

func TestAssembleKSequenceEmpty2(t *testing.T) {
	kseq := assembleKSequence(m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq)
}

func TestAssembleKSequenceEmpty3(t *testing.T) {
	kseq := assembleKSequence(m.EmptyKSequence, m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq)
}

func TestAssembleKSequenceEmpty4(t *testing.T) {
	kseq := assembleKSequence(m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq)
}

func TestAssembleKSequence1(t *testing.T) {
	kseq := assembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2))
	assertKSequenceOfInts(t, kseq, 1, 2)
}

func TestAssembleKSequence2(t *testing.T) {
	kseq := assembleKSequence(m.EmptyKSequence, m.NewIntFromInt(1), m.NewIntFromInt(2), m.EmptyKSequence)
	assertKSequenceOfInts(t, kseq, 1, 2)
}

func TestAssembleKSequence3(t *testing.T) {
	kseq := assembleKSequence(m.NewIntFromInt(1))
	assertIntOk(t, "1", kseq, nil)
}

func TestAssembleKSequence4(t *testing.T) {
	kseq := assembleKSequence(m.NewIntFromInt(1), m.EmptyKSequence)
	assertIntOk(t, "1", kseq, nil)
}

func TestAssembleKSequence5(t *testing.T) {
	kseq := assembleKSequence(m.EmptyKSequence, m.NewIntFromInt(1))
	assertIntOk(t, "1", kseq, nil)
}

func TestAssembleKSequenceNest1(t *testing.T) {
	kseq1 := assembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2))
	kseq2 := assembleKSequence(m.NewIntFromInt(3), m.NewIntFromInt(4))
	kseq3 := assembleKSequence(kseq1, kseq2)
	assertKSequenceOfInts(t, kseq3, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest2(t *testing.T) {
	kseq1 := assembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2), m.EmptyKSequence)
	kseq2 := assembleKSequence(m.NewIntFromInt(3), m.EmptyKSequence, m.NewIntFromInt(4))
	kseq3 := assembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest3(t *testing.T) {
	kseq1 := assembleKSequence(m.EmptyKSequence)
	kseq2 := assembleKSequence(m.NewIntFromInt(3), m.EmptyKSequence, m.NewIntFromInt(4))
	kseq3 := assembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, 3, 4)
}

func assertKSequenceEmpty(t *testing.T, actual m.K) {
	expected := m.EmptyKSequence
	if !actual.Equals(expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s", m.PrettyPrint(actual), m.PrettyPrint(expected))
	}
}

func assertKSequenceOfInts(t *testing.T, actual m.K, ints ...int) {
	var ks []m.K
	for _, i := range ints {
		ks = append(ks, m.NewIntFromInt(i))
	}
	expected := &m.KSequence{Ks: ks}
	if !actual.Equals(expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s", m.PrettyPrint(actual), m.PrettyPrint(expected))
	}
}
