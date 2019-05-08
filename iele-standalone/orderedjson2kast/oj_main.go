package orderedjson2kast

import (
	"path/filepath"
)

// ConvertOrderedJSONToKast ... parses data as an ordered JSON,
// assembles code if necessary
// and converts to KAST format, readable by K
func ConvertOrderedJSONToKast(data []byte, testFilePath string) (string, error) {
	jsonObj, err := parseOrderedJSON(data)
	if err != nil {
		return "", err
	}
	testDirPath := filepath.Dir(testFilePath)
	jsonObj.assembleIele(testDirPath)
	kast := jsonToKastOrdered(jsonObj)

	return kast, nil
}
