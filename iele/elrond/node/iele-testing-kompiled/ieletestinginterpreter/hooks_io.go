// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type ioHooksType int

const ioHooks ioHooksType = 0

func (ioHooksType) close(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) getc(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) open(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) putc(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) read(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) seek(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) seekEnd(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) tell(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) write(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) lock(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) unlock(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) log(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) stat(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) lstat(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) opendir(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) parse(c1 m.KReference, c2 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) parseInModule(c1 m.KReference, c2 m.KReference, c3 m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

func (ioHooksType) system(c m.KReference,lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return m.NoResult, m.GetHookNotImplementedError()
}

