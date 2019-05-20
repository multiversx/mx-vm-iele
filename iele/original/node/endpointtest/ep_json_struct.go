package endpoint

import (
	"math/big"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	endpoint "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
)

type testTopLevel struct {
	tests map[string]*test
}

type test struct {
	testName    string
	pre         world.AccountMap
	blocks      []*block
	network     string
	blockHashes []*big.Int
	postState   world.AccountMap
}

type block struct {
	results      []*blockResult
	transactions []*blockTransaction
	blockHeader  *endpoint.BlockHeader
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
	isCreate     bool
	from         *big.Int
	to           *big.Int
	function     string
	contractCode string
	arguments    []*big.Int
	gasPrice     *big.Int
	gasLimit     *big.Int
}
