package main

import (
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	oj "github.com/ElrondNetwork/elrond-vm-util/test-util/orderedjson"
)

var ielePath = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/")
var originalTestPath = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/original/tests/")
var elrondTestPath = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/elrond/tests/")

func main() {
	origPath := path.Join(originalTestPath, "tests/iele")
	processIeleTests(origPath)
}

func processIeleTests(origPath string) {
	err := filepath.Walk(origPath, func(testFilePath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(testFilePath, ".iele.json") {
			toPath := correspondingDestination(testFilePath)
			return readModifySave(testFilePath, toPath)
		} else if strings.HasSuffix(testFilePath, ".iele") {
			toPath := correspondingDestination(testFilePath)
			return copyFile(testFilePath, toPath)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func correspondingDestination(fromPath string) string {
	if !strings.HasPrefix(fromPath, originalTestPath) {
		panic("invalid source path")
	}
	return elrondTestPath + fromPath[len(originalTestPath):]
}

func copyFile(fromPath string, toPath string) error {
	source, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer source.Close()

	err = os.MkdirAll(filepath.Dir(toPath), os.ModePerm)
	if err != nil {
		return err
	}
	destination, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func readModifySave(fromPath string, toPath string) error {
	jsonFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	jobj, err := oj.ParseOrderedJSON(byteValue)
	if err != nil {
		return err
	}

	err = modifyJSON(jobj)
	if err != nil {
		return err
	}

	resultJSON := oj.JSONString(jobj)

	err = os.MkdirAll(filepath.Dir(toPath), os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(toPath, []byte(resultJSON), 0644)
	if err != nil {
		return err
	}

	return nil
}
