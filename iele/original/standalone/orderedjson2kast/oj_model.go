package orderedjson2kast

import (
	"fmt"
	"strings"
)

type jsonObject interface {
	jsonString() string
	writeKast(sb *strings.Builder)
	assembleIele(testPath string)
}

type jsonKeyValuePair struct {
	key   string
	value jsonObject
}

type jsonMap struct {
	keySet    map[string]bool
	orderedKV []*jsonKeyValuePair
}
type jsonList []jsonObject
type jsonString string
type jsonBool bool

func newMap() *jsonMap {
	keySet := make(map[string]bool)
	return &jsonMap{keySet: keySet, orderedKV: nil}
}

func (j *jsonMap) put(kvPair *jsonKeyValuePair) {
	_, alreadyInserted := j.keySet[kvPair.key]
	if !alreadyInserted {
		j.keySet[kvPair.key] = true
		j.orderedKV = append(j.orderedKV, kvPair)
	}
}

func (j *jsonMap) size() int {
	return len(j.orderedKV)
}

func (j *jsonMap) jsonString() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, child := range j.orderedKV {
		sb.WriteString(child.key)
		sb.WriteString(":")
		sb.WriteString(child.value.jsonString())
		if i < j.size()-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (j *jsonList) jsonString() string {
	var sb strings.Builder
	collection := []jsonObject(*j)
	sb.WriteString("[")
	for i, child := range collection {
		sb.WriteString(child.jsonString())
		if i < len(collection)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (j *jsonString) jsonString() string {
	return string(*j)
}

func (j *jsonBool) jsonString() string {
	return fmt.Sprintf("%v", bool(*j))
}
