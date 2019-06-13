// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:48:34.029

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// StructPrint ... returns a representation of a K item that somewhat resembles a Go declaration
func (ms *ModelState) StructPrint(k K) string {
	var sb strings.Builder
	k.structPrint(ms, &sb, 0)
	return sb.String()
}

func simplePrint(sb *strings.Builder, indent int, str string) {
	addIndent(sb, indent)
	sb.WriteString(str)
}

func (k *KApply) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	addIndent(sb, indent)
	sb.WriteString("KApply {Label:")
	sb.WriteString(k.Label.Name())
	sb.WriteString(", List:")
	if len(k.List) == 0 {
		sb.WriteString("[] }")
	} else {
		for _, childk := range k.List {
			sb.WriteRune('\n')
			childk.structPrint(ms, sb, indent+1)
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune('}')
	}
}

func (k *InjectedKLabel) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("InjectedKLabel {Label:%s}", k.Label.Name()))
}

func (k *KToken) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("KToken {Sort:%s, Value:%s}", k.Sort.Name(), k.Value))
}

func (k *KVariable) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("KVariable {Name:%s}", k.Name))
}

func (k *Map) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	addIndent(sb, indent)
	sb.WriteString("Map {Sort:")
	sb.WriteString(k.Sort.Name())
	sb.WriteString(", Label:")
	sb.WriteString(k.Label.Name())
	sb.WriteString(", Data:")
	if len(k.Data) == 0 {
		sb.WriteString(" <empty> }")
	} else {
		for k, v := range k.Data {
			sb.WriteString("\n")
			addIndent(sb, indent+1)
			sb.WriteString("key: ")
			sb.WriteString(k.String())
			sb.WriteString("  value: ")
			v.structPrint(ms, sb, 0)
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune('}')
	}
}

func (k *List) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	// TODO: print data
	simplePrint(sb, indent, fmt.Sprintf("List {Sort:%s, Label:%s}", k.Sort.Name(), k.Label.Name()))
}

func (k *Set) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	// TODO: print data
	simplePrint(sb, indent, fmt.Sprintf("Set {Sort:%s, Label:%s}", k.Sort.Name(), k.Label.Name()))
}

func (k *Array) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	// TODO: print data
	simplePrint(sb, indent, fmt.Sprintf("Array {Sort:%s}", k.Sort.Name()))
}

func (k *Int) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("Int (%s)", k.Value.String()))
}

func (k *MInt) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("MInt (%d)", k.Value))
}

func (k *Float) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("Float (%f)", k.Value))
}

func (k *String) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("String (%s)", k))
}

func (k *StringBuffer) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("StringBuffer (%s)", k.Value.String()))
}

func (k *Bytes) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("Bytes (%b)", k))
}

func (k *Bool) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, fmt.Sprintf("Bool (%t)", k.Value))
}

func (k *Bottom) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	simplePrint(sb, indent, "Bottom")
}

func (k KSequence) structPrint(ms *ModelState, sb *strings.Builder, indent int) {
	ks := ms.KSequenceToSlice(k)
	addIndent(sb, indent)
	sb.WriteString("KSequence {")
	if len(ks) == 0 {
		sb.WriteString(" <empty> }")
	} else {
		for i, childk := range ks {
			sb.WriteString("\n")
			childk.structPrint(ms, sb, indent+1)
			if i < len(ks)-1 {
				sb.WriteString(" ~>")
			}
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune('}')
	}
}
