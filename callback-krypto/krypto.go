package callbackkrypto

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	keccak "github.com/ElrondNetwork/elrond-go-sandbox/hashing/keccak"
	sha256 "github.com/ElrondNetwork/elrond-go-sandbox/hashing/sha256"
	"golang.org/x/crypto/ripemd160"
)

func hashToString(hash []byte) string {
	var sb strings.Builder
	for _, b := range hash {
		sb.WriteString(fmt.Sprintf("%02x", b))
	}
	return sb.String()
}

// Sha256 ... krypto function
func Sha256(str string) (string, error) {
	hasher := sha256.Sha256{}
	byteResult := hasher.Compute(str)
	hashStr := hashToString(byteResult)
	return hashStr, nil
}

// Keccak256 ... krypto function
func Keccak256(str string) (string, error) {
	var k keccak.Keccak
	hash := k.Compute(str)
	hashStr := hashToString(hash)
	return hashStr, nil
}

// Ripemd160 ... krypto function
func Ripemd160(str string) (string, error) {
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
func EcdsaRecover(hash string, v *big.Int, r string, s string) (string, error) {
	fmt.Println(">>>>> EcdsaRecover")

	return "testRecover", nil
}

// Bn128Point ... point on a curve
type Bn128Point struct {
	X *big.Int
	Y *big.Int
}

// Bn128G2Point ... point on a curve
type Bn128G2Point struct {
	X1 *big.Int
	X2 *big.Int
	Y1 *big.Int
	Y2 *big.Int
}

// Bn128valid ... krypto function
func Bn128valid(p Bn128Point) (bool, error) {
	return false, errors.New("Bn128valid not implemented")
}

// Bn128g2valid ... krypto function
func Bn128g2valid(p Bn128G2Point) (bool, error) {
	return false, errors.New("Bn128g2valid not implemented")
}

// Bn128add ... krypto function
func Bn128add(p1 Bn128Point, p2 Bn128Point) (Bn128Point, error) {
	return Bn128Point{}, errors.New("Bn128add not implemented")
}

// Bn128mul ... krypto function
func Bn128mul(k *big.Int, p Bn128Point) (Bn128Point, error) {
	return Bn128Point{}, errors.New("Bn128mul not implemented")
}

// Bn128ate ... krypto function
func Bn128ate(l1 []Bn128Point, l2 []Bn128G2Point) (bool, error) {
	return false, errors.New("Bn128ate not implemented")
}
