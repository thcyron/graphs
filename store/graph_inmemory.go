package store

import (
	. "github.com/noypi/graph/store/kvgraph"
	. "github.com/noypi/graph/types"
	"github.com/noypi/kv/gtreap"
)

func NewGraphInMemory(fn VertexDeserializer, bDirected bool) (g Graph, err error) {
	store := gtreap.GetDefault()

	gstore, err := NewStore(store, fn)
	if nil != err {
		return
	}

	g, err = NewGraph(gstore, bDirected)
	return

}
