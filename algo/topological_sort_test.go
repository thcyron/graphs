package algo

import (
	"testing"

	. "github.com/noypi/graph/types"
)

func TestTopologicalSort_SimpleGraph(t *testing.T) {
	graph := NewInMemoryGraphIntTest()

	graph.AddEdge(VertexInt(1), VertexInt(3), 0)
	graph.AddEdge(VertexInt(1), VertexInt(2), 0)
	topOrder, topClasses, noDAGError := TopologicalSort(graph)
	if noDAGError != nil {
		t.Error("Error was thrown: %v", noDAGError)
	}
	if topOrder.Len() != 3 {
		t.Log("Topological order should have 3 items")
		t.Fatal("Topological order should have 3 items")
	}

	// check if the topological class assigned to each node is correct
	if len(topClasses) != 3 {
		t.Log("We don't have information about each node.")
		t.Fatal("We don't have information about each node.")
	}
	if topClasses[VertexInt(1).StringID()] != 0 {
		t.Fatal("Node 1 has wrong class")
	}
	if topClasses[VertexInt(2).StringID()] != 1 {
		t.Fatal("Node 2 has wrong class")
	}
	if topClasses[VertexInt(3).StringID()] != 1 {
		t.Fatal("Node 3 has wrong class")
	}

	// check if the topological ordering is correct	e := topOrder.Front()
	e := topOrder.Front()
	if e.Value.(string) != VertexInt(1).StringID() {
		t.Fatal("First node shuld be: 1")
	}
	e = e.Next()
	if e.Value.(string) != VertexInt(2).StringID() && e.Value.(string) != VertexInt(3).StringID() {
		t.Fatal("Second node shuld be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(string) != VertexInt(2).StringID() && e.Value.(string) != VertexInt(3).StringID() {
		t.Error("Third node shuld be: 2 or 3")
	}
}

func TestTopologicalSort_Case2(t *testing.T) {
	graph := NewInMemoryGraphIntTest()

	graph.AddEdge(VertexInt(1), VertexInt(2), 0)
	graph.AddEdge(VertexInt(1), VertexInt(3), 0)
	graph.AddEdge(VertexInt(2), VertexInt(4), 0)
	graph.AddEdge(VertexInt(3), VertexInt(4), 0)
	topOrder, topClasses, noDAGError := TopologicalSort(graph)
	if noDAGError != nil {
		t.Fatal("Error was thrown: %v", noDAGError)
	}
	if topOrder.Len() != 4 {
		t.Log("Topological order should have 4 items")
		t.Fatal("Topological order should have 4 items")
	}
	// check if the topological class assigned to each node is correct
	if len(topClasses) != 4 {
		t.Log("We don't have information about each node.")
		t.Fatal("We don't have information about each node.")
	}
	if topClasses[VertexInt(1).StringID()] != 0 {
		t.Fatal("Node 1 has wrong class")
	}
	if topClasses[VertexInt(2).StringID()] != 1 {
		t.Fatal("Node 2 has wrong class")
	}
	if topClasses[VertexInt(3).StringID()] != 1 {
		t.Fatal("Node 3 has wrong class")
	}
	if topClasses[VertexInt(4).StringID()] != 2 {
		t.Fatal("Node 4 has wrong class")
	}

	// check if the topological ordering is correct
	e := topOrder.Front()
	if e.Value.(string) != VertexInt(1).StringID() {
		t.Fatal("First node should be: 1")
	}
	e = e.Next()
	if e.Value.(string) != VertexInt(2).StringID() && e.Value.(string) != VertexInt(3).StringID() {
		t.Fatal("Second node should be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(string) != VertexInt(2).StringID() && e.Value.(string) != VertexInt(3).StringID() {
		t.Fatal("Third node should be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(string) != VertexInt(4).StringID() {
		t.Fatal("Fourth node should be: 4")
	}
}

func TestTopologicalSort_NotDAG(t *testing.T) {
	graph := NewInMemoryGraphIntTest()

	graph.AddEdge(VertexInt(1), VertexInt(2), 0)
	graph.AddEdge(VertexInt(1), VertexInt(3), 0)
	graph.AddEdge(VertexInt(2), VertexInt(1), 0)
	_, _, noDagError := TopologicalSort(graph)
	if noDagError != ErrNoDAG {
		t.Error("This graph is not a DAG. TopologicalSort should have returned an Error.")
	}
}
