package orderedjson2kast

import (
	"strings"
)

// OJsonObject ... ordered JSON tree object interface
type OJsonObject interface {
	writeJSON(sb *strings.Builder, indent int)
	writeKast(sb *strings.Builder)
	assembleIele(testPath string)
}

// OJsonKeyValuePair ... since this ir ordered JSON, maps are really ordered lists of key value pairs, this class
type OJsonKeyValuePair struct {
	key   string
	value OJsonObject
}

// OJsonMap ... ordered map, actually a list of key value pairs
type OJsonMap struct {
	keySet    map[string]bool
	orderedKV []*OJsonKeyValuePair
}

// OJsonList ... JSON list
type OJsonList []OJsonObject

// OJsonString ... JSON string value
type OJsonString string

// OJsonBool ... JSON bool value
type OJsonBool bool

func newMap() *OJsonMap {
	keySet := make(map[string]bool)
	return &OJsonMap{keySet: keySet, orderedKV: nil}
}

func (j *OJsonMap) put(kvPair *OJsonKeyValuePair) {
	_, alreadyInserted := j.keySet[kvPair.key]
	if !alreadyInserted {
		j.keySet[kvPair.key] = true
		j.orderedKV = append(j.orderedKV, kvPair)
	}
}

func (j *OJsonMap) size() int {
	return len(j.orderedKV)
}
