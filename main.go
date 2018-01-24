package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var numTrades int = 500
var numTests int = 10
var startingCapital int = 10000
var percentCorrect float64 = 0.85
var targetProfit float64 = 0.01
var maxLoss float64 = 0.03

func perform_single_test() float64 {
	capital := float64(startingCapital)
	for i := 0; i < numTrades; i++ {
		randNum := rand.Int63n(100) + 1
		if randNum <= int64(100*percentCorrect) {
			capital = capital + (capital * targetProfit)
		} else {
			capital = capital - (capital * maxLoss)
		}
	}
	return capital
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var sets []float64
	for j := 0; j < numTests; j++ {
		sets = append(sets, perform_single_test())
	}
	printResults(sets)
}

func printResults(s []float64) {
	var bestCase, worstCase, aveCase float64
	var tempArray []float64
	bestCase = s[0]
	worstCase = s[0]
	aveCase = 0
	for i, _ := range s {
		tempArray = append(tempArray, s[i])
		if s[i] > bestCase {
			bestCase = s[i]
		}
		if s[i] < worstCase {
			worstCase = s[i]
		}
		aveCase = aveCase + s[i]
	}
	sort.Float64s(tempArray)
	medianCase := tempArray[(len(tempArray)/2)-1]
	aveCase = aveCase / float64(len(s))
	fmt.Printf("simulating %d trades... %d times, with $%d starting capital\n", numTrades, numTests, startingCapital)
	fmt.Printf("median %.2f, average %.2f, best %.2f, worst %.2f \n", aveCase, medianCase, bestCase, worstCase)
}
