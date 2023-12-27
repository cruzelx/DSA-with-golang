package hasher

import "testing"

func BenchmarkNaiveHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveHash("apple")
	}
}

func BenchmarkUnknownHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnknownHash("apple")
	}
}

func BenchmarkDjb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Djb2("apple")
	}
}
func BenchmarkFnv1a(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fnv1a("apple")
	}
}

func BenchmarkMurmurHash3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MurmurHash3("apple")
	}
}
