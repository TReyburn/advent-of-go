package adapter

import (
	"errors"
	"sort"
)

type Adapter struct {
	SortedArray []int
}

func (a *Adapter) Load(ns []int) {
	sort.Ints(ns)
	a.SortedArray = append(a.SortedArray, ns...)
	max := a.SortedArray[len(a.SortedArray)-1] + 3
	a.SortedArray= append(a.SortedArray, max)
}

func (a *Adapter) Summarize() (map[int]int, error) {
	res := map[int]int{1:0, 2:0, 3:0}
	prev := 0
	for _, val := range a.SortedArray[1:] {
		diff := val - prev
		if diff > 3 {
			return nil, errors.New("unable to link adapters")
		}
		if diff == 0 {
			return nil, errors.New("unable to use all adapters")
		}
		res[diff]++
		prev = val
	}
	return res, nil
}

func NewAdapter() *Adapter {
	a := Adapter{SortedArray: []int{0}}
	return &a
}