package main

import "fmt"
import "flag"

func main() {
	var A, B, RANGE int
	flag.IntVar(&A, "A", 7, "A (default: 7)")
	flag.IntVar(&B, "B", 12, "B (default: 12)")
	flag.IntVar(&RANGE, "Range", 100, "Range to search (default: 100)")
	flag.Parse()
	Calculate(A, B, RANGE)
}

func getRoot(possibilities, used []bool, RANGEIZ int) int {
	// We want to stop if there are no more roots to go off from
	// Otherwise, we want to continue with the existing root
	for i := 0; i < RANGEIZ; i++ {
		if possibilities[i] {
			if !used[i] {
				return i
			}
		}
	}
	return -1
}

func Calculate(A, B, RANGE int) {
	RANGEIZ := RANGE + 1
	// This array contains 101 boolean values, each 
	// representing if a possible number is unachievable or not.
	// At first, we don't know, so we assume all numbers are unachievable, except for 0
	var possibilities = make([]bool, RANGEIZ, RANGEIZ)
	for i := 0; i < RANGEIZ; i++ {
		possibilities[i] = false
	}
	possibilities[0] = true

	// This represents if we have used a particular number as a root
	// If we have, we don't need to use it again
	var used = make([]bool, RANGEIZ, RANGEIZ)
	for i := 0; i < RANGEIZ; i++ {
		used[i] = false
	}

	// We calculate the delta values
	var delta = [8]int{
//		A * 0 +  B * 0, // a delta of zero is meaningless
		A * 0 +  B * 1,
		A * 1 +  B * 0,
		A * 1 +  B * 1,
	}

	root := getRoot(possibilities, used, RANGEIZ)
	for root != -1 {
		used[root] = true
		for i := 0; i < len(delta); i++ {
			index := root + delta[i]
			if index <= RANGE && index >= 0 {
				possibilities[index] = true
			}
		}
		root = getRoot(possibilities, used, RANGEIZ)
	}

	finalResults := []int{}
	for i := RANGE; i >= 0; i-- {
		if !possibilities[i] {
			finalResults = append(finalResults, i)
		}
	}

	fmt.Printf("\nA=%v, B=%v, Range=%v\n\n", A, B, RANGE)
	fmt.Printf("Unreachable scores\n%v\n\n", finalResults)
}
