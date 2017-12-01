package adventofcode2017

import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getIntFromByte(b byte) int {
	result, e := strconv.Atoi(string(b))
	if e != nil {
		panic(e)
	}
	return result
}
