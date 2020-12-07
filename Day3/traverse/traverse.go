package traverse

func DumbTraverse(tm []string, search string, xInc int, yInc int) int {
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
