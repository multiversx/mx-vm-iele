package kryptoadapter

import (
	"encoding/hex"
	"errors"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// Krypto is an adapter between K and the outside world
type Krypto struct {
	Upstream vmi.CryptoHook
}

// Sha256 adapts between K model and elrond function
func (k *Krypto) Sha256(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	str, isStr := ms.GetString(c)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	byteResult, err := k.Upstream.Sha256([]byte(str))
	if err != nil {
		return m.NoResult, err
	}
	hashStr := hex.EncodeToString(byteResult)
	return ms.NewString(hashStr), nil
}

// Keccak256 adapts between K model and elrond function
func (k *Krypto) Keccak256(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	str, isStr := ms.GetString(c)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	byteResult, err := k.Upstream.Keccak256([]byte(str))
	if err != nil {
		return m.NoResult, err
	}
	hashStr := hex.EncodeToString(byteResult)
	return ms.NewString(hashStr), nil
}

// Ripemd160 adapts between K model and elrond function
func (k *Krypto) Ripemd160(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	str, isStr := ms.GetString(c)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	byteResult, err := k.Upstream.Ripemd160([]byte(str))
	if err != nil {
		return m.NoResult, err
	}
	hashStr := hex.EncodeToString(byteResult)
	return ms.NewString(hashStr), nil
}

// EcdsaRecover adapts between K model and elrond function
func (k *Krypto) EcdsaRecover(c1 m.KReference, c2 m.KReference, c3 m.KReference, c4 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	hash, hashOk := ms.GetString(c1)
	if !hashOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	v, vOk := ms.GetBigInt(c2)
	if !vOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	r, rOk := ms.GetString(c3)
	if !rOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	s, sOk := ms.GetString(c4)
	if !sOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	byteResult, err := k.Upstream.Ecrecover([]byte(hash), v.Bytes(), []byte(r), []byte(s))
	if err != nil {
		return m.NoResult, err
	}
	hashStr := hex.EncodeToString(byteResult)
	return ms.NewString(hashStr), nil
}

// Bn128valid adapts between K model and elrond function
func (k *Krypto) Bn128valid(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	return m.NoResult, errors.New("Bn128valid not implemented")
}

// Bn128g2valid adapts between K model and elrond function
func (k *Krypto) Bn128g2valid(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	return m.NoResult, errors.New("Bn128g2valid not implemented")
}

// Bn128add adapts between K model and elrond function
func (k *Krypto) Bn128add(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	return m.NoResult, errors.New("Bn128add not implemented")
}

// Bn128mul adapts between K model and elrond function
func (k *Krypto) Bn128mul(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	return m.NoResult, errors.New("Bn128mul not implemented")
}

// Bn128ate adapts between K model and elrond function
func (k *Krypto) Bn128ate(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	return m.NoResult, errors.New("Bn128ate not implemented")
}
