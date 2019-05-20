// File provided by the K Framework Go backend. Timestamp: 2019-05-21 00:58:51.823

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// KPrint ... returns a standard representation of a K item
func KPrint(k K) string {
	var sb strings.Builder
	k.kprint(&sb)
	return sb.String()
}

func kprintKLabel(sb *strings.Builder, klabel KLabel) {
	sb.WriteString("`")
	sb.WriteString(klabel.Name())
	sb.WriteString("`")
}

func kprintKApply(sb *strings.Builder, label KLabel, children []K) {
	kprintKLabel(sb, label)
	sb.WriteString("(")

	if len(children) == 0 {
		sb.WriteString(".KList")
	} else {
		for i, childk := range children {
			childk.kprint(sb)
			if i < len(children)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteRune(')')
	}
}

func (k *KApply) kprint(sb *strings.Builder) {
	kprintKApply(sb, k.Label, k.List)
}

func (k *InjectedKLabel) kprint(sb *strings.Builder) {
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

func (k *KToken) kprint(sb *strings.Builder) {
	kprintKToken(sb, k.Sort, k.Value, false)
}

func (k *KVariable) kprint(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("var %s", k.Name))
}

func (k *Map) kprint(sb *strings.Builder) {
	toK := k.collectionsToK()
	toK.kprint(sb)
}

func (k *List) kprint(sb *strings.Builder) {
	toK := k.collectionsToK()
	toK.kprint(sb)
}

func (k *Set) kprint(sb *strings.Builder) {
	toK := k.collectionsToK()
	toK.kprint(sb)
}

func (k *Array) kprint(sb *strings.Builder) {
	toK := k.collectionsToK()
	toK.kprint(sb)
}

func (k *Int) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortInt, k.Value.String(), false)
}

func (k *MInt) kprint(sb *strings.Builder) {
	panic("Not implemented")
}

func (k *Float) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortFloat, fmt.Sprintf("%f", k.Value), false)
}

func (k *String) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortString, k.Value, true)
}

func (k *StringBuffer) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortStringBuffer, k.Value.String(), true)
}

func (k *Bytes) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortBytes, string(k.Value), true)
}

func (k *Bool) kprint(sb *strings.Builder) {
	kprintKToken(sb, SortBool, fmt.Sprintf("%t", k.Value), false)
}

func (k *Bottom) kprint(sb *strings.Builder) {
	kprintKApply(sb, LblXhashBottom, []K{})
}

func (k *KSequence) kprint(sb *strings.Builder) {
	if len(k.Ks) == 0 {
		sb.WriteString(".K")
	} else if len(k.Ks) == 1 {
		k.Ks[0].kprint(sb)
	} else {
		for i, childk := range k.Ks {
			childk.kprint(sb)
			if i < len(k.Ks)-1 {
				sb.WriteString(" ~> ")
			}
		}
	}
}
