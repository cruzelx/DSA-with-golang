package hasher

// Cumulative sum of ASCII chars of the key
func NaiveHash(key string) uint32 {
	hash := uint32(0)
	for _, c := range key {
		char_code := uint32(c)
		hash += char_code
	}
	return hash
}
