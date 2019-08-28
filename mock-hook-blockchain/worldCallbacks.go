package callbackblockchain

import (
	"errors"
	"math/big"
)

var zero = big.NewInt(0)

// AccountExists note: an account with Balance = 0 and Nonce = 0 is considered to not exist
func (b *BlockchainHookMock) AccountExists(address []byte) (bool, error) {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return false, nil
	}
	return acct.Exists, nil
}

// NewAddress adapts between K model and elrond function
func (b *BlockchainHookMock) NewAddress(creatorAddress []byte, creatorNonce uint64) ([]byte, error) {
	// empty byte array signals not implemented, fallback to default
	return []byte{}, nil
}

// GetBalance should retrieve the balance of an account
func (b *BlockchainHookMock) GetBalance(address []byte) (*big.Int, error) {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return zero, nil
	}
	return acct.Balance, nil
}

// GetNonce should retrieve the nonce of an account
func (b *BlockchainHookMock) GetNonce(address []byte) (uint64, error) {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return 0, nil
	}
	return acct.Nonce, nil
}

// GetStorageData yields the storage value for a certain account and index.
// Should return an empty byte array if the key is missing from the account storage
func (b *BlockchainHookMock) GetStorageData(accountAddress []byte, index []byte) ([]byte, error) {
	acct := b.AcctMap.GetAccount(accountAddress)
	if acct == nil {
		return []byte{}, nil
	}
	return acct.StorageValue(string(index)), nil
}

// IsCodeEmpty should return whether of not an account is SC.
func (b *BlockchainHookMock) IsCodeEmpty(address []byte) (bool, error) {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return true, nil
	}
	return len(acct.Code) == 0, nil
}

// GetCode should return the compiled and assembled SC code.
// Empty byte array if the account is a wallet.
func (b *BlockchainHookMock) GetCode(address []byte) ([]byte, error) {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return []byte{}, nil
	}
	return acct.Code, nil
}

// GetBlockhash should return the hash of the nth previous blockchain.
// Offset specifies how many blocks we need to look back.
func (b *BlockchainHookMock) GetBlockhash(offset *big.Int) ([]byte, error) {
	if !offset.IsUint64() {
		return nil, errors.New("blockhash offset is too large")
	}
	offsetInt32 := int(offset.Int64())
	if offsetInt32 < 0 {
		return nil, errors.New("blockhash offset is negative")
	}
	if offsetInt32 >= len(b.Blockhashes) {
		return nil, errors.New("blockhash offset exceeds the blockhashes slice")
	}
	return b.Blockhashes[offsetInt32], nil
}
