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
	Number       = 64
)

func NewIndex(file, rank int8) Index {
	return Index(rank<<3 + file)
}

func (s Index) File() int8 {
	return int8(s) & 0b111
}

func (s Index) Rank() int8 {
	return int8(s) >> 3
}

func (s Index) Diagonal() Diagonal {
	return IndexDiagonals[s]
}

func (s Index) CounterDiagonal() CounterDiagonal {
	return IndexCounterDiagonals[s]
}

func (s Index) String() string {
	letters := "ABCDEFGH"
	return fmt.Sprintf("%c%d", letters[s.File()], s.Rank()+1)
}
