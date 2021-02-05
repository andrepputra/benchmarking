package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"regexp"
	"runtime"
)

const (
	numberOfSimulation  = 16
	numberOfInteraction = 100
	dropRate            = 0.1
	charset             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	SimulateLootRNG()
}

func SimulateLootRNG() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	nCPU := runtime.NumCPU()
	rngTests := make([]chan []int, 0, nCPU)
	for i := 0; i < nCPU; i++ {
		c := make(chan []int)
		//divide per CPU thread
		go simulateRNG(numberOfSimulation/nCPU, c)
		rngTests = append(rngTests, c)
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
	rx := regexp.MustCompile(`(?i)(.*)v(.*)|(.*)i(.*)|(.*)c(.*)|(.*)t(.*)|(.*)o(.*)|(.*)r(.*)|(.*)y(.*)`)

	monsterName := String(RandomNumber())
	nameContainsVictory := rx.MatchString(monsterName)
	isItemDrop := rand.Float64() <= dropRate

	if !nameContainsVictory {
		return 0
	}

	if isItemDrop {
		return 1
	}

	return 0
}

/**
 * Runs several interactions and retuns a slice representing the results
 */
func simulation(n int) []int {
	interactions := make([]int, 0, n)
	for i := 0; i < n; i++ {
		interactions = append(interactions, interaction())
	}
	return interactions
}

/**
 * Runs several simulations and returns the results
 */
func simulateRNG(n int, c chan []int) {
	simulations := make([]int, n)
	for i := 0; i < n; i++ {
		for _, v := range simulation(numberOfInteraction) {
			simulations = append(simulations, simulations[i]+v)
		}
	}
	c <- simulations
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		b = append(b, charset[rand.Intn(len(charset))])
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
