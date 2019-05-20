package blockchainadapter

import (
	blockchain "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestingmodel"
)

// GetBalance ... adapts between K model and elrond function
func GetBalance(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.GetBalance(acct.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// GetNonce ... adapts between K model and elrond function
func GetNonce(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.GetNonce(acct.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// IsCodeEmpty ... adapts between K model and elrond function
func IsCodeEmpty(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.IsCodeEmpty(acct.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// AccountExists ... adapts between K model and elrond function
func AccountExists(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.AccountExists(acct.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// GetStorageData ... adapts between K model and elrond function
func GetStorageData(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt1 := c1.(*m.Int)
	if !isInt1 {
		return m.NoResult, &hookInvalidArgsError{}
	}
	index, isInt2 := c2.(*m.Int)
	if !isInt2 {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.GetStorageData(acct.Value, index.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// GetCode ... adapts between K model and elrond function
func GetCode(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.GetCode(acct.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(result), nil
}

// GetBlockhash ... adapts between K model and elrond function
func GetBlockhash(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	offset, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, &hookInvalidArgsError{}
	}
	result, err := blockchain.GetBlockhash(offset.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

type hookInvalidArgsError struct {
}

func (e *hookInvalidArgsError) Error() string {
	return "Invalid argument(s) provided to blockchain hook."
}
