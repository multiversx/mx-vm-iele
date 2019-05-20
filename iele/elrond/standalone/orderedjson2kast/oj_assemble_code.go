package orderedjson2kast

import (
	"path"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
)

func (j *jsonMap) assembleIele(testPath string) {
	isCreateTx := false
	for _, keyValuePair := range j.orderedKV {
		if keyValuePair.key == "to" {
			if strVal, isStr := keyValuePair.value.(*jsonString); isStr {
				if string(*strVal) == "" {
					isCreateTx = true
					break
				}
			}
		}
	}

	for _, keyValuePair := range j.orderedKV {
		if keyValuePair.key == "code" ||
			(keyValuePair.key == "contractCode" && isCreateTx) {
			if strVal, isStr := keyValuePair.value.(*jsonString); isStr {
				assembled := assembleIeleCode(testPath, string(*strVal))
				asJObj := jsonString(assembled)
				keyValuePair.value = &asJObj
			}
		} else {
			keyValuePair.value.assembleIele(testPath)
		}
	}
}

func (j *jsonList) assembleIele(testPath string) {
	collection := []jsonObject(*j)
	for _, elem := range collection {
		elem.assembleIele(testPath)
	}
}

func (j *jsonString) assembleIele(testPath string) {
}

func (j *jsonBool) assembleIele(testPath string) {
}

func assembleIeleCode(testPath string, value string) string {
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, "0x") {
		return value
	}

	contractPathFilePath := path.Join(testPath, value)
	return compiler.AssembleIeleCode(contractPathFilePath)
}
