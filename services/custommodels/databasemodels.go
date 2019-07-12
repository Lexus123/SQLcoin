package custommodels

import "github.com/btcsuite/btcd/chaincfg/chainhash"

/*
DbOutput ...
*/
type DbOutput struct {
	BlockHash         chainhash.Hash
	TxHash            chainhash.Hash
	Index             int
	Address           string
	Amount            int64
	Timestamp         int64
	SpendingBlockHash chainhash.Hash
	SpendingTxHash    chainhash.Hash
	SpendingIndex     int
	SpendingTimestamp int64
}
