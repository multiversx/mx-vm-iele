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

// AccountAddress ... convert to account address bytes from big.Int
func AccountAddress(i *big.Int) []byte {
	if i.Sign() < 0 {
		panic("address cannot be negative")
	}
	return i.Bytes()
}
