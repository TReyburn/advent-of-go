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

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	n := NewNode("Test")
	n2 := NewNode("Test2")
	cost := 5

	g.AddEdge(n, n2, cost)

	if len(g.Nodes) != 2 {
		t.Error("Expected Nodes len of 2; got", len(g.Nodes))
	}

	if len(g.Edges) != 1 {
		t.Error("Expected edges len of 1; got", len(g.Edges))
	}
}

func TestGraph_AddEdgeDuplicateNodes(t *testing.T) {
	g := NewGraph()
	n := NewNode("Test")
	n2 := NewNode("Test2")
	n3 := NewNode("Test3")
	cost := 5

	g.AddEdge(n, n2, cost)
	g.AddEdge(n2, n3, cost)

	if len(g.Nodes) != 3 {
		t.Error("Expected Nodes len of 3; got", len(g.Nodes))
	}

	if len(g.Edges) != 2 {
		t.Error("Expected edges len of 2; got", len(g.Edges))
	}
}