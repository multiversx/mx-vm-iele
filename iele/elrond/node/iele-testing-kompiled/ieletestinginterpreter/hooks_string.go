// File provided by the K Framework Go backend. Timestamp: 2019-05-21 00:58:51.823

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"math/big"
    "strconv"
    "strings"
)

type stringHooksType int

const stringHooks stringHooksType = 0

func (stringHooksType) concat(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.NewString(k1.String() + k2.String()), nil
}

func (stringHooksType) lt(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() < k2.String()), nil
}

func (stringHooksType) le(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() <= k2.String()), nil
}

func (stringHooksType) gt(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() > k2.String()), nil
}

func (stringHooksType) ge(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() >= k2.String()), nil
}

func (stringHooksType) eq(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() == k2.String()), nil
}

func (stringHooksType) ne(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k1, ok1 := c1.(*m.String)
	k2, ok2 := c2.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return m.ToBool(k1.String() != k2.String()), nil
}

func (stringHooksType) chr(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	i, ok := c.(*m.Int)
	if !ok {
		return invalidArgsResult()
	}

	b := byte(i.Value.Uint64())
	bytes := []byte{b}
	return m.NewString(string(bytes)), nil
}

func (stringHooksType) find(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, ok1 := c1.(*m.String)
	substr, ok2 := c2.(*m.String)
	firstIdx, ok3 := c3.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if !firstIdx.Value.IsUint64() {
		return invalidArgsResult()
	}
	firstIdxInt := firstIdx.Value.Uint64()
	if firstIdxInt > uint64(len(str.String())) {
		return invalidArgsResult()
	}

	result := strings.Index(str.String()[firstIdxInt:], substr.String())
	if result == -1 {
		return m.IntMinusOne, nil
	}
	return m.NewIntFromUint64(firstIdxInt + uint64(result)), nil
}

func (stringHooksType) rfind(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, ok1 := c1.(*m.String)
	substr, ok2 := c2.(*m.String)
	lastIdx, ok3 := c3.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if !lastIdx.Value.IsUint64() {
		return invalidArgsResult()
	}
	lastIdxInt := lastIdx.Value.Uint64()
	if lastIdxInt > uint64(len(str.String())) {
		return invalidArgsResult()
	}
	result := strings.LastIndex(str.String()[0:lastIdxInt], substr.String())
	if result == -1 {
		return m.IntMinusOne, nil
	}
	return m.NewIntFromInt(result), nil
}

func (stringHooksType) length(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k, ok := c.(*m.String)
	if !ok {
		return invalidArgsResult()
	}
	return m.NewIntFromInt(len(k.String())), nil
}

func (stringHooksType) substr(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, ok1 := c1.(*m.String)
	from, ok2 := c2.(*m.Int) // from is inclusive
	to, ok3 := c3.(*m.Int)   // to is exclusive
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if !from.Value.IsUint64() || !to.Value.IsUint64() {
		return invalidArgsResult()
	}
	fromInt := from.Value.Uint64()
	toInt := to.Value.Uint64()
	length := uint64(len(str.String()))
	if fromInt > toInt || fromInt > length {
		return invalidArgsResult()
	}
	if toInt > length {
		toInt = length
	}
	return m.NewString(str.String()[fromInt:toInt]), nil
}

func (stringHooksType) ord(arg m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, ok := arg.(*m.String)
	if !ok {
		return invalidArgsResult()
	}
	asBytes := []byte(str.Value)
	if len(asBytes) == 0 {
		// TODO: HACK!!!!
		// correct implementation should throw invalidArgsResult()
		// fix after implementing lazy evaluation of && in 'requires' part of rule
		return m.IntZero, nil
	}
	return m.NewIntFromByte(asBytes[0]), nil
}

func (stringHooksType) int2string(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	i, ok := c.(*m.Int)
	if !ok {
		return invalidArgsResult()
	}
	return m.NewString(i.Value.String()), nil
}

