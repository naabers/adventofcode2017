package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLargeSmallDiff(line string) int {

	splitString := strings.Split(line, "\t")

	var largest, smallest int

	for index, numberString := range splitString {
		currentNumber, _ := strconv.Atoi(numberString)
		if index == 0 {
			largest = currentNumber
			smallest = currentNumber
		} else if currentNumber > largest {
			largest = currentNumber
		} else if currentNumber < smallest {
			smallest = currentNumber
		}
	}
	return largest - smallest
}

func day2PartA() string {
	file, err := os.Open("inputs/day2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		total += getLargeSmallDiff(currentLine)
	}

	return strconv.Itoa(total)
}

func getDivisibleAmount(line string) int {

	splitString := strings.Split(line, "\t")

	for _, currentString := range splitString {
		currentNumber, _ := strconv.Atoi(currentString)
		for _, innerString := range splitString {
			innerNumber, _ := strconv.Atoi(innerString)

			if innerNumber == currentNumber {
				continue
			}
			fmt.Printf("'%d' %d'\n", currentNumber, innerNumber)
			mod := currentNumber % innerNumber
			if mod == 0 {
				return currentNumber / innerNumber
			}
		}
	}
	return 0
}

func day2PartB() string {
	file, err := os.Open("inputs/day2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		total += getDivisibleAmount(currentLine)
	}

	return strconv.Itoa(total)
}
