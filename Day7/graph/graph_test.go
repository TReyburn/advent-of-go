package graph

import (
	"fmt"
	"github.com/TReyburn/advent-of-go/common/filehandler"
	"testing"
)

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

func TestGraph_LoadStr(t *testing.T) {
	str := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	g := NewGraph()

	err := g.LoadStr(str)
	if err != nil {
		t.Error("Unexpected err loading strings", err)
	}

	if len(g.Nodes) != 3 {
		t.Error("Expected Nodes len of 3; got", len(g.Nodes))
	}

	if len(g.Edges) != 2 {
		t.Error("Expected Edges len of 2; got", len(g.Edges))
	}
}

func TestGraph_LoadStrNoEdge(t *testing.T) {
	str := "faded blue bags contain no other bags."
	g := NewGraph()

	err := g.LoadStr(str)
	if err != nil {
		t.Error("Unexpected err loading strings", err)
	}

	if len(g.Nodes) != 1 {
		t.Error("Expected Nodes len of 3; got", len(g.Nodes))
	}

	if len(g.Edges) != 0 {
		t.Error("Expected Edges len of 2; got", len(g.Edges))
	}
}

func TestGraph_Write(t *testing.T) {
	g := NewGraph()
	err := filehandler.LoadInputFile("testdata/test.txt", g)
	if err != nil {
		t.Error("Unexpected error writing", err)
	}

	if len(g.Nodes) != 9 {
		t.Error("Expected Nodes len of 9; got", len(g.Nodes))
		for _, node := range g.Nodes {
			fmt.Println("Node:", node.Name)
		}
	}

	if len(g.Edges) != 13 {
		t.Error("Expected Edges len of 13; got", len(g.Edges))
	}
}