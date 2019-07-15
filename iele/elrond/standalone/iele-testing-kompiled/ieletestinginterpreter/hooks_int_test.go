// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:19:23.686

package ieletestinginterpreter

import (
	"fmt"
    m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
    "math"
	"testing"
)

func TestParseIntOk(t *testing.T) {
	interpreter := newTestInterpreter()
	strs := []string{
		"0", "123",
		"-123",
		"57896044618658097711785492504343953926634992332820282019728792003956564819968",
		"-57896044618658097711785492504343953926634992332820282019728792003956564819968"}
	for _, s := range strs {
		i, err := interpreter.Model.ParseInt(s)
		assertIntOk(t, s, i, err, interpreter)
	}
}

func TestParseIntError(t *testing.T) {
	interpreter := newTestInterpreter()
	strs := []string{"qwerty", "0r", ""}
	for _, s := range strs {
		_, err := interpreter.Model.ParseInt(s)
		if err == nil {
			t.Errorf("Error expected when parsing %s", s)
		}

		_, err16 := interpreter.Model.ParseIntFromBase(s, 16)
		if err16 == nil {
			t.Errorf("Error expected when parsing %s", s)
		}
	}
}

func TestIntHooks1(t *testing.T) {
	interpreter := newTestInterpreter()
	a := interpreter.Model.FromInt(1)
	b := interpreter.Model.FromInt(2)
	var z m.KReference
	var err error

	interpreter.backupInput(a, b)
	z, err = intHooks.eq(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ne(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.le(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.lt(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ge(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.gt(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.add(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "3", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.sub(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.mul(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.tdiv(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.tmod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ediv(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.emod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	z, err = intHooks.shl(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "4", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.shr(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.and(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.or(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "3", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.xor(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "3", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.not(b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-3", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.abs(b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.max(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.min(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)
}

func TestIntHooks2(t *testing.T) {
	interpreter := newTestInterpreter()
	a := interpreter.Model.FromInt(1)
	b := interpreter.Model.FromInt(1)

	var z m.KReference
	var err error

	interpreter.backupInput(a, b)
	z, err = intHooks.eq(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ne(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.le(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.lt(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ge(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.gt(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.add(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.sub(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.mul(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.tdiv(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.tmod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.ediv(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.emod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.shl(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.shr(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.and(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, m.IntZero)
	z, err = intHooks.and(a, m.IntZero, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, m.IntZero)

	interpreter.backupInput(m.IntZero, b)
	z, err = intHooks.and(m.IntZero, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, m.IntZero, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.or(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, m.IntZero)
	z, err = intHooks.or(a, m.IntZero, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, m.IntZero)

	interpreter.backupInput(m.IntZero, b)
	z, err = intHooks.or(m.IntZero, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, m.IntZero, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.xor(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.not(a, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.abs(a, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.max(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(a, b)
	z, err = intHooks.min(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)
}

func testOverflowAddSub(t *testing.T, interpreter *Interpreter, aVal int64) {
	var z, a m.KReference
	var expectedStr string
	var err error

	a = interpreter.Model.FromInt64(aVal)
	one := interpreter.Model.FromInt(1)

	// a + 1
	interpreter.backupInput(a, one)
	z, err = intHooks.add(a, one, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	expectedStr = fmt.Sprintf("%d", aVal+int64(1))
	assertIntOk(t, expectedStr, z, err, interpreter)
	interpreter.checkImmutable(t, a, one)

	// 1 + a
	interpreter.backupInput(a, one)
	z, err = intHooks.add(one, a, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	expectedStr = fmt.Sprintf("%d", aVal+int64(1))
	assertIntOk(t, expectedStr, z, err, interpreter)
	interpreter.checkImmutable(t, a, one)

	// a - 1
	interpreter.backupInput(a, one)
	z, err = intHooks.sub(a, one, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	expectedStr = fmt.Sprintf("%d", aVal-int64(1))
	assertIntOk(t, expectedStr, z, err, interpreter)
	interpreter.checkImmutable(t, a, one)

	// 1 - a
	interpreter.backupInput(a, one)
	z, err = intHooks.sub(a, one, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	expectedStr = fmt.Sprintf("%d", aVal-int64(1))
	assertIntOk(t, expectedStr, z, err, interpreter)
	interpreter.checkImmutable(t, a, one)
}

func TestIntHooksOverflowAddSub(t *testing.T) {
	interpreter := newTestInterpreter()

	// around the upper limit
	for aVal := int64(4294967290); aVal <= 4294967300; aVal++ {
		testOverflowAddSub(t, interpreter, aVal)
	}

	// around the lower limit
	for aVal := int64(-4294967290); aVal >= -4294967300; aVal-- {
		testOverflowAddSub(t, interpreter, aVal)
	}
}

func testOverflowMul(t *testing.T, interpreter *Interpreter, aVal int, bVal int) {
	var z m.KReference
	var expectedStr string
	var err error

	a := interpreter.Model.FromInt(aVal)
	b := interpreter.Model.FromInt(bVal)
	interpreter.backupInput(a, b)
	z, err = intHooks.mul(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	expectedStr = fmt.Sprintf("%d", int64(aVal)*int64(bVal))
	assertIntOk(t, expectedStr, z, err, interpreter)
	interpreter.checkImmutable(t, a, b)
}

func TestIntHooksOverflowMul(t *testing.T) {
	interpreter := newTestInterpreter()
	sqrtMaxInt32 := int(math.Sqrt(float64(math.MaxInt32)))

	// around the upper limit
	for aVal := sqrtMaxInt32 - 20; aVal < sqrtMaxInt32+10; aVal++ {
		for bVal := sqrtMaxInt32 - 20; bVal < sqrtMaxInt32+10; bVal++ {
			testOverflowMul(t, interpreter, aVal, bVal)
			testOverflowMul(t, interpreter, -aVal, -bVal)
			testOverflowMul(t, interpreter, -aVal, bVal)
			testOverflowMul(t, interpreter, aVal, -bVal)
		}
	}
}

func TestIntHooksMod(t *testing.T) {
	interpreter := newTestInterpreter()
	var a, b, z m.KReference
	var err error

	a = interpreter.Model.IntFromString("231584178474632390847141970017375815706539969331281128078915168015826259279869")
	b = interpreter.Model.FromInt(2)
	interpreter.backupInput(a, b)
	z, err = intHooks.tmod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	a = interpreter.Model.FromInt(-5)
	b = interpreter.Model.FromInt(3)
	interpreter.backupInput(a, b)
	z, err = intHooks.tmod(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-2", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)
}

func TestIntHooksPow(t *testing.T) {
	interpreter := newTestInterpreter()
	a := interpreter.Model.FromInt(2)
	b := interpreter.Model.FromInt(10)
	c := interpreter.Model.FromInt(1000)
	var z m.KReference
	var err error

	interpreter.backupInput(a, b)
	z, err = intHooks.pow(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1024", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

	interpreter.backupInput(b, a)
	z, err = intHooks.pow(b, a, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "100", z, err, interpreter)
	interpreter.checkImmutable(t, b, a)

	interpreter.backupInput(a, b, c)
	z, err = intHooks.powmod(a, b, c, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "24", z, err, interpreter)
	interpreter.checkImmutable(t, a, b, c)

}

func TestIntLog2(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var arg m.KReference

	arg = interpreter.Model.FromInt(1)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = interpreter.Model.FromInt(2)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = interpreter.Model.FromInt(3)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = interpreter.Model.FromInt(4)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = interpreter.Model.FromInt(255)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "7", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = interpreter.Model.FromInt(256)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "8", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	for i := 1000; i < 1009; i++ {
		// 1 << i
		arg1, arg2 := interpreter.Model.FromInt(1), interpreter.Model.FromInt(i)
		interpreter.backupInput(arg1, arg2)
		big, _ := intHooks.shl(arg1, arg2, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
		interpreter.checkImmutable(t, arg1, arg2)

		interpreter.backupInput(big)
		log, err = intHooks.log2(big, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
		assertIntOk(t, fmt.Sprintf("%d", i), log, err, interpreter)
		interpreter.checkImmutable(t, big)

		// (1 << i) - 1
		big, _ = intHooks.sub(big, m.IntOne, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)

		interpreter.backupInput(big)
		log, err = intHooks.log2(big, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
		assertIntOk(t, fmt.Sprintf("%d", i-1), log, err, interpreter)
		interpreter.checkImmutable(t, big)

	}
}

func TestIntBitRangeZero(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(8), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(8), interpreter.Model.FromInt(254)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(12345), interpreter.Model.FromInt(8), interpreter.Model.FromInt(0)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntBitRangePositive(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen = interpreter.Model.FromInt(5), interpreter.Model.FromInt(0), interpreter.Model.FromInt(32)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "5", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(255), interpreter.Model.FromInt(0), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "255", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(256), interpreter.Model.FromInt(0), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(256), interpreter.Model.FromInt(8), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntBitRangeNegative(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen = interpreter.Model.FromInt(-1), interpreter.Model.FromInt(0), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "255", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(-2), interpreter.Model.FromInt(0), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "254", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	// TODO: add cases with offset
}

func TestIntBitRangeExamplesFromCode(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen = interpreter.Model.FromInt(-2), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639934", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(-6), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639930", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(-5), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639931", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(-70), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639866", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "57896044618658097711785492504343953926634992332820282019728792003956564819968", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819967"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "57896044618658097711785492504343953926634992332820282019728792003956564819969", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("839073110415334749446166558033970346762825975837975101735199884115312533623"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "839073110415334749446166558033970346762825975837975101735199884115312533623", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639936"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(64)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(-1), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639935", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(1), interpreter.Model.FromInt(0), interpreter.Model.FromInt(264)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntSignExtendBitRangeZero(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(8), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(0), interpreter.Model.FromInt(8), interpreter.Model.FromInt(254)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = interpreter.Model.FromInt(12345), interpreter.Model.FromInt(8), interpreter.Model.FromInt(0)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntSignExtendBitRangeMinusOne(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	for len := 8; len <= 256; len += 8 {
		argI, argOff, argLen =
			interpreter.Model.IntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
			interpreter.Model.FromInt(0), interpreter.Model.FromInt(len)
		interpreter.backupInput(argI, argOff, argLen)
		log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
		assertIntOk(t, "-1", log, err, interpreter)
		interpreter.checkImmutable(t, argI, argOff, argLen)
	}
}

func TestIntSignExtendBitRangeExamplesFromCode(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.KReference
	var err error
	var argI, argOff, argLen m.KReference

	argI, argOff, argLen =
		interpreter.Model.IntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639934"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-2", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("1243892"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(16)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-1292", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("128"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-128", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		interpreter.Model.IntFromString("65407"),
		interpreter.Model.FromInt(0), interpreter.Model.FromInt(16)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-129", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

}

func assertIntOk(t *testing.T, expectedAsStr string, actual m.KReference, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	expectedK := interpreter.Model.IntFromString(expectedAsStr)
	if !interpreter.Model.Equals(expectedK, actual) {
		t.Errorf("Unexpected result. Got:%s Want:%s", interpreter.Model.PrettyPrint(actual), interpreter.Model.PrettyPrint(expectedK))
	}
}
