// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:53:01.019

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// KPrint returns a standard representation of a K item
func (ms *ModelState) KPrint(ref KReference) string {
	var sb strings.Builder
	ms.kprintToStringBuilder(&sb, ref)
	return sb.String()
}

func (ms *ModelState) kprintToStringBuilder(sb *strings.Builder, ref KReference) {
	// int types
	intStr, isInt := ms.GetIntAsDecimalString(ref)
	if isInt {
		kprintKToken(sb, SortInt, intStr, false)
		return
	}

	refType, dataRef, value := parseKrefBasic(ref)

	// collection types
	if isCollectionType(refType) {
		_, _, _, _, index := parseKrefCollection(ref)
		obj := ms.getData(dataRef).getReferencedObject(index)
		obj.kprint(ms, sb)
		return
	}

	switch refType {
	case boolRef:
		kprintKToken(sb, SortBool, fmt.Sprintf("%t", IsTrue(ref)), false)
	case bottomRef:
		kprintKApply(ms, sb, LblXhashBottom, []KReference{})
	case emptyKseqRef:
		sb.WriteString(".K")
	case nonEmptyKseqRef:
		ks := ms.KSequenceToSlice(ref)
		if len(ks) == 0 {
			panic("K sequences of length 0 should have type emptyKseqRef, not nonEmptyKseqRef")
		} else if len(ks) == 1 {
			ms.kprintToStringBuilder(sb, ks[0])
		} else {
			for i, child := range ks {
				ms.kprintToStringBuilder(sb, child)
				if i < len(ks)-1 {
					sb.WriteString(" ~> ")
				}
			}
		}
	case kapplyRef:
		kprintKApply(ms, sb, ms.KApplyLabel(ref), ms.kapplyArgSlice(ref))
	case stringRef:
		str, _ := ms.GetString(ref)
		kprintKToken(sb, SortString, str, true)
	case bytesRef:
		bytes, _ := ms.GetBytes(ref)
		kprintKToken(sb, SortBytes, string(bytes), true)
	case ktokenRef:
		ktoken, _ := ms.GetKTokenObject(ref)
		kprintKToken(sb, ktoken.Sort, ktoken.Value, false)
	default:
		// object types
		obj := ms.getData(dataRef).getReferencedObject(value)
		obj.kprint(ms, sb)
	}
}

func kprintKLabel(sb *strings.Builder, klabel KLabel) {
	sb.WriteString("`")
	sb.WriteString(klabel.Name())
	sb.WriteString("`")
}

func kprintKApply(ms *ModelState, sb *strings.Builder, label KLabel, args []KReference) {
	kprintKLabel(sb, label)
	sb.WriteString("(")

	if len(args) == 0 {
		sb.WriteString(".KList")
	} else {
		for i, child := range args {
			ms.kprintToStringBuilder(sb, child)
			if i < len(args)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteRune(')')
	}
}

func (k *InjectedKLabel) kprint(ms *ModelState, sb *strings.Builder) {
	sb.WriteString("#klabel(\"")
	sb.WriteString(k.Label.Name())
	sb.WriteString("\")")
}

func kprintKToken(sb *strings.Builder, sort Sort, value string, escape bool) {
	sb.WriteString("#token(\"")
	if escape {
		sb.WriteString("\\\"")
		writeEscapedString(sb, value)
		sb.WriteString("\\\"")
	} else {
		sb.WriteString(value)
	}
	sb.WriteString("\", \"")
	sb.WriteString(sort.Name())
	sb.WriteString("\")")
}

func (k *KVariable) kprint(ms *ModelState, sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("var %s", k.Name))
}

func (k *Map) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	ms.kprintToStringBuilder(sb, toK)
}

func (k *List) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	ms.kprintToStringBuilder(sb, toK)
}

func (k *Set) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	ms.kprintToStringBuilder(sb, toK)
}

func (k *Array) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	ms.kprintToStringBuilder(sb, toK)
}

func (k *MInt) kprint(ms *ModelState, sb *strings.Builder) {
	panic("Not implemented")
}

func (k *Float) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortFloat, fmt.Sprintf("%f", k.Value), false)
}

func (k *StringBuffer) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortStringBuffer, k.Value.String(), true)
}
