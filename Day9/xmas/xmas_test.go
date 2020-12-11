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

func TestDecoder_LoadNotEnough(t *testing.T) {
	ns := []int{0, 1, 2, 3, 4}
	d := NewDecoder(5)
	err := d.Load(ns)
	assert.Error(t, err, "int slice is too short to load")

}

func TestDecoder_AddKey(t *testing.T) {
	d := NewDecoder(0)
	d.AddKey(1)

	assert.Equal(t, 1, len(d.Preamble))
	assert.Equal(t, false, d.Preamble[1])
}