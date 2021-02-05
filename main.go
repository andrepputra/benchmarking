package main

import (
	"math/rand"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	numberOfSimulation  = 16
	numberOfInteraction = 100
	dropRate            = 0.1
	charset             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	SimulateLootRNG()
}

func SimulateLootRNG() {

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
	interactions := make([]int, n)
	for i := range interactions {
		interactions[i] = interactionV4()
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

func interactionV2() int {
	isItemDrop := rand.Float64() <= dropRate
	if isItemDrop {
		return 1
	}

	monsterName := String(RandomNumber())
	for _, r := range monsterName {
		if validateChar(r) {
			return 1
		}
	}

	return 0
}

func validateChar(c rune) bool {
	return c == 'v' || c == 'i' || c == 'c' || c == 't' || c == 'o' || c == 'r' || c == 'y' || c == 'V' || c == 'I' || c == 'C' || c == 'T' || c == 'O' || c == 'R' || c == 'Y'
}

var listRune = []rune{'v', 'i', 'c', 't', 'o', 'r', 'y', 'V', 'I', 'C', 'T', 'O', 'R', 'Y'}

func interactionV3() int {
	isItemDrop := rand.Float64() <= dropRate
	if isItemDrop {
		return 1
	}

	monsterName := String(RandomNumber())
	for _, r := range listRune {
		if strings.ContainsRune(monsterName, r) {
			return 1
		}
	}

	return 0
}

var mapRune = map[rune]bool{
	'v': true, 'i': true, 'c': true, 't': true, 'o': true, 'r': true, 'y': true, 'V': true, 'I': true, 'C': true, 'T': true, 'O': true, 'R': true, 'Y': true,
}

func interactionV4() int {
	isItemDrop := rand.Float64() <= dropRate
	if isItemDrop {
		return 1
	}

	monsterName := String(RandomNumber())
	for _, r := range monsterName {
		if mapRune[r] {
			return 1
		}
	}

	return 0
}
