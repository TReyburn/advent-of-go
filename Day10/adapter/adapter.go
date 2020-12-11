package adapter

import "sort"

type Adapter struct {
	SortedArray []int
}

func (a *Adapter) Load(ns []int) {
	sort.Ints(ns)
	a.SortedArray = append(a.SortedArray, ns...)
	max := a.SortedArray[len(a.SortedArray)-1] + 3
	a.SortedArray= append(a.SortedArray, max)
}

func NewAdapter() *Adapter {
	a := Adapter{SortedArray: []int{0}}
	return &a
}