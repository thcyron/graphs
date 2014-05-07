package graphs

import "testing"

func TestGraphNVertices(t *testing.T) {
	graph := NewGraph()

	if graph.NVertices() != 0 {
		t.Error("empty graph should not have any vertices")
	}

	graph.AddEdge("a", "b", 0)
	if graph.NVertices() != 2 {
		t.Error("graph should have two vertices")
	}
}

func TestVerticesIter(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)

	vertices := 0
	for _ = range graph.VerticesIter() {
		vertices++
	}

	if vertices != graph.NVertices() {
		t.Error("wrong number of vertices")
	}
}

func TestGraphNEdges(t *testing.T) {
	graph := NewGraph()

	if graph.NEdges() != 0 {
		t.Error("empty graph should not have any edges")
	}

	graph.AddEdge("a", "b", 0)
	if graph.NEdges() != 1 {
		t.Error("graph should have one edge")
	}
}

func TestGraphEquals(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()

	if !g1.Equals(g2) {
		t.Error("two empty graphs should be equal")
	}

	g1.AddEdge("a", "b", 0)
	if g1.Equals(g2) {
		t.Error("two graphs with different number of edges should not be equal")
	}

	g2.AddEdge("a", "b", 0)
	if !g1.Equals(g2) {
		t.Error("two graphs with same edges should be equal")
	}

	g1.AddEdge("b", "c", 0)
	g2.AddEdge("a", "c", 0)
	if g1.Equals(g2) {
		t.Error("two graphs with different edges should not be equal")
	}
}

func TestGraphAdjacency(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("a", "c", 0)
	graph.AddEdge("c", "d", 0)
	graph.AddEdge("d", "a", 0)

	if !graph.AdjacentVertices("a").Equals(NewSetWithElements("b", "c", "d")) {
		t.Error("bad adjacency for a")
	}

	if !graph.AdjacentVertices("b").Equals(NewSetWithElements("a")) {
		t.Error("bad adjacency for b")
	}

	if !graph.AdjacentVertices("c").Equals(NewSetWithElements("a", "d")) {
		t.Error("bad adjacency for c")
	}

	if !graph.AdjacentVertices("d").Equals(NewSetWithElements("a", "c")) {
		t.Error("bad adjacency for d")
	}
}