func (stringHooksType) string2int(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) string2base(kstr m.K, kbase m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	str, ok1 := kstr.(*m.String)
	base, ok2 := kbase.(*m.Int)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !base.Value.IsUint64() {
		return invalidArgsResult()
	}
	baseVal := base.Value.Uint64()
	if baseVal < 2 || baseVal > 16 {
		return invalidArgsResult()
	}
	i := new(big.Int)
	var parseOk bool
	i, parseOk = i.SetString(str.Value, int(baseVal))
	if !parseOk {
		return invalidArgsResult()
	}
	return m.NewInt(i), nil
}

func (stringHooksType) base2string(kint m.K, kbase m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	i, ok1 := kint.(*m.Int)
	base, ok2 := kbase.(*m.Int)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !base.Value.IsUint64() {
		return invalidArgsResult()
	}
	baseVal := base.Value.Uint64()
	if baseVal < 2 || baseVal > 16 {
		return invalidArgsResult()
	}
	str := i.Value.Text(int(baseVal))
	return m.NewString(str), nil
}

func (stringHooksType) string2token(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	k, ok := c.(*m.String)
	if !ok {
		return invalidArgsResult()
	}
	return &m.KToken{Sort: sort, Value: k.String()}, nil
}

func (stringHooksType) token2string(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	if k, typeOk := c.(*m.KToken); typeOk {
		return m.NewString(k.Value), nil
	}
	if k, typeOk := c.(*m.Bool); typeOk {
		return m.NewString(strconv.FormatBool(k.Value)), nil
	}
	if k, typeOk := c.(*m.String); typeOk {
		return k, nil // TODO: should do escaping
	}
	if k, typeOk := c.(*m.Int); typeOk {
		return m.NewString(k.Value.String()), nil
	}
	if _, typeOk := c.(*m.Float); typeOk {
		return m.NoResult, &hookNotImplementedError{}
	}

	return invalidArgsResult()
}

func (stringHooksType) float2string(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) uuid(lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) floatFormat(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) string2float(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) replace(argS m.K, argToReplace m.K, argReplacement m.K, argCount m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	kS, ok1 := argS.(*m.String)
	kToReplace, ok2 := argToReplace.(*m.String)
	kReplacement, ok3 := argReplacement.(*m.String)
	kCount, ok4 := argCount.(*m.Int)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return invalidArgsResult()
	}
	count, countOk := kCount.ToInt32()
	if !countOk {
		return invalidArgsResult()
	}

	result := strings.Replace(kS.Value, kToReplace.Value, kReplacement.Value, count)
	return m.NewString(result), nil
}

func (stringHooksType) replaceAll(argS m.K, argToReplace m.K, argReplacement m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	kS, ok1 := argS.(*m.String)
	kToReplace, ok2 := argToReplace.(*m.String)
	kReplacement, ok3 := argReplacement.(*m.String)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if kS.IsEmpty() {
		return kS, nil
	}
	count := strings.Count(kS.Value, kToReplace.Value)
	if count == 0 {
		return kS, nil
	}
	if count == 1 && strings.HasPrefix(kS.Value, kToReplace.Value) {
		if kReplacement.IsEmpty() {
			// just cut off the prefix
			return m.NewString(kS.Value[len(kToReplace.Value):]), nil
		}
		return m.NewString(kReplacement.Value + kS.Value[len(kToReplace.Value):]), nil
	}

	result := strings.ReplaceAll(kS.Value, kToReplace.Value, kReplacement.Value)
	return m.NewString(result), nil
}

func (stringHooksType) replaceFirst(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) countAllOccurrences(argS m.K, argToCount m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	kS, ok1 := argS.(*m.String)
	kToCount, ok2 := argToCount.(*m.String)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}

	result := strings.Count(kS.Value, kToCount.Value)
	return m.NewIntFromInt(result), nil
}

func (stringHooksType) category(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) directionality(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) findChar(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (stringHooksType) rfindChar(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}
