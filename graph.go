package graphs

import (
	"container/list"
	"fmt"
)

// A Vertex can be just anything.
type Vertex interface{}

// An Edge connects two vertices with a cost.
type Edge struct {
	Start Vertex
	End   Vertex
	Cost  float64
}

// A Graph is a set of vertices and a list of edges. The list
// of edges is sorted by their cost.
type Graph struct {
	Vertices *Set
	Edges    *list.List
}

// NewGraph creates a new empty graph.
func NewGraph() *Graph {
	return &Graph{
		Vertices: NewSet(),
		Edges:    list.New(),
	}
}

// AddEdge adds an edge to the graph. The edge connects
// vertex v1 and vertex v2 with cost c.
func (g *Graph) AddEdge(v1, v2 Vertex, c float64) {
	edge := &Edge{Start: v1, End: v2, Cost: c}

	g.Vertices.Add(v1)
	g.Vertices.Add(v2)

	// Just append the edge to the list if there
	// aren’t any edges yet.
	if g.Edges.Len() == 0 {
		g.Edges.PushBack(edge)
		return
	}

	// If there are already edges, insert the new
	// edge at the right position into the array
	// to keep it sorted by cost.
	for e := g.Edges.Front(); e != nil; e = e.Next() {
		ee := e.Value.(*Edge)
		if ee.Cost >= edge.Cost {
			g.Edges.InsertBefore(edge, e)
			return
		}
	}

	// The edge to be added is the one with the largest
	// cost so far, so just append it at the end.
	g.Edges.PushBack(edge)
}

// Dump prints all edges with their cost on stdout.
func (g *Graph) Dump() {
	for e := g.Edges.Front(); e != nil; e = e.Next() {
		ee := e.Value.(*Edge)
		fmt.Printf("(%v,%v,%f)\n", ee.Start, ee.End, ee.Cost)
	}
}

// NVertices returns the number of vertices.
func (g *Graph) NVertices() int {
	return g.Vertices.Len()
}

// NEdges returns the number of edges.
func (g *Graph) NEdges() int {
	return g.Edges.Len()
}

// Equals returns whether the graph is equal to the given graph.
// Two graphs are equal of their adjacency is equal.
func (g *Graph) Equals(g2 *Graph) bool {
	// Two graphs with differnet vertices aren’t equal.
	if !g.Vertices.Equals(g2.Vertices) {
		return false
	}

	// Some for number of edges.
	if g.NEdges() != g2.NEdges() {
		return false
	}

	// Check if the adjacency for each vertex is equal
	// for both graphs.
	a1 := g.Adjacency()
	a2 := g2.Adjacency()

	for k, v := range a1 {
		if !v.Equals(a2[k]) {
			return false
		}
	}

	return true
}

// Adjacency returns a map with the adjacent vertices
// of each vertex.
func (g *Graph) Adjacency() map[Vertex]*Set {
	adjacency := map[Vertex]*Set{}

	for v, _ := range *g.Vertices {
		adjacency[v] = NewSet()
	}

	for e := g.Edges.Front(); e != nil; e = e.Next() {
		edge := e.Value.(*Edge)
		adjacency[edge.Start].Add(edge.End)
		adjacency[edge.End].Add(edge.Start)
	}

	return adjacency
}
