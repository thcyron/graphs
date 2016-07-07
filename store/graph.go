package store

import (
	"fmt"
	"sort"

	"github.com/noypi/kv"

	. "github.com/noypi/graph/types"
)

// A Graph is defined by its vertices and edges stored as
// an adjacency set.
type GraphBase struct {
	adjacency Adjacency
	directed  bool
	gstore    Store
}

func NewGraph(store Store, bDirected bool) (g Graph, err error) {
	g = &GraphBase{
		adjacency: store.NewAdjacency(),
		directed:  bDirected,
		gstore:    store,
	}

	return
}

func (g GraphBase) Adj() Adjacency {
	return g.adjacency
}

func (g GraphBase) V(id string) Vertex {
	return g.adjacency.V(id)
}

// AddVertex adds the given vertex to the graph.
func (g *GraphBase) AddVertex(v Vertex) {
	if !g.adjacency.Has(v) {
		g.adjacency.Add(v)
	}
}

// AddEdge adds an edge to the graph. The edge connects
// vertex v1 and vertex v2 with cost c.
func (g *GraphBase) AddEdge(v1, v2 Vertex, c float64) {
	g.AddVertex(v1)
	g.AddVertex(v2)

	set := g.adjacency.EdgeSetOf(v1)
	if nil != set {
		set.AddHalf(HalfedgeBase{
			E: v2.StringID(),
			C: c,
		})
	}

	if !g.directed {
		if set := g.adjacency.EdgeSetOf(v2); nil != set {
			set.AddHalf(HalfedgeBase{
				E: v1.StringID(),
				C: c,
			})
		}
	}
}

// Dump prints all edges with their cost to stdout.
func (g *GraphBase) Dump() {
	fmt.Println("(+)-------- Dump EdgeIter()")
	it := g.EdgesIter()
	for ; it.Valid(); it.Next() {
		e := it.Value()
		fmt.Printf("(%v,%v,%f)\n", e.Start(), e.End(), e.Cost())
	}
	fmt.Println("(-)-------- Dump EdgeIter()")

	fmt.Println("(+)-------- Raw Dump")
	rdr, _ := g.gstore.Store().Reader()
	itkv := rdr.PrefixIterator([]byte(""))
	for ; itkv.Valid(); itkv.Next() {
		fmt.Println("k=>", string(itkv.Key()), "... v=>", string(itkv.Value()))
	}
	fmt.Println("(-)-------- Raw Dump")

}

// NVertices returns the number of vertices.
func (g *GraphBase) NVertices() int {
	return g.adjacency.VertexCount()
}

// NEdges returns the number of edges.
func (g *GraphBase) NEdges() int {
	n := 0

	rdr, _ := g.gstore.Store().Reader()
	it := rdr.PrefixIterator([]byte("\xffe\xff"))

	//it := g.adjacency.Iterator()
	for ; it.Valid(); it.Next() {
		//n += it.Value().Len()
		n++
	}

	// Don’t count a-b and b-a edges for undirected graphs
	// as two separate edges.
	if !g.directed {
		n /= 2
	}

	return n
}

// Equals returns whether the graph is equal to the given graph.
// Two graphs are equal of their adjacency is equal.
func (g *GraphBase) Equals(g2 Graph) bool {
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
	a1 := g.adjacency
	a2 := g2.Adj()

	it := a1.Iterator()
	for ; it.Valid(); it.Next() {
		if !it.Value().Equals(a2.EdgeSetOf(it.Key())) {
			return false
		}
	}

	return true
}

// VerticesIter returns a channel where all vertices
// are sent to.
func (g *GraphBase) VerticesIter() VertexIter {
	return &VertexIterImpl{
		it: g.adjacency.Iterator(),
	}
}

// HalfedgesIter returns a channel with all halfedges for
// the given start vertex.
func (g *GraphBase) HalfedgesIter(v Vertex) (it HalfEdgeIter) {
	set := g.adjacency.EdgeSetOf(v)
	if nil != set {
		it = &HalfEdgeIterImpl{
			it: set.Iterator(),
		}
	}

	return
}

// SortedEdges returns an array of edges sorted by their cost.
func (g *GraphBase) SortedEdges() SortedEdges {

	it := g.EdgesIter()
	edges := SortedEdges{}
	for ; it.Valid(); it.Next() {
		edges = append(edges, it.Value())
	}

	sort.Sort(&edges)
	return edges
}

// EdgesIter returns a channel with all edges of the graph.
func (g *GraphBase) EdgesIter() EdgeIter {
	rdr, _ := g.gstore.Store().Reader()
	it := &EdgeIterImpl{
		it: rdr.PrefixIterator([]byte("\xffe\xff")),
	}
	return it
}

func (g GraphBase) GetVertexDeserializer() VertexDeserializer {
	return g.adjacency.GetVertexDeserializer()
}

func (g GraphBase) IsDirected() bool {
	return g.directed
}

func (g GraphBase) Store() kv.KVStore {
	return g.gstore.Store()
}
