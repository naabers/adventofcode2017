package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func populateMemoryBank() []int {
	file, err := os.Open("inputs/day6.txt")
	check(err)
	defer file.Close()

	memoryBank := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		currentSplit := strings.Split(currentLine, "\t")
		for _, blockString := range currentSplit {
			block, err := strconv.Atoi(blockString)
			check(err)
			memoryBank = append(memoryBank, block)
		}
		break
	}
	return memoryBank
}

func redistributeBank(memoryBank []int) []int {
	max := 0
	for _, block := range memoryBank {
		if block > max {
			max = block
		}
	}
	maxFound := false
	redistribute := max
	index := 0
	for redistribute > 0 {
		if !maxFound {
			if memoryBank[index] == max {
				maxFound = true
				memoryBank[index] = 0
			}
		} else {
			memoryBank[index]++
			redistribute--
		}

		if index == (len(memoryBank) - 1) {
			index = 0
		} else {
			index++
		}
	}

	return memoryBank
}

//if it is not new it also returns back the index where it was found
func isNewMemoryBank(memoryBank []int, oldMemoryBanks [][]int) (bool, int) {
	for index := 0; index < len(oldMemoryBanks); index++ {
		if isIntSliceEqual(memoryBank, oldMemoryBanks[index]) {
			return false, index
		}
	}

	return true, 0
}

func requiresRedistribution(memoryBank []int) bool {
	if len(memoryBank) == 0 {
		return false
	}
	compareTo := memoryBank[0]
	for _, block := range memoryBank {
		if block != compareTo {
			return true
		}
	}
	return false
}

func addToMemoryBank(memoryBank []int, oldMemoryBanks [][]int) [][]int {
	tempBank := []int{}
	for _, block := range memoryBank {
		tempBank = append(tempBank, block)
	}
	oldMemoryBanks = append(oldMemoryBanks, tempBank)
	return oldMemoryBanks
}

func day6PartA() string {
	memoryBank := populateMemoryBank()
	oldMemoryBanks := [][]int{}

	for {
		isNewBank, _ := isNewMemoryBank(memoryBank, oldMemoryBanks)
		if !isNewBank {
			break
		}
		if !requiresRedistribution(memoryBank) {
			break
		}
		oldMemoryBanks = addToMemoryBank(memoryBank, oldMemoryBanks)
		memoryBank = redistributeBank(memoryBank)
	}
	return strconv.Itoa(len(oldMemoryBanks))
}

func day6PartB() string {
	memoryBank := populateMemoryBank()
	oldMemoryBanks := [][]int{}
	isNewBank := false
	oldBankIndex := 0
	for {
		isNewBank, oldBankIndex = isNewMemoryBank(memoryBank, oldMemoryBanks)
		if !isNewBank {
			break
		}
		if !requiresRedistribution(memoryBank) {
			break
		}
		oldMemoryBanks = addToMemoryBank(memoryBank, oldMemoryBanks)
		memoryBank = redistributeBank(memoryBank)
	}
	return strconv.Itoa(len(oldMemoryBanks) - oldBankIndex)
}
