package adventofcode2017

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isValidPassphraseA(phrase string) bool {
	splitPassphrase := strings.Split(phrase, " ")

	usedWords := make(map[string]int)

	for _, currentWord := range splitPassphrase {
		if _, existed := usedWords[currentWord]; existed {
			return false
		}
		usedWords[currentWord] = 0
	}
	return true
}

func day4PartA() string {
	file, err := os.Open("inputs/day4.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		if isValidPassphraseA(currentLine) {
			total++
		}
	}

	return strconv.Itoa(total)
}

func areAnograms(original, newWord string) bool {
	if len(original) != len(newWord) {
		return false
	}

	usedLetters := make(map[byte]int)

	for i := 0; i < len(original); i++ {
		usedLetters[original[i]] = usedLetters[original[i]] + 1
	}

	for i := 0; i < len(newWord); i++ {
		letterCount, existed := usedLetters[newWord[i]]
		if !existed || letterCount == 0 {
			return false
		}
		usedLetters[newWord[i]] = usedLetters[newWord[i]] - 1
	}

	return true
}

func isValidPassphraseB(phrase string) bool {
	splitPassphrase := strings.Split(phrase, " ")

	usedWords := make(map[string]int)

	for _, currentWord := range splitPassphrase {
		for usedWord := range usedWords {
			if areAnograms(currentWord, usedWord) {
				return false
			}
		}
		usedWords[currentWord] = 0
	}
	return true
}

func day4PartB() string {
	file, err := os.Open("inputs/day4.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		if isValidPassphraseB(currentLine) {
			total++
		}
	}

	return strconv.Itoa(total)
}
