package graphs

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := NewGraph()

	graph.AddEdge("a", "b", 1)
	graph.AddEdge("a", "c", 3)
	graph.AddEdge("b", "g", 5)
	graph.AddEdge("c", "g", 8)
	graph.AddEdge("g", "h", 6)
	graph.AddEdge("c", "d", 2)
	graph.AddEdge("g", "f", 4)
	graph.AddEdge("d", "f", 3)
	graph.AddEdge("d", "e", 5)

	path := Dijkstra(graph, "a", "e")
	result := []Vertex{"a", "c", "d", "e"}

	i := 0
	for e := path.Front(); e != nil; e = e.Next() {
		if e.Value != result[i] {
			t.Errorf("bad vertex in path at index %d", i)
		}
		i++
	}
}
