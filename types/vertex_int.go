package types

import (
	"encoding/binary"
	"strconv"
)

type VertexInt int

func (this VertexInt) StringID() string {
	return strconv.Itoa(int(this))
}

func (this VertexInt) Serialize() ([]byte, error) {
	bb := make([]byte, 32)
	binary.LittleEndian.PutUint32(bb, uint32(this))
	return bb, nil
}
