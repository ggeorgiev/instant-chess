package chess

import (
	"fmt"
)

type Square int8

func NewSquare(x, y int8) Square {
	return Square(y<<3 + x)
}

func (s Square) X() int8 {
	return int8(s) & 0b111
}

func (s Square) Y() int8 {
	return int8(s) >> 3
}

func (s Square) String() string {
	letters := "ABCDEFGH"
	return fmt.Sprintf("%c%d", letters[s.X()], s.Y()+1)
}
