package algo

import (
	"encoding/binary"
	"testing"

	. "github.com/noypi/graph/store"
	. "github.com/noypi/graph/types"
)

func NewInMemoryGraphIntTest() Graph {
	fn := func(bb []byte) (Vertex, error) {
		return VertexInt(int(binary.LittleEndian.Uint32(bb))), nil
	}

	g, _ := NewGraphInMemory(fn, true)

	/*path, _ := ioutil.TempDir(os.TempDir(), "test-graph")
	g, err := NewGraphLeveldb(path, fn, true)
	if nil != err {
		panic(err.Error())
	}*/

	return g
}

func TestBFS(t *testing.T) {
	graph := NewInMemoryGraphIntTest()

	graph.AddEdge(VertexInt(1), VertexInt(3), 0)
	graph.AddEdge(VertexInt(1), VertexInt(2), 0)
	graph.AddEdge(VertexInt(3), VertexInt(8), 0)
	graph.AddEdge(VertexInt(2), VertexInt(12), 0)
	graph.AddEdge(VertexInt(12), VertexInt(13), 0)
	graph.AddEdge(VertexInt(13), VertexInt(14), 0)

	var result Vertex
	walks := 0

	BFS(graph, VertexInt(1), func(v Vertex, stop *bool) {
		walks++
		if i := int(v.(VertexInt)); i > 10 && i%2 != 0 {
			result = v
			*stop = true
		}
	})

	if int(result.(VertexInt)) != 13 {
		t.Error("bad result")
	}

	if walks != 6 {
		t.Error("should visit 6 vertices")
	}
}
