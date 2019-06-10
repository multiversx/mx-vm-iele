package mockhookcrypto

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"crypto/sha256"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func hashToString(hash []byte) string {
	var sb strings.Builder
	for _, b := range hash {
		sb.WriteString(fmt.Sprintf("%02x", b))
	}
	return sb.String()
}

// KryptoHookMock ... krypto hook implementation that we use for VM tests
type KryptoHookMock int

// KryptoHookMockInstance ... krypto hook mock singleton
const KryptoHookMockInstance KryptoHookMock = 0

// Sha256 ... krypto function
func (KryptoHookMock) Sha256(str string) (string, error) {
	h := sha256.New()
	h.Write([]byte(str))
	byteResult := h.Sum(nil)
	hashStr := hashToString(byteResult)
	return hashStr, nil
}

// Keccak256 ... krypto function
func (KryptoHookMock) Keccak256(str string) (string, error) {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(str))
	hash := h.Sum(nil)
	hashStr := hashToString(hash)
	return hashStr, nil
}

// Ripemd160 ... krypto function
func (KryptoHookMock) Ripemd160(str string) (string, error) {
	md := ripemd160.New()
	_, err := md.Write([]byte(str))
	if err != nil {
		return "", err
	}
	byteResult := md.Sum(nil)
	hashStr := hashToString(byteResult)
	return hashStr, nil
}

// EcdsaRecover ... krypto function
func (KryptoHookMock) EcdsaRecover(hash string, v *big.Int, r string, s string) (string, error) {
	fmt.Println(">>>>> EcdsaRecover")

	return "testRecover", nil
}

// Bn128valid ... krypto function
func (KryptoHookMock) Bn128valid(p vmi.Bn128Point) (bool, error) {
	return false, errors.New("Bn128valid not implemented")
}

// Bn128g2valid ... krypto function
func (KryptoHookMock) Bn128g2valid(p vmi.Bn128G2Point) (bool, error) {
	return false, errors.New("Bn128g2valid not implemented")
}

// Bn128add ... krypto function
func (KryptoHookMock) Bn128add(p1 vmi.Bn128Point, p2 vmi.Bn128Point) (vmi.Bn128Point, error) {
	return vmi.Bn128Point{}, errors.New("Bn128add not implemented")
}

// Bn128mul ... krypto function
func (KryptoHookMock) Bn128mul(k *big.Int, p vmi.Bn128Point) (vmi.Bn128Point, error) {
	return vmi.Bn128Point{}, errors.New("Bn128mul not implemented")
}

// Bn128ate ... krypto function
func (KryptoHookMock) Bn128ate(l1 []vmi.Bn128Point, l2 []vmi.Bn128G2Point) (bool, error) {
	return false, errors.New("Bn128ate not implemented")
}
