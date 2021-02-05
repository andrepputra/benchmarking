package main

import (
	"testing"
)

func BenchmarkSimulateLootRNG(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SimulateLootRNG()
	}
}

// interaction
// BenchmarkSimulateLootRNG-4   	       4	 320766202 ns/op	45378734 B/op	  151754 allocs/op

// v2
// BenchmarkSimulateLootRNG-4   	      27	  46825735 ns/op	   92216 B/op	    2911 allocs/op

// interactionV3
// BenchmarkSimulateLootRNG-4   	      26	  41497286 ns/op	   91806 B/op	    2903 allocs/op

// v4
// BenchmarkSimulateLootRNG-4   	      19	  60782554 ns/op	   91848 B/op	    2913 allocs/op
