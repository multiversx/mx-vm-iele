// File provided by the K Framework Go backend. Timestamp: 2019-06-12 11:57:09.485

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

type kreflectionHooksType int

const kreflectionHooks kreflectionHooksType = 0

func (kreflectionHooksType) sort(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	switch k := c.(type) {
	case *m.KToken:
		return m.NewString(k.Sort.Name()), nil
	case *m.Int:
		return m.NewString("Int"), nil
	case *m.String:
		return m.NewString("String"), nil
	case *m.Bytes:
		return m.NewString("Bytes"), nil
	case *m.Bool:
		return m.NewString("Bool"), nil
	case *m.Map:
		return m.NewString(k.Sort.Name()), nil
	case *m.List:
		return m.NewString(k.Sort.Name()), nil
	case *m.Set:
		return m.NewString(k.Sort.Name()), nil
	default:
		return m.NoResult, &hookNotImplementedError{}
	}
}

func (kreflectionHooksType) getKLabel(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	if k, t := c.(*m.KApply); t {
		return &m.InjectedKLabel{Label: k.Label}, nil
	}
	return m.InternedBottom, nil
}

func (kreflectionHooksType) configuration(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return config, nil
}

var freshCounter int

func (kreflectionHooksType) fresh(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	if k, t := c.(*m.String); t {
		sort := m.ParseSort(k.Value)
		result, err := interpreter.freshFunction(sort, config, freshCounter)
		if err != nil {
		    return m.NoResult, err
		}
		freshCounter++
		return result, nil
	}
	return m.NoResult, &hookNotImplementedError{}
}

func (kreflectionHooksType) isConcrete(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.BoolTrue, nil
}

func (kreflectionHooksType) getenv(c m.K, lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}

func (kreflectionHooksType) argv(lbl m.KLabel, sort m.Sort, config m.K, interpreter *Interpreter) (m.K, error) {
	return m.NoResult, &hookNotImplementedError{}
}
