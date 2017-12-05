package adventofcode2017

import (
	"bufio"
	"os"
	"strconv"
)

func createMapOfInstructions() map[int]int {
	file, err := os.Open("inputs/day5.txt")
	check(err)
	defer file.Close()

	instructions := make(map[int]int)

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		instructionNumber, err := strconv.Atoi(currentLine)
		check(err)
		instructions[index] = instructionNumber
		index++
	}
	return instructions
}

func day5PartA() string {
	instructions := createMapOfInstructions()
	index := 0
	total := 0
	for index < len(instructions) {
		currentInstruction := instructions[index]
		instructions[index] = instructions[index] + 1
		index += currentInstruction
		total++
	}

	return strconv.Itoa(total)
}

func day5PartB() string {
	instructions := createMapOfInstructions()
	index := 0
	total := 0
	for index < len(instructions) {
		currentInstruction := instructions[index]
		if currentInstruction >= 3 {
			instructions[index] = instructions[index] - 1
		} else {
			instructions[index] = instructions[index] + 1
		}
		index += currentInstruction
		total++
	}

	return strconv.Itoa(total)
}
