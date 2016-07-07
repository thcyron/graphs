package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestDFS(t *testing.T) {
	graph := NewInMemoryGraphIntTest()

	graph.AddEdge(VertexInt(1), VertexInt(2), 0)
	graph.AddEdge(VertexInt(2), VertexInt(3), 0)
	graph.AddEdge(VertexInt(3), VertexInt(4), 0)
	graph.AddEdge(VertexInt(1), VertexInt(5), 0)
	graph.AddEdge(VertexInt(5), VertexInt(6), 0)
	graph.AddEdge(VertexInt(6), VertexInt(3), 0)
	graph.AddEdge(VertexInt(1), VertexInt(7), 0)

	walks := 0
	DFS(graph, VertexInt(1), func(v Vertex, stop *bool) {
		walks++
	})

	if walks != 7 {
		t.Errorf("should visit 7 vertices; visited %d", walks)
	}

	visited := make(map[Vertex]bool)
	DFS(graph, VertexInt(1), func(v Vertex, stop *bool) {
		visited[v] = true
		if int(v.(VertexInt)) == 5 {
			*stop = true
		}
	})
	if visited6 := visited[VertexInt(6)]; visited6 {
		t.Errorf("visited vertex 6 vertices, but should not")
	}
}
