package endpoint

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	endpoint "github.com/ElrondNetwork/elrond-vm/iele-node/endpoint"
)

func parseTopLevel(jsonString []byte) (*testTopLevel, error) {
	var raw interface{}

	jsonErr := json.Unmarshal(jsonString, &raw)
	if jsonErr != nil {
		return nil, jsonErr
	}

	topMap, isMap := raw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("top level unmarshalled object is not a map")
	}

	top := &testTopLevel{tests: make(map[string]*test)}
	for testName, testObj := range topMap {
		t, tErr := processTest(testObj)
		if tErr != nil {
			return nil, tErr
		}
		t.testName = testName
		top.tests[testName] = t
	}
	return top, nil
}

func processTest(testObj interface{}) (*test, error) {
	testMap, isTestMap := testObj.(map[string]interface{})
	if !isTestMap {
		return nil, errors.New("unmarshalled test object is not a map")
	}
	test := &test{pre: world.MakeAccountMap(), postState: world.MakeAccountMap()}

	preRaw := testMap["pre"]
	preMap, isPreMap := preRaw.(map[string]interface{})
	if !isPreMap {
		return nil, errors.New("unmarshalled pre object is not a map")
	}
	for acctAddrRaw, acctRaw := range preMap {
		acct, acctErr := processAccount(acctRaw)
		if acctErr != nil {
			return nil, acctErr
		}
		acctAddr, hexErr := processAccountAddress(acctAddrRaw)
		if hexErr != nil {
			return nil, hexErr
		}
		acct.Address = acctAddr
		test.pre.PutAccount(acct)
	}

	blocksRaw, blocksOk := testMap["blocks"].([]interface{})
	if !blocksOk {
		return nil, errors.New("unmarshalled blocks object is not a list")
	}
	for _, blRaw := range blocksRaw {
		bl, blErr := processBlock(blRaw)
		if blErr != nil {
			return nil, blErr
		}
		test.blocks = append(test.blocks, bl)
	}

	var networkOk bool
	test.network, networkOk = testMap["network"].(string)
	if !networkOk {
		return nil, errors.New("test network value not a string")
	}

	var bhsOk bool
	test.blockHashes, bhsOk = processBigIntList(testMap["blockhashes"])
	if !bhsOk {
		return nil, errors.New("unmarshalled blockHashes object is not a list")
	}

	postRaw := testMap["postState"]
	postMap, isPostMap := postRaw.(map[string]interface{})
	if !isPostMap {
		return nil, errors.New("unmarshalled postState object is not a map")
	}
	for acctAddrRaw, acctRaw := range postMap {
		acct, acctErr := processAccount(acctRaw)
		if acctErr != nil {
			return nil, acctErr
		}
		acctAddr, hexErr := processAccountAddress(acctAddrRaw)
		if hexErr != nil {
			return nil, hexErr
		}
		acct.Address = acctAddr
		test.postState.PutAccount(acct)
	}

	return test, nil
}

func processAccount(acctRaw interface{}) (*world.Account, error) {
	acctMap, isMap := acctRaw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("unmarshalled account object is not a map")
	}

	acct := &world.Account{Storage: make(map[string]*big.Int)}
	var nonceOk, balanceOk, codeOk bool

	acct.Nonce, nonceOk = parseBigInt(acctMap["nonce"])
	if !nonceOk {
		return nil, errors.New("invalid account nonce")
	}

	acct.Balance, balanceOk = parseBigInt(acctMap["balance"])
	if !balanceOk {
		return nil, errors.New("invalid account balance")
	}

	storageRaw, storageOk := acctMap["storage"].(map[string]interface{})
	if !storageOk {
		return nil, errors.New("invalid account storage")
	}
	for k, v := range storageRaw {
		intKey := big.NewInt(0)
		_, keyOk := intKey.SetString(k, 0)
		if !keyOk {
			return nil, errors.New("invalid account storage key")
		}
		intVal, valOk := parseBigInt(v)
		if !valOk {
			return nil, errors.New("invalid account storage value")
		}
		acct.Storage[intKey.Text(16)] = intVal
	}

	acct.Code, codeOk = acctMap["code"].(string)
	if !codeOk {
		return nil, errors.New("invalid account code")
	}

	return acct, nil
}

func processBlock(blockRaw interface{}) (*block, error) {
	blockMap, isMap := blockRaw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("unmarshalled block object is not a map")
	}
	bl := &block{}

	resultsRaw, resultsOk := blockMap["results"].([]interface{})
	if !resultsOk {
		return nil, errors.New("unmarshalled block results object is not a list")
	}
	for _, resRaw := range resultsRaw {
		blr, blrErr := processBlockResult(resRaw)
		if blrErr != nil {
			return nil, blrErr
		}
		bl.results = append(bl.results, blr)
	}

	transactionsRaw, transactionsOk := blockMap["transactions"].([]interface{})
	if !transactionsOk {
		return nil, errors.New("unmarshalled block transactions object is not a list")
	}
	for _, trRaw := range transactionsRaw {
		tr, trErr := processBlockTransaction(trRaw)
		if trErr != nil {
			return nil, trErr
		}
		bl.transactions = append(bl.transactions, tr)
	}

	blh, blhErr := processBlockHeader(blockMap["blockHeader"])
	if blhErr != nil {
		return nil, blhErr
	}
	bl.blockHeader = blh

	return bl, nil
}

