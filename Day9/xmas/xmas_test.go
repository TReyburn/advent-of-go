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
	err := d.AddKey(1)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.Equal(t, 1, len(d.Preamble))
	assert.Equal(t, false, d.Preamble[1])
}

func TestDecoder_AddKeyAlreadyExists(t *testing.T) {
	d := NewDecoder(0)
	err := d.AddKey(1)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	err = d.AddKey(1)
	assert.Error(t, err, "key already exists")
}

func TestDecoder_Shift(t *testing.T) {
	ns := []int{0, 1, 2, 3, 4}
	d := NewDecoder(2)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.Equal(t, 2, len(d.Preamble))
	assert.Equal(t, 2, len(d.Queue.Queue))
	assert.Equal(t, 3, len(d.Remainder))
	assert.Equal(t, map[int]bool{0:false, 1:false}, d.Preamble)
	assert.Equal(t, []int{0, 1}, d.Queue.Queue)
	assert.Equal(t, []int{2, 3, 4}, d.Remainder)

	err = d.Shift()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	assert.Equal(t, 2, len(d.Preamble))
	assert.Equal(t, 2, len(d.Queue.Queue))
	assert.Equal(t, 2, len(d.Remainder))
	assert.Equal(t, map[int]bool{1:false, 2:false}, d.Preamble)
	assert.Equal(t, []int{1, 2}, d.Queue.Queue)
	assert.Equal(t, []int{3, 4}, d.Remainder)
}

func TestDecoder_ShiftLastShift(t *testing.T) {
	ns := []int{0, 1, 2}
	d := NewDecoder(2)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.Equal(t, 2, len(d.Preamble))
	assert.Equal(t, 2, len(d.Queue.Queue))
	assert.Equal(t, 1, len(d.Remainder))
	assert.Equal(t, map[int]bool{0:false, 1:false}, d.Preamble)
	assert.Equal(t, []int{0, 1}, d.Queue.Queue)
	assert.Equal(t, []int{2}, d.Remainder)

	err = d.Shift()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	assert.Equal(t, 2, len(d.Preamble))
	assert.Equal(t, 2, len(d.Queue.Queue))
	assert.Equal(t, 0, len(d.Remainder))
	assert.Equal(t, map[int]bool{1:false, 2:false}, d.Preamble)
	assert.Equal(t, []int{1, 2}, d.Queue.Queue)
	assert.Equal(t, []int{}, d.Remainder)
}

func TestDecoder_ShiftCantShift(t *testing.T) {
	ns := []int{0, 1, 2}
	d := NewDecoder(2)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	err = d.Shift()
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	err = d.Shift()
	assert.Error(t, err, "no remainder left to shift")
}

func TestDecoder_ProcessTrue(t *testing.T) {
	ns := []int{0, 1, 2, 3}
	d := NewDecoder(3)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.True(t, d.Process())
}

func TestDecoder_ProcessFalse(t *testing.T) {
	ns := []int{0, 1, 2, 4}
	d := NewDecoder(3)
	err := d.Load(ns)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	assert.False(t, d.Process())
}