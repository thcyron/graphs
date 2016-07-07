package types

type VertexString string

func (this VertexString) StringID() string {
	return string(this)
}

func (this VertexString) Serialize() ([]byte, error) {
	return []byte(this), nil
}
