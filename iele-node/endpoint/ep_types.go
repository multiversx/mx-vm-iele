package endpoint

import (
	"math/big"
)

// BlockHeader ...
type BlockHeader struct {
	Beneficiary   *big.Int // "coinbase"
	Difficulty    *big.Int
	Number        *big.Int
	GasLimit      *big.Int
	UnixTimestamp *big.Int
}

// VMInput ... transaction!
type VMInput struct {
	CallerAddr    *big.Int
	RecipientAddr *big.Int
	InputData     string
	Function      string
	Arguments     []*big.Int
	CallValue     *big.Int
	GasPrice      *big.Int
	GasProvided   *big.Int
	BlockHeader   *BlockHeader
	Schedule      Schedule
}

// StorageUpdate ... data pertaining changes in an account storage
type StorageUpdate struct {
	Offset *big.Int
	Data   *big.Int
}

// ModifiedAccount ... data containing how an account is being modified by a contract call
type ModifiedAccount struct {
	Address        *big.Int
	Nonce          *big.Int
	Balance        *big.Int
	StorageUpdates []*StorageUpdate
	Code           string
}

// LogEntry ... contract execution log
type LogEntry struct {
	Address *big.Int
	topics  []*big.Int
	data    []byte
}

// VMOutput ...
type VMOutput struct {
	ReturnData       []*big.Int
	ReturnCode       *big.Int
	GasRemaining     *big.Int
	GasRefund        *big.Int
	Error            bool
	ModifiedAccounts []*ModifiedAccount
	DeletedAccounts  []*big.Int
	TouchedAccounts  []*big.Int
	Logs             []*LogEntry
}
