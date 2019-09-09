package main

import (
	"path/filepath"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
)

func assembleIeleCode(testPath string, value string) string {
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, "0x") {
		return value
	}

	contractPathFilePath := filepath.Join(testPath, value)
	return compiler.AssembleIeleCode(contractPathFilePath)
}
