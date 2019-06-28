package custommodels_test

import (
	. "sqlcoin/services/blockreader"
	"sqlcoin/services/fileopener"
	"testing"
)

func TestReadBlockHeaderVersion(t *testing.T) {
	expected := int32(1)

	file := fileopener.OpenFile("/Users/lex/Documents/Development/Go/privateprojects/sqlcoin/testfiles/custommodels/readblockheaderversion")

	result := file.ReadBlockHeaderVersion()
	if result != expected {
		t.Errorf("%v not equal to %v", result, expected)
	}
}

func Benchmark_ReadBlockHeader(b *testing.B) {
	file := fileopener.OpenFile("/Users/lex/Documents/Development/Go/privateprojects/sqlcoin/testfiles/custommodels/readblockheader")
	for i := 0; i < b.N; i++ {
		_ = ReadBlockHeader(file)
		file.File.Seek(0, 0)
	}
}
