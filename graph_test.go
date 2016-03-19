package graph

import (
	"fmt"
	"testing"
)

func graphSixVertices() AdjList {
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
	g := graphSixVertices()
	if g.VertexCount() != 6 {
		t.Fatalf("Expected vertex count to be 6, got %d", g.VertexCount())
	}
	if g.EdgeCount() != 8 {
		t.Fatalf("Expected vertex count to be 6, got %d", g.EdgeCount())
	}
}

func TestString(t *testing.T) {
	g := graphSixVertices()
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
	fmt.Println(actual)
	if actual != expected {
		t.Fatal("Graphviz representation not working correctly")
		t.Fatalf("actual result: \n%s\n", actual)
	}
}
