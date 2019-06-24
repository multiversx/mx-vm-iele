// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:24:14.667

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type listHooksType int

const listHooks listHooksType = 0

func (listHooksType) unit(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	data := make([]m.K, 0)
	return &m.List{Sort: sort, Label: m.CollectionFor(lbl), Data: data}, nil
}

func (listHooksType) element(e m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	data := make([]m.K, 1)
	data[0] = e
	return &m.List{Sort: sort, Label: m.CollectionFor(lbl), Data: data}, nil
}

func (listHooksType) concat(klist1 m.K, klist2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l1, isList1 := klist1.(*m.List)
	l2, isList2 := klist2.(*m.List)
	if !isList1 || !isList2 {
		return invalidArgsResult()
	}
	if len(l1.Data) == 0 {
		return l2, nil
	}
	if len(l2.Data) == 0 {
		return l1, nil
	}
	data := append(l1.Data, l2.Data...)
	return &m.List{Sort: sort, Label: m.CollectionFor(lbl), Data: data}, nil
}

func (listHooksType) in(e m.K, klist m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l, isList := klist.(*m.List)
	if !isList {
		return invalidArgsResult()
	}
	for _, x := range l.Data {
		if x == e {
			return m.BoolTrue, nil
		}
	}
	return m.BoolFalse, nil
}

func (listHooksType) get(klist m.K, kindex m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l, isList := klist.(*m.List)
	i, isInt := kindex.(*m.Int)
	if !isList || !isInt {
		return invalidArgsResult()
	}
	if !i.Value.IsInt64() {
		return invalidArgsResult()
	}
	index := int(i.Value.Int64())
	if index < 0 {
		index = len(l.Data) + index // count from the end, e.g. -1 is the last element
	}
	if index < 0 || index >= len(l.Data) {
		return m.InternedBottom, nil
	}
	return l.Data[index], nil
}

func (listHooksType) listRange(klist m.K, start m.K, end m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l, isList := klist.(*m.List)
	si, isInt1 := start.(*m.Int)
	ei, isInt2 := end.(*m.Int)
	if !isList || !isInt1 || !isInt2 || !si.Value.IsUint64() || !ei.Value.IsUint64() {
		return invalidArgsResult()
	}
	siUint := si.Value.Uint64()
	eiUint := ei.Value.Uint64()
	return &m.List{Sort: l.Sort, Label: l.Label, Data: l.Data[siUint:eiUint]}, nil
}

func (listHooksType) size(klist m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	l, isList := klist.(*m.List)
	if !isList {
		return invalidArgsResult()
	}
	return m.NewIntFromInt(len(l.Data)), nil
}

func (listHooksType) make(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (listHooksType) fill(c1 m.K, c2 m.K, c3 m.K, c4 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (listHooksType) update(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (listHooksType) updateAll(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}
