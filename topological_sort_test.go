package graphs

import (
	"testing"
)

func TestTopologicalSort_SimpleGraph(t *testing.T) {
	graph := NewDigraph[int]()

	graph.AddEdge(1, 3, 0)
	graph.AddEdge(1, 2, 0)
	topOrder, topClasses, noDAGError := TopologicalSort(graph)
	if noDAGError != nil {
		t.Errorf("Error was thrown: %v", noDAGError)
	}
	if topOrder.Len() != 3 {
		t.Log("Topological order should have 3 items")
		t.Error("Topological order should have 3 items")
	}

	// check if the topological class assigned to each node is correct
	if len(topClasses) != 3 {
		t.Log("We don't have information about each node.")
		t.Error("We don't have information about each node.")
	}
	if topClasses[1] != 0 {
		t.Error("Node 1 has wrong class")
	}
	if topClasses[2] != 1 {
		t.Error("Node 2 has wrong class")
	}
	if topClasses[3] != 1 {
		t.Error("Node 3 has wrong class")
	}

	// check if the topological ordering is correct	e := topOrder.Front()
	e := topOrder.Front()
	if e.Value.(int) != 1 {
		t.Error("First node shuld be: 1")
	}
	e = e.Next()
	if e.Value.(int) != 2 && e.Value.(int) != 3 {
		t.Error("Second node shuld be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(int) != 2 && e.Value.(int) != 3 {
		t.Error("Third node shuld be: 2 or 3")
	}
}

func TestTopologicalSort_Case2(t *testing.T) {
	graph := NewDigraph[int]()

	graph.AddEdge(1, 2, 0)
	graph.AddEdge(1, 3, 0)
	graph.AddEdge(2, 4, 0)
	graph.AddEdge(3, 4, 0)
	topOrder, topClasses, noDAGError := TopologicalSort(graph)
	if noDAGError != nil {
		t.Errorf("Error was thrown: %v", noDAGError)
	}
	if topOrder.Len() != 4 {
		t.Log("Topological order should have 4 items")
		t.Error("Topological order should have 4 items")
	}
	// check if the topological class assigned to each node is correct
	if len(topClasses) != 4 {
		t.Log("We don't have information about each node.")
		t.Error("We don't have information about each node.")
	}
	if topClasses[1] != 0 {
		t.Error("Node 1 has wrong class")
	}
	if topClasses[2] != 1 {
		t.Error("Node 2 has wrong class")
	}
	if topClasses[3] != 1 {
		t.Error("Node 3 has wrong class")
	}
	if topClasses[4] != 2 {
		t.Error("Node 4 has wrong class")
	}

	// check if the topological ordering is correct
	e := topOrder.Front()
	if e.Value.(int) != 1 {
		t.Error("First node should be: 1")
	}
	e = e.Next()
	if e.Value.(int) != 2 && e.Value.(int) != 3 {
		t.Error("Second node should be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(int) != 2 && e.Value.(int) != 3 {
		t.Error("Third node should be: 2 or 3")
	}
	e = e.Next()
	if e.Value.(int) != 4 {
		t.Error("Fourth node should be: 4")
	}
}

func TestTopologicalSort_NotDAG(t *testing.T) {
	graph := NewDigraph[int]()

	graph.AddEdge(1, 2, 0)
	graph.AddEdge(1, 3, 0)
	graph.AddEdge(2, 1, 0)
	_, _, noDagError := TopologicalSort(graph)
	if noDagError != ErrNoDAG {
		t.Error("This graph is not a DAG. TopologicalSort should have returned an Error.")
	}
}
