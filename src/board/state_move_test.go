package board

import (
	"strings"
	"testing"

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
· O-O: --, O-O-O: --, En Passant: - ·
`

	state := MustParseState(starting)

	snapshot := state.Do(peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingKingsideCastlingPlace)
	assert.Equal(t, "1·| ♖ |   |   |   |   | ♖ | ♔ |   |·1", strings.Split(state.Matrix.Sprint(), "\n")[16])

	state.Undo(snapshot, peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingKingsideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())

	snapshot = state.Do(peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingQueensideCastlingPlace)
	assert.Equal(t, "1·|   |   | ♔ | ♖ |   |   |   | ♖ |·1", strings.Split(state.Matrix.Sprint(), "\n")[16])

	state.Undo(snapshot, peaceplaces.WhiteKingStartingPlace, peaceplaces.WhiteKingQueensideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())

	snapshot = state.Do(peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingKingsideCastlingPlace)
	assert.Equal(t, "8·| ♜ |   |   |   |   | ♜ | ♚ |   |·8", strings.Split(state.Matrix.Sprint(), "\n")[2])

	state.Undo(snapshot, peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingKingsideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())

	snapshot = state.Do(peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingQueensideCastlingPlace)
	assert.Equal(t, "8·|   |   | ♚ | ♜ |   |   |   | ♜ |·8", strings.Split(state.Matrix.Sprint(), "\n")[2])

	state.Undo(snapshot, peaceplaces.BlackKingStartingPlace, peaceplaces.BlackKingQueensideCastlingPlace)
	assert.Equal(t, starting, "\n"+state.Sprint())
}
