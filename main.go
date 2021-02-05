package main

import (
	"math/rand"
	"regexp"
	"runtime"
	"sync"
	"time"
)

const (
	numberOfSimulation  = 16
	numberOfInteraction = 100
	dropRate            = 0.1
	charset             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetLen          = len(charset)
	min                 = 10
	max                 = 30
)

var (
	rx   = regexp.MustCompile(`(?i)(.*)v(.*)|(.*)i(.*)|(.*)c(.*)|(.*)t(.*)|(.*)o(.*)|(.*)r(.*)|(.*)y(.*)`)
	once sync.Once
	nCPU = runtime.NumCPU()
)

func main() {
	SimulateLootRNG()
}

func SimulateLootRNG() {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	poolSize := 10
	sem := make(chan struct{}, poolSize)

	rngTests := make([]chan []int, nCPU)
	for i := range rngTests {
		c := make(chan []int)
		//divide per CPU thread
		sem <- struct{}{}
		go simulateRNG(numberOfSimulation/nCPU, c, sem)
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
	interactions := make([]int, n)
	for i := range interactions {
		interactions[i] = interaction()
	}
	return interactions
}

/**
 * Runs several simulations and returns the results
 */
func simulateRNG(n int, c chan []int, sem chan struct{}) {
	defer func() {
		<-sem
	}()

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
		b[i] = charset[rand.Intn(charsetLen)]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func RandomNumber() int {
	return rand.Intn(max-min+1) + min
}
