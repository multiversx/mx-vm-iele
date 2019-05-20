package kryptoadapter

import (
	krypto "github.com/ElrondNetwork/elrond-vm/callback-krypto"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

// Sha256 ... adapts between K model and elrond function
func Sha256(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := krypto.Sha256(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// Keccak256 ... adapts between K model and elrond function
func Keccak256(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := krypto.Keccak256(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// Ripemd160 ... adapts between K model and elrond function
func Ripemd160(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, isStr := c.(*m.String)
	if !isStr {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := krypto.Ripemd160(str.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// EcdsaRecover ... adapts between K model and elrond function
func EcdsaRecover(c1 m.K, c2 m.K, c3 m.K, c4 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	hash, hashOk := c1.(*m.String)
	if !hashOk {
		return m.NoResult, &hookInvalidArgsError{}
	}
	v, vOk := c2.(*m.Int)
	if !vOk {
		return m.NoResult, &hookInvalidArgsError{}
	}
	r, rOk := c3.(*m.String)
	if !rOk {
		return m.NoResult, &hookInvalidArgsError{}
	}
	s, sOk := c4.(*m.String)
	if !sOk {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := krypto.EcdsaRecover(hash.Value, v.Value, r.Value, s.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

func parseBn128Point(c m.K) (krypto.Bn128Point, error) {
	kapp, isKapp := c.(*m.KApply)
	if !isKapp {
		return krypto.Bn128Point{}, &hookInvalidArgsError{}
	}
	if kapp.Label != m.LblXlparenXuXcommaXuXrparenXuKRYPTO {
		return krypto.Bn128Point{}, &hookInvalidArgsError{}
	}
	if len(kapp.List) != 2 {
		return krypto.Bn128Point{}, &hookInvalidArgsError{}
	}
	x, isInt1 := kapp.List[0].(*m.Int)
	if !isInt1 {
		return krypto.Bn128Point{}, &hookInvalidArgsError{}
	}
	y, isInt2 := kapp.List[1].(*m.Int)
	if !isInt2 {
		return krypto.Bn128Point{}, &hookInvalidArgsError{}
	}
	return krypto.Bn128Point{X: x.Value, Y: y.Value}, nil
}

func convertBn128Point(p krypto.Bn128Point) m.K {
	return &m.KApply{Label: m.LblXlparenXuXcommaXuXrparenXuKRYPTO, List: []m.K{
		m.NewInt(p.X),
		m.NewInt(p.Y)}}
}

func parseBn128G2Point(c m.K) (krypto.Bn128G2Point, error) {
	kapp, isKapp := c.(*m.KApply)
	if !isKapp {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	if kapp.Label != m.LblXlparenXuxXuXcommaXuxXuXrparenXuKRYPTO {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	if len(kapp.List) != 4 {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	x1, isInt1 := kapp.List[0].(*m.Int)
	if !isInt1 {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	x2, isInt2 := kapp.List[1].(*m.Int)
	if !isInt2 {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	y1, isInt3 := kapp.List[0].(*m.Int)
	if !isInt3 {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	y2, isInt4 := kapp.List[1].(*m.Int)
	if !isInt4 {
		return krypto.Bn128G2Point{}, &hookInvalidArgsError{}
	}
	return krypto.Bn128G2Point{X1: x1.Value, X2: x2.Value, Y1: y1.Value, Y2: y2.Value}, nil
}

// Bn128valid ... adapts between K model and elrond function
func Bn128valid(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p, err := parseBn128Point(c)
	if err != nil {
		return m.NoResult, err
	}
	result, err := krypto.Bn128valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// Bn128g2valid ... adapts between K model and elrond function
func Bn128g2valid(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p, err := parseBn128G2Point(c)
	if err != nil {
		return m.NoResult, err
	}
	result, err := krypto.Bn128g2valid(p)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// Bn128add ... adapts between K model and elrond function
func Bn128add(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	p1, err := parseBn128Point(c1)
	if err != nil {
		return m.NoResult, err
	}
	p2, err := parseBn128Point(c2)
	if err != nil {
		return m.NoResult, err
	}
	result, err := krypto.Bn128add(p1, p2)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result), nil
}

// Bn128mul ... adapts between K model and elrond function
func Bn128mul(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k, isInt := c1.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	p, err := parseBn128Point(c2)
	if err != nil {
		return m.NoResult, err
	}
	result, err := krypto.Bn128mul(k.Value, p)
	if err != nil {
		return m.NoResult, err
	}
	return convertBn128Point(result), nil
}

// Bn128ate ... adapts between K model and elrond function
func Bn128ate(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	l1, isList := c1.(*m.List)
	if !isList {
		return m.NoResult, &hookInvalidArgsError{}
	}
	l2, isList := c2.(*m.List)
	if !isList {
		return m.NoResult, &hookInvalidArgsError{}
	}
	var p1List []krypto.Bn128Point
	var p2List []krypto.Bn128G2Point
	for _, k := range l1.Data {
		p, err := parseBn128Point(k)
		if err != nil {
			return m.NoResult, &hookInvalidArgsError{}
		}
		p1List = append(p1List, p)
	}
	for _, k := range l2.Data {
		p, err := parseBn128G2Point(k)
		if err != nil {
			return m.NoResult, &hookInvalidArgsError{}
		}
		p2List = append(p2List, p)
	}
	result, err := krypto.Bn128ate(p1List, p2List)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

type hookInvalidArgsError struct {
}

func (e *hookInvalidArgsError) Error() string {
	return "Invalid argument(s) provided to krypto hook."
}
