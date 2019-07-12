package blockreader

import (
	"sqlcoin/services/custommodels"
	"sqlcoin/services/databaser"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

/*
StartReading is the function that runs the entire script basically.
It loops through a blk-file and creates blocks.
At the end of every block it is sent to the database.
*/
func StartReading(file custommodels.File) {
	for {
		// Reads the magic number and the block size, but discards them
		readBlockInfo(file)

		// Reads the block header from the file
		header := ReadBlockHeader(file)

		// Creates a new block with the blockheader already attached
		block := wire.NewMsgBlock(&header)

		// Total number of transactions is read
		totalTxs := file.ReadTotalTxs()

		// All transactions are read and stored in this variable
		// It takes the total number of txs as an argument for a loop
		block.Transactions = readTxs(file, totalTxs)

		// Makes the database insert once the entire block is read
		databaser.MakeInsert(block)
	}
}

/*
readBlockInfo reads the magic number and block size, then discards them
*/
func readBlockInfo(file custommodels.File) {
	file.ReadAndDiscard(custommodels.Magic)
	file.ReadAndDiscard(custommodels.BlockSize)
}

/*
ReadBlockHeader does what it does and returns the block's header
*/
func ReadBlockHeader(file custommodels.File) wire.BlockHeader {
	return wire.BlockHeader{
		file.ReadBlockHeaderVersion(),
		file.ReadBlockHeaderPrevBlock(),
		file.ReadBlockHeaderMerkle(),
		file.ReadBlockHeaderTimestamp(),
		file.ReadBlockHeaderBits(),
		file.ReadBlockHeaderNonce(),
	}
}

/*
readTxs is looping through all txs of a block.
It requires the file and the total number of txs of a block.
Return an array of txs, as described by btcd.wire.
*/
func readTxs(file custommodels.File, totalTxs uint64) []*wire.MsgTx {
	// Create an empty slice of txs to which we'll append txs
	var txSlice []*wire.MsgTx

	// Loop through the block's txs
	for counter := uint64(0); counter < totalTxs; counter++ {
		originTx, segwit := initNewTx()

		file.ReadAndDiscard(custommodels.TxVersion)

		totalInputs := file.ReadTotalIns()
		if totalInputs == 0 {
			segwit = true
			file.ReadAndDiscard(custommodels.SegwitFlag)
			totalInputs = file.ReadTotalIns()
		}
		originTx.TxIn = readIns(file, totalInputs, originTx)

		totalOutputs := file.ReadTotalOuts()
		originTx.TxOut = readOuts(file, totalOutputs, originTx)

		if segwit {
			readSegwit(file, len(originTx.TxIn))
		}

		originTx.LockTime = file.ReadTxLocktime()

		txSlice = append(txSlice, originTx)
	}

	return txSlice
}

func initNewTx() (*wire.MsgTx, bool) {
	return wire.NewMsgTx(wire.TxVersion), false
}

func readIns(file custommodels.File, totalInputs uint64, originTx *wire.MsgTx) []*wire.TxIn {
	for counter := uint64(0); counter < totalInputs; counter++ {
		prevOut := wire.NewOutPoint(readOutpoint(file))
		txIn := wire.NewTxIn(prevOut, file.ReadInputScriptSig(), nil)
		file.ReadAndDiscard(custommodels.TxSequence)
		originTx.AddTxIn(txIn)
	}

	return originTx.TxIn
}

func readOutpoint(file custommodels.File) (*chainhash.Hash, uint32) {
	outpointPrevTx := file.ReadOutpointPrevTx()
	outpointPrevTxIndex := file.ReadOutpointPrevTxIndex()
	return &outpointPrevTx, outpointPrevTxIndex
}

func readOuts(file custommodels.File, totalOutputs uint64, originTx *wire.MsgTx) []*wire.TxOut {
	for counter := uint64(0); counter < totalOutputs; counter++ {
		amount := file.ReadOutputAmount()
		pub := file.ReadOutputScriptPubKey()
		txOut := wire.NewTxOut(amount, pub)
		originTx.AddTxOut(txOut)
	}

	return originTx.TxOut
}

func readSegwit(file custommodels.File, inputsLen int) {
	for inputCounter := 0; inputCounter < inputsLen; inputCounter++ {
		totalStackItems := file.ReadStackItems()
		for stackCounter := uint64(0); stackCounter < totalStackItems; stackCounter++ {
			file.ReadSegwit()
		}
	}
}
