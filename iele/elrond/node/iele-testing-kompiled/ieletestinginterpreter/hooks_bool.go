// File provided by the K Framework Go backend. Timestamp: 2019-07-04 01:26:11.488

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type boolHooksType int

const boolHooks boolHooksType = 0

func (boolHooksType) and(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool(b1 && b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (h boolHooksType) andThen(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return h.and(c1, c2, lbl, sort, config, interpreter)
}

func (boolHooksType) or(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool(b1 || b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (h boolHooksType) orElse(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return h.or(c1, c2, lbl, sort, config, interpreter)
}

func (boolHooksType) not(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b, ok := m.CastToBool(c)
	if ok {
		return m.ToKBool(!b), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) implies(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool((!b1) || b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) ne(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool(b1 != b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) eq(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool(b1 == b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) xor(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b1, ok1 := m.CastToBool(c1)
	b2, ok2 := m.CastToBool(c2)
	if ok1 && ok2 {
		return m.ToKBool(b1 != b2), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}
