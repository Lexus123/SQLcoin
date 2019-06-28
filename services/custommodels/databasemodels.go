package custommodels

type DbBlock struct {
	Hash      string
	PrevHash  string
	Merkle    string
	Timestamp int64
}

type DbTx struct {
	Hash        string
	BlockHeight int64
}

type DbOutput struct {
	TxHash  string
	TxIndex int
	Amount  int64
	Address string
	Used    uint64
}

type DbInput struct {
	TxHash      string
	PrevTxHash  string
	PrevTxIndex uint32
}
