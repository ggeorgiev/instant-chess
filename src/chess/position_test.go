package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePosition(t *testing.T) {
	position := CreatePosition([8][8]Peace{})

	position.Print()

	assert.True(t, position.VerticalySymetric)
}

func TestParsePosition(t *testing.T) {
	position := ParsePosition(`
    a   b   c   d   e   f   g   h
  +---+---+---+---+---+---+---+---+
8 | ♜ | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ | ♜ | 8
  +---+---+---+---+---+---+---+---+
7 | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | 7
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   |   |   |   |   | 6
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
  +---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   | 4
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   | 3
  +---+---+---+---+---+---+---+---+
2 | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | 2
  +---+---+---+---+---+---+---+---+
1 | ♖ | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ | ♖ | 1
  +---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

	assert.True(t, position.VerticalySymetric)
}
