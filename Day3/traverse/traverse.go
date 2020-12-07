package traverse

func Traverse(tm []string, search string, xInc int, yInc int) int {
	xLoc := 0
	yLoc := 0
	xLen := len(tm)
	yLen := len(tm[0])
	count := 0
	// Multibyte chars will break this...
	bSearch := []byte(search)[0]

	for xLoc < xLen {
		if tm[xLoc][yLoc] == bSearch {
			count++
		}
		xLoc += xInc
		yLoc += yInc
		if yLoc >= yLen {
			yLoc -= yLen
		}
	}
	return count
}

func MultiTraverse(tm []string, search string, angles [][]int) int {
	total := 1
	for _, angleSet := range angles {
		xInc := angleSet[0]
		yInc := angleSet[1]
		res := Traverse(tm, search, xInc, yInc)
		total *= res
	}
	return total
}
