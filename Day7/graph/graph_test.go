package graph

import "testing"

func TestGraph_AddNode(t *testing.T) {
	g := NewGraph()
	n := NewNode("Test")
	n2 := NewNode("Test2")

	g.AddNode(n)
	g.AddNode(n2)

	if len(g.Nodes) != 2 {
		t.Error("Expected Nodes len of 2; got", len(g.Nodes))
	}
}

func TestGraph_AddNodeDuplicates(t *testing.T) {
	g := NewGraph()
	n := NewNode("Test")
	n2 := NewNode("Test")

	g.AddNode(n)
	g.AddNode(n2)

	if len(g.Nodes) != 1 {
		t.Error("Expected Nodes len of 1; got", len(g.Nodes))
	}
}