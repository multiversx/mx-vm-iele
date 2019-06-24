// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:24:14.667

package ieletestinginterpreter

import (
	"fmt"
	koreparser "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/koreparser"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
	"strconv"
)

func (i *Interpreter) convertParserModelToKModel(pk koreparser.K) m.K {
	switch v := pk.(type) {
	case koreparser.KApply:
		var convertedList []m.K
		for _, le := range v.List {
			convertedList = append(convertedList, i.convertParserModelToKModel(le))
		}
		return &m.KApply{Label: m.ParseKLabel(v.Label), List: convertedList}
	case koreparser.InjectedKLabel:
		return &m.InjectedKLabel{Label: m.ParseKLabel(v.Label)}
	case koreparser.KToken:
		return convertKToken(m.ParseSort(v.Sort), v.Value)
	case koreparser.KVariable:
		return &m.KVariable{Name: v.Name}
	case koreparser.KSequence:
		var convertedKs []m.K
		for _, ksElem := range v {
			convertedKs = append(convertedKs, i.convertParserModelToKModel(ksElem))
		}
		return i.Model.NewKSequence(convertedKs)
	default:
		panic(fmt.Sprintf("Unknown parser model K type: %#v", v))
	}
}

func convertKToken(sort m.Sort, value string) m.K {
	switch sort {
	case m.SortInt:
		i, err := m.ParseInt(value)
		if err != nil {
			panic(err)
		}
		return i
	case m.SortFloat:
		panic("Float token parse not implemented.")
	case m.SortString:
		unescapedStr := string(koreparser.UnescapeKString([]byte(value)))
		return m.NewString(unescapedStr)
	case m.SortBool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			panic("Could not parse bool token: " + value)
		}
		return m.ToBool(b)
	default:
		return &m.KToken{Value: value, Sort: sort}
	}
}
