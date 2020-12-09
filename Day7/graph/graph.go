package graph

import (
	"strconv"
	"strings"
)

type Node struct {
	Name string
}

type Edge struct {
	Parent *Node
	Child *Node
	Cost int
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func (g *Graph) AddEdge(parent *Node, child *Node, cost int) {
	edge := Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}
	g.Edges = append(g.Edges, &edge)
	g.AddNode(parent)
	g.AddNode(child)
}

func (g *Graph) AddNode(node *Node) {
	exists := false
	for _, n := range g.Nodes{
		if n.Name == node.Name {
			exists = true
		}
	}

	if !exists {
		g.Nodes = append(g.Nodes, node)
	}
}

func (g *Graph) LoadStr(rawStr string) error {
	split := strings.Split(rawStr, "s contain ")
	parent := NewNode(split[0])
	if split[1] == "no other bags." {
		g.AddNode(parent)
		return nil
	}
	children := strings.Split(split[1], ",")
	for _, childStr := range children {
		childStr = strings.TrimSuffix(childStr, "s")
		childStr = strings.TrimSuffix(childStr, "s.")
		childStr = strings.TrimSpace(childStr)
		childSlice := strings.SplitAfterN(childStr, " ", 2)
		costStr := strings.TrimSpace(childSlice[0])
		cost, err := strconv.Atoi(costStr)
		if err != nil {
			return err
		}
		child := NewNode(childSlice[1])
		g.AddEdge(parent, child, cost)
	}
	return nil
}

func NewNode(name string) *Node {
	n := Node{Name: name}
	return &n
}

func NewGraph() *Graph {
	g := Graph{
		Nodes: make([]*Node, 0),
		Edges: make([]*Edge, 0),
	}
	return &g
}