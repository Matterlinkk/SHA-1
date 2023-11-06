package operations

import "encoding/binary"

func ProcessBlock(block []byte, h0, h1, h2, h3, h4 uint32) (uint32, uint32, uint32, uint32, uint32) {

	words := make([]uint32, 80)
	for i := 0; i < 16; i++ {
		words[i] = binary.BigEndian.Uint32(block[i*4 : (i+1)*4])
	}

	for i := 16; i < 80; i++ {
		words[i] = LeftRotate(words[i-3]^words[i-8]^words[i-14]^words[i-16], 1)
	}

	a, b, c, d, e := h0, h1, h2, h3, h4

	for i := 0; i < 80; i++ {
		var f, k uint32

		if i < 20 {
			f = (b & c) | ((^b) & d)
			k = 0x5A827999
		} else if i < 40 {
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		} else if i < 60 {
			f = (b & c) | (b & d) | (c & d)
			k = 0x8F1BBCDC
		} else {
			f = b ^ c ^ d
			k = 0xCA62C1D6
		}

		temp := LeftRotate(a, 5) + f + e + k + words[i]
		e = d
		d = c
		c = LeftRotate(b, 30)
		b = a
		a = temp
	}

	h0 += a
	h1 += b
	h2 += c
	h3 += d
	h4 += e

	return h0, h1, h2, h3, h4
}

func LeftRotate(value, shift uint32) uint32 {
	return (value << shift) | (value >> (32 - shift))
}
