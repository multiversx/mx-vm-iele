package endpoint

import (
	"math/big"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
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

// LogEntry ... contract execution log
type LogEntry struct {
	Address *big.Int
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
