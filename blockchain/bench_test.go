package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkHashWithNonce(b *testing.B) {
	nonceCounts := []int{1, 2, 3, 4, 5, 6, 7}

	b.N = 1

	for _, nc := range nonceCounts {
		b.Run(fmt.Sprintf("nonceCount=%d", nc), func(b *testing.B) {
			bb := Block{
				nonce: 0,
				index: 0,
				data: "test data",
				hash2: "",
				time: time.Now(),
			}

			bb.hash1 = Hash(bb)

			b.ResetTimer()

			HashWithNonce(bb, nc)
		})
	}
}