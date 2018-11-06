package utils

import (
	"encoding/binary"
	"math"
)

//func _UInt32toBytes(v uint32) []byte {
//	b := make([]byte, 4)
//	b[0] = byte(0xff & v)
//	b[1] = byte(0xff00 & v >> 8)
//	b[2] = byte(0xff0000 & v >> 16)
//	b[3] = byte(0xff000000  & v >> 24)
//	return b
//}

func UInt16ToBytes(value uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, value)
	return bytes
}

func UInt16FromBytes(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func UInt32ToBytes(value uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, value)
	return bytes
}

func UInt32FromBytes(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func UInt64ToBytes(value uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, value)
	return bytes
}

func UInt64FromBytes(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Float32ToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Float32FromBytes(b []byte) float32 {
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits)
}

func Float64ToBytes(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func Float64FromBytes(d []byte) float64 {
	bits := binary.LittleEndian.Uint64(d)
	float := math.Float64frombits(bits)
	return float
}
