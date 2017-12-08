package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getInstructionValue(registers map[string]int, instructionPieces string) int {

	instructionValue, err := strconv.Atoi(instructionPieces)
	if err == nil {
		return instructionValue
	}
	return registers[instructionPieces]
}

func updateRegister(registers map[string]int, instructionPieces []string) {
	adjustAmount, err := strconv.Atoi(instructionPieces[2])
	check(err)
	if instructionPieces[1] == "inc" {
		registers[instructionPieces[0]] = registers[instructionPieces[0]] + adjustAmount
	} else {
		registers[instructionPieces[0]] = registers[instructionPieces[0]] - adjustAmount
	}
}

func processInstruction(registers map[string]int, instruction string) {
	instructionPieces := strings.Split(instruction, " ")
	if len(instructionPieces) != 7 {
		panic("Instruction is invalid")
	}

	leftHalfValue := getInstructionValue(registers, instructionPieces[4])
	rightHalfValue := getInstructionValue(registers, instructionPieces[6])

	switch instructionPieces[5] {
	case ">":
		if leftHalfValue > rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	case ">=":
		if leftHalfValue >= rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	case "<":
		if leftHalfValue < rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	case "<=":
		if leftHalfValue <= rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	case "==":
		if leftHalfValue == rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	case "!=":
		if leftHalfValue != rightHalfValue {
			updateRegister(registers, instructionPieces)
		}
	default:
		panic("Unknown operator")
	}
}

func day8PartA() string {
	file, err := os.Open("inputs/day8.txt")
	check(err)
	defer file.Close()

	registers := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		processInstruction(registers, currentLine)
	}

	largest := 0
	for _, value := range registers {
		if value > largest {
			largest = value
		}
	}

	return strconv.Itoa(largest)
}

func day8PartB() string {
	file, err := os.Open("inputs/day8.txt")
	check(err)
	defer file.Close()

	registers := make(map[string]int)
	largest := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		processInstruction(registers, currentLine)
		for _, value := range registers {
			if value > largest {
				largest = value
			}
		}
	}

	return strconv.Itoa(largest)
}
