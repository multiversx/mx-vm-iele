package ielecompiler

import (
	"go/build"
	"log"
	"os/exec"
	"path"
)

// AssembleIeleCode ... calls the Haskell compiler to assemble contract code
func AssembleIeleCode(contractPathFilePath string) string {
	cmd := exec.Command("stack", "exec", "iele-assemble", contractPathFilePath)
	cmd.Dir = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/compiler/compiler")
	compiledBytes, err := cmd.Output()
	if err != nil {
		log.Fatal("compile iele error " + err.Error())
	}

	return string(compiledBytes)
}
