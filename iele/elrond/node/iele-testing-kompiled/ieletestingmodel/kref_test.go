// File provided by the K Framework Go backend. Timestamp: 2019-07-30 16:33:19.058

package ieletestingmodel

import (
	"testing"
)

func TestKrefBasic(t *testing.T) {
	testKrefBasic(t, bottomRef, mainDataRef, 0)
	testKrefBasic(t, floatRef, constDataRef, 0)

	testKrefBasic(t, bottomRef, memoDataRef, refBasicDataMask)
	testKrefBasic(t, boolRef, constDataRef, refBasicDataMask)
}

func testKrefBasic(t *testing.T, refType kreferenceType, dataRef modelDataReference, rest uint64) {
	ref := createKrefBasic(refType, dataRef, rest)
	decodedType, decodedDataRef, decodedRest := parseKrefBasic(ref)
	if decodedType != refType {
		t.Error("testKrefBasic mismatch")
	}
	if decodedDataRef != dataRef {
		t.Error("testKrefBasic mismatch")
	}
	if decodedRest != rest {
		t.Error("testKrefBasic mismatch")
	}
}

func TestKrefBigInt(t *testing.T) {
	testKrefBigInt(t, mainDataRef, 0, 0)
	testKrefBigInt(t, constDataRef, 0, 0)

	testKrefBigInt(t, memoDataRef, 100, 50)
	testKrefBigInt(t, mainDataRef, 1000, 3)
}

func testKrefBigInt(t *testing.T, dataRef modelDataReference, recycleCount uint64, index uint64) {
	ref := createKrefBigInt(dataRef, recycleCount, index)
	isBigInt, dataRefOut, recycleCountOut, indexOut := parseKrefBigInt(ref)
	if !isBigInt {
		t.Error("testKrefBigInt bad refType")
	}
	if dataRefOut != dataRef {
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
	testKrefCollection(t, listRef, memoDataRef, 5, 7, 123)
	testKrefList(t, 2, 4)
}

func testKrefCollection(t *testing.T, refType kreferenceType, dataRef modelDataReference, sortInt uint64, labelInt uint64, index uint64) {
	ref := createKrefCollection(refType, dataRef, sortInt, labelInt, index)
	refTypeOut, dataRefOut, sortOut, labelOut, indexOut := parseKrefCollection(ref)
	if refTypeOut != refType {
		t.Error("testKrefCollection bad refType")
	}
	if dataRefOut != dataRef {
		t.Error("testKrefCollection mismatch")
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
	refTypeOut, dataRefOut, sortOut, labelOut, _ := parseKrefCollection(ref)
	if refTypeOut != listRef {
		t.Error("testKrefList bad refType")
	}
	if dataRefOut != mainDataRef {
		t.Error("testKrefList mismatch")
	}
	if sortOut != sortInt {
		t.Error("testKrefList mismatch")
	}
	if labelOut != labelInt {
		t.Error("testKrefList mismatch")
	}
}
