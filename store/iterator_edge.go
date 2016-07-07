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

//------------------------------------------------------

type EdgeIterImpl2orig struct {
	adjiter     AdjIterator
	currSetiter EdgeSetIterator
}

func (this *EdgeIterImpl2orig) Next() {
	if nil == this.currSetiter {
		if this.adjiter.Valid() {
			this.currSetiter = this.adjiter.Value().Iterator()
		}

	} else if !this.currSetiter.Valid() {
		this.adjiter.Next()
		if this.adjiter.Valid() {
			this.currSetiter = this.adjiter.Value().Iterator()
		}

	} else {
		this.currSetiter.Next()
		if !this.currSetiter.Valid() {
			for !this.currSetiter.Valid() && this.adjiter.Valid() {
				this.adjiter.Next()
				if this.adjiter.Valid() {
					this.currSetiter = this.adjiter.Value().Iterator()
				}
			}
		}

	}

}

func (this *EdgeIterImpl2orig) Value() Edge {
	he := HalfEdge(this.currSetiter.Value())
	e := &EdgeBase{S: this.adjiter.Key().StringID()}
	e.E = he.End()
	e.C = he.Cost()
	return e
}

func (this *EdgeIterImpl2orig) Valid() bool {
	return this.adjiter.Valid() && (nil != this.currSetiter) && this.currSetiter.Valid()
}

func (this *EdgeIterImpl2orig) Close() (err error) {
	var err1, err2 error
	if nil != this.currSetiter {
		err1 = this.currSetiter.Close()
	}
	if nil != this.adjiter {
		err2 = this.adjiter.Close()
	}
	if nil != err1 || nil != err2 {
		err = fmt.Errorf("%v. %v", err1, err2)
	}
	return
}
