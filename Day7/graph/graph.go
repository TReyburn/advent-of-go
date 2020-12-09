package graph

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