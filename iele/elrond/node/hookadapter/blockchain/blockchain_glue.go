package blockchainadapter

import (
	"errors"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// Blockchain is an adapter between K and the outside world
type Blockchain struct {
	Upstream vmi.BlockchainHook
}

// GetBalance ... adapts between K model and elrond function
func (b *Blockchain) GetBalance(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetBalance(acct.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// GetNonce ... adapts between K model and elrond function
func (b *Blockchain) GetNonce(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetNonce(acct.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// IsCodeEmpty ... adapts between K model and elrond function
func (b *Blockchain) IsCodeEmpty(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.IsCodeEmpty(acct.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// AccountExists ... adapts between K model and elrond function
func (b *Blockchain) AccountExists(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.AccountExists(acct.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// GetStorageData ... adapts between K model and elrond function
func (b *Blockchain) GetStorageData(c1 m.K, c2 m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt1 := c1.(*m.Int)
	if !isInt1 {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	index, isInt2 := c2.(*m.Int)
	if !isInt2 {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetStorageData(acct.Value.Bytes(), index.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.NewIntFromBytes(result), nil
}

// GetCode ... adapts between K model and elrond function
func (b *Blockchain) GetCode(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acct, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetCode(acct.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(string(result)), nil
}

// GetBlockhash ... adapts between K model and elrond function
func (b *Blockchain) GetBlockhash(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	offset, isInt := c.(*m.Int)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetBlockhash(offset.Value)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewIntFromBytes(result), nil
}
