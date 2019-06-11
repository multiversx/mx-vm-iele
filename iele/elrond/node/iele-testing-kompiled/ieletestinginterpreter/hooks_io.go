// File provided by the K Framework Go backend. Timestamp: 2019-06-12 02:37:11.223

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type ioHooksType int

const ioHooks ioHooksType = 0

func (ioHooksType) close(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) getc(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) open(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) putc(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) read(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) seek(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) seekEnd(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) tell(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) write(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) lock(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) unlock(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) log(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) stat(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) lstat(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) opendir(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) parse(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) parseInModule(c1 m.K, c2 m.K, c3 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (ioHooksType) system(c m.K,lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

