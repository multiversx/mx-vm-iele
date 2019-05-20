package orderedjson

import (
	"strings"
)

// OJsonObject ... ordered JSON tree object interface
type OJsonObject interface {
	writeJSON(sb *strings.Builder, indent int)
}

// OJsonKeyValuePair ... since this ir ordered JSON, maps are really ordered lists of key value pairs, this class
type OJsonKeyValuePair struct {
	Key   string
	Value OJsonObject
}

// OJsonMap ... ordered map, actually a list of key value pairs
type OJsonMap struct {
	KeySet    map[string]bool
	OrderedKV []*OJsonKeyValuePair
}

// OJsonList ... JSON list
type OJsonList []OJsonObject

// OJsonString ... JSON string value
type OJsonString string

// OJsonBool ... JSON bool value
type OJsonBool bool

func newMap() *OJsonMap {
	KeySet := make(map[string]bool)
	return &OJsonMap{KeySet: KeySet, OrderedKV: nil}
}

func (j *OJsonMap) put(kvPair *OJsonKeyValuePair) {
	_, alreadyInserted := j.KeySet[kvPair.Key]
	if !alreadyInserted {
		j.KeySet[kvPair.Key] = true
		j.OrderedKV = append(j.OrderedKV, kvPair)
	}
}

// Size ... size of ordered map
func (j *OJsonMap) Size() int {
	return len(j.OrderedKV)
}
