package store

import (
	. "github.com/noypi/graph/types"
)

type VertexIterImpl struct {
	it AdjIterator
}

func (this *VertexIterImpl) Next() {
	this.it.Next()
}

func (this *VertexIterImpl) Value() Vertex {
	return this.it.Key()
}

func (this *VertexIterImpl) Valid() bool {
	return this.it.Valid()
}

func (this *VertexIterImpl) Close() error {
	return this.it.Close()
}
