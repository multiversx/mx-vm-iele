package endpoint

import (
	"errors"
	"math/big"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	vmi "github.com/ElrondNetwork/elrond-vm/iele/vm-interface"
)

// UpdateWorldStateBefore ... performs gas payment, before transaction
func UpdateWorldStateBefore(w world.WorldState, fromAddr []byte, gasLimit *big.Int, gasPrice *big.Int) error {
	acct, err := w.GetAccount(fromAddr)
	if err != nil {
		return err
	}
	acct.Nonce.Add(acct.Nonce, big.NewInt(1))
	gasPayment := big.NewInt(0).Mul(gasLimit, gasPrice)
	if acct.Balance.Cmp(gasPayment) < 0 {
		return errors.New("not enough balance to pay gas upfront")
	}
	acct.Balance.Sub(acct.Balance, gasPayment)
	w.UpdateBalance(fromAddr, acct.Balance)

	return nil
}

// UpdateWorldStateAfter ... updates accounts, after transaction
func UpdateWorldStateAfter(w world.WorldState, trOut *vmi.VMOutput) error {
	return w.UpdateAccounts(trOut.ModifiedAccounts, trOut.DeletedAccounts)
}
