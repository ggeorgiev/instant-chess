package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/alignment"
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) SquareWhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsWhite() {
		return nil
	}

	if pc == peace.WhiteKing {
		return m.WhiteKingTos(s)
	}

	original := m[s]
	m[s] = peace.Null
	maybeCheckedVector := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch pc {
	case peace.WhiteBishop:
		if maybeCheckedVector.IsLeaniar() {
			return nil
		}
		return m.WhiteBishopNoCheckedTos(s, maybeCheckedVector)
	case peace.WhiteRook:
		if maybeCheckedVector.IsDiagonalic() {
			return nil
		}
		return m.WhiteRookNoCheckedTos(s, maybeCheckedVector)
	case peace.WhiteKnight:
		if maybeCheckedVector != alignment.NoVector {
			return nil
		}
		return m.WhiteKnightNoCheckedTos(s)
	case peace.WhiteQueen:
		return m.WhiteQueenNoCheckedTos(s, maybeCheckedVector)
	}

	return nil
}

func (m *Matrix) SquareCaptureWhiteTos(s square.Index, kingSquare square.Index, capture square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsWhite() {
		return nil
	}

	if !m.IsSquareUnderAttackFrom(capture, s) {
		return nil
	}

	if pc == peace.WhiteKing {
		return square.Indexes{capture}
	}

	original := m[s]
	m[s] = peace.Null
	maybeCheckedVector := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch pc {
	case peace.WhiteBishop:
		if maybeCheckedVector.IsLeaniar() {
			return nil
		}
		return m.WhiteBishopCapture(s, maybeCheckedVector, capture)
	case peace.WhiteRook:
		if maybeCheckedVector.IsDiagonalic() {
			return nil
		}
		return m.WhiteRookCapture(s, maybeCheckedVector, capture)
	case peace.WhiteKnight:
		if maybeCheckedVector != alignment.NoVector {
			return nil
		}
		return square.Indexes{capture}
	}

	return nil
}

func (m *Matrix) SquareBlockWhiteTos(s square.Index, kingSquare square.Index, attacker square.Index) square.Indexes {
	pc := m[s]
	if !pc.IsWhite() || pc == peace.WhiteKing {
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
	maybeCheckedVector := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	if maybeCheckedVector != alignment.NoVector {
		return nil
	}

	return square.ConvertBitboardMaskIntoIndexes(overlap)
}

func (m *Matrix) WhiteTos(king square.Index) move.Halfs {
	var moves move.Halfs
	checked, attacker, block := m.IsWhiteCheckedToMoveCaptureOrBlock(king)
	if checked {
		tos := m.WhiteKingTos(king)
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
					tos = append(tos, m.SquareBlockWhiteTos(s, king, attacker)...)
				}
				tos = append(tos, m.SquareCaptureWhiteTos(s, king, attacker)...)
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
		tos := m.SquareWhiteTos(s, king)
		if len(tos) > 0 {
			moves = append(moves, move.Half{
				From: s,
				Tos:  tos,
			})
		}
	}

	return moves
}
