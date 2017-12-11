package main

import (
	"fmt"
	"time"
)

type fn func() string

func main() {
	fmt.Printf("Advent of Code 2017\n\n")
	startTotal := time.Now()

	dayRunner("day1A", day1PartA)
	dayRunner("day1B", day1PartB)

	dayRunner("day2A", day2PartA)
	dayRunner("day2B", day2PartB)

	dayRunner("day3A", day3PartA)
	dayRunner("day3B", day3PartB)

	dayRunner("day4A", day4PartA)
	dayRunner("day4B", day4PartB)

	dayRunner("day5A", day5PartA)
	dayRunner("day5B", day5PartB)

	dayRunner("day6A", day6PartA)
	dayRunner("day6B", day6PartB)

	dayRunner("day7A", day7PartA)
	dayRunner("day7B", day7PartB)

	dayRunner("day8A", day8PartA)
	dayRunner("day8B", day8PartB)

	dayRunner("day9A", day9PartA)
	dayRunner("day9B", day9PartB)

	fmt.Printf("\nDone in %s\n", time.Since(startTotal))
}

func dayRunner(day string, f fn) {
	startIndividual := time.Now()
	result := f()
	fmt.Printf("%s: '%v' in %v\n", day, result, time.Since(startIndividual))
}
