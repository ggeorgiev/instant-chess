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

func TestPeaceIsEmptyOrNot(t *testing.T) {
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

	assert.True(t, Empty.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhitePawn.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhiteBishop.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhiteKnight.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhiteRook.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhiteQueen.IsEmptyOrNot(PeaceColorWhite))
	assert.False(t, WhiteKing.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackPawn.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackBishop.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackKnight.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackRook.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackQueen.IsEmptyOrNot(PeaceColorWhite))
	assert.True(t, BlackKing.IsEmptyOrNot(PeaceColorWhite))

	assert.True(t, Empty.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhitePawn.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhiteBishop.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhiteKnight.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhiteRook.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhiteQueen.IsEmptyOrNot(PeaceColorBlack))
	assert.True(t, WhiteKing.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackPawn.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackBishop.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackKnight.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackRook.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackQueen.IsEmptyOrNot(PeaceColorBlack))
	assert.False(t, BlackKing.IsEmptyOrNot(PeaceColorBlack))
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
