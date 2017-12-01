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
	for _, inputNumber := range input[:len(input)] {
		if prev == inputNumber {
			currentNumber, _ := strconv.Atoi(string(inputNumber))
			result += currentNumber
		}
		prev = inputNumber
	}

	if prev == input[0] {
		currentNumber, _ := strconv.Atoi(string(prev))
		result += currentNumber
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
			currentNumber, _ := strconv.Atoi(string(input[index]))
			result += currentNumber * 2
		}
	}
	return strconv.Itoa(result)
}
