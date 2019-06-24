// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:19:21.575

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type arrayHooksType int

const arrayHooks arrayHooksType = 0

func (arrayHooksType) make(maxSize m.K, defValue m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	maxSizeInt, ok := maxSize.(*m.Int)
	if !ok {
		return invalidArgsResult()
	}
	if !maxSizeInt.Value.IsUint64() {
		return invalidArgsResult()
	}
	maxSizeUint := maxSizeInt.Value.Uint64()
	return &m.Array{Sort: sort, Data: interpreter.Model.MakeDynamicArray(maxSizeUint, defValue)}, nil
}

func (t arrayHooksType) makeEmpty(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return t.make(c, m.InternedBottom, lbl, sort, config, interpreter)
}

func (t arrayHooksType) ctor(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return t.makeEmpty(c2, lbl, sort, config, interpreter)
}

func (arrayHooksType) lookup(karr m.K, kidx m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	arr, ok1 := karr.(*m.Array)
	idx, ok2 := kidx.(*m.Int)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !idx.Value.IsUint64() {
		return invalidArgsResult()
	}
	idxUint := idx.Value.Uint64()
	return arr.Data.Get(idxUint)
}

func (arrayHooksType) remove(karr m.K, kidx m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	arr, ok1 := karr.(*m.Array)
	idx, ok2 := kidx.(*m.Int)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !idx.Value.IsUint64() {
		return invalidArgsResult()
	}
	idxUint := idx.Value.Uint64()
	err := arr.Data.Set(idxUint, arr.Data.Default)
	if err != nil {
		return m.NoResult, err
	}
	return arr, nil
}

func (arrayHooksType) update(karr m.K, kidx m.K, newVal m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	arr, ok1 := karr.(*m.Array)
	idx, ok2 := kidx.(*m.Int)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !idx.Value.IsUint64() {
		return invalidArgsResult()
	}
	idxUint := idx.Value.Uint64()
	err := arr.Data.Set(idxUint, newVal)
	if err != nil {
		return m.NoResult, err
	}
	return arr, nil
}

func (arrayHooksType) updateAll(karr m.K, kidx m.K, klist m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	arr, ok1 := karr.(*m.Array)
	idx, ok2 := kidx.(*m.Int)
	list, ok3 := klist.(*m.List)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if !idx.Value.IsUint64() {
		return invalidArgsResult()
	}
	idxUint := idx.Value.Uint64()
	listLen := uint64(len(list.Data))
	arr.Data.UpgradeSize(idxUint + listLen - 1) // upgrade size all at once
	for i := uint64(0); i < listLen && idxUint+i < arr.Data.MaxSize; i++ {
		err := arr.Data.Set(idxUint+i, list.Data[i])
		if err != nil {
			return m.NoResult, err
		}
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (arrayHooksType) fill(karr m.K, kfrom m.K, kto m.K, elt m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	arr, ok1 := karr.(*m.Array)
	from, ok2 := kfrom.(*m.Int)
	to, ok3 := kto.(*m.Int)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	if !from.Value.IsUint64() || !to.Value.IsUint64() {
		return invalidArgsResult()
	}
	fromInt := from.Value.Uint64()
	toInt := to.Value.Uint64()
	for i := fromInt; i < toInt && i < arr.Data.MaxSize; i++ {
		arr.Data.Set(i, elt)
	}
	return arr, nil
}

func (arrayHooksType) inKeys(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	idx, ok2 := c1.(*m.Int)
	arr, ok1 := c2.(*m.Array)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	if !idx.Value.IsUint64() {
		return invalidArgsResult()
	}
	idxUint := idx.Value.Uint64()
	val, err := arr.Data.Get(idxUint)
	if err != nil {
		return m.NoResult, err
	}
	hasValue := !interpreter.Model.Equals(val, arr.Data.Default)
	return m.ToBool(hasValue), nil
}
