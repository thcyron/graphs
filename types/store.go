package types

import (
	"github.com/noypi/kv"
)

type VertexDeserializer func([]byte) (Vertex, error)

type StoreConstructor func(kv.KVStore, fn VertexDeserializer) (Store, error)

type Store interface {
	NewAdjacency() Adjacency
	Store() kv.KVStore
}
