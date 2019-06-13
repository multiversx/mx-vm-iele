// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:50:56.636

package koreparser

import (
	"fmt"
	"strings"
)

// K ... Defines a parsed K entity, this is either a KItem, or a KSequence of KItems
type K interface {
	PrettyTreePrint(indent int) string
}

// KSequence ... a sequence of K items
type KSequence []K

// KApply ... raw KApply object, as parsed from KAST
type KApply struct {
	Label string
	List  []K
}

// InjectedKLabel ... raw InjectedKLabel, as parsed from KAST
type InjectedKLabel struct {
	Label string
}

// KToken ... raw KToken, as parsed from KAST
type KToken struct {
	Value string
	Sort  string
}

// KVariable ... raw KVariable, as parsed from KAST
type KVariable struct {
	Name string
}

func addIndent(sb *strings.Builder, indent int) {
	for i := 0; i < indent; i++ {
		sb.WriteString("    ")
	}
}

func simplePrint(indent int, str string) string {
	var sb strings.Builder
	addIndent(&sb, indent)
	sb.WriteString(str)
	return sb.String()
}

// PrettyTreePrint ... A tree representation of a parsed K object
func (k KApply) PrettyTreePrint(indent int) string {
	var sb strings.Builder
	addIndent(&sb, indent)
	sb.WriteString("KApply {label:")
	sb.WriteString(k.Label)
	sb.WriteString(", list:")
	if len(k.List) == 0 {
		sb.WriteString("[] }")
	} else {
		for _, childk := range k.List {
			sb.WriteRune('\n')
			sb.WriteString(childk.PrettyTreePrint(indent + 1))
		}
		sb.WriteRune('\n')
		addIndent(&sb, indent)
		sb.WriteRune('}')
	}

	return sb.String()
}

// PrettyTreePrint ... A tree representation of a parsed K object
func (k InjectedKLabel) PrettyTreePrint(indent int) string {
	return simplePrint(indent, fmt.Sprintf("InjectedKLabel {label:%s}", k.Label))
}

// PrettyTreePrint ... A tree representation of a parsed K object
func (k KToken) PrettyTreePrint(indent int) string {
	return simplePrint(indent, fmt.Sprintf("KToken {value:%s, sort:%s}", k.Value, k.Sort))
}

// PrettyTreePrint ... A tree representation of a parsed K object
func (k KVariable) PrettyTreePrint(indent int) string {
	return simplePrint(indent, fmt.Sprintf("KVariable {name:%s}", k.Name))
}

// PrettyTreePrint ... A tree representation of a parsed K object
func (k KSequence) PrettyTreePrint(indent int) string {
	var sb strings.Builder
	addIndent(&sb, indent)
	sb.WriteString("KSequence {")
	if len(k) == 0 {
		sb.WriteString(" <empty> }")
	} else {
		for i, childk := range k {
			sb.WriteString("\n")
			sb.WriteString(childk.PrettyTreePrint(indent + 1))
			if i < len(k)-1 {
				sb.WriteString(" ~>")
			}
		}
		sb.WriteRune('\n')
		addIndent(&sb, indent)
		sb.WriteRune('}')
	}

	return sb.String()
}
