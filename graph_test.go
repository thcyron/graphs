package graphs

import "testing"

func TestGraphNVertices(t *testing.T) {
	graph := NewGraph[string]()

	if graph.NVertices() != 0 {
		t.Error("empty graph should not have any vertices")
	}

	graph.AddEdge("a", "b", 0)
	if graph.NVertices() != 2 {
		t.Error("graph should have two vertices")
	}
}

func TestEachVertex(t *testing.T) {
	graph := NewGraph[string]()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)

	vertices := 0
	graph.EachVertex(func(_ string, _ func()) {
		vertices++
	})

	if vertices != graph.NVertices() {
		t.Error("wrong number of vertices")
	}
}

func TestEachVertexStop(t *testing.T) {
	graph := NewGraph[string]()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)

	vertices := 0
	graph.EachVertex(func(_ string, stop func()) {
		vertices++
		stop()
	})

	if vertices != 1 {
		t.Error("wrong number of vertices")
	}
}

func TestEachEdge(t *testing.T) {
	graph := NewGraph[string]()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)

	edges := 0
	graph.EachEdge(func(_ Edge[string], _ func()) {
		edges++
	})

	if edges != graph.NEdges()*2 {
		t.Errorf("wrong number of edges: %d", edges)
	}
}

func TestEachEdgeStop(t *testing.T) {
	graph := NewGraph[string]()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)

	edges := 0
	graph.EachEdge(func(_ Edge[string], stop func()) {
		edges++
		stop()
	})

	if edges != 1 {
		t.Errorf("wrong number of edges: %d", edges)
	}
}

func TestGraphNEdges(t *testing.T) {
	graph := NewGraph[string]()

	if graph.NEdges() != 0 {
		t.Error("empty graph should not have any edges")
	}

	graph.AddEdge("a", "b", 0)
	if graph.NEdges() != 1 {
		t.Error("graph should have one edge")
	}
}

func TestGraphEquals(t *testing.T) {
	g1 := NewGraph[string]()
	g2 := NewGraph[string]()

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
