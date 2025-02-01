package dbgo

import "encoding/binary"

func uint64toBytes(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)
	return b
}

func uint64FromBytes(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}
