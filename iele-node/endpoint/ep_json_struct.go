package endpoint

import (
	"math/big"

	eh "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
)

type testTopLevel struct {
	tests map[string]*test
}

type test struct {
	testName    string
	pre         eh.AccountMap
	blocks      []*block
	network     string
	blockHashes []string
	postState   eh.AccountMap
}

type block struct {
	results      []*blockResult
	transactions []*blockTransaction
	blockHeader  *BlockHeader
}

type blockResult struct {
	out    []*big.Int
	status *big.Int
	gas    *big.Int
	refund *big.Int
	logs   string
}

type blockTransaction struct {
	nonce        *big.Int
	value        *big.Int
	from         *big.Int
	to           *big.Int
	function     string
	contractCode string
	arguments    []*big.Int
	gasPrice     *big.Int
	gasLimit     *big.Int
}

/*
type blockHeader struct {
	gasLimit   string
	number     string
	difficulty string
	timestamp  string
	coinbase   string
}*/
