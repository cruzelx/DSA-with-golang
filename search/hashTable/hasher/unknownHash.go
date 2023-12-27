package hasher

// Because I forgot where I got this algorithm from 😅
func UnknownHash(key string) uint32 {
	hash := uint32(0)
	mul := uint32(1)
	for i, c := range key {
		char_code := uint32(c)
		if i%4 == 0 {
			mul = 1
		} else {
			mul *= 256
		}
		hash += char_code * mul
	}
	return hash
}
