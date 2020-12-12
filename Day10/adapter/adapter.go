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

func (a *Adapter) GetEdges(n *Node) []*Node {
	res := make([]*Node, 0)
	if n.Index >= len(a.SortedArray) - 1 {
		return res
	}
	for idx, val := range a.SortedArray[n.Index+1:] {
		if val > n.Value + 3 {
			return res
		}
		nn := NewNode(idx + n.Index+1, val)
		res = append(res, nn)
	}
	return res
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

func (a *Adapter) BFSSummarize() int {
	counter := 0
	startNode := NewNode(0, 0)
	q := NewQueue()
	q.Push(startNode)
	end := a.SortedArray[len(a.SortedArray) - 1]

	for len(q.Items) > 0 {
		n := q.Pop()
		if n.Value == end {
			counter++
		} else {
			edges := a.GetEdges(n)
			q.BulkPush(edges)
		}
	}
	return counter
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

type Node struct {
	Index int
	Value int
	Visited bool
}

type LiFoQueue struct {
	Items []*Node
}

func (q *LiFoQueue) Pop() *Node {
	i := q.Items[0]
	if len(q.Items) > 1{
		q.Items = q.Items[1:]
	} else {
		q.Items = make([]*Node, 0)
	}
	return i
}

func (q *LiFoQueue) Push(n *Node) {
	q.Items = append([]*Node{n}, q.Items...)
}

func (q *LiFoQueue) BulkPush(ns []*Node) {
	for _, n := range ns {
		q.Push(n)
	}
}

func NewAdapter() *Adapter {
	a := Adapter{SortedArray: []int{0}}
	return &a
}

func NewQueue() *LiFoQueue {
	q := LiFoQueue{Items: make([]*Node, 0)}
	return &q
}

func NewNode(idx int, val int) *Node {
	n := Node{
		Index:   idx,
		Value:   val,
		Visited: false,
	}
	return &n
}