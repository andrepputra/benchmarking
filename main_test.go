package main

import "testing"

func BenchmarkSimulateLootRNG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimulateLootRNG()
	}
}
