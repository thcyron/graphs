package store

import (
	. "github.com/noypi/graph/types"
)

type HalfEdgeIterImpl struct {
	it EdgeSetIterator
}

func (this *HalfEdgeIterImpl) Next() {
	this.it.Next()
}

func (this *HalfEdgeIterImpl) Value() HalfEdge {
	return HalfEdge(this.it.Value())
}

func (this *HalfEdgeIterImpl) Valid() bool {
	return this.it.Valid()
}

func (this *HalfEdgeIterImpl) Close() (err error) {
	return this.it.Close()
}
