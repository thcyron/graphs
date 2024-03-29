package graphs

import (
	"fmt"
	"sort"
)

// A Vertex can be just anything.
type Vertex interface {
	comparable
}

// An Edge connects two vertices with a cost.
type Edge[T Vertex] struct {
	Start T
	End   T
	Cost  float64
}

// A Halfedge is an edge where just the end vertex is
// stored. The start vertex is inferred from the context.
type Halfedge[T Vertex] struct {
	End  T
	Cost float64
}

// A Graph is defined by its vertices and edges stored as
// an adjacency set.
type Graph[T Vertex] struct {
	Adjacency map[T]*Set[Halfedge[T]]
	Directed  bool
}

// NewGraph creates a new empty graph.
func NewGraph[T Vertex]() *Graph[T] {
	return &Graph[T]{
		Adjacency: map[T]*Set[Halfedge[T]]{},
		Directed:  false,
	}
}

// NewDigraph creates a new empty directed graph.
func NewDigraph[T Vertex]() *Graph[T] {
	graph := NewGraph[T]()
	graph.Directed = true
	return graph
}

// AddVertex adds the given vertex to the graph.
func (g *Graph[T]) AddVertex(v T) {
	if _, exists := g.Adjacency[v]; !exists {
		g.Adjacency[v] = NewSet[Halfedge[T]]()
	}
}

// AddEdge adds an edge to the graph. The edge connects
// vertex v1 and vertex v2 with cost c.
func (g *Graph[T]) AddEdge(v1, v2 T, c float64) {
	g.AddVertex(v1)
	g.AddVertex(v2)

	g.Adjacency[v1].Add(Halfedge[T]{
		End:  v2,
		Cost: c,
	})

	if !g.Directed {
		g.Adjacency[v2].Add(Halfedge[T]{
			End:  v1,
			Cost: c,
		})
	}
}

// Dump prints all edges with their cost to stdout.
func (g *Graph[T]) Dump() {
	g.EachEdge(func(e Edge[T], _ func()) {
		fmt.Printf("(%v,%v,%f)\n", e.Start, e.End, e.Cost)
	})
}

// NVertices returns the number of vertices.
func (g *Graph[T]) NVertices() int {
	return len(g.Adjacency)
}

// NEdges returns the number of edges.
func (g *Graph[T]) NEdges() int {
	n := 0

	for _, v := range g.Adjacency {
		n += v.Len()
	}

	// Don’t count a-b and b-a edges for undirected graphs
	// as two separate edges.
	if !g.Directed {
		n /= 2
	}

	return n
}

// Equals returns whether the graph is equal to the given graph.
// Two graphs are equal of their adjacency is equal.
func (g *Graph[T]) Equals(g2 *Graph[T]) bool {
	// Two graphs with different number of vertices aren’t equal.
	if g.NVertices() != g2.NVertices() {
		return false
	}

	// Some for number of edges.
	if g.NEdges() != g2.NEdges() {
		return false
	}

	// Check if the adjacency for each vertex is equal
	// for both graphs.
	a1 := g.Adjacency
	a2 := g2.Adjacency

	for k, v := range a1 {
		if !v.Equals(a2[k]) {
			return false
		}
	}

	return true
}

// EachVertex calls f for every vertex.
func (g *Graph[T]) EachVertex(f func(T, func())) {
	var stopped bool
	stop := func() { stopped = true }
	for k, _ := range g.Adjacency {
		f(k, stop)
		if stopped {
			break
		}
	}
}

// SortedEdges is an array of edges that can be sorted
// by their cost.
type SortedEdges[T Vertex] []Edge[T]

func (se SortedEdges[T]) Len() int {
	return len(se)
}

func (se SortedEdges[T]) Less(i, j int) bool {
	return se[i].Cost < se[j].Cost
}

func (se SortedEdges[T]) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

// SortedEdges returns an array of edges sorted by their cost.
func (g *Graph[T]) SortedEdges() SortedEdges[T] {
	set := NewSet[Edge[T]]()

	for v := range g.Adjacency {
		g.EachHalfedge(v, func(he Halfedge[T], _ func()) {
			set.Add(Edge[T]{
				Start: v,
				End:   he.End,
				Cost:  he.Cost,
			})
		})
	}

	edges := make(SortedEdges[T], set.Len())
	set.Each(func(e Edge[T], _ func()) {
		edges = append(edges, e)
	})

	sort.Sort(&edges)
	return edges
}

// EachEdge calls f for every edge.
func (g *Graph[T]) EachEdge(f func(Edge[T], func())) {
	var stopped bool
	stop := func() { stopped = true }
	for v, s := range g.Adjacency {
		s.Each(func(he Halfedge[T], innerStop func()) {
			edge := Edge[T]{v, he.End, he.Cost}
			f(edge, stop)
			if stopped {
				innerStop()
			}
		})
		if stopped {
			break
		}
	}
}

// EachHalfedge calls f for every halfedge for
// the given start vertex.
func (g *Graph[T]) EachHalfedge(v T, f func(Halfedge[T], func())) {
	if s, exists := g.Adjacency[v]; exists {
		var stopped bool
		stop := func() { stopped = true }
		s.Each(func(he Halfedge[T], innerStop func()) {
			f(he, stop)
			if stopped {
				innerStop()
			}
		})
	}
}
