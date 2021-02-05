package main

import "testing"

func BenchmarkSimulateLootRNG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SimulateLootRNG()
	}
}