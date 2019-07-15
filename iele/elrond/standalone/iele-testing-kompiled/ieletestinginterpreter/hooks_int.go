// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:19:23.686

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
	"math/big"
)

type intHooksType int

const intHooks intHooksType = 0

var bigIntZero = big.NewInt(0)
var bigIntOne = big.NewInt(1)

func (intHooksType) eq(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntEquals(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (intHooksType) ne(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntEquals(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(!result), nil
}

func (intHooksType) le(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntLe(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (intHooksType) lt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntLt(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (intHooksType) ge(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntGe(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (intHooksType) gt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntGt(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return m.ToKBool(result), nil
}

func (intHooksType) add(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntAdd(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) sub(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntSub(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) mul(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntMul(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

// Integer division. The result is truncated towards zero and obeys the rule of signs.
func (t intHooksType) tdiv(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if interpreter.Model.IsZero(c2) {
		return m.NoResult, &hookDivisionByZeroError{}
	}
	result, typesOk := interpreter.Model.IntDiv(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

// Integer remainder. The result of rem a b has the sign of a, and its absolute value is strictly smaller than the absolute value of b.
// The result satisfies the equality a = b * div a b + rem a b.
func (t intHooksType) tmod(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if interpreter.Model.IsZero(c2) {
		return m.NoResult, &hookDivisionByZeroError{}
	}
	result, typesOk := interpreter.Model.IntMod(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

// Euclidian division
func (intHooksType) ediv(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if interpreter.Model.IsZero(c2) {
		return m.NoResult, &hookDivisionByZeroError{}
	}
	result, typesOk := interpreter.Model.IntEuclidianDiv(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

// Euclidian remainder
func (intHooksType) emod(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	if interpreter.Model.IsZero(c2) {
		return m.NoResult, &hookDivisionByZeroError{}
	}
	result, typesOk := interpreter.Model.IntEuclidianMod(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) pow(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntPow(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) powmod(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntPowMod(c1, c2, c3)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) shl(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntShl(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) shr(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntShr(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) and(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntAnd(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) or(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntOr(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) xor(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typesOk := interpreter.Model.IntXor(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) not(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typeOk := interpreter.Model.IntNot(c)
	if !typeOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) abs(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typeOk := interpreter.Model.IntAbs(c)
	if !typeOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) max(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	gt, typesOk := interpreter.Model.IntGt(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	if gt {
		return c1, nil
	}
	return c2, nil
}

func (intHooksType) min(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	gt, typesOk := interpreter.Model.IntGt(c1, c2)
	if !typesOk {
		return invalidArgsResult()
	}
	if gt {
		return c2, nil
	}
	return c1, nil
}

func (intHooksType) log2(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	result, typeOk := interpreter.Model.IntLog2(c)
	if !typeOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) bitRange(argI m.KReference, argOffset m.KReference, argLen m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// rule bitRangeInt(I::Int, IDX::Int, LEN::Int) => (I >>Int IDX) modInt (1 <<Int LEN)
	result, typesOk := interpreter.Model.IntBitRange(argI, argOffset, argLen)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) signExtendBitRange(argI m.KReference, argOffset m.KReference, argLen m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	// rule signExtendBitRangeInt(I::Int, IDX::Int, LEN::Int) => (bitRangeInt(I, IDX, LEN) +Int (1 <<Int (LEN -Int 1))) modInt (1 <<Int LEN) -Int (1 <<Int (LEN -Int 1))
	result, typesOk := interpreter.Model.IntSignExtendBitRange(argI, argOffset, argLen)
	if !typesOk {
		return invalidArgsResult()
	}
	return result, nil
}

func (intHooksType) rand(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (intHooksType) srand(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, &hookNotImplementedError{}
}
