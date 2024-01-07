package board

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestMatrixMovesStartPosition(t *testing.T) {
	text := `
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
· O-O: --, O-O-O: --, En Passant: - ·
`
	state, err := ParseState(text)
	assert.NoError(t, err)

	expected := HalfMoves{
		peacemoves.FromTo{From: 1, Tos: square.Indexes{16, 18}},
		peacemoves.FromTo{From: 6, Tos: square.Indexes{21, 23}},
	}

	result := state.Matrix.WhiteTos()
	assert.Equal(t, expected, result)
}

func TestMatrixMovesRooksFileBlockingEachOther(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   | ♚ |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   | ♜ |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   | ♖ |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   | ♔ |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	state, err := ParseState(text)
	assert.NoError(t, err)

	expected := HalfMoves{
		peacemoves.FromTo{From: 4, Tos: square.Indexes{3, 5, 11, 12, 13}},
		peacemoves.FromTo{From: 20, Tos: square.Indexes{12, 28, 36, 44}},
	}

	result := state.Matrix.WhiteTos()
	assert.Equal(t, expected, result)
}

func TestMatrixMovesRooksRankBlockingEachOther(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·| ♚ |   | ♜ |   |   | ♖ |   | ♔ |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	state, err := ParseState(text)
	assert.NoError(t, err)

	expected := HalfMoves{
		peacemoves.FromTo{From: 29, Tos: square.Indexes{28, 27, 26, 30}},
		peacemoves.FromTo{From: 31, Tos: square.Indexes{22, 23, 30, 38, 39}},
	}

	result := state.Matrix.WhiteTos()
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveOrCapture(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·| ♚ |   |   |   |   | ♖ |   | ♔ |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   | ♞ |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	state, err := ParseState(text)
	assert.NoError(t, err)

	expected := HalfMoves{
		peacemoves.FromTo{From: 29, Tos: square.Indexes{28, 27, 26, 30}},
		peacemoves.FromTo{From: 31, Tos: square.Indexes{22, 23, 30, 38, 39}},
	}

	result := state.Matrix.WhiteTos()
	assert.NotEqual(t, expected, result)
}
