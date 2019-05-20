package orderedjson2kast

import (
	"path"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
)

func (j *OJsonMap) assembleIele(testPath string) {
	isCreateTx := false
	for _, keyValuePair := range j.orderedKV {
		if keyValuePair.key == "to" {
			if strVal, isStr := keyValuePair.value.(*OJsonString); isStr {
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
			if strVal, isStr := keyValuePair.value.(*OJsonString); isStr {
				assembled := assembleIeleCode(testPath, string(*strVal))
				asJObj := OJsonString(assembled)
				keyValuePair.value = &asJObj
			}
		} else {
			keyValuePair.value.assembleIele(testPath)
		}
	}
}

func (j *OJsonList) assembleIele(testPath string) {
	collection := []OJsonObject(*j)
	for _, elem := range collection {
		elem.assembleIele(testPath)
	}
}

func (j *OJsonString) assembleIele(testPath string) {
}

func (j *OJsonBool) assembleIele(testPath string) {
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
