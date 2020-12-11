package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdapter_Load(t *testing.T) {
	ns := []int{3, 2, 1, 4}
	a := NewAdapter()
	a.Load(ns)

	assert.Equal(t, []int{0,1,2,3,4,7}, a.SortedArray)
}

func TestAdapter_Summarize(t *testing.T) {
	ns := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	a := NewAdapter()
	a.Load(ns)

	res, err := a.Summarize()
	if err != nil {
		t.Error("Unexpected error while summarizing jolt jumps:", err)
	}

	assert.Equal(t, 7, res[1])
	assert.Equal(t, 5, res[3])
}