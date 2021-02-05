package main

import "testing"

func BenchmarkGacha(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SimulateLootRNG()
	}
}
