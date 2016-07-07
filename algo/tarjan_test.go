package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestTarjan(t *testing.T) {
	graph := NewInMemoryGraphTest()

	graph.AddEdge(VertexString("a"), VertexString("b"), 0)
	graph.AddEdge(VertexString("b"), VertexString("c"), 0)
	graph.AddEdge(VertexString("b"), VertexString("e"), 0)
	graph.AddEdge(VertexString("b"), VertexString("f"), 0)
	graph.AddEdge(VertexString("c"), VertexString("d"), 0)
	graph.AddEdge(VertexString("d"), VertexString("h"), 0)
	graph.AddEdge(VertexString("d"), VertexString("i"), 0)
	graph.AddEdge(VertexString("d"), VertexString("c"), 0)
	graph.AddEdge(VertexString("e"), VertexString("a"), 0)
	graph.AddEdge(VertexString("f"), VertexString("g"), 0)
	graph.AddEdge(VertexString("g"), VertexString("f"), 0)
	graph.AddEdge(VertexString("h"), VertexString("d"), 0)
	graph.AddEdge(VertexString("h"), VertexString("g"), 0)
	graph.AddEdge(VertexString("i"), VertexString("c"), 0)
	graph.AddEdge(VertexString("i"), VertexString("k"), 0)

	ccList := TarjanStrongCC(graph)

	if ccList.Len() != 4 {
		t.Errorf("should return 4 strongly connected components; got %d", ccList.Len())
	}
}
