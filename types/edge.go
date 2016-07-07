package types

import (
	"encoding/json"
	"fmt"
)

// A Halfedge is an edge where just the end vertex is
// stored. The start vertex is inferred from the context.
type HalfedgeBase struct {
	E string //id
	C float64
}

// An Edge connects two vertices with a cost.
type EdgeBase struct {
	S string //id
	HalfedgeBase
}

func NewEdge(s, e string, c float64) Edge {
	edge := new(EdgeBase)
	edge.S = s
	edge.E = e
	edge.C = c
	return edge
}

func (this EdgeBase) Start() string {
	return this.S
}

func (this EdgeBase) StringID() string {
	return fmt.Sprintf("!%s!%s!%f", this.S, this.E, this.C)
}

func (this HalfedgeBase) End() string {
	return this.E
}

func (this HalfedgeBase) Cost() float64 {
	return this.C
}

func (this HalfedgeBase) StringID() string {
	return fmt.Sprintf("!%s!%s!%f", "", this.E, this.C)
}

func (this HalfedgeBase) Serialize() ([]byte, error) {
	return json.Marshal(&this)
}

func DeserializeEdge(bb []byte) (Edge, error) {
	e := new(EdgeBase)
	err := json.Unmarshal(bb, e)
	return e, err
}

// SortedEdges is an array of edges that can be sorted
// by their cost.
type SortedEdges []Edge

func (se SortedEdges) Len() int {
	return len(se)
}

func (se SortedEdges) Less(i, j int) bool {
	return se[i].Cost() < se[j].Cost()
}

func (se SortedEdges) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}
