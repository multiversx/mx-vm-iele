// File generated by the K Framework Go backend. Timestamp: 2019-06-13 13:37:26.737

package ieletestinginterpreter 

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) makeStuck(c m.K, config m.K) (m.K, error) {
	// rule #-1
	// source: ? @?
	// {| rule `<generatedTop>`(_0,`<s>`(``.K=>#STUCK(.KList)``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token("true","Bool") ensures #token("true","Bool") [] |}
	// LHS
	if kapp0, t := c.(*m.KApply); t && kapp0.Label == m.LblXltgeneratedTopXgt && len(kapp0.List) == 9 { // `<generatedTop>`(_0,`<s>`(DotVar1),_1,_2,_3,_4,_5,_6,_7)
		varXu0 := kapp0.List[0] // lhs KVariable _0
		if kapp1, t := kapp0.List[1].(*m.KApply); t && kapp1.Label == m.LblXltsXgt && len(kapp1.List) == 1 { // `<s>`(DotVar1)
			// KSequence, size 1:DotVar1
			varDotVar1 := kapp1.List[0] // lhs KVariable DotVar1
			varXu1 := kapp0.List[2] // lhs KVariable _1
			varXu2 := kapp0.List[3] // lhs KVariable _2
			varXu3 := kapp0.List[4] // lhs KVariable _3
			varXu4 := kapp0.List[5] // lhs KVariable _4
			varXu5 := kapp0.List[6] // lhs KVariable _5
			varXu6 := kapp0.List[7] // lhs KVariable _6
			varXu7 := kapp0.List[8] // lhs KVariable _7
			// RHS
			i.traceRuleApply("STEP", -1, "{| rule `<generatedTop>`(_0,`<s>`(``.K=>#STUCK(.KList)``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token(\"true\",\"Bool\") ensures #token(\"true\",\"Bool\") [] |}")
			return &m.KApply{Label: m.LblXltgeneratedTopXgt, List: []m.K{ // as-is <generatedTop>
				varXu0,
				&m.KApply{Label: m.LblXltsXgt, List: []m.K{ // as-is <s>
					i.Model.AssembleKSequence(
						&m.KApply{Label: m.LblXhashSTUCK, List: []m.K{ // as-is #STUCK
						}},
						varDotVar1,
					),
				}},
				varXu1,
				varXu2,
				varXu3,
				varXu4,
				varXu5,
				varXu6,
				varXu7,
			}}, nil
		}
	}

	return c, nil
}

func (i *Interpreter) makeUnstuck(c m.K, config m.K) (m.K, error) {
	// rule #-1
	// source: ? @?
	// {| rule `<generatedTop>`(_0,`<s>`(``#STUCK(.KList)=>.K``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token("true","Bool") ensures #token("true","Bool") [] |}
	// LHS
	if kapp0, t := c.(*m.KApply); t && kapp0.Label == m.LblXltgeneratedTopXgt && len(kapp0.List) == 9 { // `<generatedTop>`(_0,`<s>`(#STUCK(.KList)~>DotVar1),_1,_2,_3,_4,_5,_6,_7)
		varXu0 := kapp0.List[0] // lhs KVariable _0
		if kapp1, t := kapp0.List[1].(*m.KApply); t && kapp1.Label == m.LblXltsXgt && len(kapp1.List) == 1 { // `<s>`(#STUCK(.KList)~>DotVar1)
			if ok, kseq2Head, kseq2Tail := i.Model.TrySplitToHeadTail(kapp1.List[0]); ok { // #STUCK(.KList)~>DotVar1
				if kapp3, t := kseq2Head.(*m.KApply); t && kapp3.Label == m.LblXhashSTUCK && len(kapp3.List) == 0 { // #STUCK(.KList)
					varDotVar1 := kseq2Tail // lhs KVariable DotVar1
					varXu1 := kapp0.List[2] // lhs KVariable _1
					varXu2 := kapp0.List[3] // lhs KVariable _2
					varXu3 := kapp0.List[4] // lhs KVariable _3
					varXu4 := kapp0.List[5] // lhs KVariable _4
					varXu5 := kapp0.List[6] // lhs KVariable _5
					varXu6 := kapp0.List[7] // lhs KVariable _6
					varXu7 := kapp0.List[8] // lhs KVariable _7
					// RHS
					i.traceRuleApply("STEP", -1, "{| rule `<generatedTop>`(_0,`<s>`(``#STUCK(.KList)=>.K``~>DotVar1),_1,_2,_3,_4,_5,_6,_7) requires #token(\"true\",\"Bool\") ensures #token(\"true\",\"Bool\") [] |}")
					return &m.KApply{Label: m.LblXltgeneratedTopXgt, List: []m.K{ // as-is <generatedTop>
						varXu0,
						&m.KApply{Label: m.LblXltsXgt, List: []m.K{ // as-is <s>/* rhs KSequence size=1 */ 
							varDotVar1,
						}},
						varXu1,
						varXu2,
						varXu3,
						varXu4,
						varXu5,
						varXu6,
						varXu7,
					}}, nil
				}
			}
		}
	}

	return c, nil
}

