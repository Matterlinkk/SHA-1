package sha1func

import (
	"Project/operations"
	"encoding/binary"
)

func Sha1hashing(message string) [20]byte {
	h0 := uint32(0x67452301)
	h1 := uint32(0xEFCDAB89)
	h2 := uint32(0x98BADCFE)
	h3 := uint32(0x10325476)
	h4 := uint32(0xC3D2E1F0)

	data := []byte(message)

	messageLength := uint64(len(data) * 8)

	data = append(data, 0x80)

	for (len(data)+8)%64 != 0 {
		data = append(data, 0x00)
	}

	messageLengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLengthBytes, messageLength)
	data = append(data, messageLengthBytes...)

	for i := 0; i < len(data); i += 64 {
		block := data[i : i+64]
		h0, h1, h2, h3, h4 = operations.ProcessBlock(block, h0, h1, h2, h3, h4)
	}

	hash := [20]byte{}
	binary.BigEndian.PutUint32(hash[0:], h0)
	binary.BigEndian.PutUint32(hash[4:], h1)
	binary.BigEndian.PutUint32(hash[8:], h2)
	binary.BigEndian.PutUint32(hash[12:], h3)
	binary.BigEndian.PutUint32(hash[16:], h4)

	return hash
}
