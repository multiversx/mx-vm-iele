// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:46:43.258

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
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
	kseq := interpreter.Model.AssembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2))
	assertKSequenceOfInts(t, kseq, interpreter, 1, 2)
}

func TestAssembleKSequence2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, m.NewIntFromInt(1), m.NewIntFromInt(2), m.EmptyKSequence)
	assertKSequenceOfInts(t, kseq, interpreter, 1, 2)
}

func TestAssembleKSequence3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.NewIntFromInt(1))
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequence4(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.NewIntFromInt(1), m.EmptyKSequence)
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequence5(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.AssembleKSequence(m.EmptyKSequence, m.NewIntFromInt(1))
	assertIntOk(t, "1", kseq, nil, interpreter)
}

func TestAssembleKSequenceNest1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2))
	kseq2 := interpreter.Model.AssembleKSequence(m.NewIntFromInt(3), m.NewIntFromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(m.NewIntFromInt(1), m.NewIntFromInt(2), m.EmptyKSequence)
	kseq2 := interpreter.Model.AssembleKSequence(m.NewIntFromInt(3), m.EmptyKSequence, m.NewIntFromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 1, 2, 3, 4)
}

func TestAssembleKSequenceNest3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq1 := interpreter.Model.AssembleKSequence(m.EmptyKSequence)
	kseq2 := interpreter.Model.AssembleKSequence(m.NewIntFromInt(3), m.EmptyKSequence, m.NewIntFromInt(4))
	kseq3 := interpreter.Model.AssembleKSequence(kseq1, m.EmptyKSequence, kseq2)
	assertKSequenceOfInts(t, kseq3, interpreter, 3, 4)
}

func assertKSequenceEmpty(t *testing.T, actual m.K, interpreter *Interpreter) {
	expected := m.EmptyKSequence
	if !interpreter.Model.Equals(actual, expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
}

func TestSubSequence1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 0)
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequence2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 1)
	assertKSequenceOfInts(t, sub, interpreter, 2)
}

func TestSubSequence3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
	sub := interpreter.Model.KSequenceSub(kseq, 2)
	assertKSequenceEmpty(t, sub, interpreter)
}

func TestSubSequenceAssemble1(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
	sub := interpreter.Model.AssembleKSequence(interpreter.Model.KSequenceSub(kseq, 0))
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequenceAssemble2(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
	sub := interpreter.Model.AssembleKSequence(interpreter.Model.KSequenceGet(kseq, 0), interpreter.Model.KSequenceSub(kseq, 1))
	assertKSequenceOfInts(t, sub, interpreter, 1, 2)
}

func TestSubSequenceAssemble3(t *testing.T) {
	interpreter := newTestInterpreter()
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})
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
	kseq := interpreter.Model.NewKSequence([]m.K{m.NewIntFromInt(1), m.NewIntFromInt(2)})

	length := interpreter.Model.KSequenceLength(kseq)
	if length != 2 {
		t.Errorf("Unexpected result length. Got:%d Want:%d", length, 2)
	}
}

func assertKSequenceOfInts(t *testing.T, actual m.K, interpreter *Interpreter, ints ...int) {
	var ks []m.K
	for _, i := range ints {
		ks = append(ks, m.NewIntFromInt(i))
	}
	expected := interpreter.Model.NewKSequence(ks)
	if !interpreter.Model.Equals(actual, expected) {
		t.Errorf("Unexpected result. Got:%s Want:%s",
			interpreter.Model.PrettyPrint(actual),
			interpreter.Model.PrettyPrint(expected))
	}
}
