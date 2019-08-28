// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type floatHooksType int

const floatHooks floatHooksType = 0

func (floatHooksType) isNaN(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) maxValue(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) minValue(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) round(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) abs(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) ceil(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) floor(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) acos(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) asin(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) atan(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) cos(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) sin(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) tan(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) exp(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) log(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) neg(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) add(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) sub(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) mul(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) div(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) pow(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) eq(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) lt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) le(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) gt(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) ge(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) precision(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) exponentBits(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) float2int(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) int2float(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) min(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) max(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) rem(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) root(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) sign(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) significand(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) atan2(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (floatHooksType) exponent(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

