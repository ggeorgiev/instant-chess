package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardMove(t *testing.T) {
	position := ParsePosition(`
    a   b   c   d   e   f   g   h
  +---+---+---+---+---+---+---+---+
8 |   |   |   |   |   |   | ♔ |   | 8
  +---+---+---+---+---+---+---+---+
7 |   |   |   |   |   |   |   |   | 7
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   |   | ♚ |   |   | 6
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
  +---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   | 4
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   | 3
  +---+---+---+---+---+---+---+---+
2 |   |   |   |   |   |   |   |   | 2
  +---+---+---+---+---+---+---+---+
1 |   |   |   |   |   |   |   |   | 1
  +---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

	assert.Equal(t, []Move{
		{
			WhiteForm: Square{X: 6, Y: 7},
			ToAnswers: []ToAnswer{
				{
					WhiteTo: Square{X: 5, Y: 7},
				},
				{
					WhiteTo: Square{X: 7, Y: 6},
				},
				{
					WhiteTo: Square{X: 7, Y: 7},
				},
			},
		},
	}, position.Moves)
}
