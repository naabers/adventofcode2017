package main

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

func isIntSliceEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
