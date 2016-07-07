package kvgraph

import (
	"fmt"

	. "github.com/noypi/graph/types"

	"github.com/noypi/kv"
)

type EdgeSetDb struct {
	store         kv.KVStore
	v             Vertex
	fnDeserialize VertexDeserializer
}

func EdgeSetOf(store kv.KVStore, v Vertex, fn VertexDeserializer) *EdgeSetDb {
	if nil == v {
		panic("should not be nil")
	}
	return &EdgeSetDb{store: store, v: v, fnDeserialize: fn}
}

func (this EdgeSetDb) genEdgeKey(e HasID) string {
	var s2 string
	if nil != e {
		s2 = e.StringID()
	}
	return fmt.Sprintf("\xffe\xff%s\xff%s", this.v.StringID(), s2)
}

func (this *EdgeSetDb) addEdgeEntity(e Entity) bool {
	if this.containsEdgeEntity(e) {
		return false
	}
	wrtr, _ := this.store.Writer()
	defer wrtr.Close()

	batch := wrtr.NewBatch()
	defer batch.Close()

	bb, err := e.Serialize()
	if nil != err {
		return false
	}

	batch.Set([]byte(this.genEdgeKey(e)), bb)
	err = wrtr.ExecuteBatch(batch)
	return nil == err
}

// Add adds an element to the set. It returns true if the
// element has been added and false if the set already contained
// that element.
func (this *EdgeSetDb) Add(e Edge) bool {
	return this.addEdgeEntity(e)
}

func (this *EdgeSetDb) AddHalf(e HalfEdge) bool {
	return this.addEdgeEntity(e)
}

// Len returns the number of elements.
func (this *EdgeSetDb) Len() int {
	rdr, _ := this.store.Reader()
	defer rdr.Close()

	prefix := fmt.Sprintf("\xffe\xff%s\xff", this.v.StringID())
	it := rdr.PrefixIterator([]byte(prefix))
	var n int
	for ; it.Valid(); it.Next() {
		n++
	}
	return n
}

func (this *EdgeSetDb) containsEdgeEntity(e Entity) bool {
	rdr, _ := this.store.Reader()
	defer rdr.Close()

	bb, _ := rdr.Get([]byte(this.genEdgeKey(e)))
	return 0 < len(bb)
}

// Contains returns whether the set contains the given element.
func (this *EdgeSetDb) Contains(e Edge) bool {
	return this.containsEdgeEntity(e)
}

// Remove removes the given element from the set and returns
// whether the element was removed from the set.
func (this *EdgeSetDb) Remove(e Edge) bool {
	wrtr, _ := this.store.Writer()
	defer wrtr.Close()

	batch := wrtr.NewBatch()
	defer batch.Close()

	batch.Delete([]byte(this.genEdgeKey(e)))
	err := wrtr.ExecuteBatch(batch)
	return nil != err
}

func (this *EdgeSetDb) Iterator() EdgeSetIterator {
	rdr, _ := this.store.Reader()
	kprefix := []byte(this.genEdgeKey(nil))

	return &EdgeSetIterImpl{
		it:            rdr.PrefixIterator(kprefix),
		kprefix:       kprefix,
		fnDeserialize: this.fnDeserialize,
	}

}

func (this EdgeSetDb) Equals(s2 EdgeSet) bool {
	if s2 == nil || this.Len() != s2.Len() {
		return false
	}

	it := this.Iterator()
	for ; it.Valid(); it.Next() {
		if !s2.Contains(it.Value()) {
			return false
		}
	}

	return true
}

func (this EdgeSetDb) Store() kv.KVStore {
	return this.store
}
