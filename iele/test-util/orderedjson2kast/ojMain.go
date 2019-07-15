package orderedjson2kast

import (
	"path/filepath"

	oj "github.com/ElrondNetwork/elrond-vm/iele/test-util/orderedjson"
)

// ConvertOrderedJSONToKast parses data as an ordered JSON,
// assembles code if necessary
// and converts to KAST format, readable by K
func ConvertOrderedJSONToKast(data []byte, testFilePath string) (string, error) {
	jsonObj, err := oj.ParseOrderedJSON(data)
	if err != nil {
		return "", err
	}
	testDirPath := filepath.Dir(testFilePath)
	assembleIele(jsonObj, testDirPath)
	kast := jsonToKastOrdered(jsonObj)

	return kast, nil
}
