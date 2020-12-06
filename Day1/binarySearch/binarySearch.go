package binarySearch

//func IntSearch(ns []int, sumV int) (Solution, error) {
//	maxV := sumV - ns[0]
//	for i, v := range ns {
//		if v >= maxV {
//			return Solution{}, errors.New("no solution found")
//		} else {
//
//		}
//
//	}
//
//}

func binarySearch(ns []int, searchV int) (bool, int) {
	nsLen := len(ns)
	searchLoc := nsLen/2
	previousLoc := 0
	for {
		foundV := ns[searchLoc]
		if foundV == searchV {
			return true, searchLoc
		} else {
			if foundV < searchV {
				previousLoc = searchLoc
				searchLoc = searchLoc + (searchLoc/2)
				if searchLoc == previousLoc {
					searchLoc++
					if searchLoc == nsLen {
						return false, 0
					}
				}
			} else {
				previousLoc = searchLoc
				searchLoc = searchLoc - (searchLoc/2)
				if searchLoc == previousLoc {
					searchLoc--
					if searchLoc < 0 {
						return false, 0
					}
				}
			}
		}
	}
}
