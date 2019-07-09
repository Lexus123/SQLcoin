package custommodels

import "github.com/btcsuite/btcd/chaincfg/chainhash"

type DbBlock struct {
	Hash      chainhash.Hash
	PrevHash  [32]byte
	Merkle    [32]byte
	Timestamp int64
}

type DbTx struct {
	Hash    [32]byte
	BlockId int64
}

type DbOutput struct {
	BlockId         int64
	TxHash          [32]byte
	Index           int
	Address         string
	Amount          int64
	SpendingBlockId int64
	SpendingTxHash  [32]byte
	SpendingIndex   int
}

type DbInput struct {
	TxHash      string
	PrevTxHash  string
	PrevTxIndex uint32
}
