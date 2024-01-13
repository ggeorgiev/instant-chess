package square

import (
	"fmt"
	"math"
)

type Index int8
type Indexes []Index

const (
	InvalidIndex = Index(-1)
	ZeroIndex    = A1
	LastIndex    = H8
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

func (s Index) Min(i Index) Index {
	if i < s {
		return i
	}
	return s
}

func (s Index) Max(i Index) Index {
	if i > s {
		return i
	}
	return s
}

func (s Index) String() string {
	letters := "ABCDEFGH"
	return fmt.Sprintf("%c%d", letters[s.File()], s.Rank()+1)
}

func (i Indexes) Min() Index {
	if len(i) == 0 {
		return InvalidIndex
	}
	min := Index(math.MaxInt8)
	for _, s := range i {
		if s < min {
			min = s
		}
	}
	return min
}

func (i Indexes) Max() Index {
	max := InvalidIndex
	for _, s := range i {
		if s > max {
			max = s
		}
	}
	return max
}

func (i Indexes) MaxLess(currentMax Index) Index {
	max := InvalidIndex
	for _, s := range i {
		if s > max && s < currentMax {
			max = s
		}
	}
	return max
}

func (i Indexes) MaxBase() Index {
	max := i.Max()
	for next := i.MaxLess(max); next != InvalidIndex && next == max-1; next = i.MaxLess(max) {
		max--
	}
	return max
}
