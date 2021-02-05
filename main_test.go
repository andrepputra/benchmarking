package main

import "testing"

func BenchmarkInteraction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		interaction()
	}
}

// before
// goos: linux
// goarch: amd64
// pkg: github.com/andrepputra/benchmarking
// BenchmarkInteraction-4   	   45376	     36800 ns/op	   18450 B/op	      94 allocs/op
// PASS
// ok  	github.com/andrepputra/benchmarking	1.959s

// after
// goos: linux
// goarch: amd64
// pkg: github.com/andrepputra/benchmarking
// BenchmarkInteraction-4   	 8534157	       189 ns/op	       7 B/op	       0 allocs/op
// PASS
// ok  	github.com/andrepputra/benchmarking	1.761s
