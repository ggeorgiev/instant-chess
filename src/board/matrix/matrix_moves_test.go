package matrix

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 1, Tos: square.Indexes{16, 18}},
		move.Half{From: 6, Tos: square.Indexes{21, 23}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesBishopDiagonalBlockingEachOther(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   | ♚ |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   | ♝ |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   | ♗ |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♔ |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 0, Tos: square.Indexes{1, 8, 9}},
		move.Half{From: 18, Tos: square.Indexes{9, 27, 36, 45}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesBishopCounterDiagonalBlockingEachOther(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♚ |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   | ♝ |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   | ♗ |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   | ♔ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 7, Tos: square.Indexes{6, 14, 15}},
		move.Half{From: 21, Tos: square.Indexes{14, 28, 35, 42}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 4, Tos: square.Indexes{3, 5, 11, 12, 13}},
		move.Half{From: 20, Tos: square.Indexes{12, 28, 36, 44}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 29, Tos: square.Indexes{28, 27, 26, 30}},
		move.Half{From: 31, Tos: square.Indexes{22, 23, 30, 38, 39}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveOrCaptureWithBishop(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♚ |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   | ♗ |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   | ♞ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   | ♔ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 7, Tos: square.Indexes{6, 14, 15}},
		move.Half{From: 31, Tos: square.Indexes{13}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveOrCaptureWithKnight(t *testing.T) {
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
4·| ♚ |   |   | ♘ |   |   |   | ♔ |·4
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 31, Tos: square.Indexes{22, 23, 30, 39}},
		move.Half{From: 27, Tos: square.Indexes{21}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveOrCaptureWithRook(t *testing.T) {
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 31, Tos: square.Indexes{22, 23, 30, 39}},
		move.Half{From: 29, Tos: square.Indexes{21}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveBlockOrCaptureWithKnight(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♚ |   |   |   |   |   |   | ♔ |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   | ♘ |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♝ |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 63, Tos: square.Indexes{55, 62}},
		move.Half{From: 10, Tos: square.Indexes{27, 0}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveButCannotCaptureWithBishop(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♚ |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   | ♜ |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   | ♗ |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   | ♞ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   | ♔ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 7, Tos: square.Indexes{6, 14, 15}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveButCannotCaptureWithKnight(t *testing.T) {
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
4·| ♚ |   | ♜ | ♘ |   |   |   | ♔ |·4
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
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 31, Tos: square.Indexes{22, 23, 30, 39}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}

func TestMatrixMovesNeedToMoveButCannotCaptureWithRook(t *testing.T) {
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
3·|   |   |   |   |   | ♞ |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`
	matrix, err := Parse(text)
	assert.NoError(t, err)

	expected := move.Halfs{
		move.Half{From: 31, Tos: square.Indexes{22, 23, 30, 39}},
	}

	king := matrix.FindSinglePeace(peace.WhiteKing)
	result := matrix.WhiteTos(king)
	assert.Equal(t, expected, result)
}
