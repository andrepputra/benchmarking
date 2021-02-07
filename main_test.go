package main

import "testing"

func BenchmarkSimulateLootRNG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SimulateLootRNG()
	}
}

//before
// goos: linux
// goarch: amd64
// pkg: github.com/santosomarvin/benchmarking
// BenchmarkSimulateLootRNG-8           500          10121086 ns/op        30191035 B/op     150534 allocs/op
// PASS
// ok      github.com/santosomarvin/benchmarking   5.078s

// after
// goos: linux
// goarch: amd64
// pkg: github.com/santosomarvin/benchmarking
// BenchmarkSimulateLootRNG-8           500           3164997 ns/op          104259 B/op       3234 allocs/op
// PASS
// ok      github.com/santosomarvin/benchmarking   1.588s
