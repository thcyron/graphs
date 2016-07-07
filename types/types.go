package types

type HasID interface {
	StringID() string
}

type Entity interface {
	HasID
	Serialize() ([]byte, error)
}

type EdgeSetIterator interface {
	Next()
	Valid() bool
	Key() string
	Value() Edge
	Close() error
}

type AdjIterator interface {
	Next()
	Valid() bool
	Key() Vertex
	Value() EdgeSet
	Close() error
}

type Adjacency interface {
	V(id string) Vertex
	Add(Vertex) error
	Has(Vertex) bool
	EdgeSetOf(Vertex) EdgeSet
	Iterator() AdjIterator
	VertexCount() int
	GetVertexDeserializer() VertexDeserializer
}
