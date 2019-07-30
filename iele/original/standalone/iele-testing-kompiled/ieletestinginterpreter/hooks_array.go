// File provided by the K Framework Go backend. Timestamp: 2019-07-30 16:35:04.814

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type arrayHooksType int

const arrayHooks arrayHooksType = 0

func (arrayHooksType) make(maxSize m.KReference, defValue m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	maxSizeUint, ok := interpreter.Model.GetUint64(maxSize)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewArray(sort, interpreter.Model.MakeDynamicArray(maxSizeUint, defValue)), nil
}

func (t arrayHooksType) makeEmpty(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return t.make(c, m.InternedBottom, lbl, sort, config, interpreter)
}

func (t arrayHooksType) ctor(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return t.makeEmpty(c2, lbl, sort, config, interpreter)
}

func (arrayHooksType) lookup(karr m.KReference, kidx m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	arr, ok1 := interpreter.Model.GetArrayObject(karr)
	idx, ok2 := interpreter.Model.GetUint64(kidx)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	return arr.Data.Get(idx)
}

func (arrayHooksType) remove(karr m.KReference, kidx m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	arr, ok1 := interpreter.Model.GetArrayObject(karr)
	idx, ok2 := interpreter.Model.GetUint64(kidx)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	err := arr.Data.Set(idx, arr.Data.Default)
	if err != nil {
		return m.NoResult, err
	}
	return karr, nil
}

func (arrayHooksType) update(karr m.KReference, kidx m.KReference, newVal m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	arr, ok1 := interpreter.Model.GetArrayObject(karr)
	idx, ok2 := interpreter.Model.GetUint64(kidx)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	err := arr.Data.Set(idx, newVal)
	if err != nil {
		return m.NoResult, err
	}
	return karr, nil
}

func (arrayHooksType) updateAll(karr m.KReference, kidx m.KReference, klist m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	arr, ok1 := interpreter.Model.GetArrayObject(karr)
	idx, ok2 := interpreter.Model.GetUint64(kidx)
	list, ok3 := interpreter.Model.GetListObject(klist)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	listLen := uint64(len(list.Data))
	arr.Data.UpgradeSize(idx + listLen - 1) // upgrade size all at once
	for i := uint64(0); i < listLen && idx+i < arr.Data.MaxSize; i++ {
		err := arr.Data.Set(idx+i, list.Data[i])
		if err != nil {
			return m.NoResult, err
		}
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (arrayHooksType) fill(karr m.KReference, kfrom m.KReference, kto m.KReference, elt m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	arr, ok1 := interpreter.Model.GetArrayObject(karr)
	from, ok2 := interpreter.Model.GetUint64(kfrom)
	to, ok3 := interpreter.Model.GetUint64(kto)
	if !ok1 || !ok2 || !ok3 {
		return invalidArgsResult()
	}
	for i := from; i < to && i < arr.Data.MaxSize; i++ {
		arr.Data.Set(i, elt)
	}
	return karr, nil
}

func (arrayHooksType) inKeys(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	idx, ok2 := interpreter.Model.GetUint64(c1)
	arr, ok1 := interpreter.Model.GetArrayObject(c2)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	val, err := arr.Data.Get(idx)
	if err != nil {
		return m.NoResult, err
	}
	hasValue := !interpreter.Model.Equals(val, arr.Data.Default)
	return m.ToKBool(hasValue), nil
}
