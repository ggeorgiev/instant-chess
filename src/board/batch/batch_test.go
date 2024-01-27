package batch

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/stretchr/testify/assert"
)

func TestBatchCreateNoRights(t *testing.T) {
	batch := Create(peace.MustParseFigures("♚♔♖♖"), move.NoRights)
	assert.Equal(t, uint64(635376), batch.CountBitsets())
}

func TestBatchCreateWhiteBothCastlingRights(t *testing.T) {
	batch := Create(peace.MustParseFigures("♚♔♖♖"), move.WhiteBothCastlingRights)
	assert.Equal(t, uint64(64), batch.CountBitsets())
}

func TestBatchCreateWhiteKingsideCastlingRights(t *testing.T) {
	batch := Create(peace.MustParseFigures("♚♔♖♖"), move.WhiteKingsideCastlingRights)
	assert.Equal(t, uint64(2016), batch.CountBitsets())
}

func TestBatchCreateBlackBothCastlingRights(t *testing.T) {
	batch := Create(peace.MustParseFigures("♚♔♜♜"), move.BlackBothCastlingRights)
	assert.Equal(t, uint64(64), batch.CountBitsets())
}
