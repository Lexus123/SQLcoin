package converter

import (
	"encoding/binary"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
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
