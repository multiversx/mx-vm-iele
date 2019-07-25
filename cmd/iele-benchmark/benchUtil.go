package main

import (
	"encoding/hex"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

var lastReturnCode vmi.ReturnCode

func hexToBigInt(hexRepresentation string) *big.Int {
	result, ok := big.NewInt(0).SetString(hexRepresentation, 16)
	if !ok {
		panic("invalid hex: " + hexRepresentation)
	}
	return result
}

func storageKey(hexRepresentation string) string {
	decoded, err := hex.DecodeString(hexRepresentation)
	if err != nil {
		panic("invalid hex: " + hexRepresentation)
	}
	return string(decoded)
}
