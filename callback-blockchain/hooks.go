package callbackblockchain

import (
	"errors"
	"math/big"
)

// HookWorldState ... world state to be used by the BLOCKCHAIN hook
var HookWorldState WorldState

// GetBalance ... balance of account, hook endpoint
func GetBalance(kaddr *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return nil, err
	}
	if acct == nil {
		return nil, errors.New("account not found")
	}
	return acct.Balance, nil
}

// GetNonce ... nonce of account, hook endpoint
func GetNonce(kaddr *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return nil, err
	}
	if acct == nil {
		return nil, errors.New("account not found")
	}
	return acct.Nonce, nil
}

// AccountExists ... true if account exists, hook endpoint
func AccountExists(kaddr *big.Int) (bool, error) {
	if HookWorldState == nil {
		return false, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return false, err
	}
	exist := acct.Balance.Sign() > 0 && acct.Nonce.Sign() > 0
	return exist, nil
}

// GetStorageData ... data at address/offset, hook endpoint
func GetStorageData(kaddr *big.Int, index *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	return HookWorldState.GetStorageData(addr, index)
}

// IsCodeEmpty ... true if account code is empty, hook endpoint
func IsCodeEmpty(kaddr *big.Int) (bool, error) {
	if HookWorldState == nil {
		return false, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	code, err := HookWorldState.GetCode(addr)
	if err != nil {
		return true, err
	}
	codeEmpty := len(code) == 0
	return codeEmpty, nil
}

// GetCode ... account code, hook endpoint
func GetCode(kaddr *big.Int) (string, error) {
	if HookWorldState == nil {
		return "", errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	code, err := HookWorldState.GetCode(addr)
	if err != nil {
		return "", err
	}
	return string(code), nil
}

// GetBlockhash ... nth previous blockhash, hook endpoint
func GetBlockhash(offset *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	return HookWorldState.GetBlockhash(offset)
}
