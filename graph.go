package graphs

import (
	"fmt"
	"sort"
)

// A Vertex can be just anything.
type Vertex interface{}

// An Edge connects two vertices with a cost.
type Edge struct {
	Start Vertex
	End   Vertex
	Cost  float64
}

// A Halfedge is an edge where just the end vertex is
// stored. The start vertex is inferred from the context.
type Halfedge struct {
	End  Vertex
	Cost float64
}

// A Graph is defined by its vertices and edges stored as
// an adjacency set.
type Graph struct {
	Adjacency map[Vertex]*Set
	Vertices  *Set
	Directed  bool
}

// NewGraph creates a new empty graph.
func NewGraph() *Graph {
	return &Graph{
		Adjacency: map[Vertex]*Set{},
		Vertices:  NewSet(),
		Directed:  false,
	}
}

// NewDigraph creates a new empty directed graph.
func NewDigraph() *Graph {
	graph := NewGraph()
	graph.Directed = true
	return graph
}

// AddVertex adds the given vertex to the graph.
func (g *Graph) AddVertex(v Vertex) {
	g.Vertices.Add(v)
}

// AddEdge adds an edge to the graph. The edge connects
// vertex v1 and vertex v2 with cost c.
func (g *Graph) AddEdge(v1, v2 Vertex, c float64) {
	g.Vertices.Add(v1)
	g.Vertices.Add(v2)

	if _, exists := g.Adjacency[v1]; !exists {
		g.Adjacency[v1] = NewSet()
	}

	g.Adjacency[v1].Add(Halfedge{
		End:  v2,
		Cost: c,
	})

	if !g.Directed {
		if _, exists := g.Adjacency[v2]; !exists {
			g.Adjacency[v2] = NewSet()
		}

		g.Adjacency[v2].Add(Halfedge{
			End:  v1,
			Cost: c,
		})
	}
}

// Dump prints all edges with their cost to stdout.
func (g *Graph) Dump() {
	for v, s := range g.Adjacency {
		s.Each(func(e interface{}, stop *bool) {
			he := e.(Halfedge)
			fmt.Printf("(%v,%v,%f)\n", v, he.End, he.Cost)
		})
	}
}

// NVertices returns the number of vertices.
func (g *Graph) NVertices() int {
	return g.Vertices.Len()
}

// NEdges returns the number of edges.
func (g *Graph) NEdges() int {
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
	a1 := g.Adjacency
	a2 := g2.Adjacency

	for k, v := range a1 {
		if !v.Equals(a2[k]) {
			return false
		}
	}

	return true
}

// VerticesIter returns a channel where all vertices
// are sent to.
func (g *Graph) VerticesIter() chan Vertex {
	ch := make(chan Vertex)
	go func() {
		for e := range g.Vertices.Iter() {
			ch <- e.(Vertex)
		}
		close(ch)
	}()
	return ch
}

// AdjacentVertices returns a set containing all
// adjacent vertices for a given vertex.
func (g *Graph) AdjacentVertices(v Vertex) *Set {
	vertices := NewSet()
	if _, exists := g.Adjacency[v]; exists {

		g.Adjacency[v].Each(func(e interface{}, stop *bool) {
			vertices.Add(e.(Halfedge).End)
		})
	}
	return vertices
}

// SortedEdges is an array of edges that can be sorted
// by their cost.
type SortedEdges []Edge

func (se SortedEdges) Len() int {
	return len(se)
}

func (se SortedEdges) Less(i, j int) bool {
	return se[i].Cost < se[j].Cost
}

func (se SortedEdges) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

// SortedEdges returns an array of edges sorted by their cost.
func (g *Graph) SortedEdges() SortedEdges {
	set := NewSet()

	for v, s := range g.Adjacency {
		s.Each(func(e interface{}, stop *bool) {
			he := e.(Halfedge)
			set.Add(Edge{
				Start: v,
				End:   he.End,
				Cost:  he.Cost,
			})
		})
	}

	edges := make(SortedEdges, set.Len())
	set.Each(func(e interface{}, stop *bool) {
		edges = append(edges, e.(Edge))
	})

	sort.Sort(&edges)
	return edges
}
