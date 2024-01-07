package chess

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestSquareUnderAttackLineOnLeft(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+····
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+····
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+····
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+····
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+····
4·|   | ♜ |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+····
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+····
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+····
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+····
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineOnLeftObstructed(t *testing.T) {
	position := ParsePosition(`
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
4·|   | ♜ | ♞ | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineOnRight(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   | ♛ |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineOnRightObstructed(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   | ♞ | ♛ |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineUnder(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   | ♜ |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineUnderObstructed(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   | ♞ |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   | ♜ |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineAbove(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   | ♛ |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackLineAboveObstructed(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   | ♛ |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   | ♞ |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalOnAboveLeft(t *testing.T) {
	position := ParsePosition(`
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
4·| ♝ |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   | ♚ |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   | ♔ |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(2, 1), 10))
}

func TestSquareUnderAttackDiagonalOnAboveLeftObstructed(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   | ♝ |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ♞ |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalOnAboveRight(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   | ♛ |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalOnAboveRightObstructed(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   | ♛ |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   | ♞ |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalUnderLeft(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   | ♚ |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♝ |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalUnderLeftObstructed(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   | ♔ |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   | ♞ |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   | ♛ |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(4, 5), 10))
}

func TestSquareUnderAttackDiagonalUnderRight(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   | ♛ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackDiagonalUnderRightObstructed(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   | ♞ |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   | ♛ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 0, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackKnight_m1m2(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ | ♞ |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackKnight_p1m2(t *testing.T) {
	position := ParsePosition(`
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
4·|   |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   | ♞ |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 1, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 10))
}

func TestSquareUnderAttackMutiple(t *testing.T) {
	position := ParsePosition(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   | ♛ | ♞ | ♛ | ♞ | ♛ |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   | ♞ |   |   |   | ♞ |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   | ♛ |   | ♔ |   | ♛ |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   | ♞ |   |   |   | ♞ |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♛ | ♞ | ♛ | ♞ | ♛ |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   | ♚ |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`)

	assert.Equal(t, 16, position.BoardState.Matrix.SquareUnderAttackFromBlack(square.NewIndex(3, 3), 20))
}
