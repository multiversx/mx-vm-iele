// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestinginterpreter

import (
	"fmt"
	koreparser "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/koreparser"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
	"strconv"
)

func (i *Interpreter) convertParserModelToKModel(pk koreparser.K) m.KReference {
	switch v := pk.(type) {
	case koreparser.KApply:
		var convertedList []m.KReference
		for _, le := range v.List {
			convertedList = append(convertedList, i.convertParserModelToKModel(le))
		}
		return i.Model.NewKApply(m.ParseKLabel(v.Label), convertedList...)
	case koreparser.InjectedKLabel:
		return i.Model.NewInjectedKLabel(m.ParseKLabel(v.Label))
	case koreparser.KToken:
		return i.convertKToken(m.ParseSort(v.Sort), v.Value)
	case koreparser.KVariable:
		return i.Model.NewKVariable(v.Name)
	case koreparser.KSequence:
		var convertedKs []m.KReference
		for _, ksElem := range v {
			convertedKs = append(convertedKs, i.convertParserModelToKModel(ksElem))
		}
		return i.Model.NewKSequence(convertedKs)
	default:
		panic(fmt.Sprintf("Unknown parser model K type: %#v", v))
	}
}

func (i *Interpreter) convertKToken(sort m.Sort, value string) m.KReference {
	switch sort {
	case m.SortInt:
		i, err := i.Model.ParseInt(value)
		if err != nil {
			panic(err)
		}
		return i
	case m.SortFloat:
		panic("Float token parse not implemented.")
	case m.SortString:
		unescapedStr := string(koreparser.UnescapeKString([]byte(value)))
		return i.Model.NewString(unescapedStr)
	case m.SortBool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			panic("Could not parse bool token: " + value)
		}
		return m.ToKBool(b)
	default:
		return i.Model.NewKToken(sort, value)
	}
}
