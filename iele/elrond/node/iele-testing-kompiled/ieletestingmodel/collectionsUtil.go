// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestingmodel

import (
	"fmt"
	"sort"
)

// ExtractKApplyArgs checks that a K item is a KApply and returns its arguments if so
func (ms *ModelState) ExtractKApplyArgs(subject KReference, expectedLabel KLabel, expectedArity int) ([]KReference, bool) {
	kapp, isKapp := ms.GetKApplyObject(subject)
	if !isKapp {
		return nil, false
	}
	if kapp.Label != expectedLabel || len(kapp.List) != expectedArity {
		return nil, false
	}
	return kapp.List, true
}

// ExtractListData checks that a K item is a list and returns its contents if so
func (ms *ModelState) ExtractListData(subject KReference, expectedSort Sort, expectedLabel KLabel) ([]KReference, bool) {
	klist, isList := ms.GetListObject(subject)
	if !isList {
		return nil, false
	}
	if klist.Sort != expectedSort || klist.Label != expectedLabel {
		return nil, false
	}
	return klist.Data, true
}

// ExtractMapData ... checks that a K item is a map and returns its contents if so
// func (ms *ModelState) ExtractMapData(subject KReference, expectedSort Sort, expectedLabel KLabel) (map[KMapKey]KReference, bool) {
// 	kmap, isMap := ms.GetMapObject(subject)
// 	if !isMap {
// 		return nil, false
// 	}
// 	if kmap.Sort != expectedSort || kmap.Label != expectedLabel {
// 		return nil, false
// 	}
// 	return kmap.Data, true
// }

// MapKeyValuePair is just a pair of key and value that was stored in a map
type MapKeyValuePair struct {
	KeyAsString string
	Key         KReference
	Value       KReference
}

// mapOrderedKeyValuePairs yields a list of key-value pairs,
// ordered by the pretty print representation of the keys.
func (ms *ModelState) mapOrderedKeyValuePairs(ref KReference) []MapKeyValuePair {
	refType, _, _, _, _, length := parseKrefCollection(ref)
	if refType != mapRef {
		panic("mapOrderedKeyValuePairs argument not a map")
	}

	result := make([]MapKeyValuePair, 0, int(length))
	var keysAsString []string
	stringKeysToPair := make(map[string]MapKeyValuePair)
	ms.MapForEach(ref, func(k KReference, v KReference) bool {
		keyAsString := ms.PrettyPrint(k)
		keysAsString = append(keysAsString, keyAsString)
		pair := MapKeyValuePair{KeyAsString: keyAsString, Key: k, Value: v}
		stringKeysToPair[keyAsString] = pair
		return false
	})
	if len(keysAsString) != int(length) {
		panic(fmt.Sprintf("map length mismatch. Reference length: %d. Data length: %d",
			length,
			len(keysAsString)))
	}
	sort.Strings(keysAsString)
	for _, keyAsString := range keysAsString {
		pair, _ := stringKeysToPair[keyAsString]
		result = append(result, pair)
	}

	return result
}

// setOrderedElements yields a list of the items in the set,
// ordered by the pretty print representation of the elements
func (ms *ModelState) setOrderedElements(ref KReference) []KReference {
	refType, _, _, _, _, length := parseKrefCollection(ref)
	if refType != setRef {
		panic("setOrderedElements argument not a set")
	}

	result := make([]KReference, 0, int(length))
	var keysAsString []string
	stringKeysToElem := make(map[string]KReference)
	ms.SetForEach(ref, func(elem KReference) bool {
		keyAsString := ms.PrettyPrint(elem)
		keysAsString = append(keysAsString, keyAsString)
		stringKeysToElem[keyAsString] = elem
		return false
	})
	if len(keysAsString) != int(length) {
		panic(fmt.Sprintf("set length mismatch. Reference length: %d. Data length: %d",
			length,
			len(keysAsString)))
	}
	sort.Strings(keysAsString)
	for _, keyAsString := range keysAsString {
		elem, _ := stringKeysToElem[keyAsString]
		result = append(result, elem)
	}

	return result
}

// checks that the real length matches the declared length, only use for debugging
// returns input, fon convenience
func (ms *ModelState) checkCollectionSize(mp KReference) KReference {
	if !ms.collectionSizeOk(mp) {
		panic("collection real length doesn't match declared length")
	}
	return mp
}

func (ms *ModelState) collectionSizeOk(mp KReference) bool {
	refType, _, _, _, _, length := parseKrefCollection(mp)
	if refType == mapRef {
		realLength := uint64(0)
		ms.MapForEach(mp, func(_, _ KReference) bool {
			realLength++
			return false
		})
		if realLength != length {
			return false
		}
	}
	if refType == setRef {
		realLength := uint64(0)
		ms.SetForEach(mp, func(_ KReference) bool {
			realLength++
			return false
		})
		if realLength != length {
			return false
		}
	}

	return true
}
