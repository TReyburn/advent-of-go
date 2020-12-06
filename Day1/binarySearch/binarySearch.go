package binarySearch

import (
	"errors"
)

type Solution struct {
	N1 int
	N2 int
	MultVal int64
}

func IntSearch(ns []int, sumV int) (Solution, error) {
	maxV := sumV - ns[0]
	for _, v := range ns {
		if v >= maxV {
			return Solution{}, errors.New("no solution found")
		} else {
			search := sumV - v
			b, _ := binarySearch(ns, search)
			if b == true {
				return Solution{N1: v, N2: search, MultVal: int64(v * search)}, nil
			}
		}
	}
	return Solution{}, errors.New("no solution found")
}

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