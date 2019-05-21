package main

import (
	"strings"

	oj "github.com/ElrondNetwork/elrond-vm/iele/test-util/orderedjson"
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
