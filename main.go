package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Advent of Code 2017\n\n")
	startTotal := time.Now()
	result := day1PartA()
	fmt.Printf("day1A: %v in %v\n", result, time.Since(startTotal))

	startIndividual := time.Now()
	result = day1PartB()
	fmt.Printf("day1B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day2PartA()
	fmt.Printf("day2A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day2PartB()
	fmt.Printf("day2B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day3PartA()
	fmt.Printf("day3A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day3PartB()
	fmt.Printf("day3B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day4PartA()
	fmt.Printf("day4A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day4PartB()
	fmt.Printf("day4B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day5PartA()
	fmt.Printf("day5A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day5PartB()
	fmt.Printf("day5B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day6PartA()
	fmt.Printf("day6A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day6PartB()
	fmt.Printf("day6B: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day7PartA()
	fmt.Printf("day7A: %v in %v\n", result, time.Since(startIndividual))

	startIndividual = time.Now()
	result = day7PartB()
	fmt.Printf("day7B: %v in %v\n", result, time.Since(startIndividual))

	fmt.Printf("\nDone in %s\n", time.Since(startTotal))
}
