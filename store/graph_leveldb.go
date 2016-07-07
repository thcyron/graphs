package store

import (
	. "github.com/noypi/graph/store/kvgraph"
	. "github.com/noypi/graph/types"
	"github.com/noypi/kv/leveldb"
)

func NewGraphLeveldb(path string, fn VertexDeserializer, bDirected bool) (g Graph, err error) {
	store, err := leveldb.GetDefault(path)
	if nil != err {
		return
	}

	gstore, err := NewStore(store, fn)
	if nil != err {
		return
	}

	g, err = NewGraph(gstore, bDirected)
	return

}
