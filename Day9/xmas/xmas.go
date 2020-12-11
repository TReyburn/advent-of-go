package xmas

import "errors"

type Decoder struct {
	Queue *FixedQueue
	Preamble map[int]bool
	Remainder []int
}

func (d *Decoder) Load(ns []int) error {
	l := d.Queue.GetMaxLen()

	if l >= len(ns) {
		return errors.New("int slice is too short to load")
	}
	for _, n := range ns[:l] {
		d.Queue.Load(n)
		d.Preamble[n] = false
	}
	d.Remainder = append(d.Remainder, ns[l:]...)
	return nil
}

type FixedQueue struct {
	Len int
	Queue []int
}

func (q *FixedQueue) GetMaxLen() int {
	return q.Len
}

func (q *FixedQueue) Load(n int) {
	if len(q.Queue) == q.Len {
		q.Queue = append(q.Queue[1:], n)
	} else {
		q.Queue = append(q.Queue, n)
	}
}

func NewDecoder(n int) *Decoder {
	d := Decoder{
		Queue:     NewFixedQueue(n),
		Preamble:  make(map[int]bool),
		Remainder: make([]int, 0),
	}
	return &d
}

func NewFixedQueue(n int) *FixedQueue {
	q := FixedQueue{
		Len:   n,
		Queue: make([]int, 0),
	}
	return &q
}