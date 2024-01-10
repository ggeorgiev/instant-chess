package board

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) SquareBlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsBlack() {
		return nil
	}

	if figure == peace.BlackKing {
		return m.BlackKingTos(s)
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch figure {
	case peace.BlackBishop:
		if maybeCheckedVector.IsLeaniar() {
			return nil
		}
		return m.BlackBishopNoCheckedTos(s, maybeCheckedVector)
	case peace.BlackRook:
		if maybeCheckedVector.IsDiagonalic() {
			return nil
		}
		return m.BlackRookNoCheckedTos(s, maybeCheckedVector)
	case peace.BlackKnight:
		if maybeCheckedVector != peacealignment.NoVector {
			return nil
		}
		return m.BlackKnightNoCheckedTos(s)
	}

	return nil
}

func (m Matrix) SquareCaptureBlackTos(s square.Index, kingSquare square.Index, capture square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsBlack() {
		return nil
	}

	if !m.IsSquareUnderAttackFrom(capture, s) {
		return nil
	}

	if figure == peace.BlackKing {
		return square.Indexes{capture}
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch figure {
	case peace.BlackBishop:
		if maybeCheckedVector.IsLeaniar() {
			return nil
		}
		return m.BlackBishopCapture(s, maybeCheckedVector, capture)
	case peace.BlackRook:
		if maybeCheckedVector.IsDiagonalic() {
			return nil
		}
		return m.BlackRookCapture(s, maybeCheckedVector, capture)
	case peace.BlackKnight:
		if maybeCheckedVector != peacealignment.NoVector {
			return nil
		}
		return square.Indexes{capture}
	}

	return nil
}

func (m Matrix) SquareBlockBlackTos(s square.Index, kingSquare square.Index, attacker square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsBlack() || figure == peace.BlackKing {
		return nil
	}

	mask := peacealignment.BlockRelationMasksList[attacker][kingSquare]
	peaceMask := peaceattacks.PeaceBitboardMasks(figure)[s]

	overlap := mask & peaceMask
	if overlap == bitboard.Empty {
		return nil
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	if maybeCheckedVector != peacealignment.NoVector {
		return nil
	}

	return square.ConvertBitboardMaskIntoIndexes(overlap)
}

func (m Matrix) BlackTos(king square.Index) peacemoves.Halfs {
	var moves peacemoves.Halfs
	checked, attacker, block := m.IsBlackCheckedToMoveCaptureOrBlock(king)
	if checked {
		tos := m.BlackKingTos(king)
		if len(tos) > 0 {
			moves = append(moves, peacemoves.Half{
				From: king,
				Tos:  tos,
			})
		}
		if attacker != square.InvalidIndex {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				var tos square.Indexes
				if block {
					tos = append(tos, m.SquareBlockBlackTos(s, king, attacker)...)
				}
				tos = append(tos, m.SquareCaptureBlackTos(s, king, attacker)...)
				if len(tos) > 0 {
					moves = append(moves, peacemoves.Half{
						From: s,
						Tos:  tos,
					})
				}
			}
		}
		return moves
	}

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		tos := m.SquareBlackTos(s, king)
		if len(tos) > 0 {
			moves = append(moves, peacemoves.Half{
				From: s,
				Tos:  tos,
			})
		}
	}

	return moves
}
