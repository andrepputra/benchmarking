gobench:
	@go test -bench=. -benchtime=109x

gobench_mem:
	@go test -bench=. -benchtime=109x -benchmem 

gobench_cpuout:
	@go test -bench=. -benchtime=109x -cpuprofile=cpu.out
	@go tool pprof --pdf benchmarking.test cpu.out > cpu0.pdf