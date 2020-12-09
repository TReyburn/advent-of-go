package graph

import (
	"bytes"
	"errors"
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

type Queue struct {
	Nodes []*Node
}

func (q *Queue) Enqueue(node *Node) {
	q.Nodes = append(q.Nodes, node)
}

func (q *Queue) Dequeue() *Node {
	node := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return node
}

func (q *Queue) Size() int {
	return len(q.Nodes)
}

func (g *Graph) AddEdge(parent *Node, child *Node, cost int) {
	edge := Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}
	// If we have duplicate edges then we are kinda screwed. What about circular edges?
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

func (g *Graph) BFSTraverse(startNode *Node, endNodeName string) bool {
	q := NewQueue()
	vistedNodes := []*Node{startNode}
	startEdges := g.GetNodeEdges(startNode)
	for _, edge := range startEdges {
		q.Enqueue(edge.Child)
	}

	for {
		if q.Size() == 0 {
			break
		}
		node := q.Dequeue()
		if node.Name == endNodeName {
			return true
		}
		visited := false
		for _, vnode := range vistedNodes {
			if vnode == node {
				visited = true
				break
			}
		}
		if !visited {
			vistedNodes = append(vistedNodes, node)
			edges := g.GetNodeEdges(node)
			for _, edge := range edges {
				q.Enqueue(edge.Child)
			}
		}
	}
	return false
}

func (g *Graph) GetNodeEdges(node *Node) []*Edge {
	edges := make([]*Edge, 0)
	for _, edge := range g.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *Graph) GetNodeByName(name string) (*Node, error) {
	for _, node := range g.Nodes {
		if node.Name == name {
			return node, nil
		}
	}
	return NewNode(""), errors.New("could not find node")
}

func (g *Graph) LoadStr(rawStr string) error {
	split := strings.Split(rawStr, "s contain ")
	parent, err := g.GetNodeByName(split[0])
	if err != nil {
		parent = NewNode(split[0])
	}
	if split[1] == "no other bags." {
		g.AddNode(parent)
		return nil
	}
	children := strings.Split(split[1], ",")
	for _, childStr := range children {
		childStr = strings.TrimSuffix(childStr, ".")
		childStr = strings.TrimSuffix(childStr, "s")
		childStr = strings.TrimSpace(childStr)
		childSlice := strings.SplitAfterN(childStr, " ", 2)
		costStr := strings.TrimSpace(childSlice[0])
		cost, err := strconv.Atoi(costStr)
		if err != nil {
			return err
		}
		child, err := g.GetNodeByName(childSlice[1])
		if err != nil {
			child = NewNode(childSlice[1])
		}
		g.AddEdge(parent, child, cost)
	}
	return nil
}

func (g *Graph) Write(p []byte) (int, error) {
	rb := len(p)
	bss := bytes.Split(p, []byte{13})
	for _, bString := range bss {
		rawStr := string(bString)
		rawStr = strings.Trim(rawStr, "\n")
		err := g.LoadStr(rawStr)
		if err != nil {
			return 0, err
		}
	}
	return rb, nil
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

func NewQueue() *Queue {
	q := Queue{Nodes: make([]*Node, 0)}
	return &q
}