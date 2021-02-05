package main

import (
	"math/rand"
	"regexp"
	"runtime"
	"time"
)

const (
	numberOfSimulation  = 16
	numberOfInteraction = 100
	dropRate            = 0.1
	charset             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	rx = regexp.MustCompile(`(?i)(.*)v(.*)|(.*)i(.*)|(.*)c(.*)|(.*)t(.*)|(.*)o(.*)|(.*)r(.*)|(.*)y(.*)`)
)

func main() {
	SimulateLootRNG()
}

func SimulateLootRNG() {
	rand.Seed(time.Now().UnixNano())

	nCPU := runtime.NumCPU()
	rngTests := make([]chan []int, nCPU)
	for i := range rngTests {
		c := make(chan []int)
		//divide per CPU thread
		go simulateRNG(numberOfSimulation/nCPU, c)
		rngTests[i] = c
	}

	// Concatentate the test results
	results := make([]int, numberOfSimulation)
	for i, c := range rngTests {
		start := (numberOfSimulation / nCPU) * i
		stop := (numberOfSimulation / nCPU) * (i + 1)
		copy(results[start:stop], <-c)
	}

	// fmt.Println("RNG Loot Results: ", results)
}

/*
	Simulates a single interaction with a monster
	Returns 1 if the monster dropped an item and 0 otherwise
	But if monster name doesn't contain any of character from `victory`, it will be treated as 0
*/
func interaction() int {
	if !rx.MatchString(String(RandomNumber())) {
		return 0
	}

	if rand.Float64() <= dropRate {
		return 1
	}

	return 0
}

/**
 * Runs several interactions and retuns a slice representing the results
 */
func simulation(n int) []int {
	interactions := make([]int, n)
	for i := range interactions {
		interactions[i] = interaction()
	}
	return interactions
}

/**
 * Runs several simulations and returns the results
 */
func simulateRNG(n int, c chan []int) {
	simulations := make([]int, n)
	for i := range simulations {
		for _, v := range simulation(numberOfInteraction) {
			simulations[i] += v
		}
	}
	c <- simulations
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func RandomNumber() int {
	min := 10
	max := 30
	return rand.Intn(max-min+1) + min
}
