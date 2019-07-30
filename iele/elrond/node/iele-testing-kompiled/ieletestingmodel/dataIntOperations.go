package ieletestingmodel

import (
    "math"
	"math/big"
)

var bigOne = big.NewInt(1)

// helper function for writing operations in fewer lines
func (*ModelState) bothSmall(ref1 KReference, ref2 KReference) (int64, int64, bool) {
	small1, isSmall1 := parseKrefSmallInt(ref1)
	if !isSmall1 {
		return 0, 0, false
	}
	small2, isSmall2 := parseKrefSmallInt(ref2)
	if !isSmall2 {
		return 0, 0, false
	}
	return small1, small2, true
}

// helper function for writing operations in fewer lines
func (ms *ModelState) bothBig(ref1 KReference, ref2 KReference) (*big.Int, *big.Int, bool) {
	big1, isInt1 := ms.GetBigIntUnsafe(ref1)
	if !isInt1 {
		return nil, nil, false
	}
	big2, isInt2 := ms.GetBigIntUnsafe(ref2)
	if !isInt2 {
		return nil, nil, false
	}
	return big1, big2, true
}

// IntEquals returns ref1 == ref2, if types ok.
// Also compares big ints with small ints.
func (ms *ModelState) IntEquals(ref1 KReference, ref2 KReference) (bool, bool) {
	if small1, isSmall1 := parseKrefSmallInt(ref1); isSmall1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// small == small
			return small1 == small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// small == big
			if big2.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return false, true
			}
			if big2.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return false, true
			}
			return small1 == big2.bigValue.Int64(), true
		}
		return false, false
	} else if big1, isBig1 := ms.getBigIntObject(ref1); isBig1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// big == small
			if big1.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return false, true
			}
			if big1.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return false, true
			}
			return big1.bigValue.Int64() == small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// big == big
			return big1.bigValue.Cmp(big2.bigValue) == 0, true
		}
	}
	return false, false
}

// IntGt returns ref1 > ref2, if types ok.
// Also compares big ints with small ints.
func (ms *ModelState) IntGt(ref1 KReference, ref2 KReference) (bool, bool) {
	if small1, isSmall1 := parseKrefSmallInt(ref1); isSmall1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// small > small
			return small1 > small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// small > big
			if big2.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return false, true
			}
			if big2.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return true, true
			}
			return small1 > big2.bigValue.Int64(), true
		}
		return false, false
	} else if big1, isBig1 := ms.getBigIntObject(ref1); isBig1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// big > small
			if big1.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return true, true
			}
			if big1.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return false, true
			}
			return big1.bigValue.Int64() > small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// big > big
			return big1.bigValue.Cmp(big2.bigValue) > 0, true
		}
	}
	return false, false
}

// IntGe returns ref1 >= ref2, if types ok.
// Also compares big ints with small ints.
func (ms *ModelState) IntGe(ref1 KReference, ref2 KReference) (bool, bool) {
	if small1, isSmall1 := parseKrefSmallInt(ref1); isSmall1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// small >= small
			return small1 >= small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// small >= big
			if big2.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return false, true
			}
			if big2.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return true, true
			}
			return small1 >= big2.bigValue.Int64(), true
		}
		return false, false
	} else if big1, isBig1 := ms.getBigIntObject(ref1); isBig1 {
		if small2, isSmall2 := parseKrefSmallInt(ref2); isSmall2 {
			// big >= small
			if big1.bigValue.Cmp(maxSmallIntAsBigInt) > 0 {
				return true, true
			}
			if big1.bigValue.Cmp(minSmallIntAsBigInt) < 0 {
				return false, true
			}
			return big1.bigValue.Int64() >= small2, true
		} else if big2, isBig2 := ms.getBigIntObject(ref2); isBig2 {
			// big >= big
			return big1.bigValue.Cmp(big2.bigValue) >= 0, true
		}
	}
	return false, false
}

// IntLt returns ref1 < ref2, if types ok.
// Also compares big ints with small ints.
func (ms *ModelState) IntLt(ref1 KReference, ref2 KReference) (bool, bool) {
	return ms.IntGt(ref2, ref1)
}

// IntLe returns ref1 <= ref2, if types ok.
// Also compares big ints with small ints.
func (ms *ModelState) IntLe(ref1 KReference, ref2 KReference) (bool, bool) {
	return ms.IntGe(ref2, ref1)
}

