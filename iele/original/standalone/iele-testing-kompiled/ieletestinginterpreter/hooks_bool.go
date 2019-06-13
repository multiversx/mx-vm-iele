// File provided by the K Framework Go backend. Timestamp: 2019-06-14 00:50:56.636

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type boolHooksType int

const boolHooks boolHooksType = 0

func (boolHooksType) and(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool(b1.Value && b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (h boolHooksType) andThen(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return h.and(c1, c2, lbl, sort, config, interpreter)
}

func (boolHooksType) or(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool(b1.Value || b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (h boolHooksType) orElse(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return h.or(c1, c2, lbl, sort, config, interpreter)
}

func (boolHooksType) not(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b, ok := c.(*m.Bool)
	if ok {
		return m.ToBool(!b.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) implies(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool((!b1.Value) || b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) ne(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool(b1.Value != b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) eq(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool(b1.Value == b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (boolHooksType) xor(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	b1, ok1 := c1.(*m.Bool)
	b2, ok2 := c2.(*m.Bool)
	if ok1 && ok2 {
		return m.ToBool(b1.Value != b2.Value), nil
	}
	return m.NoResult, &hookNotImplementedError{}
}
