package kryptoadapter

import (
	"errors"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
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
	result, err := k.Upstream.Sha256(str)
	if err != nil {
		return m.NoResult, err
	}
	return ms.NewString(result), nil
}

// Keccak256 adapts between K model and elrond function
func (k *Krypto) Keccak256(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	str, isStr := ms.GetString(c)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.Keccak256(str)
	if err != nil {
		return m.NoResult, err
	}
	return ms.NewString(result), nil
}

// Ripemd160 adapts between K model and elrond function
func (k *Krypto) Ripemd160(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	str, isStr := ms.GetString(c)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.Ripemd160(str)
	if err != nil {
		return m.NoResult, err
	}
	return ms.NewString(result), nil
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
	result, err := k.Upstream.EcdsaRecover(hash, v, r, s)
	if err != nil {
		return m.NoResult, err
	}
	return ms.NewString(result), nil
}

func parseBn128Point(c m.KReference, ms *m.ModelState) (vmi.Bn128Point, error) {
	kapp, isKapp := ms.GetKApplyObject(c)
	if !isKapp {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if kapp.Label != m.LblXlparenXuXcommaXuXrparenXuKRYPTO {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if len(kapp.List) != 2 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x, isInt1 := ms.GetBigInt(kapp.List[0])
	if !isInt1 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y, isInt2 := ms.GetBigInt(kapp.List[1])
	if !isInt2 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	return vmi.Bn128Point{X: x, Y: y}, nil
}

func convertBn128Point(p vmi.Bn128Point, ms *m.ModelState) m.KReference {
	return ms.NewKApply(m.LblXlparenXuXcommaXuXrparenXuKRYPTO,
		ms.FromBigInt(p.X),
		ms.FromBigInt(p.Y))
}

func parseBn128G2Point(c m.KReference, ms *m.ModelState) (vmi.Bn128G2Point, error) {
	kapp, isKapp := ms.GetKApplyObject(c)
	if !isKapp {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if kapp.Label != m.LblXlparenXuxXuXcommaXuxXuXrparenXuKRYPTO {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if len(kapp.List) != 4 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x1, isInt1 := ms.GetBigInt(kapp.List[0])
	if !isInt1 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x2, isInt2 := ms.GetBigInt(kapp.List[1])
	if !isInt2 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y1, isInt3 := ms.GetBigInt(kapp.List[2])
	if !isInt3 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y2, isInt4 := ms.GetBigInt(kapp.List[3])
	if !isInt4 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	return vmi.Bn128G2Point{X1: x1, X2: x2, Y1: y1, Y2: y2}, nil
}

// Bn128valid adapts between K model and elrond function
func (k *Krypto) Bn128valid(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	p, err := parseBn128Point(c, ms)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToKBool(result), nil
}

// Bn128g2valid adapts between K model and elrond function
func (k *Krypto) Bn128g2valid(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	p, err := parseBn128G2Point(c, ms)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128g2valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToKBool(result), nil
}

// Bn128add adapts between K model and elrond function
func (k *Krypto) Bn128add(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	p1, err := parseBn128Point(c1, ms)
	if err != nil {
		return m.NoResult, err
	}
	p2, err := parseBn128Point(c2, ms)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128add(p1, p2)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result, ms), nil
}

// Bn128mul adapts between K model and elrond function
func (k *Krypto) Bn128mul(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	k1, isInt := ms.GetBigInt(c1)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	p, err := parseBn128Point(c2, ms)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128mul(k1, p)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result, ms), nil
}

// Bn128ate adapts between K model and elrond function
func (k *Krypto) Bn128ate(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	l1, isList := ms.GetListObject(c1)
	if !isList {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	l2, isList := ms.GetListObject(c2)
	if !isList {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	var p1List []vmi.Bn128Point
	var p2List []vmi.Bn128G2Point
	for _, k := range l1.Data {
		p, err := parseBn128Point(k, ms)
		if err != nil {
			return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
		}
		p1List = append(p1List, p)
	}
	for _, k := range l2.Data {
		p, err := parseBn128G2Point(k, ms)
		if err != nil {
			return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
		}
		p2List = append(p2List, p)
	}
	result, err := k.Upstream.Bn128ate(p1List, p2List)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToKBool(result), nil
}
