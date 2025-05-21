package main

type Graph struct {
	Nodes map[string]map[string]int
}

func NewGraph() *Graph {
	return &Graph{Nodes: make(map[string]map[string]int)}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	if g.Nodes[from] == nil {
		g.Nodes[from] = make(map[string]int)
	}
	g.Nodes[from][to] = weight

	if g.Nodes[to] == nil {
		g.Nodes[to] = make(map[string]int)
	}
	g.Nodes[to][from] = weight // Se for bidirecional
}
