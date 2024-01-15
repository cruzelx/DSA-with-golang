package main

func MurmurHash3(key string) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)

	var (
		length  = len(key)
		h1      = uint32(length)
		nblocks = length / 4
	)

	for i := 0; i < nblocks; i++ {
		k1 := uint32(key[i*4+0]) | uint32(key[i*4+1])<<8 | uint32(key[i*4+2])<<16 | uint32(key[i*4+3])<<24
		k1 *= c1
		k1 = (k1 << r1) | (k1 >> (32 - r1))
		k1 *= c2
		h1 ^= k1
		h1 = (h1 << r2) | (h1 >> (32 - r2))
		h1 = h1*m + n
	}

	tail := length & 3
	switch tail {
	case 3:
		h1 ^= uint32(key[nblocks*4+2]) << 16
		fallthrough
	case 2:
		h1 ^= uint32(key[nblocks*4+1]) << 8
		fallthrough
	case 1:
		h1 ^= uint32(key[nblocks*4])
		h1 *= c1
		h1 = (h1 << r1) | (h1 >> (32 - r1))
		h1 *= c2
	}

	h1 ^= uint32(length)
	h1 = fmix(h1)

	return h1
}

func fmix(h1 uint32) uint32 {
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
	return h1
}
