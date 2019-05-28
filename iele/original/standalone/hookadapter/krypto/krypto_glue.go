package kryptoadapter

import (
	"errors"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

// Krypto is an adapter between K and the outside world
type Krypto struct {
	Upstream vmi.CryptoHook
}

// Sha256 adapts between K model and elrond function
func (k *Krypto) Sha256(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.Sha256(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// Keccak256 adapts between K model and elrond function
func (k *Krypto) Keccak256(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.Keccak256(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// Ripemd160 adapts between K model and elrond function
func (k *Krypto) Ripemd160(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.Ripemd160(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// EcdsaRecover adapts between K model and elrond function
func (k *Krypto) EcdsaRecover(c1 m.K, c2 m.K, c3 m.K, c4 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	hash, hashOk := c1.(*m.String)
	if !hashOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	v, vOk := c2.(*m.Int)
	if !vOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	r, rOk := c3.(*m.String)
	if !rOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	s, sOk := c4.(*m.String)
	if !sOk {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	result, err := k.Upstream.EcdsaRecover(hash.Value, v.Value, r.Value, s.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

func parseBn128Point(c m.K) (vmi.Bn128Point, error) {
	kapp, isKapp := c.(*m.KApply)
	if !isKapp {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if kapp.Label != m.LblXlparenXuXcommaXuXrparenXuKRYPTO {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if len(kapp.List) != 2 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x, isInt1 := kapp.List[0].(*m.Int)
	if !isInt1 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y, isInt2 := kapp.List[1].(*m.Int)
	if !isInt2 {
		return vmi.Bn128Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	return vmi.Bn128Point{X: x.Value, Y: y.Value}, nil
}

func convertBn128Point(p vmi.Bn128Point) m.K {
	return &m.KApply{Label: m.LblXlparenXuXcommaXuXrparenXuKRYPTO, List: []m.K{
		m.NewInt(p.X),
		m.NewInt(p.Y)}}
}

func parseBn128G2Point(c m.K) (vmi.Bn128G2Point, error) {
	kapp, isKapp := c.(*m.KApply)
	if !isKapp {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if kapp.Label != m.LblXlparenXuxXuXcommaXuxXuXrparenXuKRYPTO {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	if len(kapp.List) != 4 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x1, isInt1 := kapp.List[0].(*m.Int)
	if !isInt1 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	x2, isInt2 := kapp.List[1].(*m.Int)
	if !isInt2 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y1, isInt3 := kapp.List[0].(*m.Int)
	if !isInt3 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	y2, isInt4 := kapp.List[1].(*m.Int)
	if !isInt4 {
		return vmi.Bn128G2Point{}, errors.New("invalid argument(s) provided to krypto hook")
	}
	return vmi.Bn128G2Point{X1: x1.Value, X2: x2.Value, Y1: y1.Value, Y2: y2.Value}, nil
}

// Bn128valid adapts between K model and elrond function
func (k *Krypto) Bn128valid(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p, err := parseBn128Point(c)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// Bn128g2valid adapts between K model and elrond function
func (k *Krypto) Bn128g2valid(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p, err := parseBn128G2Point(c)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128g2valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// Bn128add adapts between K model and elrond function
func (k *Krypto) Bn128add(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p1, err := parseBn128Point(c1)
	if err != nil {
		return m.NoResult, err
	}
	p2, err := parseBn128Point(c2)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128add(p1, p2)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result), nil
}

// Bn128mul adapts between K model and elrond function
func (k *Krypto) Bn128mul(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, isInt := c1.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	p, err := parseBn128Point(c2)
	if err != nil {
		return m.NoResult, err
	}
	result, err := k.Upstream.Bn128mul(k1.Value, p)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result), nil
}

// Bn128ate adapts between K model and elrond function
func (k *Krypto) Bn128ate(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	l1, isList := c1.(*m.List)
	if !isList {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	l2, isList := c2.(*m.List)
	if !isList {
		return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
	}
	var p1List []vmi.Bn128Point
	var p2List []vmi.Bn128G2Point
	for _, k := range l1.Data {
		p, err := parseBn128Point(k)
		if err != nil {
			return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
		}
		p1List = append(p1List, p)
	}
	for _, k := range l2.Data {
		p, err := parseBn128G2Point(k)
		if err != nil {
			return m.NoResult, errors.New("invalid argument(s) provided to krypto hook")
		}
		p2List = append(p2List, p)
	}
	result, err := k.Upstream.Bn128ate(p1List, p2List)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}
