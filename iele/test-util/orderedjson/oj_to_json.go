package orderedjson

import (
	"fmt"
	"strings"
)

// JSONString ... returns a formatted string representation of an ordered JSON
func JSONString(j OJsonObject) string {
	var sb strings.Builder
	j.writeJSON(&sb, 0)
	return sb.String()
}

func addIndent(sb *strings.Builder, indent int) {
	for i := 0; i < indent; i++ {
		sb.WriteString("    ")
	}
}

func (j *OJsonMap) writeJSON(sb *strings.Builder, indent int) {
	sb.WriteString("{")
	for i, child := range j.OrderedKV {
		sb.WriteString("\n")
		addIndent(sb, indent+1)
		sb.WriteString(child.Key)
		sb.WriteString(": ")
		child.Value.writeJSON(sb, indent+1)
		if i < len(j.OrderedKV)-1 {
			sb.WriteString(",")
		}
	}
	addIndent(sb, indent)
	sb.WriteString("\n}")
}

func (j *OJsonList) writeJSON(sb *strings.Builder, indent int) {
	collection := []OJsonObject(*j)
	sb.WriteString("[\n")
	for i, child := range collection {
		sb.WriteString("\n")
		addIndent(sb, indent+1)
		child.writeJSON(sb, indent+1)
		if i < len(collection)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("\n]")
}

func (j *OJsonString) writeJSON(sb *strings.Builder, indent int) {
	sb.WriteString(string(*j))
}

func (j *OJsonBool) writeJSON(sb *strings.Builder, indent int) {
	sb.WriteString(fmt.Sprintf("%v", bool(*j)))
}
