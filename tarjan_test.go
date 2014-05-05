package graphs

import "testing"

func TestTarjan(t *testing.T) {
	graph := NewDigraph()

	graph.AddEdge("a", "b", 0)
	graph.AddEdge("b", "c", 0)
	graph.AddEdge("b", "e", 0)
	graph.AddEdge("b", "f", 0)
	graph.AddEdge("c", "d", 0)
	graph.AddEdge("d", "h", 0)
	graph.AddEdge("d", "i", 0)
	graph.AddEdge("d", "c", 0)
	graph.AddEdge("e", "a", 0)
	graph.AddEdge("f", "g", 0)
	graph.AddEdge("g", "f", 0)
	graph.AddEdge("h", "d", 0)
	graph.AddEdge("h", "g", 0)
	graph.AddEdge("i", "c", 0)
	graph.AddEdge("i", "k", 0)

	ccList := TarjanStrongCC(graph)

	if ccList.Len() != 4 {
		t.Errorf("should return 4 strongly connected components; got %d", ccList.Len())
	}
}
