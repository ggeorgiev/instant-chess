package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardMoveKings(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   | ♔ |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   | ♚ |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	actial := "\n" + position.BoardState.Matrix.Moves().String()

	assert.Equal(t, `
White: G8:H7
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, E7, F7]
White: G8:F8
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6]
White: G8:H8
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6, E7, F7]
`, actial)
}

func TestBoardMoveRooks(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   | ♔ | ♖ |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   | ♚ |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	actial := "\n" + position.BoardState.Matrix.Moves().String()

	assert.Equal(t, `
White: G8:H7
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, E7, F7]
White: G8:F8
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6]

White: H8:H7
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6]
White: H8:H6
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E7]
White: H8:H5
  Answers: BlackFrom: F6, BlackTos: [E6, G6, E7]
White: H8:H4
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6, E7]
White: H8:H3
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6, E7]
White: H8:H2
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6, E7]
White: H8:H1
  Answers: BlackFrom: F6, BlackTos: [E5, F5, G5, E6, G6, E7]
`, actial)

}
