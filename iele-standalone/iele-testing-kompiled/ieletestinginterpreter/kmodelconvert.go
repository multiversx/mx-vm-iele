package ieletestinginterpreter

import (
	"fmt"
	koreparser "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/koreparser"
	m "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestingmodel"
	"strconv"
)

func convertParserModelToKModel(pk koreparser.K) m.K {
	switch v := pk.(type) {
	case koreparser.KApply:
		var convertedList []m.K
		for _, le := range v.List {
			convertedList = append(convertedList, convertParserModelToKModel(le))
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
			convertedKs = append(convertedKs, convertParserModelToKModel(ksElem))
		}
		return &m.KSequence{Ks: convertedKs}
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
