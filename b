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

