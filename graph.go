package graph

import (
	"bytes"
	"fmt"
)

// Each edge has a destination (y), a weight (optional) and
// forms a linked list of incident edges
type edge struct {
	y      int
	weight int
	next   *edge
}

// Each vertex is identified by the index into the edges array
//    AdjList
//      0 -> { y: 1 } -> { y: 3}
//      1 -> { y: 2 } -> { y: 3}
//      2 -> { y: 0 }
type AdjList struct {
	edges       []*edge
	edgeCount   int
	vertexCount int
	directed    bool
	vertexset   []bool
}

// Builds a new adjaceny list, storing the boolean flag that
// distinguishes a directed from an undirected graph.
func New(directed bool) AdjList {
	edges := make([]*edge, 0)
	return AdjList{edges: edges, directed: directed}
}

// Returns the size of the set of vertices
// Runs in O(1) time
func (g AdjList) VertexCount() int {
	return g.vertexCount
}

// Returns the size of the set of edges
// Runs in O(1) time
func (g AdjList) EdgeCount() int {
	return g.edgeCount
}

// Add edge adds an edge node to the x-th spot in the edges slice.
// It resizes if necessary.
func (g *AdjList) AddEdge(x, y int) {
	if max(x,y) >= len(g.edges) {
		g.resizeEdges(max(x,y))
	}
	if g.edges[x].y == -1 {
		g.edges[x].y = y
		g.edges[x].next = nil
	} else {
		newnext := g.edges[x]
		g.edges[x] = &edge{y: y, next: newnext}
	}
	g.edgeCount += 1
	if !g.directed && g.edges[y].y != x {
		g.AddEdge(y, x)
	}
	// record the vertices
	if !g.vertexset[x] {
		g.vertexset[x] = true
		g.vertexCount += 1
	}
	if !g.vertexset[y] {
		g.vertexset[y] = true
		g.vertexCount += 1
	}
}

func (g *AdjList) resizeEdges(size int) {
	diff := size - len(g.edges)
	for i := 0; i <= diff; i++ {
		g.edges = append(g.edges, &edge{y: -1})
		g.vertexset = append(g.vertexset, false)
	}
}

// Builds a graphviz string representing the graph 
func (g AdjList) String() string {
	var buffer bytes.Buffer 
	buffer.WriteString("digraph {\n")
	for x, e := range g.edges {
		if x > -1 && e.y > -1 {
			for c := e; c != nil; c = c.next {
				buffer.WriteString(fmt.Sprintf("  %d -> %d;\n", x, c.y))
			}
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
