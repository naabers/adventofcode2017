package adventofcode2017

import "strconv"

func day3PartA() string {
	input := 265149
	bottomRightCorner := 1
	count := 1
	for bottomRightCorner <= input {
		bottomRightCorner += 8 * count
		count++
	}
	heightWidth := 2*count - 1
	toMiddle := heightWidth / 2
	bottomLeftCorner := bottomRightCorner - heightWidth + 1
	topLeftCorner := bottomLeftCorner - heightWidth + 1
	topRightCorner := topLeftCorner - heightWidth + 1

	// fmt.Printf("count            %d\n", count)
	// fmt.Printf("bottomRight      %d\n", bottomRightCorner)
	// fmt.Printf("bottomLeftCorner %d\n", bottomLeftCorner)
	// fmt.Printf("topLeftCorner    %d\n", topLeftCorner)
	// fmt.Printf("topRightCorner   %d\n", topRightCorner)
	// fmt.Printf("heightWidth      %d\n", heightWidth)
	// fmt.Printf("toMiddle         %d\n", toMiddle)
	// fmt.Printf("input            %d\n", input)

	//TODO fix for corners
	result := toMiddle
	if input > bottomLeftCorner {
		// fmt.Printf("bottom\n")
		if input > bottomRightCorner-toMiddle {
			result += input - (bottomRightCorner - toMiddle)
		} else {
			result += (bottomRightCorner - heightWidth + 1 + toMiddle) - input
		}
	} else if input > topLeftCorner {
		// fmt.Printf("left\n")
		if input > bottomLeftCorner-toMiddle {
			result += input - (bottomLeftCorner - toMiddle)
		} else {
			result += (bottomLeftCorner - heightWidth + 1 + toMiddle) - input
		}
	} else if input > topRightCorner {
		// fmt.Printf("top\n")
		if input > topLeftCorner-toMiddle {
			result += input - (topLeftCorner - toMiddle)
		} else {
			result += (topLeftCorner - heightWidth + 1 + toMiddle) - input
		}
	} else if input > (bottomRightCorner - 4*heightWidth + 3) {
		// fmt.Printf("right\n")
		if input > topRightCorner-toMiddle {
			result += input - (topRightCorner - toMiddle)
		} else {
			result += (topRightCorner - heightWidth + 1 + toMiddle) - input
		}
	} else {
		panic("whoops\n")
	}

	return strconv.Itoa(result)
}
