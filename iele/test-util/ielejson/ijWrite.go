package ielejson

import (
	"encoding/hex"
	"math/big"

	oj "github.com/ElrondNetwork/elrond-vm/iele/test-util/orderedjson"
)

// ToJSONString ...
func ToJSONString(testTopLevel []*Test) string {
	jobj := ToOrderedJSON(testTopLevel)
	return oj.JSONString(jobj)
}

// ToOrderedJSON ...
func ToOrderedJSON(testTopLevel []*Test) oj.OJsonObject {

	result := oj.NewMap()
	for _, test := range testTopLevel {
		result.Put(test.TestName, testToOJ(test))
	}

	return result
}

func testToOJ(test *Test) oj.OJsonObject {
	testOJ := oj.NewMap()

	if !test.CheckGas {
		ojFalse := oj.OJsonBool(false)
		testOJ.Put("checkGas", &ojFalse)
	}

	testOJ.Put("pre", accountsToOJ(test.Pre))

	var blockList []oj.OJsonObject
	for _, block := range test.Blocks {
		blockList = append(blockList, blockToOJ(block))
	}
	blocksOJ := oj.OJsonList(blockList)
	testOJ.Put("blocks", &blocksOJ)

	testOJ.Put("network", stringToOJ(test.Network))

	var blockhashesList []oj.OJsonObject
	for _, blh := range test.BlockHashes {
		blockhashesList = append(blockhashesList, stringToOJ(byteArrayToString(blh)))
	}
	blockHashesOJ := oj.OJsonList(blockhashesList)
	testOJ.Put("blockhashes", &blockHashesOJ)

	testOJ.Put("postState", accountsToOJ(test.PostState))
	return testOJ
}

func accountsToOJ(accounts []*Account) oj.OJsonObject {
	acctsOJ := oj.NewMap()
	for _, account := range accounts {
		acctOJ := oj.NewMap()
		acctOJ.Put("nonce", intToOJ(account.Nonce))
		acctOJ.Put("balance", intToOJ(account.Balance))
		storageOJ := oj.NewMap()
		for _, st := range account.Storage {
			storageOJ.Put(intToString(st.Key), intToOJ(st.Value))
		}
		acctOJ.Put("storage", storageOJ)
		acctOJ.Put("code", stringToOJ(account.OriginalCode))

		acctsOJ.Put(byteArrayToString(account.Address), acctOJ)
	}

	return acctsOJ
}

func blockToOJ(block *Block) oj.OJsonObject {
	blockOJ := oj.NewMap()

	var resultList []oj.OJsonObject
	for _, blr := range block.Results {
		resultList = append(resultList, resultToOJ(blr))
	}
	resultsOJ := oj.OJsonList(resultList)
	blockOJ.Put("results", &resultsOJ)

	var txList []oj.OJsonObject
	for _, tx := range block.Transactions {
		txList = append(txList, transactionToOJ(tx))
	}
	txsOJ := oj.OJsonList(txList)
	blockOJ.Put("transactions", &txsOJ)

	blockHeaderOJ := oj.NewMap()
	blockHeaderOJ.Put("gasLimit", intToOJ(block.BlockHeader.GasLimit))
	blockHeaderOJ.Put("number", intToOJ(block.BlockHeader.Number))
	blockHeaderOJ.Put("difficulty", intToOJ(block.BlockHeader.Difficulty))
	blockHeaderOJ.Put("timestamp", intToOJ(block.BlockHeader.UnixTimestamp))
	blockHeaderOJ.Put("coinbase", intToOJ(block.BlockHeader.Beneficiary))
	blockOJ.Put("blockHeader", blockHeaderOJ)

	return blockOJ
}

func transactionToOJ(tx *Transaction) oj.OJsonObject {
	transactionOJ := oj.NewMap()
	transactionOJ.Put("nonce", intToOJ(tx.Nonce))
	transactionOJ.Put("function", stringToOJ(tx.Function))
	transactionOJ.Put("gasLimit", intToOJ(tx.GasLimit))
	transactionOJ.Put("value", intToOJ(tx.Value))
	transactionOJ.Put("to", accountAddressToOJ(tx.To))

	var argList []oj.OJsonObject
	for _, arg := range tx.Arguments {
		argList = append(argList, intToOJ(arg))
	}
	argOJ := oj.OJsonList(argList)
	transactionOJ.Put("arguments", &argOJ)

	transactionOJ.Put("contractCode", stringToOJ(tx.ContractCode))
	transactionOJ.Put("gasPrice", intToOJ(tx.GasPrice))
	transactionOJ.Put("from", accountAddressToOJ(tx.From))

	return transactionOJ
}

func resultToOJ(res *TransactionResult) oj.OJsonObject {
	resultOJ := oj.NewMap()

	var outList []oj.OJsonObject
	for _, out := range res.Out {
		outList = append(outList, intToOJ(out))
	}
	outOJ := oj.OJsonList(outList)
	resultOJ.Put("out", &outOJ)

	resultOJ.Put("status", intToOJ(res.Status))
	resultOJ.Put("gas", intToOJ(res.Gas))
	if res.IgnoreLogs {
		resultOJ.Put("logs", stringToOJ("*"))
	} else {
		if len(res.LogHash) > 0 {
			resultOJ.Put("logs", stringToOJ(res.LogHash))
		} else {
			resultOJ.Put("logs", logsToOJ(res.Logs))
		}
	}
	resultOJ.Put("refund", intToOJ(res.Refund))

	return resultOJ
}

// LogToString returns a json representation of a log entry, we use it for debugging
func LogToString(logEntry *LogEntry) string {
	logOJ := logToOJ(logEntry)
	return oj.JSONString(logOJ)
}

func logToOJ(logEntry *LogEntry) oj.OJsonObject {
	logOJ := oj.NewMap()
	logOJ.Put("address", accountAddressToOJ(logEntry.Address))

	var topicsList []oj.OJsonObject
	for _, topic := range logEntry.Topics {
		topicsList = append(topicsList, intToOJ(topic))
	}
	topicsOJ := oj.OJsonList(topicsList)
	logOJ.Put("topics", &topicsOJ)

	dataAsInt := big.NewInt(0).SetBytes(logEntry.Data)
	logOJ.Put("data", intToOJ(dataAsInt))

	return logOJ
}

func logsToOJ(logEntries []*LogEntry) oj.OJsonObject {
	var logList []oj.OJsonObject
	for _, logEntry := range logEntries {
		logOJ := logToOJ(logEntry)
		logList = append(logList, logOJ)
	}
	logOJList := oj.OJsonList(logList)
	return &logOJList
}

func byteArrayToString(byteArray []byte) string {
	if len(byteArray) == 0 {
		return "0x00"
	}
	return "0x" + hex.EncodeToString(byteArray)
}

func accountAddressToOJ(address []byte) oj.OJsonObject {
	if len(address) == 0 {
		return stringToOJ("")
	}
	return stringToOJ(byteArrayToString(address))
}

func intToString(i *big.Int) string {
	if i == nil {
		return ""
	}
	if i.Sign() == 0 {
		return "0x00"
	}

	isNegative := i.Sign() == -1
	str := i.Text(16)
	if isNegative {
		str = str[1:] // drop the minus in front
	}
	if len(str)%2 != 0 {
		str = "0" + str
	}
	str = "0x" + str
	if isNegative {
		str = "-" + str
	}
	return str
}

func intToOJ(i *big.Int) oj.OJsonObject {
	return &oj.OJsonString{Value: intToString(i)}
}

func stringToOJ(str string) oj.OJsonObject {
	return &oj.OJsonString{Value: str}
}
