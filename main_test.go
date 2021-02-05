package main

import "testing"

func BenchmarkSimulateLootRNG(bench *testing.B) {
	for n := 0; n < bench.N; n++ {
		SimulateLootRNG()
	}
}
