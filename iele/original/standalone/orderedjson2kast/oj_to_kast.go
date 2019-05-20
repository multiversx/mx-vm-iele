package orderedjson2kast

import (
	"fmt"
	"strings"
)

func jsonToKastOrdered(j jsonObject) string {
	var sb strings.Builder
	j.writeKast(&sb)
	return sb.String()
}

func (j *jsonString) writeKast(sb *strings.Builder) {
	value := string(*j)
	writeStringKast(sb, value)
}

func writeStringKast(sb *strings.Builder, value string) {
	sb.WriteString(fmt.Sprintf("#token(\"\\\"%s\\\"\",\"String\")", value))
}

func (j *jsonBool) writeKast(sb *strings.Builder) {
	value := bool(*j)
	sb.WriteString(fmt.Sprintf("#token(\"%t\",\"Bool\")", value))
}

func (j *jsonMap) writeKast(sb *strings.Builder) {

	sb.WriteString("`{_}_IELE-DATA`(")
	for _, keyValuePair := range j.orderedKV {
		sb.WriteString("`_,__IELE-DATA`(`_:__IELE-DATA`(")
		writeStringKast(sb, keyValuePair.key)
		sb.WriteString(",")
		keyValuePair.value.writeKast(sb)
		sb.WriteString("),")
	}
	sb.WriteString("`.List{\"_,__IELE-DATA\"}`(.KList)")
	for i := 0; i < j.size(); i++ {
		sb.WriteString(")")
	}
	sb.WriteString(")")
}

func (j *jsonList) writeKast(sb *strings.Builder) {
	collection := []jsonObject(*j)

	sb.WriteString("`[_]_IELE-DATA`(")
	for _, elem := range collection {
		sb.WriteString("`_,__IELE-DATA`(")
		elem.writeKast(sb)
		sb.WriteString(",")
	}
	sb.WriteString("`.List{\"_,__IELE-DATA\"}`(.KList)")
	for i := 0; i < len(collection); i++ {
		sb.WriteString(")")
	}
	sb.WriteString(")")
}
