package kvgraph

import (
	"bytes"

	. "github.com/noypi/graph/types"
	"github.com/noypi/kv"
)

type EdgeSetIterImpl struct {
	it            kv.KVIterator
	kprefix       []byte
	fnDeserialize VertexDeserializer
}

func (this *EdgeSetIterImpl) Next() {
	this.it.Next()
}
func (this *EdgeSetIterImpl) Valid() bool {
	return this.it.Valid()
}
func (this *EdgeSetIterImpl) Key() string {
	return string(bytes.Replace(this.it.Key(), this.kprefix, []byte{}, 1))
}
func (this *EdgeSetIterImpl) Value() Edge {
	e, _ := DeserializeEdge(this.it.Value())
	return e
}
func (this *EdgeSetIterImpl) Close() error {
	return this.it.Close()
}
