package intSearch

import (
	"errors"
)

type Solution struct {
	Values []int
	MultVal int64
}

func IntSearch(ns []int, sumV int, numVals int) (Solution, error) {
	maxV := sumV - ns[0]
	remainVals := numVals - 1
	for _, v := range ns {
		if v >= maxV {
			return Solution{}, errors.New("no solution found")
		}
		search := sumV - v
		if remainVals > 1 {
			sol, err := IntSearch(ns, search, remainVals)
			if err == nil {
				sv := append(sol.Values, v)
				mv := sol.MultVal * int64(v)
				return Solution{Values: sv, MultVal: mv}, nil
			}
		} else {
			b, _ := binarySearch(ns, search)
			if b == true {
				return Solution{Values: []int{v, search}, MultVal: int64(v * search)}, nil
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