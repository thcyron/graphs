package graphs

import "testing"

func TestDFS(t *testing.T) {
	graph := NewDigraph[int]()

	graph.AddEdge(1, 2, 0)
	graph.AddEdge(2, 3, 0)
	graph.AddEdge(3, 4, 0)
	graph.AddEdge(1, 5, 0)
	graph.AddEdge(5, 6, 0)
	graph.AddEdge(6, 3, 0)
	graph.AddEdge(1, 7, 0)

	walks := 0
	DFS(graph, 1, func(v int, stop *bool) {
		walks++
	})

	if walks != 7 {
		t.Errorf("should visit 7 vertices; visited %d", walks)
	}

	visited := make(map[int]bool)
	DFS(graph, 1, func(v int, stop *bool) {
		visited[v] = true
		if v == 5 {
			*stop = true
		}
	})
	if visited6 := visited[6]; visited6 {
		t.Errorf("visited vertex 6 vertices, but should not")
	}
}
