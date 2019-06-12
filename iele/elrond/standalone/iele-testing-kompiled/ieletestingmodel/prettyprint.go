// File provided by the K Framework Go backend. Timestamp: 2019-06-12 17:39:27.917

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// PrettyPrint ... returns a representation of a K item that tries to be as readable as possible
// designed for debugging purposes only
func (ms *ModelState) PrettyPrint(k K) string {
	var sb strings.Builder
	k.prettyPrint(ms, &sb, 0)
	return sb.String()
}

func (k *KApply) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	lblName := k.Label.Name()
	isKCell := strings.HasPrefix(lblName, "<") && strings.HasSuffix(lblName, ">")

	// begin
	sb.WriteString(lblName)
	if !isKCell {
		sb.WriteString("(")
	}

	// contents
	done := false
	if len(k.List) == 0 {
		done = true
	}
	if !done && len(k.List) == 1 {
		var tempSb strings.Builder
		k.List[0].prettyPrint(ms, &tempSb, 0)
		childStr := tempSb.String()
		if !strings.Contains(childStr, "\n") {
			// if only one child and its representation not too big, just put everything in one row
			if isKCell {
				sb.WriteString(" ")
			}
			sb.WriteString(childStr)
			if isKCell {
				sb.WriteString(" ")
			}
			done = true
		}
	}
	if !done {
		for i, childk := range k.List {
			sb.WriteRune('\n')
			addIndent(sb, indent+1)
			childk.prettyPrint(ms, sb, indent+1)
			if !isKCell && i < len(k.List)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
	}

	// end
	if isKCell {
		sb.WriteString("</")
		sb.WriteString(strings.TrimPrefix(lblName, "<"))
	} else {
		sb.WriteString(")")
	}
}

func (k *InjectedKLabel) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("InjectedKLabel(%s)", k.Label.Name()))
}

func (k *KToken) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("%s: %s", k.Sort.Name(), k.Value))
}

func (k *KVariable) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("var %s", k.Name))
}

func (k *Map) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("Map Sort:")
	sb.WriteString(k.Sort.Name())
	sb.WriteString(", Label:")
	sb.WriteString(k.Label.Name())
	if len(k.Data) == 0 {
		sb.WriteString(" <empty>")
	} else {
		sb.WriteString(", Data: (")
		orderedKVPairs := k.ToOrderedKeyValuePairs()
		for _, pair := range orderedKVPairs {
			sb.WriteString("\n")
			addIndent(sb, indent+1)
			pair.Key.prettyPrint(ms, sb, indent+1)
			sb.WriteString(" => ")
			pair.Value.prettyPrint(ms, sb, indent+1)
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune(')')
	}
}

func (k *List) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("List Sort:")
	sb.WriteString(k.Sort.Name())
	sb.WriteString(", Label:")
	sb.WriteString(k.Label.Name())
	if len(k.Data) == 0 {
		sb.WriteString(" <empty>")
	} else {
		sb.WriteString(", Data: [")
		for _, item := range k.Data {
			sb.WriteString("\n")
			addIndent(sb, indent+1)
			if item == nil {
				sb.WriteString("nil")
			} else {
				item.prettyPrint(ms, sb, indent+1)
			}
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune(']')
	}
}

func (k *Set) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("Set Sort:")
	sb.WriteString(k.Sort.Name())
	sb.WriteString(", Label:")
	sb.WriteString(k.Label.Name())
	if len(k.Data) == 0 {
		sb.WriteString(" <empty>")
	} else {
		sb.WriteString(", Data: {")
		orderedElems := k.ToOrderedElements()
		for _, elem := range orderedElems {
			sb.WriteString("\n")
			addIndent(sb, indent+1)
			elem.prettyPrint(ms, sb, indent+1)
		}

		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune('}')
	}
}

func (k *Array) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("Array Sort:")
	sb.WriteString(k.Sort.Name())
	slice := k.Data.ToSlice()
	if len(slice) == 0 {
		sb.WriteString(" <empty>")
	} else {
		sb.WriteString(", Data: [")
		for i, item := range slice {
			sb.WriteString("\n")
			addIndent(sb, indent+1)
			sb.WriteString(fmt.Sprintf("[%d] => ", i))
			item.prettyPrint(ms, sb, indent+1)
		}
		sb.WriteRune('\n')
		addIndent(sb, indent)
		sb.WriteRune(']')
	}
}

func (k *Int) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("Int (0x%s | %s)", k.Value.Text(16), k.Value.String()))
}

func (k *MInt) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("MInt (%d)", k.Value))
}

func (k *Float) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("Float (%f)", k.Value))
}

func (k *String) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("String(\"")
	writeEscapedString(sb, k.Value)
	sb.WriteString("\")")
}

func (k *StringBuffer) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("StringBuffer(\"")
	writeEscapedString(sb, k.Value.String())
	sb.WriteString("\")")
}

func (k *Bytes) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("Bytes(")
	if len(k.Value) == 0 {
		sb.WriteString("empty")
	} else {
		for i, b := range k.Value {
			sb.WriteString(fmt.Sprintf("%02x", b))
			if i < len(k.Value)-1 {
				sb.WriteByte(' ')
			}
		}
	}
	sb.WriteString(")")
}

func (k *Bool) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("Bool (%t)", k.Value))
}

func (k *Bottom) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	sb.WriteString("Bottom")
}

func (k KSequence) prettyPrint(ms *ModelState, sb *strings.Builder, indent int) {
	ks := ms.KSequenceToSlice(k)
	if len(ks) == 0 {
		sb.WriteString(" .K ")
	} else {
		for i, childk := range ks {
			if i > 0 {
				addIndent(sb, indent)
			}
			childk.prettyPrint(ms, sb, indent)
			if i < len(ks)-1 {
				sb.WriteString(" ~>\n")
			} else {
				sb.WriteString(" ~> . ")
			}
		}
	}
}
