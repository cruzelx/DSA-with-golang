package hasher

func Djb2(key string) uint32 {
	// Initial prime value
	hash := uint32(5381)
	for _, c := range key {
		char_code := uint32(c)

		// (hash<<5) means hash*(2^5)
		hash = ((hash << 5) + hash) + char_code
	}
	return hash
}
