package ielejson

import (
	"bytes"
	"math/big"
)

// Test ...
type Test struct {
	TestName    string
	Pre         []*Account
	Blocks      []*Block
	Network     string
	BlockHashes []*big.Int
	PostState   []*Account
}

// Account ...
type Account struct {
	Address      []byte
	Nonce        *big.Int
	Balance      *big.Int
	Storage      []*StorageKeyValuePair
	Code         string
	OriginalCode string
}

// StorageKeyValuePair ...
type StorageKeyValuePair struct {
	Key   *big.Int
	Value *big.Int
}

// Block ...
type Block struct {
	Results      []*TransactionResult
	Transactions []*Transaction
	BlockHeader  *BlockHeader
}

// BlockHeader ...
type BlockHeader struct {
	Beneficiary   *big.Int // "coinbase"
	Difficulty    *big.Int
	Number        *big.Int
	GasLimit      *big.Int
	UnixTimestamp *big.Int
}

// TransactionResult ...
type TransactionResult struct {
	Out    []*big.Int
	Status *big.Int
	Gas    *big.Int
	Refund *big.Int
	Logs   string
}

// Transaction ...
type Transaction struct {
	Nonce        *big.Int
	Value        *big.Int
	IsCreate     bool
	From         []byte
	To           []byte
	Function     string
	ContractCode string
	Arguments    []*big.Int
	GasPrice     *big.Int
	GasLimit     *big.Int
}

// FindAccount ... by address
func FindAccount(accounts []*Account, address []byte) *Account {
	for _, acct := range accounts {
		if bytes.Equal(acct.Address, address) {
			return acct
		}
	}
	return nil
}
