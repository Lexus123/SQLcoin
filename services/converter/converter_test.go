package converter_test

import (
	. "sqlcoin/services/converter"
	"testing"
)

func TestConvertToUint32(t *testing.T) {
	var resultConvertToUint32, expectedConvertToUint32 uint32

	type checkConvertToUint32 struct {
		input    []byte
		expected uint32
	}

	var table = []checkConvertToUint32{}

	table = append(table, checkConvertToUint32{[]byte{0xA1, 0xD3, 0x08, 0x00}, 578465})
	table = append(table, checkConvertToUint32{[]byte{0x0B, 0xD4, 0x08, 0x00}, 578571})
	table = append(table, checkConvertToUint32{[]byte{0x4E, 0x00, 0x00, 0x00}, 78})

	for _, j := range table {
		resultConvertToUint32 = ConvertToUint32(j.input)
		expectedConvertToUint32 = j.expected
		if resultConvertToUint32 != expectedConvertToUint32 {
			t.Errorf("%v not equal to %v", resultConvertToUint32, expectedConvertToUint32)
		}
	}
}

func Benchmark_ConvertToUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ConvertToUint32([]byte{0xA1, 0xD3, 0x08, 0x00})
	}
}

func Benchmark_Convert32Bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Convert32Bytes([]byte{
			0xBA, 0xA2, 0x97, 0x97, 0x16, 0x33,
			0x18, 0xE0, 0x9C, 0x70, 0xEA, 0x01,
			0x2B, 0x9F, 0xD5, 0x9C, 0x1C, 0xD8,
			0x57, 0xE9, 0x59, 0xE6, 0x8F, 0xD9,
			0xE0, 0xF6, 0xBA, 0x1F, 0xD8, 0xD6,
			0x94, 0x17,
		})
	}
}