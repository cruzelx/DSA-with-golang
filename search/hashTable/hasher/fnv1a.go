package hasher

func Fnv1a(key string) uint32 {
	hash := uint32(2166136261)
	for _, c := range key {
		hash ^= uint32(c)
		hash *= uint32(16777619)
	}
	return hash

}
