package orderedjson2kast

import (
	"path"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	oj "github.com/ElrondNetwork/elrond-vm/iele/test-util/orderedjson"
)

func assembleIele(jobj oj.OJsonObject, testPath string) {
	switch j := jobj.(type) {
	case *oj.OJsonMap:
		isCreateTx := false
		for _, keyValuePair := range j.OrderedKV {
			if keyValuePair.Key == "to" {
				if strVal, isStr := keyValuePair.Value.(*oj.OJsonString); isStr {
					if strVal.Value == "" {
						isCreateTx = true
						break
					}
				}
			}
		}

		for _, keyValuePair := range j.OrderedKV {
			if keyValuePair.Key == "code" ||
				(keyValuePair.Key == "contractCode" && isCreateTx) {
				if strVal, isStr := keyValuePair.Value.(*oj.OJsonString); isStr {
					strVal.Value = assembleIeleCode(testPath, strVal.Value)
				}
			} else {
				assembleIele(keyValuePair.Value, testPath)
			}
		}
	case *oj.OJsonList:
		collection := []oj.OJsonObject(*j)
		for _, elem := range collection {
			assembleIele(elem, testPath)
		}
	default:
	}
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
