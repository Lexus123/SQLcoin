package converter

import (
	"encoding/binary"
	"time"

	"sqlcoin/services/errorchecker"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

/*
ConvertToUint32 receives
*/
func ConvertToUint32(buffer []byte) uint32 {
	return binary.LittleEndian.Uint32(buffer)
}

/*
ConvertToInt32 receives
*/
func ConvertToInt32(buffer []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buffer))
}

/*
ConvertToInt64 receives
*/
func ConvertToInt64(buffer []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buffer))
}

/*
ConvertToInt receives
*/
func ConvertToInt(buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer))
}

/*
Convert32Bytes receives
*/
func Convert32Bytes(buffer []byte) chainhash.Hash {
	var raw [32]byte
	copy(raw[:], buffer[:32])
	return chainhash.Hash(raw)
}

/*
ConvertBlockHeaderTimestamp receives
*/
func ConvertBlockHeaderTimestamp(buffer []byte) time.Time {
	return time.Unix(int64(binary.LittleEndian.Uint32(buffer)), 0)
}

/*
ConvertAddress receives
*/
func ConvertAddress(output *wire.TxOut) string {
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, &chaincfg.MainNetParams)
	errorchecker.CheckFileError(err)
	var address string
	if len(addresses) > 0 {
		address = addresses[0].String()
	}
	return address
}
