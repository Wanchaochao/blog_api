package stringutil

import (
	"testing"
)

func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(8)
	}
}

func TestGetRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := RandomString(8)
		if len(s) != 8 {
			t.Error("string length error:" + s)
		}
	}
}