// File provided by the K Framework Go backend. Timestamp: 2019-07-15 14:09:18.513

package ieletestingmodel

import (
	"testing"
)

func TestKrefBasic(t *testing.T) {
	testKrefBasic(t, bottomRef, true, 0)
	testKrefBasic(t, floatRef, false, 0)

	testKrefBasic(t, bottomRef, true, refBasicDataMask)
	testKrefBasic(t, boolRef, false, refBasicDataMask)
}

func testKrefBasic(t *testing.T, refType kreferenceType, constant bool, rest uint64) {
	ref := createKrefBasic(refType, constant, rest)
	decodedType, decodedConstant, decodedRest := parseKrefBasic(ref)
	if decodedType != refType {
		t.Error("testKrefBasic mismatch")
	}
	if decodedConstant != constant {
		t.Error("testKrefBasic mismatch")
	}
	if decodedRest != rest {
		t.Error("testKrefBasic mismatch")
	}
}

func TestKrefBigInt(t *testing.T) {
	testKrefBigInt(t, true, 0, 0)
	testKrefBigInt(t, false, 0, 0)

	testKrefBigInt(t, true, 100, 50)
	testKrefBigInt(t, false, 1000, 3)
}

func testKrefBigInt(t *testing.T, constant bool, recycleCount uint64, index uint64) {
	ref := createKrefBigInt(constant, recycleCount, index)
	isBigInt, constantOut, recycleCountOut, indexOut := parseKrefBigInt(ref)
	if !isBigInt {
		t.Error("testKrefBigInt bad refType")
	}
	if constantOut != constant {
		t.Error("testKrefBigInt mismatch")
	}
	if recycleCountOut != recycleCount {
		t.Error("testKrefBigInt mismatch")
	}
	if indexOut != index {
		t.Error("testKrefBigInt mismatch")
	}
}

func TestKrefCollection(t *testing.T) {
	testKrefCollection(t, listRef, 5, 7, 123)
	testKrefList(t, 2, 4)
}

func testKrefCollection(t *testing.T, refType kreferenceType, sortInt uint64, labelInt uint64, index uint64) {
	ref := createKrefCollection(refType, sortInt, labelInt, index)
	refTypeOut, sortOut, labelOut, indexOut := parseKrefCollection(ref)
	if refTypeOut != refType {
		t.Error("testKrefCollection bad refType")
	}
	if sortOut != sortInt {
		t.Error("testKrefCollection mismatch")
	}
	if labelOut != labelInt {
		t.Error("testKrefCollection mismatch")
	}
	if indexOut != index {
		t.Error("testKrefCollection mismatch")
	}
}

func testKrefList(t *testing.T, sortInt uint64, labelInt uint64) {
	ms := NewModel()
	ref := ms.NewList(Sort(sortInt), KLabel(labelInt), nil)
	refTypeOut, sortOut, labelOut, _ := parseKrefCollection(ref)
	if refTypeOut != listRef {
		t.Error("testKrefList bad refType")
	}
	if sortOut != sortInt {
		t.Error("testKrefList mismatch")
	}
	if labelOut != labelInt {
		t.Error("testKrefList mismatch")
	}
}
