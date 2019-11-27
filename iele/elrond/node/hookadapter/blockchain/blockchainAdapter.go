package blockchainadapter

import (
	"bytes"
	"errors"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

// Blockchain is an adapter between K and the outside world
// This class is specific to only 1 generated interpreter
type Blockchain struct {
	// Upstream is the world state callback, which is common to all VMs
	Upstream vmi.BlockchainHook

	// VMType is a VM indentifier used for generating new addresses.
	VMType []byte

	// AddressLength is the expected length of an address, in bytes
	AddressLength int

	// initialBalances stores the balances of accounts as hook GetBalance is called.
	// It acts as a cache. It is also used to compute balance delta.
	initialBalances map[string]*big.Int

	// the sender gets the call value added to the initial balance,
	// which then gets subtracted by the VM
	senderAddress []byte
	callValue     *big.Int

	// LogToConsole when set to true causes the adapter to also print all operations to console
	LogToConsole bool

	// inputAccounts is a structure where we store account data as it was loaded in the interpreter,
	// if logging is enabled.
	// Used for logging only.
	inputAccounts map[string]*vmi.OutputAccount
}

// InitAdapter should be called before each SC execution.
func (b *Blockchain) InitAdapter(senderAddress []byte, callValue *big.Int) {
	b.initialBalances = make(map[string]*big.Int)
	b.senderAddress = senderAddress
	b.callValue = callValue
}

// ConvertKIntToAddress takes a K Int and converts it to an address with the correct number of bytes,
// will pad left with zeroes, based on the configured address length
func (b *Blockchain) ConvertKIntToAddress(addrAsK m.KReference, ms *m.ModelState) ([]byte, bool) {
	addrInt, isInt := ms.GetBigInt(addrAsK)
	if !isInt {
		return []byte{}, false
	}
	addrBytes := addrInt.Bytes()
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

// AccountExists adapts between K model and elrond function
func (b *Blockchain) AccountExists(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook AccountExists")
	}
	result, err := b.Upstream.AccountExists(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToKBool(result), nil
}

// NewAddress adapts between K model and elrond function
func (b *Blockchain) NewAddress(kaddr, knonce m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	creatorAddr, isAddr := b.ConvertKIntToAddress(kaddr, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid creator address provided to blockchain hook NewAddress")
	}
	creatorNonce, nonceOk := ms.GetUint64(knonce)
	if !nonceOk {
		return m.NoResult, errors.New("invalid creator nonce provided to blockchain hook NewAddress")
	}

	newAddr, err := b.Upstream.NewAddress(creatorAddr, creatorNonce, b.VMType)
	if err != nil {
		return m.NoResult, err
	}

	b.logNewAddress(creatorAddr, creatorNonce, newAddr)

	if len(newAddr) == 0 {
		// signal the interpreter that the alternate K implementation should be used
		return m.NoResult, m.GetHookNotImplementedError()
	}
	return ms.IntFromBytes(newAddr), nil
}

// GetBalance adapts between K model and elrond function
func (b *Blockchain) GetBalance(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetBalance")
	}
	if initialBalance, balLoaded := b.initialBalances[string(acctAddr)]; balLoaded {
		return ms.FromBigInt(initialBalance), nil
	}
	result, err := b.Upstream.GetBalance(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	b.initialBalances[string(acctAddr)] = result
	b.logBalance(acctAddr, result)
	return ms.FromBigInt(result), nil
}

// ComputeDelta computes the delta from a final balance of an account
func (b *Blockchain) ComputeDelta(acctAddr []byte, finalBalance *big.Int) (*big.Int, error) {
	initialBalance, initialBalanceExists := b.initialBalances[string(acctAddr)]
	if !initialBalanceExists {
		return nil, errors.New("output account balance does not have a corresponding input balance")
	}
	balanceDelta := big.NewInt(0).Sub(finalBalance, initialBalance)
	if bytes.Equal(acctAddr, b.senderAddress) {
		// discard call value subtraction from balance delta
		// by performing the opposite operation (add)
		balanceDelta = balanceDelta.Add(balanceDelta, b.callValue)
	}
	return balanceDelta, nil
}

// GetNonce adapts between K model and elrond function
func (b *Blockchain) GetNonce(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetNonce")
	}
	result, err := b.Upstream.GetNonce(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	b.logNonce(acctAddr, result)
	return ms.FromUint64(result), nil
}

// IsCodeEmpty adapts between K model and elrond function
func (b *Blockchain) IsCodeEmpty(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook IsCodeEmpty")
	}
	result, err := b.Upstream.IsCodeEmpty(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return m.ToKBool(result), nil
}

// GetStorageData adapts between K model and elrond function
func (b *Blockchain) GetStorageData(kaddr m.KReference, kindex m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(kaddr, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetStorageData")
	}
	index, isInt2 := ms.GetBigInt(kindex)
	if !isInt2 {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	indexBytes := index.Bytes()
	result, err := b.Upstream.GetStorageData(acctAddr, indexBytes)
	if err != nil {
		return m.NoResult, err
	}
	b.logStorage(acctAddr, indexBytes, result)
	return ms.IntFromBytes(result), nil
}

// GetCode adapts between K model and elrond function
func (b *Blockchain) GetCode(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	acctAddr, isAddr := b.ConvertKIntToAddress(c, ms)
	if !isAddr {
		return m.NoResult, errors.New("invalid account address provided to blockchain hook GetCode")
	}
	result, err := b.Upstream.GetCode(acctAddr)
	if err != nil {
		return m.NoResult, err
	}
	return ms.NewString(string(result)), nil
}

// GetBlockhash adapts between K model and elrond function
func (b *Blockchain) GetBlockhash(c m.KReference, lbl m.KLabel, sort m.Sort, config m.KReference, ms *m.ModelState) (m.KReference, error) {
	nonce, isInt := ms.GetUint64(c)
	if !isInt {
		return m.NoResult, errors.New("invalid argument(s) provided to blockchain hook")
	}
	result, err := b.Upstream.GetBlockhash(nonce)
	if err != nil {
		return m.NoResult, err
	}
	return ms.IntFromBytes(result), nil
}
