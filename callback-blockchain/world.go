package callbackblockchain

import (
	"math/big"
)

// WorldState ...
type WorldState interface {
	GetAccount(address []byte) (*Account, error)
	GetStorageData(address []byte, offset *big.Int) (*big.Int, error)
	GetCode(address []byte) ([]byte, error)
	GetBlockhash(address []byte, offset int) (*big.Int, error)
}

// InMemoryWorldState ... world state model that is entirely kept in memory, used in tests
type InMemoryWorldState struct {
	AcctMap AccountMap
}

func MakeInMemoryWorldState() *InMemoryWorldState {
	return &InMemoryWorldState{AcctMap: MakeAccountMap()}
}

var zero = big.NewInt(0)

func (w *InMemoryWorldState) GetAccount(address []byte) (*Account, error) {
	acct := w.AcctMap.GetAccount(address)
	if acct == nil {
		// create default and save that default in map
		acct = &Account{
			Address: address,
			Nonce:   zero,
			Balance: zero,
			Storage: make(map[string]*big.Int),
			Code:    "",
		}
		w.AcctMap.PutAccount(acct)
	}
	return acct, nil
}

func (w *InMemoryWorldState) GetStorageData(address []byte, offset *big.Int) (*big.Int, error) {
	acct := w.AcctMap.GetAccount(address)
	if acct == nil {
		return zero, nil
	}
	offsetKey := offset.Text(16)
	val, found := acct.Storage[offsetKey]
	if !found {
		val = zero
		acct.Storage[offsetKey] = val
	}
	return val, nil
}

func (w *InMemoryWorldState) GetCode(address []byte) ([]byte, error) {
	panic("not implemented")
}

func (w *InMemoryWorldState) GetBlockhash(address []byte, offset int) (*big.Int, error) {
	panic("not implemented")
}
