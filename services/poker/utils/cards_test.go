package utils

import (
	"testing"
)

func BenchmarkSelectMathRandom(b *testing.B) {
	s := NewCardSelector()
	for i := 0; i < b.N; i++ {
		s.Reset()
		for i := 0; i < 10; i++ {
			_, _ = SelectRandomN(0)
		}
	}
}
