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
	high := len(ns) - 1
	low := 0

	for low <= high {
		mid := (low + high) / 2

		if ns[mid] == searchV {
			return true, mid
		}

		if ns[mid] < searchV {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false, 0
}