package callbackblockchain

import (
	"errors"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

func defaultAccount(address []byte) *Account {
	return &Account{
		Exists:  false,
		Address: address,
		Nonce:   zero,
		Balance: zero,
		Storage: make(map[string][]byte),
		Code:    nil,
	}
}

// UpdateBalance changes balance of an account
func (b *BlockchainHookMock) UpdateBalance(address []byte, newBalance *big.Int) error {
	acct := b.AcctMap.GetAccount(address)
	if acct == nil {
		return errors.New("method UpdateBalance expects an existing address")
	}
	acct.Balance = newBalance
	return nil
}

// UpdateWorldStateBefore performs gas payment, before transaction
func (b *BlockchainHookMock) UpdateWorldStateBefore(
	fromAddr []byte,
	gasLimit *big.Int,
	gasPrice *big.Int) error {

	acct := b.AcctMap.GetAccount(fromAddr)
	if acct == nil {
		return errors.New("method UpdateBalance expects an existing address")
	}
	acct.Nonce.Add(acct.Nonce, big.NewInt(1))
	gasPayment := big.NewInt(0).Mul(gasLimit, gasPrice)
	if acct.Balance.Cmp(gasPayment) < 0 {
		return errors.New("not enough balance to pay gas upfront")
	}
	acct.Balance.Sub(acct.Balance, gasPayment)
	return nil
}

// UpdateAccounts should be called after the VM test has run, to update world state
func (b *BlockchainHookMock) UpdateAccounts(modifiedAccounts []*vmi.OutputAccount, accountsToDelete [][]byte) error {
	for _, modAcct := range modifiedAccounts {
		acct := b.AcctMap.GetAccount(modAcct.Address)
		if acct == nil {
			acct = defaultAccount(modAcct.Address)
			b.AcctMap.PutAccount(acct)
		}
		acct.Exists = true
		acct.Balance = modAcct.Balance
		acct.Nonce = modAcct.Nonce
		acct.Code = modAcct.Code

		for _, stu := range modAcct.StorageUpdates {
			acct.Storage[string(stu.Offset)] = stu.Data
		}
	}

	for _, delAddr := range accountsToDelete {
		b.AcctMap.DeleteAccount(delAddr)
	}

	return nil

}
