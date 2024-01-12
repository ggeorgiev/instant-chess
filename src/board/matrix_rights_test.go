package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRightsCastling(t *testing.T) {
	matrix := MustParseMatrix(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♜ | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ | ♜ |·8
··+---+---+---+---+---+---+---+---+··
7·| ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·| ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ |·2
··+---+---+---+---+---+---+---+---+··
1·| ♖ | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ | ♖ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`)

	assert.Len(t, matrix.MoveRights(), 16)
}

func TestMoveRightsEnPassand(t *testing.T) {
	matrix := MustParseMatrix(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ |   |·7
··+---+---+---+---+---+---+---+---+··
6·| ♜ |   |   |   |   |   |   | ♜ |·6
··+---+---+---+---+---+---+---+---+··
5·| ♟︎ | ♙ |   |   |   |   | ♙ | ♟︎ |·5
··+---+---+---+---+---+---+---+---+··
4·| ♙ |   |   |   |   |   |   | ♙ |·4
··+---+---+---+---+---+---+---+---+··
3·| ♖ |   |   |   |   |   |   | ♖ |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   | ♙ | ♙ | ♙ | ♙ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`)

	assert.Len(t, matrix.MoveRights(), 3)
}

func TestMoveRightsEnPassandEdges(t *testing.T) {
	matrix := MustParseMatrix(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ |   |·8
··+---+---+---+---+---+---+---+---+··
7·| ♜ |   | ♟︎ | ♟︎ | ♟︎ | ♟︎ |   | ♜ |·7
··+---+---+---+---+---+---+---+---+··
6·| ♟︎ |   |   |   |   |   |   | ♟︎ |·6
··+---+---+---+---+---+---+---+---+··
5·| ♙ | ♟︎ |   |   |   |   | ♟︎ | ♙ |·5
··+---+---+---+---+---+---+---+---+··
4·|   | ♙ |   |   |   |   | ♙ |   |·4
··+---+---+---+---+---+---+---+---+··
3·| ♖ |   |   |   |   |   |   | ♖ |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   | ♙ | ♙ | ♙ | ♙ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`)

	assert.Len(t, matrix.MoveRights(), 3)
}