func processBlockResult(blrRaw interface{}) (*blockResult, error) {
	blrMap, isMap := blrRaw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("unmarshalled block result is not a map")
	}

	blr := &blockResult{}
	var outOk, statusOk, gasOk, logsOk, refundOk bool

	blr.out, outOk = processBigIntList(blrMap["out"])
	if !outOk {
		return nil, errors.New("invalid block result out")
	}

	blr.status, statusOk = parseBigInt(blrMap["status"])
	if !statusOk {
		return nil, errors.New("invalid block result status")
	}

	blr.gas, gasOk = parseBigInt(blrMap["gas"])
	if !gasOk {
		return nil, errors.New("invalid block result gas")
	}

	blr.logs, logsOk = blrMap["logs"].(string)
	if !logsOk {
		return nil, errors.New("invalid block result logs")
	}

	blr.refund, refundOk = parseBigInt(blrMap["refund"])
	if !refundOk {
		return nil, errors.New("invalid block result refund")
	}

	return blr, nil
}

func processBlockTransaction(blrRaw interface{}) (*blockTransaction, error) {
	bltMap, isMap := blrRaw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("unmarshalled block transaction is not a map")
	}

	blt := &blockTransaction{}
	var nonceOk, functionOk, valueOk, toOk, argumentsOk, contractCodeOk, gasPriceOk, gasLimitOk, fromOk bool

	blt.nonce, nonceOk = parseBigInt(bltMap["nonce"])
	if !nonceOk {
		return nil, errors.New("invalid block transaction nonce")
	}

	blt.function, functionOk = bltMap["function"].(string)
	if !functionOk {
		return nil, errors.New("invalid block transaction function")
	}

	if ccRaw, ccPresent := bltMap["contractCode"]; ccPresent {
		blt.contractCode, contractCodeOk = ccRaw.(string)
		if !contractCodeOk {
			return nil, errors.New("invalid block transaction contract code")
		}
	}

	blt.value, valueOk = parseBigInt(bltMap["value"])
	if !valueOk {
		return nil, errors.New("invalid block transaction value")
	}

	blt.to, toOk = parseBigInt(bltMap["to"])
	if !toOk {
		return nil, errors.New("invalid block transaction to")
	}

	// note "to": "0x00" has to yield isCreate=false, even though it parses to 0, just like the 2 cases below
	blt.isCreate = bltMap["to"] == "" || bltMap["to"] == "0x"

	blt.arguments, argumentsOk = processBigIntList(bltMap["arguments"])
	if !argumentsOk {
		return nil, errors.New("invalid block transaction arguments")
	}

	blt.contractCode, contractCodeOk = bltMap["contractCode"].(string)
	if !contractCodeOk {
		return nil, errors.New("invalid block transaction contractCode")
	}

	blt.gasPrice, gasPriceOk = parseBigInt(bltMap["gasPrice"])
	if !gasPriceOk {
		return nil, errors.New("invalid block transaction gasPrice")
	}

	blt.gasLimit, gasLimitOk = parseBigInt(bltMap["gasLimit"])
	if !gasLimitOk {
		return nil, errors.New("invalid block transaction gasLimit")
	}

	blt.from, fromOk = parseBigInt(bltMap["from"])
	if !fromOk {
		return nil, errors.New("invalid block transaction from")
	}

	return blt, nil
}

func processBlockHeader(blhRaw interface{}) (*endpoint.BlockHeader, error) {
	blhMap, isMap := blhRaw.(map[string]interface{})
	if !isMap {
		return nil, errors.New("unmarshalled block header is not a map")
	}

	blh := &endpoint.BlockHeader{}
	var gasLimitOk, numberOk, difficultyOk, timestampOk, coinbaseOk bool

	blh.GasLimit, gasLimitOk = parseBigInt(blhMap["gasLimit"])
	if !gasLimitOk {
		return nil, errors.New("invalid block header gasLimit")
	}

	blh.Number, numberOk = parseBigInt(blhMap["number"])
	if !numberOk {
		return nil, errors.New("invalid block header number")
	}

	blh.Difficulty, difficultyOk = parseBigInt(blhMap["difficulty"])
	if !difficultyOk {
		return nil, errors.New("invalid block header difficulty")
	}

	blh.UnixTimestamp, timestampOk = parseBigInt(blhMap["timestamp"])
	if !timestampOk {
		return nil, errors.New("invalid block header timestamp")
	}

	blh.Beneficiary, coinbaseOk = parseBigInt(blhMap["coinbase"])
	if !coinbaseOk {
		return nil, errors.New("invalid block header coinbase")
	}

	return blh, nil
}

func processAccountAddress(addrRaw string) ([]byte, error) {
	if len(addrRaw) == 0 {
		return []byte{}, errors.New("missing account address")
	}
	if !(strings.HasPrefix(addrRaw, "0x") || strings.HasPrefix(addrRaw, "0X")) {
		return []byte{}, errors.New("account address should be hex representation starting with '0x'")
	}
	return hex.DecodeString(addrRaw[2:])
}

func processStringList(obj interface{}) ([]string, bool) {
	listRaw, listOk := obj.([]interface{})
	if !listOk {
		return nil, false
	}
	var result []string
	for _, elemRaw := range listRaw {
		str, strOk := elemRaw.(string)
		if !strOk {
			return nil, false
		}
		result = append(result, str)
	}
	return result, true
}

func processBigIntList(obj interface{}) ([]*big.Int, bool) {
	listRaw, listOk := obj.([]interface{})
	if !listOk {
		return nil, false
	}
	var result []*big.Int
	for _, elemRaw := range listRaw {
		i, iOk := parseBigInt(elemRaw)
		if !iOk {
			return nil, false
		}
		result = append(result, i)
	}
	return result, true
}

func parseBigInt(obj interface{}) (*big.Int, bool) {
	str, isStr := obj.(string)
	if !isStr {
		return nil, false
	}
	result := new(big.Int)
	var parseOk bool
	if len(str) > 0 { // interpret "" as 0
		result, parseOk = result.SetString(str, 0)
		if !parseOk {
			return nil, false
		}
	}

	return result, true
}
