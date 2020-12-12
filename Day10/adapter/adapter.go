package adapter

import (
	"errors"
	"sort"
	"strconv"
	"strings"
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

func (a *Adapter) Write(p []byte) (int, error) {
	rb := len(p)
	ns := make([]int, 0)
	rawStr := string(p)
	split := strings.Split(rawStr, "\r\n")
	for _, str := range split {
		if str != "" {
			n, err := strconv.Atoi(str)
			if err != nil {
				return 0, err
			}
			ns = append(ns, n)
		}
	}
	a.Load(ns)
	return rb, nil
}

type LiFoQueue struct {
	Items []int
}

func (q *LiFoQueue) Pop() int {
	i := q.Items[0]
	if len(q.Items) > 1{
		q.Items = q.Items[1:]
	} else {
		q.Items = make([]int, 0)
	}
	return i
}

func (q *LiFoQueue) Push(n int) {
	q.Items = append([]int{n}, q.Items...)
}

func NewAdapter() *Adapter {
	a := Adapter{SortedArray: []int{0}}
	return &a
}