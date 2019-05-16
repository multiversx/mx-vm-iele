package callbackblockchain

import (
	"errors"
	"math/big"
)

// InMemoryWorldState ... world state model that is entirely kept in memory, used in tests
type InMemoryWorldState struct {
	AcctMap     AccountMap
	Blockhashes []*big.Int
}

// MakeInMemoryWorldState ... create new InMemoryWorldState instance
func MakeInMemoryWorldState() *InMemoryWorldState {
	return &InMemoryWorldState{AcctMap: MakeAccountMap()}
}

var zero = big.NewInt(0)

func defaultAccount(address []byte) *Account {
	return &Account{
		Address: address,
		Nonce:   zero,
		Balance: zero,
		Storage: make(map[string]*big.Int),
		Code:    "",
	}
}

// GetAccount ... get basic account data (no storage or code required here)
func (w *InMemoryWorldState) GetAccount(address []byte) (*Account, error) {
	acct := w.AcctMap.GetAccount(address)
	if acct == nil {
		// return default account
		// do not save it in the map
		// nothing gets saved before the VM is done, this account might get removed by #trimAccounts
		return defaultAccount(address), nil
	}
	return acct, nil
}

// GetStorageData ... load storage data for an account, for given storage data key (or offset)
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

// GetCode ... load account code and return
func (w *InMemoryWorldState) GetCode(address []byte) ([]byte, error) {
	acct := w.AcctMap.GetAccount(address)
	if acct == nil {
		return []byte{}, nil
	}
	if acct == nil {
		return []byte{}, nil
	}
	return []byte(acct.Code), nil
}

// GetBlockhash ... get nth previous blockhash, offset indicates how many blockhashes to go back
func (w *InMemoryWorldState) GetBlockhash(offset *big.Int) (*big.Int, error) {
	if !offset.IsUint64() {
		return nil, errors.New("blockhash offset is too large")
	}
	offsetInt32 := int(offset.Int64())
	if offsetInt32 < 0 {
		return nil, errors.New("blockhash offset is negative")
	}
	if offsetInt32 >= len(w.Blockhashes) {
		return nil, errors.New("blockhash offset exceeds the blockhashes slice")
	}
	return w.Blockhashes[offsetInt32], nil
}

// UpdateBalance ... change balance of an account
func (w *InMemoryWorldState) UpdateBalance(address []byte, newBalance *big.Int) error {
	acct := w.AcctMap.GetAccount(address)
	acct.Balance = newBalance
	w.AcctMap.PutAccount(acct)
	return nil
}

// UpdateAccounts ... this method should be called after the VM has run to update world state
func (w *InMemoryWorldState) UpdateAccounts(modifiedAccounts []*ModifiedAccount, accountsToDelete [][]byte) error {
	for _, modAcct := range modifiedAccounts {
		acct := w.AcctMap.GetAccount(modAcct.Address)
		if acct == nil {
			acct = defaultAccount(modAcct.Address)
			w.AcctMap.PutAccount(acct)
		}
		acct.Balance = modAcct.Balance
		acct.Nonce = modAcct.Nonce
		acct.Code = modAcct.Code

		for _, stu := range modAcct.StorageUpdates {
			acct.Storage[stu.Offset.Text(16)] = stu.Data
		}
	}

	for _, delAddr := range accountsToDelete {
		w.AcctMap.DeleteAccount(delAddr)
	}

	return nil

}
