package types

type GraphConstructor func(Store, bDirected bool) (Graph, error)

// A Vertex can be just anything.
type Vertex interface {
	Entity
}

// An Edge connects two vertices with a cost.
type Edge interface {
	Start() string
	HalfEdge
}

// A Halfedge is an edge where just the end vertex is
// stored. The start vertex is inferred from the context.
type HalfEdge interface {
	End() string
	Cost() float64
	Entity
}

type EdgeSet interface {
	// Add adds an element to the set. It returns true if the
	// element has been added and false if the set already contained
	// that element.
	Add(Edge) bool

	AddHalf(HalfEdge) bool

	// Len returns the number of elements.
	Len() int

	// Contains returns whether the set contains the given element.
	Contains(Edge) bool

	// Remove removes the given element from the set and returns
	// whether the element was removed from the set.
	Remove(Edge) bool

	Equals(EdgeSet) bool

	Iterator() EdgeSetIterator
}

type VertexIter interface {
	Next()
	Value() Vertex
	Valid() bool
	Close() error
}

type EdgeIter interface {
	Next()
	Value() Edge
	Valid() bool
	Close() error
}

type HalfEdgeIter interface {
	Next()
	Value() HalfEdge
	Valid() bool
	Close() error
}

type Graph interface {
	AddVertex(v Vertex)

	// AddEdge adds an edge to the graph. The edge connects
	// vertex v1 and vertex v2 with cost c.
	AddEdge(v1, v2 Vertex, c float64)

	V(id string) Vertex

	IsDirected() bool

	// Dump prints all edges with their cost to stdout.
	Dump()

	// NVertices returns the number of vertices.
	NVertices() int

	// NEdges returns the number of edges.
	NEdges() int

	// Equals returns whether the graph is equal to the given graph.
	// Two graphs are equal of their adjacency is equal.
	Equals(g2 Graph) bool

	Adj() Adjacency

	VerticesIter() VertexIter

	EdgesIter() EdgeIter

	HalfedgesIter(Vertex) HalfEdgeIter

	GetVertexDeserializer() VertexDeserializer

	SortedEdges() SortedEdges
}
