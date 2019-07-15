package ielejson

import (
	"bytes"
	"math/big"
)

// Test is a json object representing a test
type Test struct {
	TestName    string
	CheckGas    bool
	Pre         []*Account
	Blocks      []*Block
	Network     string
	BlockHashes [][]byte
	PostState   []*Account
}

// Account is a json object representing an account
type Account struct {
	Address      []byte
	Nonce        *big.Int
	Balance      *big.Int
	Storage      []*StorageKeyValuePair
	Code         string
	OriginalCode string
}

// StorageKeyValuePair is a json key value pair in the storage map
type StorageKeyValuePair struct {
	Key   *big.Int
	Value *big.Int
}

// Block is a json object representing a block
type Block struct {
	Results      []*TransactionResult
	Transactions []*Transaction
	BlockHeader  *BlockHeader
}

// BlockHeader is a json object representing the block header
type BlockHeader struct {
	Beneficiary   *big.Int // "coinbase"
	Difficulty    *big.Int
	Number        *big.Int
	GasLimit      *big.Int
	UnixTimestamp *big.Int
}

// TransactionResult is a json object representing an expected transaction result
type TransactionResult struct {
	Out        []*big.Int
	Status     *big.Int
	Gas        *big.Int
	Refund     *big.Int
	IgnoreLogs bool
	LogHash    string
	Logs       []*LogEntry
}

// LogEntry is a json object representing an expected transaction result log entry
type LogEntry struct {
	Address []byte
	Topics  []*big.Int
	Data    []byte
}

// Transaction is a json object representing a transaction
type Transaction struct {
	Nonce         *big.Int
	Value         *big.Int
	IsCreate      bool
	From          []byte
	To            []byte
	Function      string
	ContractCode  string
	AssembledCode string
	Arguments     []*big.Int
	GasPrice      *big.Int
	GasLimit      *big.Int
}

// FindAccount by address
func FindAccount(accounts []*Account, address []byte) *Account {
	for _, acct := range accounts {
		if bytes.Equal(acct.Address, address) {
			return acct
		}
	}
	return nil
}
