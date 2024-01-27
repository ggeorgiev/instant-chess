package batch

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board/matrix"
	"github.com/ggeorgiev/instant-chess/src/board/state"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Batch struct {
	Peaces peace.Codes
	Rights move.Rights
}

type Batches []*Batch

func Create(peaces peace.Codes, rights move.Rights) *Batch {
	peaces = peaces.Copy()
	for _, f := range peace.GetRightsMap(rights) {
		peaces = peaces.RemoveOne(f)
	}
	return &Batch{
		Peaces: peaces,
		Rights: rights,
	}
}

func (b *Batch) CountBitsets() uint64 {
	return math.CountBitsets(uint64(len(b.Peaces)))
}

func (b *Batch) GenerateState(index uint64) (*state.State, uint64, square.Index) {
	bitset := math.IndexToBitset(uint64(len(b.Peaces)), index)
	indexes := square.ConvertBitboardMaskIntoIndexes(bitboard.Mask(bitset))

	matrix := &matrix.Matrix{}

	// set the peaces
	for p, s := range indexes {
		matrix[s] = b.Peaces[p]
	}

	for s, f := range peace.GetRightsMap(b.Rights) {
		if matrix[s] != peace.NoFigure {
			return nil, bitset, s
		}
		matrix[s] = f
	}

	return &state.State{
		Matrix: matrix,
		Rights: b.Rights,
	}, bitset, square.InvalidIndex
}

func (b *Batch) SkipOffender(i uint64, bitset uint64, offender square.Index, stats *Stats) uint64 {
	skipTo := math.NextValidIndex(uint64(len(b.Peaces)), i, bitset, uint64(offender))
	skip := skipTo - i
	stats.States += skip
	stats.Invalid += skip
	stats.Skipped += skip
	return skipTo
}
