// File provided by the K Framework Go backend. Timestamp: 2019-06-12 11:57:09.485

package ieletestinginterpreter

import (
	"fmt"
    m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
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
		i, err := m.ParseInt(s)
		assertIntOk(t, s, i, err, interpreter)
	}
}

func TestParseIntError(t *testing.T) {
	strs := []string{"abc", "-0", ""}
	for _, s := range strs {
		_, err := m.ParseInt(s)
		if err == nil {
			t.Errorf("Error expected when parsing %s", s)
		}
	}
}

func TestIntHooks1(t *testing.T) {
	interpreter := newTestInterpreter()
	a := m.NewIntFromInt(1)
	b := m.NewIntFromInt(2)
	var z m.K
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
	a := m.NewIntFromInt(1)
	b := m.NewIntFromInt(1)

	var z m.K
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

	interpreter.backupInput(a, b)
	z, err = intHooks.or(a, b, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", z, err, interpreter)
	interpreter.checkImmutable(t, a, b)

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

func TestIntHooksPow(t *testing.T) {
	interpreter := newTestInterpreter()
	a := m.NewIntFromInt(2)
	b := m.NewIntFromInt(10)
	c := m.NewIntFromInt(1000)
	var z m.K
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
	var log m.K
	var err error
	var arg m.K

	arg = m.NewIntFromInt(1)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = m.NewIntFromInt(2)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = m.NewIntFromInt(3)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = m.NewIntFromInt(4)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "2", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = m.NewIntFromInt(255)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "7", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	arg = m.NewIntFromInt(256)
	interpreter.backupInput(arg)
	log, err = intHooks.log2(arg, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "8", log, err, interpreter)
	interpreter.checkImmutable(t, arg)

	for i := 1000; i < 1009; i++ {
		// 1 << i
		arg1, arg2 := m.NewIntFromInt(1), m.NewIntFromInt(i)
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
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(8), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(8), m.NewIntFromInt(254)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(12345), m.NewIntFromInt(8), m.NewIntFromInt(0)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntBitRangePositive(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen = m.NewIntFromInt(5), m.NewIntFromInt(0), m.NewIntFromInt(32)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "5", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(255), m.NewIntFromInt(0), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "255", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(256), m.NewIntFromInt(0), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(256), m.NewIntFromInt(8), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntBitRangeNegative(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen = m.NewIntFromInt(-1), m.NewIntFromInt(0), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "255", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(-2), m.NewIntFromInt(0), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "254", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	// TODO: add cases with offset
}

func TestIntBitRangeExamplesFromCode(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen = m.NewIntFromInt(-2), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639934", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(-6), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639930", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(-5), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639931", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(-70), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639866", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819968"),
		m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "57896044618658097711785492504343953926634992332820282019728792003956564819968", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("-57896044618658097711785492504343953926634992332820282019728792003956564819967"),
		m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "57896044618658097711785492504343953926634992332820282019728792003956564819969", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("839073110415334749446166558033970346762825975837975101735199884115312533623"),
		m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "839073110415334749446166558033970346762825975837975101735199884115312533623", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639936"),
		m.NewIntFromInt(0), m.NewIntFromInt(64)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(-1), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "115792089237316195423570985008687907853269984665640564039457584007913129639935", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(1), m.NewIntFromInt(0), m.NewIntFromInt(264)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.bitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "1", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntSignExtendBitRangeZero(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(8), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(0), m.NewIntFromInt(8), m.NewIntFromInt(254)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen = m.NewIntFromInt(12345), m.NewIntFromInt(8), m.NewIntFromInt(0)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "0", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)
}

func TestIntSignExtendBitRangeMinusOne(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	for len := 8; len <= 256; len += 8 {
		argI, argOff, argLen =
			m.NewIntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935"),
			m.NewIntFromInt(0), m.NewIntFromInt(len)
		interpreter.backupInput(argI, argOff, argLen)
		log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
		assertIntOk(t, "-1", log, err, interpreter)
		interpreter.checkImmutable(t, argI, argOff, argLen)
	}
}

func TestIntSignExtendBitRangeExamplesFromCode(t *testing.T) {
	interpreter := newTestInterpreter()
	var log m.K
	var err error
	var argI, argOff, argLen m.K

	argI, argOff, argLen =
		m.NewIntFromString("115792089237316195423570985008687907853269984665640564039457584007913129639934"),
		m.NewIntFromInt(0), m.NewIntFromInt(256)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-2", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("1243892"),
		m.NewIntFromInt(0), m.NewIntFromInt(16)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-1292", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("128"),
		m.NewIntFromInt(0), m.NewIntFromInt(8)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-128", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

	argI, argOff, argLen =
		m.NewIntFromString("65407"),
		m.NewIntFromInt(0), m.NewIntFromInt(16)
	interpreter.backupInput(argI, argOff, argLen)
	log, err = intHooks.signExtendBitRange(argI, argOff, argLen, m.LblDummy, m.SortInt, m.InternedBottom, interpreter)
	assertIntOk(t, "-129", log, err, interpreter)
	interpreter.checkImmutable(t, argI, argOff, argLen)

}

func assertIntOk(t *testing.T, expectedAsStr string, actual m.K, err error, interpreter *Interpreter) {
	if err != nil {
		t.Error(err, interpreter)
	}
	expectedK := m.NewIntFromString(expectedAsStr)
	if !interpreter.Model.Equals(expectedK, actual) {
		t.Errorf("Unexpected result. Got:%s Want:%s", interpreter.Model.PrettyPrint(actual), interpreter.Model.PrettyPrint(expectedK))
	}
}
