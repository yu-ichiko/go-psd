package util

import (
	"encoding/binary"
	"unicode/utf16"
	"bytes"
)

func ReadString(buf []byte, offset int, limit int) string {
	return string(buf[offset:limit])
}

func ReadUint8(buf []byte, offset int) uint8 {
	return uint8(buf[offset])
}

func ReadUint16(buf []byte, offset int) uint16 {
	return binary.BigEndian.Uint16(buf[offset : offset+2])
}

func ReadInt16(buf []byte, offset int) int16 {
	return int16(ReadUint16(buf, offset))
}

func ReadUint32(buf []byte, offset int) uint32 {
	return binary.BigEndian.Uint32(buf[offset : offset+4])
}

func ReadInt32(buf []byte, offset int) int32 {
	return int32(ReadUint32(buf, offset))
}

func ReadUint64(buf []byte, offset int) uint64 {
	return binary.BigEndian.Uint64(buf[offset : offset+8])
}

func ByteString(str string) []byte {
	return []byte(str)
}

func ByteUint16(n int) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(n))
	return b
}

func ByteUint32(n int) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(n))
	return b
}

func BytePascalString(str string) []byte {
	n := len(str)
	if n == 0 {
		return []byte{0}
	}

	buf := &bytes.Buffer{}
	buf.Write([]byte{uint8(n)})
	buf.Write([]byte(str))
	return buf.Bytes()
}

func PascalString(buf []byte, offset int) (string, int) {
	size := int(buf[offset])
	if size == 0 {
		return "", 1
	}
	offset += 1
	return string(buf[offset : offset+size]), size
}

func AdjustAlign2(offset int) int {
	if offset&1 != 0 {
		return 1
	}
	return 0
}

func UnicodeString(buf []byte) string {
	size := ReadUint32(buf, 0)
	if size == 0 {
		return ""
	}
	data := make([]uint16, size)
	for i := range data {
		data[i] = ReadUint16(buf, 4+i<<1)
	}
	return string(utf16.Decode(data))
}

func GetSize(isPSB bool) int {
	if isPSB {
		return 8
	}
	return 4
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
