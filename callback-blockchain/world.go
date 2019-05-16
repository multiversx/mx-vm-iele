package callbackblockchain

import (
	"math/big"
)

// StorageUpdate ... data pertaining changes in an account storage
type StorageUpdate struct {
	Offset *big.Int
	Data   *big.Int
}

// ModifiedAccount ... data indicating how an account must be modified after a contract call
type ModifiedAccount struct {
	Address        []byte
	Nonce          *big.Int
	Balance        *big.Int
	StorageUpdates []*StorageUpdate
	Code           string
}

// WorldState ...
type WorldState interface {
	GetAccount(address []byte) (*Account, error)
	GetStorageData(address []byte, offset *big.Int) (*big.Int, error)
	GetCode(address []byte) ([]byte, error)
	GetBlockhash(offset *big.Int) (*big.Int, error)

	UpdateBalance(address []byte, newBalance *big.Int) error
	UpdateAccounts(modifiedAccounts []*ModifiedAccount, accountsToDelete [][]byte) error
}
