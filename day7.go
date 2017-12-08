package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProgramInfo struct {
	name            string
	weight          int
	childrensWeight int
	children        []string
}

func getNewProgramInfo(programInfoString string) ProgramInfo {
	var progInfo ProgramInfo
	var err error
	programInfoSplit := strings.Split(programInfoString, " ")
	for index, info := range programInfoSplit {
		if index == 0 {
			progInfo.name = info
		} else if index == 1 {
			progInfo.weight, err = strconv.Atoi(strings.Replace(strings.Replace(info, ")", "", -1), "(", "", -1))
			check(err)
		} else if index > 2 {
			progInfo.children = append(progInfo.children, strings.Trim(info, ","))
		}
	}

	return progInfo
}

func canCalculateChildrensWeight(progInfo ProgramInfo, leftToPopulate map[string]interface{}) bool {
	for _, child := range progInfo.children {
		_, found := leftToPopulate[child]
		if found {
			return false
		}
	}
	return true
}

func calculateChildrensWeight(progInfo ProgramInfo, tree map[string]ProgramInfo) int {
	totalWeight := 0
	for _, child := range progInfo.children {
		childInfo := tree[child]
		totalWeight += childInfo.weight
		totalWeight += childInfo.childrensWeight
	}
	return totalWeight
}

func populateChildrensWeight(tree map[string]ProgramInfo) map[string]ProgramInfo {
	leftToPopulate := make(map[string]interface{})
	for key := range tree {
		if len(tree[key].children) > 0 {
			leftToPopulate[key] = nil
		}
	}
	for len(leftToPopulate) > 0 {
		for left := range leftToPopulate {
			leftInfo := tree[left]
			if canCalculateChildrensWeight(leftInfo, leftToPopulate) {
				leftInfo.childrensWeight = calculateChildrensWeight(leftInfo, tree)
				tree[left] = leftInfo
				delete(leftToPopulate, left)
			}
		}
	}

	return tree
}

func buildTree() map[string]ProgramInfo {
	file, err := os.Open("inputs/day7.txt")
	check(err)
	defer file.Close()

	tree := make(map[string]ProgramInfo)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		progInfo := getNewProgramInfo(currentLine)
		tree[progInfo.name] = progInfo
	}

	tree = populateChildrensWeight(tree)

	return tree
}

func findBottomOfTower(tree map[string]ProgramInfo) string {
	possibleBottomSlice := []string{}
	for key := range tree {
		possibleBottomSlice = append(possibleBottomSlice, key)
	}

	for _, value := range tree {
		for _, child := range value.children {
			for currentBottomSliceIndex, currentBottomSliceEntry := range possibleBottomSlice {
				if currentBottomSliceEntry == child {
					possibleBottomSlice = append(possibleBottomSlice[:currentBottomSliceIndex], possibleBottomSlice[currentBottomSliceIndex+1:]...)
					break
				}
			}
		}
	}
	return possibleBottomSlice[0]
}

func day7PartA() string {
	tree := buildTree()
	return findBottomOfTower(tree)
}

func areChildrenBalanced(progInfo ProgramInfo, tree map[string]ProgramInfo) (bool, string, int) {
	neededWeight := 0
	problemChild := ""
	balanced := true

	//if no children there is nothing to check
	if len(progInfo.children) <= 1 {
		fmt.Printf("   Not enough children to check '%s'\n", progInfo.name)
		return balanced, problemChild, neededWeight
	}

	weightBaseline := tree[progInfo.children[0]].weight + tree[progInfo.children[0]].childrensWeight
	missMatchCount := 0
	missMatchIndex := 0
	//compare childrens weights to see if they are balanced
	for index, child := range progInfo.children {
		childInfo := tree[child]
		childTotalWeight := childInfo.weight + childInfo.childrensWeight
		if childTotalWeight != weightBaseline {
			//if more than one of these print the baseline is the incorrect value instead
			balanced = false
			missMatchCount++
			missMatchIndex = index
		}
	}

	validIndex := 0
	//first entry is wrong
	if missMatchCount > 1 {
		missMatchIndex = 0
		validIndex = 1
	}

	validTotalWeight := tree[progInfo.children[validIndex]].weight + tree[progInfo.children[validIndex]].childrensWeight
	invalidTotalWeight := tree[progInfo.children[missMatchIndex]].weight + tree[progInfo.children[missMatchIndex]].childrensWeight

	neededWeightChange := invalidTotalWeight - validTotalWeight
	neededWeight = tree[progInfo.children[missMatchIndex]].weight - neededWeightChange

	return balanced, tree[progInfo.children[missMatchIndex]].name, neededWeight
}

//currently is broken and will return back the first problem it encounters
//when there could be more than one (which spot in the chain needs to change to fix it all)
func findIncorrectWeight(tree map[string]ProgramInfo, attemptValidation bool) (string, int) {
	onesWithProblems := make(map[string]int)
	for _, progInfo := range tree {
		if len(progInfo.children) == 0 {
			continue
		}
		isBalanced, problemChild, neededWeight := areChildrenBalanced(progInfo, tree)
		if !isBalanced {
			onesWithProblems[problemChild] = neededWeight
		}
	}

	if attemptValidation {
		// one of these are the problems
		for key, value := range onesWithProblems {
			// fmt.Printf("Name: '%s'\n", key)
			// fmt.Printf("Weight: '%v'\n", tree[key].weight)
			// fmt.Printf("Needed Weight: '%d'\n", value)
			// fmt.Printf("Children: '%v'\n", tree[key].children)

			tempMap := make(map[string]ProgramInfo)
			for k, v := range tree {
				tempMap[k] = v
			}

			tempProgramInfo := tempMap[key]
			tempProgramInfo.weight = value
			tempMap[key] = tempProgramInfo
			tempMap = populateChildrensWeight(tempMap)

			_, tempNeededWeight := findIncorrectWeight(tempMap, false)
			if tempNeededWeight == 0 {
				return key, value
			}
		}
	}

	//this only gets reached on the nested calls
	if len(onesWithProblems) > 0 {
		for key, value := range onesWithProblems {
			return key, value
		}
	}

	return "", 0
}

func day7PartB() string {
	tree := buildTree()
	_, neededWeight := findIncorrectWeight(tree, true)

	return strconv.Itoa(neededWeight)
}
