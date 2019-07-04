// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

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

	switch ref.refType {
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
	default:
		// object types
		obj := ms.getReferencedObject(ref)
		obj.kprint(ms, sb)
	}
}

func kprintKLabel(sb *strings.Builder, klabel KLabel) {
	sb.WriteString("`")
	sb.WriteString(klabel.Name())
	sb.WriteString("`")
}

func kprintKApply(ms *ModelState, sb *strings.Builder, label KLabel, children []KReference) {
	kprintKLabel(sb, label)
	sb.WriteString("(")

	if len(children) == 0 {
		sb.WriteString(".KList")
	} else {
		for i, child := range children {
			ms.kprintToStringBuilder(sb, child)
			if i < len(children)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteRune(')')
	}
}

func (k *KApply) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKApply(ms, sb, k.Label, k.List)
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

func (k *KToken) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, k.Sort, k.Value, false)
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

func (k *String) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortString, k.Value, true)
}

func (k *StringBuffer) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortStringBuffer, k.Value.String(), true)
}

func (k *Bytes) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortBytes, string(k.Value), true)
}
