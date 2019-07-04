// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

type bufferHooksType int

const bufferHooks bufferHooksType = 0

func (bufferHooksType) empty(lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	return interpreter.Model.NewStringBuffer(), nil
}

func (bufferHooksType) concat(kbuf m.KReference, kstr m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	buf, ok1 := interpreter.Model.GetStringBufferObject(kbuf)
	str, ok2 := interpreter.Model.GetStringObject(kstr)
	if !ok1 || !ok2 {
		return invalidArgsResult()
	}
	buf.Value.WriteString(str.Value)
	return kbuf, nil
}

func (bufferHooksType) toString(kbuf m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, interpreter *Interpreter) (m.KReference, error) {
	buf, ok := interpreter.Model.GetStringBufferObject(kbuf)
	if !ok {
		return invalidArgsResult()
	}
	return interpreter.Model.NewString(buf.Value.String()), nil
}
