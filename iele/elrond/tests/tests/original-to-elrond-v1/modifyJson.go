package main

import (
	"strings"

	oj "github.com/ElrondNetwork/elrond-vm-util/test-util/orderedjson"
)

func modifyJSON(jobj oj.OJsonObject) error {
	tests, _ := jobj.(*oj.OJsonMap)
	for _, kvtest := range tests.OrderedKV {
		testMap, _ := kvtest.Value.(*oj.OJsonMap)
		for _, testComponent := range testMap.OrderedKV {
			if testComponent.Key == "pre" || testComponent.Key == "postState" {
				acctMap, _ := testComponent.Value.(*oj.OJsonMap)
				for _, acctKV := range acctMap.OrderedKV {
					acctKV.Key = expandAddress(acctKV.Key)
				}
				acctMap.RefreshKeySet()
			}
			if testComponent.Key == "blocks" {
				blocks, _ := testComponent.Value.(*oj.OJsonList)
				for _, block := range blocks.AsList() {
					blockCmps, _ := block.(*oj.OJsonMap)
					for _, blockCmp := range blockCmps.OrderedKV {
						if blockCmp.Key == "transactions" {
							txs, _ := blockCmp.Value.(*oj.OJsonList)
							for _, txRaw := range txs.AsList() {
								tx, _ := txRaw.(*oj.OJsonMap)
								for _, txCmp := range tx.OrderedKV {
									if txCmp.Key == "from" {
										from, _ := txCmp.Value.(*oj.OJsonString)
										from.Value = expandAddress(from.Value)
									}
									if txCmp.Key == "to" {
										to, _ := txCmp.Value.(*oj.OJsonString)
										if len(to.Value) == 42 {
											to.Value = expandAddress(to.Value)
										}
									}
									if txCmp.Key == "arguments" {
										args, _ := txCmp.Value.(*oj.OJsonList)
										for _, arg := range args.AsList() {
											argStr, _ := arg.(*oj.OJsonString)
											if len(argStr.Value) == 42 {
												argStr.Value = expandAddress(argStr.Value)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func expandAddress(oldAddr string) string {
	if !strings.HasPrefix(oldAddr, "0x") {
		panic("unexpected address, no 0x prefix")
	}

	if len(oldAddr) != 42 {
		panic("unexpected original address length, should be 0x + 2*20bytes")
	}

	return oldAddr + "000000000000000000000000"
}
