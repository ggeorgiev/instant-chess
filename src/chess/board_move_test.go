package chess

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func init() {
	spew.Config.DisableCapacities = true
}

func TestBoardMoveKings(t *testing.T) {
	position := ParsePosition(`
    a   b   c   d   e   f   g   h
---+---+---+---+---+---+---+---+
8 |   |   |   |   |   |   | ♔ |   | 8
---+---+---+---+---+---+---+---+
7 |   |   |   |   |   |   |   |   | 7
---+---+---+---+---+---+---+---+
6 |   |   |   |   |   | ♚ |   |   | 6
---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   | 4
---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   | 3
---+---+---+---+---+---+---+---+
2 |   |   |   |   |   |   |   |   | 2
---+---+---+---+---+---+---+---+
1 |   |   |   |   |   |   |   |   | 1
---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

	actial := "\n" + position.Moves.String()

	assert.Equal(t, `
White: F7:E7
  Answers: BlackFrom: E5, BlackTos: [D5, F5, D4, E4, F4]
White: F7:G7
  Answers: BlackFrom: E5, BlackTos: [D6, E6, D5, F5, D4, E4, F4]
White: F7:G6
  Answers: BlackFrom: E5, BlackTos: [D6, E6, D5, D4, E4, F4]
`, actial)
}

func TestBoardMoveRooks(t *testing.T) {
	position := ParsePosition(`
    a   b   c   d   e   f   g   h
---+---+---+---+---+---+---+---+
8 |   |   |   |   |   |   | ♔ | ♖ | 8
---+---+---+---+---+---+---+---+
7 |   |   |   |   |   |   |   |   | 7
---+---+---+---+---+---+---+---+
6 |   |   |   |   |   | ♚ |   |   | 6
---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   | 4
---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   | 3
---+---+---+---+---+---+---+---+
2 |   |   |   |   |   |   |   |   | 2
---+---+---+---+---+---+---+---+
1 |   |   |   |   |   |   |   |   | 1
---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

	actial := "\n" + position.Moves.String()

	assert.Equal(t, `
White: F7:E7
  Answers: BlackFrom: E5, BlackTos: [D5, F5, D4, E4, F4]
White: F7:G6
  Answers: BlackFrom: E5, BlackTos: [D6, E6, D5, D4, E4, F4]

White: G7:G6
  Answers: BlackFrom: E5, BlackTos: [D5, F5, D4, E4, F4]
White: G7:G5
  Answers: BlackFrom: E5, BlackTos: [D6, D4, E4, F4]
White: G7:G4
  Answers: BlackFrom: E5, BlackTos: [D6, D5, F5]
White: G7:G3
  Answers: BlackFrom: E5, BlackTos: [D6, D5, F5, D4, E4, F4]
White: G7:G2
  Answers: BlackFrom: E5, BlackTos: [D6, D5, F5, D4, E4, F4]
White: G7:G1
  Answers: BlackFrom: E5, BlackTos: [D6, D5, F5, D4, E4, F4]
White: G7:Invalid Square
  Answers: BlackFrom: E5, BlackTos: [D6, D5, F5, D4, E4, F4]
`, actial)

}
