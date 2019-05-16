package callbackblockchain

import "math/big"

// AccountMap ... a map from address to account
type AccountMap map[string]*Account

// Account ... account info
type Account struct {
	Address []byte
	Nonce   *big.Int
	Balance *big.Int
	Storage map[string]*big.Int
	Code    string
}

var storageDefaultValue = big.NewInt(0)

// MakeAccountMap ... new AccountMap instance
func MakeAccountMap() AccountMap {
	return AccountMap(make(map[string]*Account))
}

// PutAccount ... insert account based on address
func (am AccountMap) PutAccount(acct *Account) {
	mp := (map[string]*Account)(am)
	mp[string(acct.Address)] = acct
}

// GetAccount ... retrieve account based on address
func (am AccountMap) GetAccount(address []byte) *Account {
	mp := (map[string]*Account)(am)
	return mp[string(address)]
}

// DeleteAccount ... remove account based on address
func (am AccountMap) DeleteAccount(address []byte) {
	mp := (map[string]*Account)(am)
	delete(mp, string(address))
}

// StorageValue ... storage value for key, default 0
func (a *Account) StorageValue(key string) *big.Int {
	value, found := a.Storage[key]
	if !found {
		return storageDefaultValue
	}
	return value
}

// AccountAddress ... convert to account address bytes from big.Int
func AccountAddress(i *big.Int) []byte {
	if i.Sign() < 0 {
		panic("address cannot be negative")
	}
	return i.Bytes()
}
