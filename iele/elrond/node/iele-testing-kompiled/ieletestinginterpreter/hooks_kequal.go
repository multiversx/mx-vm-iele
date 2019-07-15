// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:14:14.526

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type kequalHooksType int

const kequalHooks kequalHooksType = 0

// equals
func (kequalHooksType) eq(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.ToKBool(interpreter.Model.Equals(c1, c2)), nil
}

// not equals
func (kequalHooksType) ne(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.ToKBool(!interpreter.Model.Equals(c1, c2)), nil
}

// kbool ? k1 : k2
func (kequalHooksType) ite(kbool m.KReference, k1 m.KReference, k2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	b, ok := m.CastToBool(kbool)
	if !ok {
		return invalidArgsResult()
	}
	if b {
		return k1, nil
	}
	return k2, nil
}

