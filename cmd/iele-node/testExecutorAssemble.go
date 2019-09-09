package main

import (
	"encoding/hex"
	"path/filepath"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
)

// make the tests run faster by not repeating code assembly over and over again
var assembledCodeCache = make(map[string]string)

func assembleIeleCode(testPath string, value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if strings.HasPrefix(value, "0x") {
		code, _ := hex.DecodeString(value[2:])
		return string(code), nil
	}

	contractPathFilePath := filepath.Join(testPath, value)

	cached, foundInCache := assembledCodeCache[contractPathFilePath]
	if foundInCache {
		return cached, nil
	}

	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		return "", err
	}

	result := string(decoded)
	assembledCodeCache[contractPathFilePath] = result
	return result, nil
}
