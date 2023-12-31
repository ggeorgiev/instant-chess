package square

import (
	"fmt"
)

type Index int8
type Indexes []Index

const (
	InvalidIndex = Index(-1)
	ZeroIndex    = Index(0)
	LastIndex    = Index(63)
)

func NewIndex(x, y int8) Index {
	return Index(y<<3 + x)
}

func (s Index) X() int8 {
	return int8(s) & 0b111
}

func (s Index) Y() int8 {
	return int8(s) >> 3
}

func (s Index) String() string {
	letters := "ABCDEFGH"
	return fmt.Sprintf("%c%d", letters[s.X()], s.Y()+1)
}
