package graph

import (
	"bytes"
	"fmt"
	"github.com/tlehman/ds"
)


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
func (g *AdjList) addEdge(x, y int) {
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

func (g *AdjList) AddEdge(x, y int) {
	g.addEdge(x,y)
	if !g.directed {
		g.addEdge(y,x)
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
	var arrow string
	if g.directed {
	    buffer.WriteString("di")
		arrow = "->"
	} else {
		arrow = "--"
	}
	buffer.WriteString("graph {\n")
	for x, e := range g.edges {
		if x > -1 && e.y > -1 {
			for c := e; c != nil; c = c.next {
				buffer.WriteString(fmt.Sprintf("  %d %s %d;\n", x, arrow, c.y))
			}
		}
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Find connected components in an undirected graph.
// The return value of Components() is a slice of integers
// that associates each vertex with a component number.
func (g *AdjList) Components() []int {
	// each index v is a vertex, and the value comps[v]
	// is the number of the component v is in
	comps := make([]int, len(g.edges))
	discovered := make([]bool, len(g.edges))
	compNum := 0
	for v, _ := range g.edges {
		if !discovered[v] {
			// indices of component containing v
			idcsComp := g.bfs(v)
			for _, x := range idcsComp {
				comps[x] = compNum
				discovered[x] = true
			}
			compNum++
		}
	}

	return comps
}


const (
	undiscovered = iota
	discovered = iota
	processed = iota
)
	
// do a Breadth-First Search from v, return indices in same
// component as v
func (g *AdjList) bfs(v int) []int {
	state := make([]int, len(g.edges)) // all initially undiscovered
	indices := make([]int, 0)
	q := ds.NewQueue()
	q.Enqueue(v)
	for q.Len() > 0 {
		x := q.Dequeue().(int)
		indices = append(indices, x)
		for e := g.edges[x]; e != nil; e = e.next {
			if e.y > -1 && state[e.y] == undiscovered {
				state[x] = discovered 
				q.Enqueue(e.y)
			}
		}
		state[x] = processed
	}
	return indices
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
