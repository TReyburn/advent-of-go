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