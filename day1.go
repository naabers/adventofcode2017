package adventofcode2017

import (
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func day1PartA() string {
	input, err := ioutil.ReadFile("inputs/day1.txt")
	check(err)

	var prev byte
	result := 0
	for _, nextByte := range input {
		if prev == nextByte {
			result += getIntFromByte(nextByte)
		}
		prev = nextByte
	}

	if prev == input[0] {
		result += getIntFromByte(prev)
	}
	return strconv.Itoa(result)
}

func day1PartB() string {
	input, err := ioutil.ReadFile("inputs/day1.txt")
	check(err)

	result := 0
	middleIndex := len(input) / 2
	for index := 0; index < len(input)/2; index, middleIndex = index+1, middleIndex+1 {
		if input[index] == input[middleIndex] {
			result += getIntFromByte(input[index]) * 2
		}
	}
	return strconv.Itoa(result)
}

func getIntFromByte(b byte) int {
	result, e := strconv.Atoi(string(b))
	if e != nil {
		panic(e)
	}
	return result
}
