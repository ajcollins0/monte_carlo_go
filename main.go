package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type set struct {
	Results  float64
	DrawDown float64
	MaxGain  float64
}

var numTrades int = 500
var numTests int = 1000
var startingCapital int = 10000
var percentCorrect float64 = 0.85
var targetProfit float64 = 0.01
var maxLoss float64 = 0.03

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var capital, drawDown, maxGain float64
	var sets []set
	var wins, losses int
	for j := 0; j < numTests; j++ {
		wins = 0
		losses = 0
		capital = float64(startingCapital)
		drawDown = capital
		maxGain = capital
		for i := 0; i < numTrades; i++ {
			randNum := rand.Int63n(100) + 1
			if randNum <= int64(100*percentCorrect) {
				capital = capital + (capital * targetProfit)
				if capital > maxGain {
					maxGain = capital
				}
				wins++
			} else {
				capital = capital - (capital * maxLoss)
				if capital < drawDown {
					drawDown = capital
				}
				losses++
			}
		}
		sets = append(sets, set{capital, drawDown, maxGain})
	}
	printResults(sets)
}

func printResults(s []set) {
	var bestCase, worstCase, aveCase float64
	var tempArray []float64
	bestCase = s[0].Results
	worstCase = s[0].Results
	aveCase = 0
	for i, _ := range s {
		tempArray = append(tempArray, s[i].Results)
		if s[i].Results > bestCase {
			bestCase = s[i].Results
		}
		if s[i].Results < worstCase {
			worstCase = s[i].Results
		}
		aveCase = aveCase + s[i].Results
	}
	sort.Float64s(tempArray)
	medianCase := tempArray[(len(tempArray)/2)-1]
	aveCase = aveCase / float64(len(s))
	fmt.Printf("simulating %d trades... %d times, with $%d starting capital\n", numTrades, numTests, startingCapital)
	fmt.Printf("median %.2f, average %.2f, best %.2f, worst %.2f \n", aveCase, medianCase, bestCase, worstCase)
}
