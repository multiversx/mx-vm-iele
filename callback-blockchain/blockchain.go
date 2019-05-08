package callbackblockchain

import (
	"errors"
	"math/big"
)

// HookWorldState ... world state to be used by the BLOCKCHAIN hook
var HookWorldState WorldState

// GetBalance ... balance of account
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

// GetNonce ... nonce of account
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
	return acct.Balance, nil
}

// IsCodeEmpty ... true if account code is empty
func IsCodeEmpty(kaddr *big.Int) (bool, error) {
	if HookWorldState == nil {
		return false, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return true, err
	}
	if acct == nil {
		return true, errors.New("account not found")
	}
	codeEmpty := len(acct.Code) == 0
	return codeEmpty, nil
}

// AccountExists ... true if account exists
func AccountExists(kaddr *big.Int) (bool, error) {
	if HookWorldState == nil {
		return false, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return false, err
	}
	return acct != nil, nil
}

func GetStorageData(kaddr *big.Int, index *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	return HookWorldState.GetStorageData(addr, index)
}

func GetCode(kaddr *big.Int) (string, error) {
	if HookWorldState == nil {
		return "", errors.New("world state manager not initialized")
	}
	addr := AccountAddress(kaddr)
	acct, err := HookWorldState.GetAccount(addr)
	if err != nil {
		return "", err
	}
	if acct == nil {
		return "", errors.New("account not found")
	}
	return acct.Code, nil
}

func GetBlockhash(offset *big.Int) (*big.Int, error) {
	if HookWorldState == nil {
		return nil, errors.New("world state manager not initialized")
	}
	return big.NewInt(0), errors.New("GetBlockhash not implemented")
}
