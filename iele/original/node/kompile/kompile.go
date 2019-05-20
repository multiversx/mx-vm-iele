package main

// Tool to more easily generate IELE sources.
// Works by calling the K framework.
// For this to work, K framework needs to be installed (https://github.com/ElrondNetwork/k)
// Also, "<K location>/k/k-distribution/target/release/k/bin/kompile" needs to be in the PATH variable.

import (
	"fmt"
	"go/build"
	"log"
	"os/exec"
	"path"
)

func main() {
	hookPackages := "github.com/ElrondNetwork/elrond-vm/iele/original/node/hookadapter/krypto " +
		"github.com/ElrondNetwork/elrond-vm/iele/original/node/hookadapter/blockchain"
	goSrcPath := path.Join(build.Default.GOPATH, "src")

	fmt.Println("Starting to generate iele-node go sources ...")
	cmd := exec.Command("kompile", "iele-testing.k",
		"--backend", "go",
		"--main-module", "IELE-TESTING",
		"--syntax-module", "IELE-SYNTAX",
		"--go-hook-packages", fmt.Sprintf("\"%s\"", hookPackages),
		"--go-src-only",
		"--go-src-path", fmt.Sprintf("\"%s\"", goSrcPath),
	)
	cmd.Dir = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/original/node")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("kompile iele-node error " + err.Error())
	}
	fmt.Println(string(output))
	fmt.Println("Done generating iele-node go sources.")
}
