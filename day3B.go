package adventofcode2017

import (
	"fmt"
	"strconv"
)

const (
	left = iota
	up
	right
	down
)

type gridHelper struct {
	direction int
	x         int
	y         int
	minX      int
	minY      int
	maxX      int
	maxY      int
}

func getValue(grid [20][20]int, x, y int) int {
	result := 0
	if x > 0 {
		result += grid[x-1][y]
		if y < len(grid)-1 {
			result += grid[x-1][y+1]
		}
		if y > 0 {
			result += grid[x-1][y-1]
		}
	}
	if x < len(grid)-1 {
		result += grid[x+1][y]
		if y < len(grid)-1 {
			result += grid[x+1][y+1]
		}
		if y > 0 {
			result += grid[x+1][y-1]
		}
	}
	if y < len(grid)-1 {
		result += grid[x][y+1]
	}
	if y > 0 {
		result += grid[x][y-1]
	}
	return result
}

func goRight(router gridHelper) gridHelper {
	router.x++
	if router.x > router.maxX {
		router.direction = up
		router.maxX = router.x
	}
	return router
}

func goLeft(router gridHelper) gridHelper {
	router.x--
	if router.x < router.minX {
		router.direction = down
		router.minX = router.x
	}
	return router
}

func goUp(router gridHelper) gridHelper {
	router.y++
	if router.y > router.maxY {
		router.direction = left
		router.maxY = router.y
	}
	return router
}

func goDown(router gridHelper) gridHelper {
	router.y--
	if router.y < router.minY {
		router.direction = right
		router.minY = router.y
	}
	return router
}

func createGrid(goal int) int {
	router := gridHelper{right, 10, 10, 10, 10, 10, 10}

	grid := [20][20]int{}

	grid[router.x][router.y] = 1

	for i := 1; i < goal; i++ {
		if router.direction == right {
			router = goRight(router)
		} else if router.direction == left {
			router = goLeft(router)
		} else if router.direction == up {
			router = goUp(router)
		} else if router.direction == down {
			router = goDown(router)
		} else {
			panic("Unknown direction")
		}
		value := getValue(grid, router.x, router.y)
		grid[router.x][router.y] = value
		if value > goal {
			return value
		}
	}

	fmt.Printf("Last Value: '%d'\n", grid[router.x][router.y])
	panic("Failed to find next value")
}

func day3PartB() string {
	input := 265149
	result := createGrid(input)
	return strconv.Itoa(result)
}