// IntAdd returns ref1 + ref2, if types ok
func (ms *ModelState) IntAdd(ref1 KReference, ref2 KReference) (KReference, bool) {
	small1, small2, smallOk := ms.bothSmall(ref1, ref2)
	// TODO: overflow handling
	if smallOk && small1 < math.MaxInt32 && small1 > math.MinInt32 && small2 < math.MaxInt32 && small2 > math.MinInt32 {
		result := int64(small1) + int64(small2)
		if fitsInSmallIntReference(result) {
			return createKrefSmallInt(result), true
		}
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		if big1.Sign() == 0 {
			return ref2, true
		}
		if big2.Sign() == 0 {
			return ref1, true
		}

		ref, obj := ms.newBigIntObject()
		obj.bigValue.Add(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntSub returns ref1 - ref2, if types ok
func (ms *ModelState) IntSub(ref1 KReference, ref2 KReference) (KReference, bool) {
	small1, small2, smallOk := ms.bothSmall(ref1, ref2)
	// TODO: overflow handling
    if smallOk && small1 < math.MaxInt32 && small1 > math.MinInt32 && small2 < math.MaxInt32 && small2 > math.MinInt32 {
		result := int64(small1) - int64(small2)
		if fitsInSmallIntReference(result) {
			return createKrefSmallInt(result), true
		}
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		if big2.Sign() == 0 {
			return ref1, true
		}

		ref, obj := ms.newBigIntObject()
		obj.bigValue.Sub(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntMul returns ref1 x ref2, if types ok
func (ms *ModelState) IntMul(ref1 KReference, ref2 KReference) (KReference, bool) {
	small1, small2, smallOk := ms.bothSmall(ref1, ref2)
	if smallOk && smallMultiplicationSafe(small1, small2) {
		return createKrefSmallInt(small1 * small2), true
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		if big1.Sign() == 0 || big2.Sign() == 0 {
			return IntZero, true
		}
		if big1.Cmp(bigOne) == 0 {
			return ref2, true
		}
		if big2.Cmp(bigOne) == 0 {
			return ref1, true
		}

		ref, obj := ms.newBigIntObject()
		obj.bigValue.Mul(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntDiv performs integer division.
// The result is truncated towards zero and obeys the rule of signs.
func (ms *ModelState) IntDiv(ref1 KReference, ref2 KReference) (KReference, bool) {
	small1, small2, smallOk := ms.bothSmall(ref1, ref2)
	if smallOk {
		return createKrefSmallInt(small1 / small2), true
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		resultPositive := true
		if big1.Sign() < 0 {
			resultPositive = !resultPositive
			big1 = big.NewInt(0).Neg(big1)
		}
		if big2.Sign() < 0 {
			resultPositive = !resultPositive
			big2 = big.NewInt(0).Neg(big2)
		}

		ref, obj := ms.newBigIntObject()
		obj.bigValue.Div(big1, big2)
		if !resultPositive {
			obj.bigValue.Neg(obj.bigValue)
		}
		return ref, true
	}

	return NullReference, false
}

// IntMod performs integer remainder.
// The result of rem a b has the sign of a, and its absolute value is strictly smaller than the absolute value of b.
// The result satisfies the equality a = b * div a b + rem a b.
func (ms *ModelState) IntMod(ref1 KReference, ref2 KReference) (KReference, bool) {
	small1, small2, smallOk := ms.bothSmall(ref1, ref2)
	if smallOk {
		return createKrefSmallInt(small1 % small2), true
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		arg1Negative := false
		if big1.Sign() < 0 {
			arg1Negative = true
			big1 = big.NewInt(0).Neg(big1)
		}
		if big2.Sign() < 0 {
			big2 = big.NewInt(0).Neg(big2)
		}

		ref, obj := ms.newBigIntObject()
		obj.bigValue.Mod(big1, big2)
		if arg1Negative {
			obj.bigValue.Neg(obj.bigValue)
		}
		return ref, true
	}

	return NullReference, false
}

// IntEuclidianDiv performs Euclidian division.
func (ms *ModelState) IntEuclidianDiv(ref1 KReference, ref2 KReference) (KReference, bool) {
	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Div(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntEuclidianMod performs Euclidian remainder.
func (ms *ModelState) IntEuclidianMod(ref1 KReference, ref2 KReference) (KReference, bool) {
	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Mod(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntPow returns ref1 ^ ref2, if types ok
func (ms *ModelState) IntPow(ref1 KReference, ref2 KReference) (KReference, bool) {
	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Exp(big1, big2, nil)
		return ref, true
	}

	return NullReference, false
}

// IntPowMod returns (ref1 ^ ref2) mod ref3, if types ok
func (ms *ModelState) IntPowMod(ref1 KReference, ref2 KReference, ref3 KReference) (KReference, bool) {
	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		big3, big3Ok := ms.GetBigInt(ref3)
		if big3Ok {
			ref, obj := ms.newBigIntObject()
			obj.bigValue.Exp(big1, big2, big3)
			return ref, true
		}
	}

	return NullReference, false
}

// IntShl returns ref1 << ref2, if types ok
func (ms *ModelState) IntShl(ref1 KReference, ref2 KReference) (KReference, bool) {
	arg2, arg2Ok := ms.GetUint(ref2)
	if !arg2Ok {
		return NullReference, false
	}

	arg1, arg1Ok := ms.GetBigInt(ref1)
	if arg1Ok {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Lsh(arg1, arg2)
		return ref, true
	}

	return NullReference, false
}

// IntShr returns ref1 >> ref2, if types ok
func (ms *ModelState) IntShr(ref1 KReference, ref2 KReference) (KReference, bool) {
	arg2, arg2Ok := ms.GetUint(ref2)
	if !arg2Ok {
		return NullReference, false
	}

	arg1, arg1Ok := ms.GetBigInt(ref1)
	if arg1Ok {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Rsh(arg1, arg2)
		return ref, true
	}

	return NullReference, false
}

// IntAnd returns bitwise and, ref1 & ref2, if types ok
func (ms *ModelState) IntAnd(ref1 KReference, ref2 KReference) (KReference, bool) {
	if ms.IsZero(ref1) || ms.IsZero(ref2) {
		return IntZero, true
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.And(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntOr returns bitwise or, ref1 | ref2, if types ok
func (ms *ModelState) IntOr(ref1 KReference, ref2 KReference) (KReference, bool) {
	if ms.IsZero(ref1) {
		return ref2, true
	}
	if ms.IsZero(ref2) {
		return ref1, true
	}

	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Or(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntXor returns bitwise xor, ref1 xor ref2, if types ok
func (ms *ModelState) IntXor(ref1 KReference, ref2 KReference) (KReference, bool) {
	big1, big2, bigOk := ms.bothBig(ref1, ref2)
	if bigOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Xor(big1, big2)
		return ref, true
	}

	return NullReference, false
}

// IntNot returns bitwise not, if type ok
func (ms *ModelState) IntNot(ref KReference) (KReference, bool) {
	arg, argOk := ms.GetBigInt(ref)
	if argOk {
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Not(arg)
		return ref, true
	}

	return NullReference, false
}

// IntAbs returns the absoute value, if type ok
func (ms *ModelState) IntAbs(ref KReference) (KReference, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if small >= 0 {
			return ref, true
		}
		return createKrefSmallInt(-small), true
	}

	bigArg, bigOk := ms.GetBigInt(ref)
	if bigOk {
		if bigArg.Sign() >= 0 {
			return ref, true
		}
		ref, obj := ms.newBigIntObject()
		obj.bigValue.Neg(bigArg)
		return ref, true
	}

	return NullReference, false
}

// IntLog2 basically counts the number of bits after the most significant bit.
// It is equal to a a truncated log2 of the number.
// Argument must be strictly positive.
func (ms *ModelState) IntLog2(ref KReference) (KReference, bool) {
	if small, isSmall := parseKrefSmallInt(ref); isSmall {
		if small <= 0 {
			return NullReference, false
		}
		nrBits := 0
		for small > 0 {
			small = small >> 1
			nrBits++
		}
		return ms.FromInt(nrBits - 1), true
	} else if bigArg, isBig := ms.getBigIntObject(ref); isBig {
		if bigArg.bigValue.Sign() <= 0 {
			return NullReference, false
		}
		nrBits := bigArg.bigValue.BitLen()
		return ms.FromInt(nrBits - 1), true
	}

	return NullReference, false
}
