package batch

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board/matrix"
	"github.com/ggeorgiev/instant-chess/src/board/state"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Batch struct {
	Figures peace.Figures
	Rights  move.Rights
	Matrix  *matrix.Matrix
}

type Batches []*Batch

func Create(figures peace.Figures, rights move.Rights) (*Batch, error) {
	matrix := matrix.Matrix{}

	figures = figures.Copy()

	// if white castling remove the white king since we need to keep it in its exact position
	if rights.IsWhiteAnyCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.WhiteKing)
		if err != nil {
			return nil, err
		}

		matrix[peaceplaces.WhiteKingStartingPlace] = peace.WhiteKing
	}

	if rights.IsWhiteKingsideCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.WhiteRook)
		if err != nil {
			return nil, err
		}
		matrix[peaceplaces.WhiteRookKingsideStartingPlace] = peace.WhiteRook
	}

	if rights.IsWhiteQueensideCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.WhiteRook)
		if err != nil {
			return nil, err
		}
		matrix[peaceplaces.WhiteRookQueensideStartingPlace] = peace.WhiteRook
	}

	// if black castling remove the black king since we need to keep it in its exact position
	if rights.IsBlackAnyCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.BlackKing)
		if err != nil {
			return nil, err
		}
		matrix[peaceplaces.BlackKingStartingPlace] = peace.BlackKing
	}

	if rights.IsBlackKingsideCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.BlackRook)
		if err != nil {
			return nil, err
		}
		matrix[peaceplaces.BlackRookKingsideStartingPlace] = peace.BlackRook
	}

	if rights.IsBlackQueensideCastling() {
		var err error
		figures, err = figures.RemoveOne(peace.BlackRook)
		if err != nil {
			return nil, err
		}
		matrix[peaceplaces.BlackRookQueensideStartingPlace] = peace.BlackRook
	}

	return &Batch{
		Figures: figures,
		Rights:  rights,
		Matrix:  &matrix,
	}, nil
}

func (b *Batch) CountBitsets() uint64 {
	return math.CountBitsets(uint64(len(b.Figures)))
}

func (b *Batch) GenerateState(index uint64) (*state.State, uint64, square.Index) {
	bitset := math.IndexToBitset(uint64(len(b.Figures)), index)
	indexes := square.ConvertBitboardMaskIntoIndexes(bitboard.Mask(bitset))

	// copy the starting matrix
	matrix := b.Matrix.Copy()

	// set the figures
	for p, s := range indexes {
		if matrix[s] != peace.NoFigure {
			return nil, bitset, s
		}
		matrix[s] = b.Figures[p]
	}

	return &state.State{
		Matrix: matrix,
		Rights: b.Rights,
	}, bitset, square.InvalidIndex
}

func (b *Batch) SkipOffender(i uint64, bitset uint64, offender square.Index, stats *Stats) uint64 {
	skipTo := math.NextValidIndex(uint64(len(b.Figures)), i, bitset, uint64(offender))
	skip := skipTo - i
	stats.States += skip
	stats.Invalid += skip
	stats.Skipped += skip
	return skipTo
}
