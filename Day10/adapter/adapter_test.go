package adapter

import (
	"github.com/TReyburn/advent-of-go/common/filehandler"
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

func TestAdapter_Write(t *testing.T) {
	a := NewAdapter()
	err := filehandler.LoadInputFile("testdata/test.txt", a)
	if err != nil {
		t.Error("Unexpected error loading file:", err)
	}
	res, err := a.Summarize()
	if err != nil {
		t.Error("Unexpected error while summarizing:", err)
	}

	assert.Equal(t, 22, res[1])
	assert.Equal(t, 10, res[3])
}

func TestAdapter_GetEdges(t *testing.T) {
	ns := []int{1, 2, 4}
	a := NewAdapter()
	a.Load(ns)

	n := NewNode(0, 0)
	res := a.GetEdges(n)
	assert.Equal(t, 2, len(res))
}

func TestAdapter_GetEdgesLastNode(t *testing.T) {
	ns := []int{1, 2, 4}
	a := NewAdapter()
	a.Load(ns)

	n := NewNode(4, 7)
	res := a.GetEdges(n)
	assert.Equal(t, 0, len(res))
}

func TestLiFoQueue_BulkPush(t *testing.T) {
	ns := []int{1, 2, 4}
	a := NewAdapter()
	a.Load(ns)

	n := NewNode(0, 0)
	res := a.GetEdges(n)
	q := NewQueue()
	q.BulkPush(res)

	assert.Equal(t, 2, len(q.Items))
}

func TestAdapter_BFSSummarize(t *testing.T) {
	ns := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	a := NewAdapter()
	a.Load(ns)

	res := a.BFSSummarize()
	assert.Equal(t, 8, res)
}