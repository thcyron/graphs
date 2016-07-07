package store

import (
	"fmt"
	"strings"

	. "github.com/noypi/graph/types"
	"github.com/noypi/kv"
)

type EdgeIterImpl struct {
	it kv.KVIterator
}

func (this *EdgeIterImpl) Next() {
	this.it.Next()
}

func (this *EdgeIterImpl) Value() Edge {
	edge, err := DeserializeEdge(this.it.Value())
	if nil != err {
		panic(err.Error())
	}
	sk := string(this.it.Key()[len("\xffe\xff"):])
	ssk := strings.SplitN(sk, "\xff", 2)
	if 2 != len(ssk) {
		fmt.Println("ssk=", ssk, ";raw=>", []byte(ssk[0]))
		panic("unexpected size")
	}
	edge.(*EdgeBase).S = ssk[0]
	return edge
}

func (this *EdgeIterImpl) Valid() bool {
	return this.it.Valid()
}

func (this *EdgeIterImpl) Close() (err error) {
	return this.it.Close()
}
