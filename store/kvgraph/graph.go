package kvgraph

import (
	. "github.com/noypi/graph/types"
	"github.com/noypi/kv"
)

type GKVStore struct {
	store         kv.KVStore
	fnDeserialize VertexDeserializer
}

func NewStore(kvstore kv.KVStore, fn VertexDeserializer) (gstore Store, err error) {
	gstore = &GKVStore{
		store:         kvstore,
		fnDeserialize: fn,
	}
	return
}

func (this GKVStore) NewAdjacency() Adjacency {
	return &AdjImpl{
		store:         this.store,
		fnDeserialize: this.fnDeserialize,
	}
}

func (this GKVStore) Store() kv.KVStore {
	return this.store
}
