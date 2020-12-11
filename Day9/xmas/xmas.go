package xmas

import (
	"errors"
)

type Decoder struct {
	Queue *FixedQueue
	Preamble map[int]bool
	Remainder []int
}

func (d *Decoder) AddKey(n int) {
	d.Preamble[n] = false
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

func (d *Decoder) Shift() error {
	l := len(d.Remainder)
	if l == 0 {
		return errors.New("no remainder left to shift")
	}
	addKey := d.Remainder[0]
	d.AddKey(addKey)
	if l >= 2 {
		d.Remainder = append(d.Remainder[:0], d.Remainder[1:]...)
	} else {
		d.Remainder = make([]int, 0)
	}
	removeKey := d.Queue.Queue[0]
	delete(d.Preamble, removeKey)
	d.Queue.Load(addKey)
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