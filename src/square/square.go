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

func NewIndex(file File, rank Rank) Index {
	return Index(int8(rank)<<3 + int8(file))
}

func (s Index) File() File {
	return File(s) & 0b111
}

func (s Index) Rank() Rank {
	return Rank(s) >> 3
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

func (i Indexes) Min() Index {
	min := i[0]
	for _, s := range i {
		if s < min {
			min = s
		}
	}
	return min
}
