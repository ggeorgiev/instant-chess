package board

import (
	"strings"
	"testing"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/stretchr/testify/assert"
)

func TestCastlingMoves(t *testing.T) {
	starting := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♜ |   |   |   | ♚ |   |   | ♜ |·8
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
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♖ |   |   |   | ♔ |   |   | ♖ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: wb, O-O-O: wb, En Passant: - ·
`

	state := MustParseState(starting)

	snapshot := state.DoWhite(peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingKingsideCastlingPlace)
	assert.Equal(t, "1·| ♖ |   |   |   |   | ♖ | ♔ |   |·1", strings.Split(state.Matrix.Sprint(), "\n")[16])
	assert.Equal(t, move.NoCastling, state.Rights.WhiteCastling)

	state.UndoWhite(snapshot, peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingKingsideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())
	assert.Equal(t, move.BothCastlings, state.Rights.WhiteCastling)

	snapshot = state.DoWhite(peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingQueensideCastlingPlace)
	assert.Equal(t, "1·|   |   | ♔ | ♖ |   |   |   | ♖ |·1", strings.Split(state.Matrix.Sprint(), "\n")[16])
	assert.Equal(t, move.NoCastling, state.Rights.WhiteCastling)

	state.UndoWhite(snapshot, peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingQueensideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())
	assert.Equal(t, move.BothCastlings, state.Rights.WhiteCastling)

	snapshot = state.DoBlack(peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingKingsideCastlingPlace)
	assert.Equal(t, "8·| ♜ |   |   |   |   | ♜ | ♚ |   |·8", strings.Split(state.Matrix.Sprint(), "\n")[2])
	assert.Equal(t, move.NoCastling, state.Rights.BlackCastling)

	state.UndoBlack(snapshot, peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingKingsideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())
	assert.Equal(t, move.BothCastlings, state.Rights.BlackCastling)

	snapshot = state.DoBlack(peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingQueensideCastlingPlace)
	assert.Equal(t, "8·|   |   | ♚ | ♜ |   |   |   | ♜ |·8", strings.Split(state.Matrix.Sprint(), "\n")[2])
	assert.Equal(t, move.NoCastling, state.Rights.BlackCastling)

	state.UndoBlack(snapshot, peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingQueensideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())
	assert.Equal(t, move.BothCastlings, state.Rights.BlackCastling)
}
