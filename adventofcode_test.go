package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1PartA(t *testing.T) {
	result := day1PartA()
	fmt.Printf("Day 1A: '%s'\n", result)
	assert.Equal(t, "1031", result, "Day 1 Part A is broken")
}

func TestDay1PartB(t *testing.T) {
	result := day1PartB()
	fmt.Printf("Day 1B: '%s'\n", result)
	assert.Equal(t, "1080", result, "Day 1 Part B is broken")
}

func TestDay2PartA(t *testing.T) {
	result := day2PartA()
	fmt.Printf("Day 2A: '%s'\n", result)
	assert.Equal(t, "54426", result, "Day 2 Part A is broken")
}

func TestDay2PartB(t *testing.T) {
	result := day2PartB()
	fmt.Printf("Day 2B: '%s'\n", result)
	assert.Equal(t, "333", result, "Day 2 Part B is broken")
}

func TestDay3PartA(t *testing.T) {
	result := day3PartA()
	fmt.Printf("Day 3A: '%s'\n", result)
	assert.Equal(t, "438", result, "Day 3 Part A is broken")
}

func TestDay3PartB(t *testing.T) {
	result := day3PartB()
	fmt.Printf("Day 3B: '%s'\n", result)
	assert.Equal(t, "266330", result, "Day 3 Part B is broken")
}

func TestDay4PartA(t *testing.T) {
	result := day4PartA()
	fmt.Printf("Day 4A: '%s'\n", result)
	assert.Equal(t, "325", result, "Day 4 Part A is broken")
}

func TestDay4PartB(t *testing.T) {
	result := day4PartB()
	fmt.Printf("Day 4B: '%s'\n", result)
	assert.Equal(t, "119", result, "Day 4 Part B is broken")
}

func TestDay5PartA(t *testing.T) {
	result := day5PartA()
	fmt.Printf("Day 5A: '%s'\n", result)
	assert.Equal(t, "388611", result, "Day 5 Part A is broken")
}

func TestDay5PartB(t *testing.T) {
	result := day5PartB()
	fmt.Printf("Day 5B: '%s'\n", result)
	assert.Equal(t, "27763113", result, "Day 5 Part B is broken")
}

func TestDay6PartA(t *testing.T) {
	result := day6PartA()
	fmt.Printf("Day 6A: '%s'\n", result)
	assert.Equal(t, "7864", result, "Day 6 Part A is broken")
}

func TestDay6PartB(t *testing.T) {
	result := day6PartB()
	fmt.Printf("Day 6B: '%s'\n", result)
	assert.Equal(t, "1695", result, "Day 6 Part B is broken")
}

func TestDay7PartA(t *testing.T) {
	result := day7PartA()
	fmt.Printf("Day 7A: '%s'\n", result)
	assert.Equal(t, "cyrupz", result, "Day 7 Part A is broken")
}

func TestDay7PartB(t *testing.T) {
	result := day7PartB()
	fmt.Printf("Day 7B: '%s'\n", result)
	assert.Equal(t, "193", result, "Day 7 Part B is broken")
}

func TestDay8PartA(t *testing.T) {
	result := day8PartA()
	fmt.Printf("Day 8A: '%s'\n", result)
	assert.Equal(t, "7787", result, "Day 8 Part A is broken")
}

func TestDay8PartB(t *testing.T) {
	result := day8PartB()
	fmt.Printf("Day 8B: '%s'\n", result)
	assert.Equal(t, "8997", result, "Day 8 Part B is broken")
}
