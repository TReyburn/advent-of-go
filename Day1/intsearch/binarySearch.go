package intsearch

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type DataManager struct {
	Data []int
}

func (dm *DataManager) Write(p []byte) (int, error) {
	rb := len(p)
	bss := bytes.Split(p, []byte{13})
	// Iterating over [][]byte, converting each []byte to str, converting str to int, and then appending int to []int
	for _, bString := range bss {
		sInt := string(bString)
		// Removing newline chars
		sInt = strings.Trim(sInt, "\n")
		n, err := strconv.Atoi(sInt)
		if err != nil {
			return 0, err
		}
		dm.Data = append(dm.Data, n)
	}
	return rb, nil
}

func NewDataManager() *DataManager {
	dm := DataManager{Data: make([]int, 0)}
	return &dm
}

type Solution struct {
	Values []int
	MultVal int64
}

func IntSearch(ns []int, sumV int, numVals int) (Solution, error) {
	maxV := sumV - ns[0]
	remainVals := numVals - 1
	// loop over every value in our []int
	for _, v := range ns {
		// no point in searching past here - just return an empty struct and an error
		if v >= maxV {
			return Solution{}, errors.New("no solution found")
		}
		search := sumV - v
		// remainVals == 1 if there's no recursion needed because 1 is the constant default
		if remainVals > 1 {
			// Recursion and creating a new Solution (instead of trying to modify the returned sol)
			sol, err := IntSearch(ns, search, remainVals)
			if err == nil {
				sv := append(sol.Values, v)
				mv := sol.MultVal * int64(v)
				return Solution{Values: sv, MultVal: mv}, nil
			}
		} else {
			b, _ := binarySearch(ns, search)
			if b {
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