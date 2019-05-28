package callbackblockchain

// BlockchainHookMock
type BlockchainHookMock struct {
	AcctMap     AccountMap
	Blockhashes [][]byte
}

func NewMock() *BlockchainHookMock {
	return &BlockchainHookMock{
		AcctMap:     NewAccountMap(),
		Blockhashes: nil,
	}
}

func (b *BlockchainHookMock) Clear() {
	b.AcctMap = NewAccountMap()
	b.Blockhashes = nil
}
