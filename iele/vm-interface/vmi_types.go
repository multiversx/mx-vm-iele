package vminterface

import (
	"math/big"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
)

// Schedule ... IELE gas model type
type Schedule int

const (
	// Default ... IELE default gas model
	Default Schedule = iota

	// Albe ... IELE "ALBE" gas model, this was their first version
	Albe

	// Danse ... IELE "DANSE" gas model, this is the latest version
	Danse
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
	IsCreate      bool
	CallerAddr    []byte
	RecipientAddr []byte
	InputData     string
	Function      string
	Arguments     []*big.Int
	CallValue     *big.Int
	GasPrice      *big.Int
	GasProvided   *big.Int
	BlockHeader   *BlockHeader
	Schedule      Schedule
}

// LogEntry ... contract execution log
type LogEntry struct {
	Address []byte
	Topics  []*big.Int
	Data    []byte
}

// VMOutput ...
type VMOutput struct {
	ReturnData       []*big.Int
	ReturnCode       *big.Int
	GasRemaining     *big.Int
	GasRefund        *big.Int
	Error            bool
	ModifiedAccounts []*world.ModifiedAccount
	DeletedAccounts  [][]byte
	TouchedAccounts  [][]byte
	Logs             []*LogEntry
}
