package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/alignment"
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) SquareBlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsBlack() {
		return nil
	}

	if pc == peace.BlackKing {
		return m.BlackKingTos(s)
	}

	original := m[s]
	m[s] = peace.Null
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch pc {
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
		if maybeCheckedVector != alignment.NoVector {
			return nil
		}
		return m.BlackKnightNoCheckedTos(s)
	case peace.BlackQueen:
		return m.BlackQueenNoCheckedTos(s, maybeCheckedVector)
	}

	return nil
}

func (m *Matrix) SquareCaptureBlackTos(s square.Index, kingSquare square.Index, capture square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsBlack() {
		return nil
	}

	if !m.IsSquareUnderAttackFrom(capture, s) {
		return nil
	}

	if pc == peace.BlackKing {
		return square.Indexes{capture}
	}

	original := m[s]
	m[s] = peace.Null
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch pc {
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
		if maybeCheckedVector != alignment.NoVector {
			return nil
		}
		return square.Indexes{capture}
	}

	return nil
}

func (m *Matrix) SquareBlockBlackTos(s square.Index, kingSquare square.Index, attacker square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsBlack() || pc == peace.BlackKing {
		return nil
	}

	mask := alignment.BlockRelationMasksList[attacker][kingSquare]
	peaceMask := attack.PeaceBitboardMasks(pc)[s]

	overlap := mask & peaceMask
	if overlap == bitboard.Empty {
		return nil
	}

	original := m[s]
	m[s] = peace.Null
	maybeCheckedVector := m.IsBlackMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	if maybeCheckedVector != alignment.NoVector {
		return nil
	}

	return square.ConvertBitboardMaskIntoIndexes(overlap)
}

func (m *Matrix) BlackTos(king square.Index) move.Halfs {
	var moves move.Halfs
	checked, attacker, block := m.IsBlackCheckedToMoveCaptureOrBlock(king)
	if checked {
		tos := m.BlackKingTos(king)
		if len(tos) > 0 {
			moves = append(moves, move.Half{
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
					moves = append(moves, move.Half{
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
			moves = append(moves, move.Half{
				From: s,
				Tos:  tos,
			})
		}
	}

	return moves
}
