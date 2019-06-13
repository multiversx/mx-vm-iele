// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:50:56.636

package ieletestingmodel

import (
	"fmt"
	"strings"
)


// KPrint ... returns a standard representation of a K item
func (ms *ModelState) KPrint(k K) string {
	var sb strings.Builder
	k.kprint(ms, &sb)
	return sb.String()
}

func kprintKLabel(sb *strings.Builder, klabel KLabel) {
	sb.WriteString("`")
	sb.WriteString(klabel.Name())
	sb.WriteString("`")
}

func kprintKApply(ms *ModelState, sb *strings.Builder, label KLabel, children []K) {
	kprintKLabel(sb, label)
	sb.WriteString("(")

	if len(children) == 0 {
		sb.WriteString(".KList")
	} else {
		for i, childk := range children {
			childk.kprint(ms, sb)
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
	toK.kprint(ms, sb)
}

func (k *List) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	toK.kprint(ms, sb)
}

func (k *Set) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	toK.kprint(ms, sb)
}

func (k *Array) kprint(ms *ModelState, sb *strings.Builder) {
	toK := k.collectionsToK(ms)
	toK.kprint(ms, sb)
}

func (k *Int) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortInt, k.Value.String(), false)
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

func (k *Bool) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKToken(sb, SortBool, fmt.Sprintf("%t", k.Value), false)
}

func (k *Bottom) kprint(ms *ModelState, sb *strings.Builder) {
	kprintKApply(ms, sb, LblXhashBottom, []K{})
}

func (k KSequence) kprint(ms *ModelState, sb *strings.Builder) {
	ks := ms.KSequenceToSlice(k)
	if len(ks) == 0 {
		sb.WriteString(".K")
	} else if len(ks) == 1 {
		ks[0].kprint(ms, sb)
	} else {
		for i, childk := range ks {
			childk.kprint(ms, sb)
			if i < len(ks)-1 {
				sb.WriteString(" ~> ")
			}
		}
	}
}
