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
type OJsonString struct {
	Value string
}

// OJsonBool ... JSON bool value
type OJsonBool bool

// NewMap ... create new ordered "map" instance
func NewMap() *OJsonMap {
	KeySet := make(map[string]bool)
	return &OJsonMap{KeySet: KeySet, OrderedKV: nil}
}

// Put ... put into map. Nothing if key exists in map
func (j *OJsonMap) Put(key string, value OJsonObject) {
	_, alreadyInserted := j.KeySet[key]
	if !alreadyInserted {
		j.KeySet[key] = true
		keyValuePair := &OJsonKeyValuePair{Key: key, Value: value}
		j.OrderedKV = append(j.OrderedKV, keyValuePair)
	}
}

// Size ... size of ordered map
func (j *OJsonMap) Size() int {
	return len(j.OrderedKV)
}

// RefreshKeySet ... recreate the key set from the key value pairs
func (j *OJsonMap) RefreshKeySet() {
	j.KeySet = make(map[string]bool)
	for _, kv := range j.OrderedKV {
		j.KeySet[kv.Key] = true
	}
}

// AsList ... returns it represented as a slice of objects
func (j *OJsonList) AsList() []OJsonObject {
	return []OJsonObject(*j)
}
