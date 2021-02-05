go test -bench=. -benchtime=1000x

before
goos: linux
goarch: amd64
pkg: github.com/tokopedia/benchmarking-ori
BenchmarkSimulateLootRNG-8   	    1000	  11951104 ns/op
PASS
ok  	github.com/tokopedia/benchmarking-ori	11.967s



after
goos: linux
goarch: amd64
pkg: github.com/tokopedia/benchmarking
BenchmarkSimulateLootRNG-8   	    1000	   3809833 ns/op
PASS
ok  	github.com/tokopedia/benchmarking	3.823s



go test -bench=. -benchtime=1000x -benchmem

before
goos: linux
goarch: amd64
pkg: github.com/tokopedia/benchmarking-ori
BenchmarkSimulateLootRNG-8   	    1000	  11730376 ns/op	30176610 B/op	  150532 allocs/op
PASS
ok  	github.com/tokopedia/benchmarking-ori	11.747s

after
goos: linux
goarch: amd64
pkg: github.com/tokopedia/benchmarking
BenchmarkSimulateLootRNG-8   	    1000	   3413953 ns/op	  103871 B/op	    3235 allocs/op
PASS
ok  	github.com/tokopedia/benchmarking	3.423s
