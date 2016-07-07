package kvgraph

import (
	"encoding/json"
	"fmt"

	. "github.com/noypi/graph/types"
	"github.com/noypi/kv"
)

type AdjImpl struct {
	store         kv.KVStore
	fnDeserialize VertexDeserializer
}

type AdjImplEl struct {
	V   Vertex
	Set EdgeSet
}

type AdjImplIter struct {
	kv.KVIterator
	currV         Vertex
	store         kv.KVStore
	fnDeserialize VertexDeserializer
}

func genDbKey(v Vertex) string {
	return fmt.Sprintf("\xffv\xff%s", v.StringID())
}

func (this AdjImpl) V(id string) (v Vertex) {
	rdr, _ := this.store.Reader()
	defer rdr.Close()

	vid := fmt.Sprintf("\xffv\xff%s", id)

	bb, _ := rdr.Get([]byte(vid))
	if 0 < len(bb) {
		v, _ = this.fnDeserialize(bb)
	}
	return
}

func (this AdjImpl) Add(v Vertex) (err error) {
	wrtr, _ := this.store.Writer()
	defer wrtr.Close()

	batch := wrtr.NewBatch()
	defer batch.Close()

	bb, err := v.Serialize()
	if nil != err {
		return
	}
	batch.Set([]byte(genDbKey(v)), bb)
	err = wrtr.ExecuteBatch(batch)
	return
}

func (this AdjImpl) Has(v Vertex) bool {
	rdr, _ := this.store.Reader()
	defer rdr.Close()

	bb, _ := rdr.Get([]byte(genDbKey(v)))
	return 0 < len(bb)
}

func (this AdjImpl) Iterator() AdjIterator {
	rdr, _ := this.store.Reader()
	defer rdr.Close()

	it := &AdjImplIter{
		KVIterator:    rdr.PrefixIterator([]byte("\xffv\xff")),
		store:         this.store,
		fnDeserialize: this.fnDeserialize,
	}

	if it.Valid() {
		e, _ := this.fnDeserialize(it.KVIterator.Value())
		it.currV = e.(Vertex)
	}

	return it
}

func (this AdjImpl) EdgeSetOf(v Vertex) EdgeSet {
	return EdgeSetOf(this.store, v, this.fnDeserialize)
}

func (this AdjImpl) VertexCount() int {
	var n int
	it := this.Iterator()
	for ; it.Valid(); it.Next() {
		n++
	}

	return n
}

func (this AdjImpl) GetVertexDeserializer() VertexDeserializer {
	return this.fnDeserialize
}

func (this *AdjImplIter) Valid() bool {
	return this.KVIterator.Valid()
}

func (this *AdjImplIter) Next() {
	this.KVIterator.Next()
	if this.KVIterator.Valid() {
		e, _ := this.fnDeserialize(this.KVIterator.Value())
		this.currV = e.(Vertex)
	} else {
		this.currV = nil
	}

}

func (this *AdjImplIter) Key() Vertex {
	return this.currV
}

func (this *AdjImplIter) Value() EdgeSet {
	if nil == this.currV && this.Valid() {
		panic("got nil currV, but have a valid() iterator")
	}
	return EdgeSetOf(this.store, this.currV, this.fnDeserialize)
}

func (this *AdjImplEl) Value() interface{} {
	return this.V
}

func (this *AdjImplEl) StringID() string {
	return this.V.StringID()
}

func (this *AdjImplEl) Serialize() []byte {
	bb, _ := json.Marshal(this)
	return bb
}
