package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeaceColor(t *testing.T) {
	assert.Equal(t, WhitePawn.Color(), PeaceColorWhite)
	assert.Equal(t, WhiteBishop.Color(), PeaceColorWhite)
	assert.Equal(t, WhiteKnight.Color(), PeaceColorWhite)
	assert.Equal(t, WhiteRook.Color(), PeaceColorWhite)
	assert.Equal(t, WhiteQueen.Color(), PeaceColorWhite)
	assert.Equal(t, WhiteKing.Color(), PeaceColorWhite)

	assert.Equal(t, BlackPawn.Color(), PeaceColorBlack)
	assert.Equal(t, BlackBishop.Color(), PeaceColorBlack)
	assert.Equal(t, BlackKnight.Color(), PeaceColorBlack)
	assert.Equal(t, BlackRook.Color(), PeaceColorBlack)
	assert.Equal(t, BlackQueen.Color(), PeaceColorBlack)
	assert.Equal(t, BlackKing.Color(), PeaceColorBlack)
}

func TestPeaceColorOponent(t *testing.T) {
	assert.Equal(t, PeaceColorWhite.Oponent(), PeaceColorBlack)
	assert.Equal(t, PeaceColorBlack.Oponent(), PeaceColorWhite)
}

func TestPeaceIsWhite(t *testing.T) {
	assert.False(t, Empty.IsWhite())
	assert.True(t, WhitePawn.IsWhite())
	assert.True(t, WhiteBishop.IsWhite())
	assert.True(t, WhiteKnight.IsWhite())
	assert.True(t, WhiteRook.IsWhite())
	assert.True(t, WhiteQueen.IsWhite())
	assert.True(t, WhiteKing.IsWhite())
	assert.False(t, BlackPawn.IsWhite())
	assert.False(t, BlackBishop.IsWhite())
	assert.False(t, BlackKnight.IsWhite())
	assert.False(t, BlackRook.IsWhite())
	assert.False(t, BlackQueen.IsWhite())
	assert.False(t, BlackKing.IsWhite())
}

func TestPeaceIsBlack(t *testing.T) {
	assert.False(t, Empty.IsBlack())
	assert.False(t, WhitePawn.IsBlack())
	assert.False(t, WhiteBishop.IsBlack())
	assert.False(t, WhiteKnight.IsBlack())
	assert.False(t, WhiteRook.IsBlack())
	assert.False(t, WhiteQueen.IsBlack())
	assert.False(t, WhiteKing.IsBlack())
	assert.True(t, BlackPawn.IsBlack())
	assert.True(t, BlackBishop.IsBlack())
	assert.True(t, BlackKnight.IsBlack())
	assert.True(t, BlackRook.IsBlack())
	assert.True(t, BlackQueen.IsBlack())
	assert.True(t, BlackKing.IsBlack())
}

func TestPeaceIsEmptyOr(t *testing.T) {
	assert.True(t, Empty.IsEmptyOrWhite())
	assert.True(t, WhitePawn.IsEmptyOrWhite())
	assert.True(t, WhiteBishop.IsEmptyOrWhite())
	assert.True(t, WhiteKnight.IsEmptyOrWhite())
	assert.True(t, WhiteRook.IsEmptyOrWhite())
	assert.True(t, WhiteQueen.IsEmptyOrWhite())
	assert.True(t, WhiteKing.IsEmptyOrWhite())
	assert.False(t, BlackPawn.IsEmptyOrWhite())
	assert.False(t, BlackBishop.IsEmptyOrWhite())
	assert.False(t, BlackKnight.IsEmptyOrWhite())
	assert.False(t, BlackRook.IsEmptyOrWhite())
	assert.False(t, BlackQueen.IsEmptyOrWhite())
	assert.False(t, BlackKing.IsEmptyOrWhite())

	assert.True(t, Empty.IsEmptyOrBlack())
	assert.False(t, WhitePawn.IsEmptyOrBlack())
	assert.False(t, WhiteBishop.IsEmptyOrBlack())
	assert.False(t, WhiteKnight.IsEmptyOrBlack())
	assert.False(t, WhiteRook.IsEmptyOrBlack())
	assert.False(t, WhiteQueen.IsEmptyOrBlack())
	assert.False(t, WhiteKing.IsEmptyOrBlack())
	assert.True(t, BlackPawn.IsEmptyOrBlack())
	assert.True(t, BlackBishop.IsEmptyOrBlack())
	assert.True(t, BlackKnight.IsEmptyOrBlack())
	assert.True(t, BlackRook.IsEmptyOrBlack())
	assert.True(t, BlackQueen.IsEmptyOrBlack())
	assert.True(t, BlackKing.IsEmptyOrBlack())

	assert.True(t, Empty.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhitePawn.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhiteBishop.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhiteKnight.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhiteRook.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhiteQueen.IsEmptyOr(PeaceColorWhite))
	assert.True(t, WhiteKing.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackPawn.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackBishop.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackKnight.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackRook.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackQueen.IsEmptyOr(PeaceColorWhite))
	assert.False(t, BlackKing.IsEmptyOr(PeaceColorWhite))

	assert.True(t, Empty.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhitePawn.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhiteBishop.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhiteKnight.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhiteRook.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhiteQueen.IsEmptyOr(PeaceColorBlack))
	assert.False(t, WhiteKing.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackPawn.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackBishop.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackKnight.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackRook.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackQueen.IsEmptyOr(PeaceColorBlack))
	assert.True(t, BlackKing.IsEmptyOr(PeaceColorBlack))
}

func TestPeaceIsKing(t *testing.T) {
	assert.False(t, Empty.IsKing())
	assert.False(t, WhitePawn.IsKing())
	assert.False(t, WhiteBishop.IsKing())
	assert.False(t, WhiteKnight.IsKing())
	assert.False(t, WhiteRook.IsKing())
	assert.False(t, WhiteQueen.IsKing())
	assert.True(t, WhiteKing.IsKing())
	assert.False(t, BlackPawn.IsKing())
	assert.False(t, BlackBishop.IsKing())
	assert.False(t, BlackKnight.IsKing())
	assert.False(t, BlackRook.IsKing())
	assert.False(t, BlackQueen.IsKing())
	assert.True(t, BlackKing.IsKing())
}

func TestPeaceIsRook(t *testing.T) {
	assert.False(t, Empty.IsRook())
	assert.False(t, WhitePawn.IsRook())
	assert.False(t, WhiteBishop.IsRook())
	assert.False(t, WhiteKnight.IsRook())
	assert.True(t, WhiteRook.IsRook())
	assert.False(t, WhiteQueen.IsRook())
	assert.False(t, WhiteKing.IsRook())
	assert.False(t, BlackPawn.IsRook())
	assert.False(t, BlackBishop.IsRook())
	assert.False(t, BlackKnight.IsRook())
	assert.True(t, BlackRook.IsRook())
	assert.False(t, BlackQueen.IsRook())
	assert.False(t, BlackKing.IsRook())
}
