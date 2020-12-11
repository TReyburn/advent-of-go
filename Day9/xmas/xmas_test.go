package xmas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedQueue_GetLen(t *testing.T) {
	q := NewFixedQueue(5)

	l := q.GetMaxLen()
	assert.Equal(t, 5, l)
}

func TestFixedQueue_Load(t *testing.T) {
	ns := []int{0, 1, 2, 3, 4, 5}
	q := NewFixedQueue(3)
	for _, n := range ns {
		q.Load(n)
	}

	assert.Equal(t, []int{3, 4, 5}, q.Queue)
}

func TestDecoder_Load(t *testing.T) {
	ns := []int{0, 1, 2, 3, 4}
	d := NewDecoder(3)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.Equal(t, 3, len(d.Queue.Queue))
	assert.Equal(t, 3, len(d.Preamble))
	assert.Equal(t, 2, len(d.Remainder))
}