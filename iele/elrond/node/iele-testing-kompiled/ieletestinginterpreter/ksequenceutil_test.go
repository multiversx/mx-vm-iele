// File provided by the K Framework Go backend. Timestamp: 2019-07-30 16:33:19.058

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

func TestAssembleKSequenceEmpty1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence()
	assertKSequenceEmpty(t, kseq, interpreter)
}

func TestAssembleKSequenceEmpty2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq, interpreter)
}

func TestAssembleKSequenceEmpty3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq, interpreter)
}

func TestAssembleKSequenceEmpty4(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence, m.EmptyKSequence)
	assertKSequenceEmpty(t, kseq, interpreter)
}

func TestAssembleKSequence1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1), interpreter.Model.FromInt(2))
	assertKSequenceOfInts(t, kseq, interpreter, 1, 2)
}

func TestAssembleKSequence2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, interpreter.Model.FromInt(1), interpreter.Model.FromInt(2), m.EmptyKSequence)
	assertKSequenceOfInts(t, kseq, interpreter, 1, 2)
}

func TestAssembleKSequence3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1))
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequence4(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1), m.EmptyKSequence)
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequence5(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, interpreter.Model.FromInt(1))
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequenceNest1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1), interpreter.Model.FromInt(2))
	kseq2 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(3), interpreter.Model.FromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1), interpreter.Model.FromInt(2), m.EmptyKSequence)
	kseq2 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(3), m.EmptyKSequence, interpreter.Model.FromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(m.EmptyKSequence)
	kseq2 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(3), m.EmptyKSequence, interpreter.Model.FromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 3, 4)
}

func TestAssembleKSequenceNest4(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(2), interpreter.Model.FromInt(3))
	kseq2 := interpreter.Model.AssembleKSequence(interpreter.Model.FromInt(1), kseq1, interpreter.Model.FromInt(4))
	assertKSequenceOfInts(t, kseq2, interpreter, 1, 2, 3, 4)
}

func assertKSequenceEmpty(t *testing.T, actual m.KReference, interpreter *Interpreter) {
	expected := m.EmptyKSequence
	if !interpreter.Model.Equals(actual, expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
}

func TestSubSequence1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 0)
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequence2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 1)
	assertKSequenceOfInts(t, sub, interpreter, 2)
}

func TestSubSequence3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 2)
	assertKSequenceEmpty(t, sub, interpreter)
}

func TestSubSequenceAssemble1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.AssembleKSequence(interpreter.Model.KSequenceSub(kseq, 0))
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequenceAssemble2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.AssembleKSequence(interpreter.Model.KSequenceGet(kseq, 0), interpreter.Model.KSequenceSub(kseq, 1))
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequenceAssemble3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})
	sub := interpreter.Model.AssembleKSequence(
		interpreter.Model.KSequenceGet(kseq, 0),
		interpreter.Model.KSequenceGet(kseq, 1),
		interpreter.Model.KSequenceSub(kseq, 2))
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequenceEmpty(t *testing.T) {
	interpreter := newTestInterpreter()
	sub := interpreter.Model.KSequenceSub(m.EmptyKSequence, 0)
	assertKSequenceEmpty(t, sub, interpreter)

	// sub = m.EmptyKSequence.SubSequence(1)
	// assertKSequenceEmpty(t, sub)
}

func TestKSequenceLength(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.KReference{interpreter.Model.FromInt(1), interpreter.Model.FromInt(2)})

	length := interpreter.Model.KSequenceLength(kseq)
	if length != 2 {
		t.Errorf("Unexpected result length. Got:%d Want:%d", length, 2)
	}
}

func assertKSequenceOfInts(t *testing.T, actual m.KReference, interpreter *Interpreter, ints ...int) {
	var ks []m.KReference
	for _, i := range ints {
		ks = append(ks, interpreter.Model.FromInt(i))
	}
	expected := interpreter.Model.NewKSequence(ks)
	if !interpreter.Model.Equals(actual, expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
	if len(ints) >= 2 {
		// some additional checks related to slicing and length
		actualSlice := interpreter.Model.KSequenceToSlice(actual)
		if len(actualSlice) != len(ints) {
			t.Error("Bad K sequence length when converted to slice")
		}
		for i, expectedElem := range ints {
			if !interpreter.Model.Equals(actualSlice[i], interpreter.Model.FromInt(expectedElem)) {
				t.Error("Bad K sequence length when converted to slice")
			}
		}
	}
}
