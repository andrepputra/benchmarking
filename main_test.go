package main

import "testing"

func BenchmarkSimulateLootRNG(b *testing.B) {
	for n := 10; n <= b.N; n++ {
		SimulateLootRNG()
	}
}
