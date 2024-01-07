package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) SquareWhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsWhite() {
		return nil
	}

	if figure == peace.WhiteKing {
		return m.WhiteKingTos(s)
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeCheckedVector := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch figure {
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
		if maybeCheckedVector != peacealignment.NoVector {
			return nil
		}
		return m.WhiteKnightNoCheckedTos(s)
	}

	return nil
}

func (m Matrix) SquareCaptureWhiteTos(s square.Index, kingSquare square.Index, capture square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsWhite() {
		return nil
	}

	if !m.IsSquareUnderAttackFrom(capture, s) {
		return nil
	}

	if figure == peace.WhiteKing {
		return square.Indexes{capture}
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeCheckedVector := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	switch figure {
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
		if maybeCheckedVector != peacealignment.NoVector {
			return nil
		}
		return square.Indexes{capture}
	}

	return nil
}

func (m Matrix) WhiteTos() HalfMoves {
	var moves HalfMoves
	king := m.FindSinglePeace(peace.WhiteKing)

	checked, attacker, block := m.IsWhiteCheckedToMoveCaptureOrBlock(king)
	if checked {
		moves = append(moves, peacemoves.FromTo{
			From: king,
			Tos:  m.WhiteKingTos(king),
		})
		if attacker != square.InvalidIndex {
			if block {
				// TODO: find if the attacker can be captured or blocked
			} else {
				for s := square.ZeroIndex; s <= square.LastIndex; s++ {
					tos := m.SquareCaptureWhiteTos(s, king, attacker)
					if len(tos) > 0 {
						moves = append(moves, peacemoves.FromTo{
							From: s,
							Tos:  tos,
						})
					}
				}
			}
		}
		return moves
	}

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		tos := m.SquareWhiteTos(s, king)
		if len(tos) > 0 {
			moves = append(moves, peacemoves.FromTo{
				From: s,
				Tos:  tos,
			})
		}
	}

	return moves
}
