package blockchainadapter

import (
	"errors"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// Blockchain is an adapter between K and the outside world
// This class is specific to only 1 generated interpreter
type Blockchain struct {
	// Upstream is the world state callback, which is common to all VMs
	Upstream vmi.BlockchainHook

	// AddressLength is the expected length of an address, in bytes
	AddressLength int
}

// ConvertKIntToAddress takes a K Int and converts it to an address with the correct number of bytes,
// will pad left with zeroes, based on the configured address length
func (b *Blockchain) ConvertKIntToAddress(addrAsK m.K) ([]byte, bool) {
	addrInt, isInt := addrAsK.(*m.Int)
	if !isInt {
		return []byte{}, false
	}
	addrBytes := addrInt.Value.Bytes()
	if len(addrBytes) > b.AddressLength {
		return []byte{}, false
	}
	result := make([]byte, b.AddressLength)

	i := len(addrBytes) - 1
	j := b.AddressLength - 1
	for i >= 0 {
		result[j] = addrBytes[i]
		i--
		j--
	}

	return result, true
}

// GetBalance adapts between K model and elrond function
func (b *Blockchain) GetBalance(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetBalance")
	}
	result, err := b.Upstream.GetBalance(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// GetNonce adapts between K model and elrond function
func (b *Blockchain) GetNonce(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetNonce")
	}
	result, err := b.Upstream.GetNonce(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewInt(result), nil
}

// IsCodeEmpty adapts between K model and elrond function
func (b *Blockchain) IsCodeEmpty(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook IsCodeEmpty")
	}
	result, err := b.Upstream.IsCodeEmpty(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// AccountExists adapts between K model and elrond function
func (b *Blockchain) AccountExists(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook AccountExists")
	}
	result, err := b.Upstream.AccountExists(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToBool(result), nil
}

// GetStorageData adapts between K model and elrond function
func (b *Blockchain) GetStorageData(kaddr m.K, kindex m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(kaddr)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetStorageData")
	}
	index, isInt2 := kindex.(*m.Int)
	if !isInt2 {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetStorageData(acctAddr, index.Value.Bytes())
	if err != nil {
		return m.NoResult, err
	}
	return m.NewIntFromBytes(result), nil
}

// GetCode adapts between K model and elrond function
func (b *Blockchain) GetCode(c m.K, lbl m.KLabel, sort m.Sort, config m.K) (m.K, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetCode")
	}
	result, err := b.Upstream.GetCode(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.NewString(string(result)), nil
}

// GetBlockhash adapts between K model and elrond function
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
