package graph

import (
	"testing"
)

func digraphSixVertices() AdjList {
	directed := true
	g := New(directed)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 2)
	g.AddEdge(4, 2)
	g.AddEdge(2, 3)
	g.AddEdge(0, 5)
	return g
}

func TestAddEdge(t *testing.T) {
	g := digraphSixVertices()
	if g.VertexCount() != 6 {
		t.Fatalf("Expected vertex count to be 6, got %d", g.VertexCount())
	}
	if g.EdgeCount() != 8 {
		t.Fatalf("Expected vertex count to be 6, got %d", g.EdgeCount())
	}
}

func TestString(t *testing.T) {
	g := digraphSixVertices()
	expected := `digraph {
  0 -> 5;
  0 -> 2;
  1 -> 4;
  1 -> 3;
  2 -> 3;
  4 -> 2;
  4 -> 5;
  5 -> 2;
}`
	actual := g.String()
	if actual != expected {
		t.Fatal("Graphviz representation not working correctly")
		t.Fatalf("actual result: \n%s\n", actual)
	}
}

func graphSevenVertices() AdjList {
	directed := false
	g := New(directed)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 4)
	return g
}

func TestStringUndirected(t *testing.T) {
	var g AdjList = graphSevenVertices()
	expected := `strict graph {
  0 -- 2;
  1 -- 3;
  1 -- 2;
  2 -- 3;
  2 -- 1;
  2 -- 0;
  3 -- 2;
  3 -- 1;
  4 -- 6;
  4 -- 5;
  5 -- 6;
  5 -- 4;
  6 -- 4;
  6 -- 5;
}`
	actual := g.String()
	if actual != expected {
		t.Fatal("Graphviz representation not working correctly")
		t.Fatalf("actual result: \n%s\n", actual)
	}
}

// Find connected components in an undirected graph.
// The return value of Components() is a slice of integers
// that associates each vertex with a component number.
func TestComponents(t *testing.T) {
	var g AdjList = graphSevenVertices()
	var comps []int = g.Components()
	var compOne bool = true
	var compTwo bool = true
	for i := 0; i < 4; i++ {
		if comps[i] != 0 {
			compOne = false
		}
	}
	for i := 4; i < 7; i++ {
		if comps[i] != 1 {
			compTwo = false
		}
	}
	if !compOne || !compTwo {
		t.Fatalf("%v", comps)
	}
}
