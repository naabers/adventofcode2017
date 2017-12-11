package main

import (
	"io/ioutil"
	"strconv"
)

func day9Processor(input []byte) (int, int) {
	score := 0
	currentGroupScore := 0
	skip := false
	inGarbage := false
	garbageCount := 0
	for _, nextByte := range input {
		if skip {
			skip = false
			continue
		}
		switch string(nextByte) {
		case "{":
			if !inGarbage {
				currentGroupScore++
			} else {
				garbageCount++
			}
		case "}":
			if !inGarbage {
				score += currentGroupScore
				currentGroupScore--
			} else {
				garbageCount++
			}
		case "<":
			if inGarbage {
				garbageCount++
			}
			inGarbage = true
		case ">":
			inGarbage = false
		case "!":
			skip = true
		default:
			if inGarbage {
				garbageCount++
			}
		}
	}
	return score, garbageCount
}

func day9PartA() string {
	input, err := ioutil.ReadFile("inputs/day9.txt")
	check(err)

	score, _ := day9Processor(input)

	return strconv.Itoa(score)
}

func day9PartB() string {
	input, err := ioutil.ReadFile("inputs/day9.txt")
	check(err)

	_, skipped := day9Processor(input)

	return strconv.Itoa(skipped)
}
