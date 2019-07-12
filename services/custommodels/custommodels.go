package custommodels

import (
	// "fmt"
	"os"
	"sqlcoin/services/converter"
	// "sqlcoin/services/errorchecker"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

const (
	Magic                = 4
	BlockSize            = 4
	BlockHeaderVersion   = 4
	BlockHeaderPrevBlock = 32
	BlockHeaderMerkle    = 32
	BlockHeaderTimestamp = 4
	BlockHeaderBits      = 4
	BlockHeaderNonce     = 4
	TxVersion            = 4
	SegwitFlag           = 1
	OutpointPrevTx       = 32
	OutpointPrevTxIndex  = 4
	TxSequence           = 4
	OutputAmount         = 8
	TxLocktime           = 4
)

/*
BlockFile acts as an interface to provide custom
methods to File struct
*/
type BlockFile interface {
	readStaticSize(size int)
	ReadAndDiscard(bufferSize int)
	ReadBlockHeaderVersion()
	ReadBlockHeaderPrevBlock()
	ReadBlockHeaderMerkle()
	ReadBlockHeaderTimestamp()
	ReadBlockHeaderBits()
	ReadBlockHeaderNonce()
	ReadTotalTxs()
	ReadTotalIns()
	ReadOutpointPrevTx()
	ReadOutpointPrevTxIndex()
	ReadInputScriptSig()
	ReadTotalOuts()
	ReadOutputAmount()
	ReadOutputScriptPubKey()
	ReadTxLocktime()
	ReadStackItems()
	ReadSegwit()
}

// File is a structure with custom methods added to it
type File struct {
	*os.File
}

/*
ReadStaticSize reads a predefined amount of bytes
and then returns it
*/
func (blk *File) readStaticSize(bufferSize int) ([]byte, error) {
	buffer := make([]byte, bufferSize)
	_, err := blk.Read(buffer)
	return buffer, err
}

/*
ReadAndDiscard reads some amount of bytes
and then throws it away. This method
is only used to skip some byte of the file
*/
func (blk *File) ReadAndDiscard(bufferSize int) error {
	_, err := blk.readStaticSize(bufferSize)
	return err
}

/*
ReadBlockHeaderVersion reads the version number of the
block. The byte array is directly converted to a int32
and returned.
*/
func (blk *File) ReadBlockHeaderVersion() int32 {
	buffer, _ := blk.readStaticSize(BlockHeaderVersion)
	return converter.ConvertToInt32(buffer)
}

/*
ReadBlockHeaderPrevBlock reads the previous block hash
from the blockheader. The byte array is directly converted
to a [32]byte as defined in chainhash lib by btcd and returned.
*/
func (blk *File) ReadBlockHeaderPrevBlock() chainhash.Hash {
	buffer, _ := blk.readStaticSize(BlockHeaderPrevBlock)
	return converter.Convert32Bytes(buffer)
}

/*
ReadBlockHeaderMerkle reads the merkle hash
from the blockheader. The byte array is directly converted
to a [32]byte as defined in chainhash lib by btcd and returned.
*/
func (blk *File) ReadBlockHeaderMerkle() chainhash.Hash {
	buffer, _ := blk.readStaticSize(BlockHeaderMerkle)
	return converter.Convert32Bytes(buffer)
}

/*
ReadBlockHeaderTimestamp reads the time from the
blockheader and converts it to a time.Time struct.
*/
func (blk *File) ReadBlockHeaderTimestamp() time.Time {
	buffer, _ := blk.readStaticSize(BlockHeaderTimestamp)
	return converter.ConvertBlockHeaderTimestamp(buffer)
}

/*
ReadBlockHeaderBits reads the bits field from the
blockheader and converts it to a uint32 before returning.
*/
func (blk *File) ReadBlockHeaderBits() uint32 {
	buffer, _ := blk.readStaticSize(BlockHeaderBits)
	return converter.ConvertToUint32(buffer)
}

/*
ReadBlockHeaderNonce reads
*/
func (blk *File) ReadBlockHeaderNonce() uint32 {
	buffer, _ := blk.readStaticSize(BlockHeaderNonce)
	// fmt.Printf("%v\n", buffer)
	return converter.ConvertToUint32(buffer)
}

/*
ReadTotalTxs reads
*/
func (blk *File) ReadTotalTxs() uint64 {
	totalTxs, _ := wire.ReadVarInt(blk.File, 1)
	return totalTxs
}

/*
ReadTotalIns reads
*/
func (blk *File) ReadTotalIns() uint64 {
	totalIns, _ := wire.ReadVarInt(blk.File, 1)
	return totalIns
}

/*
ReadOutpointPrevTx reads
*/
func (blk *File) ReadOutpointPrevTx() chainhash.Hash {
	buffer, _ := blk.readStaticSize(OutpointPrevTx)
	return converter.Convert32Bytes(buffer)
}

/*
ReadOutpointPrevTxIndex reads
*/
func (blk *File) ReadOutpointPrevTxIndex() uint32 {
	buffer, _ := blk.readStaticSize(OutpointPrevTxIndex)
	return converter.ConvertToUint32(buffer)
}

/*
ReadInputScriptSig reads
*/
func (blk *File) ReadInputScriptSig() []byte {
	scriptSigLen, _ := wire.ReadVarInt(blk.File, 1)
	buffer, _ := blk.readStaticSize(int(scriptSigLen))
	return buffer
}

/*
ReadTotalOuts reads
*/
func (blk *File) ReadTotalOuts() uint64 {
	totalOuts, _ := wire.ReadVarInt(blk.File, 1)
	return totalOuts
}

/*
ReadOutputAmount reads
*/
func (blk *File) ReadOutputAmount() int64 {
	buffer, _ := blk.readStaticSize(OutputAmount)
	return converter.ConvertToInt64(buffer)
}

/*
ReadOutputScriptPubKey reads
*/
func (blk *File) ReadOutputScriptPubKey() []byte {
	scriptPubKeyLen, _ := wire.ReadVarInt(blk.File, 1)
	buffer, _ := blk.readStaticSize(int(scriptPubKeyLen))
	return buffer
}

/*
ReadTxLocktime reads
*/
func (blk *File) ReadTxLocktime() uint32 {
	buffer, _ := blk.readStaticSize(TxLocktime)
	return converter.ConvertToUint32(buffer)
}

/*
ReadStackItems reads
*/
func (blk *File) ReadStackItems() uint64 {
	totalStackItems, _ := wire.ReadVarInt(blk.File, 1)
	return totalStackItems
}

/*
ReadSegwit reads
*/
func (blk *File) ReadSegwit() {
	segwitLen, _ := wire.ReadVarInt(blk.File, 1)
	_, _ = blk.readStaticSize(int(segwitLen))
}
