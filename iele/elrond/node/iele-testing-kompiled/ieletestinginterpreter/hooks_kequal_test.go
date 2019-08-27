// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

var constKToken1 = m.NewKTokenConstant(m.SortInt, "KToken1")

func TestKTokenEquals(t *testing.T) {
	interpreter := newTestInterpreter()
	var z m.KReference
	var err error

	var ktok1 = interpreter.Model.NewKToken(m.SortInt, "KToken1")

	z, err = kequalHooks.eq(ktok1, constKToken1, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = kequalHooks.eq(constKToken1, ktok1, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = kequalHooks.eq(ktok1, ktok1, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	z, err = kequalHooks.eq(constKToken1, constKToken1, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, true, z, err, interpreter)

	var ktok2 = interpreter.Model.NewKToken(m.SortInt, "KToken2")

	z, err = kequalHooks.eq(ktok1, ktok2, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = kequalHooks.eq(constKToken1, ktok2, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	var ktok3 = interpreter.Model.NewKToken(m.SortBool, "KToken1")

	z, err = kequalHooks.eq(ktok1, ktok3, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

	z, err = kequalHooks.eq(constKToken1, ktok3, m.LblDummy, m.SortBool, m.InternedBottom, interpreter)
	assertBoolOk(t, false, z, err, interpreter)

}
