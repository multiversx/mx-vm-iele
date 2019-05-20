// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:40:30.522

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type kequalHooksType int

const kequalHooks kequalHooksType = 0

// equals
func (kequalHooksType) eq(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.ToBool(c1.Equals(c2)), nil
}

// not equals
func (kequalHooksType) ne(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	return m.ToBool(!c1.Equals(c2)), nil
}

// kbool ? k1 : k2
func (kequalHooksType) ite(kbool m.K, k1 m.K, k2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	b, ok := kbool.(*m.Bool)
	if !ok {
		return invalidArgsResult()
	}
	if b.Value {
		return k1, nil
	}
	return k2, nil
}
