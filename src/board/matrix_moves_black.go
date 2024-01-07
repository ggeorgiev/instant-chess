package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
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

func (m Matrix) BlackTos() HalfMoves {
	var moves HalfMoves
	king := m.FindSinglePeace(peace.BlackKing)

	checked, attacker, block := m.IsBlackCheckedToMoveCaptureOrBlock(king)
	if checked {
		moves = append(moves, peacemoves.FromTo{
			From: king,
			Tos:  m.BlackKingTos(king),
		})
		if attacker != square.InvalidIndex {
			if block {
				// TODO: find if the attacker can be captured or blocked
			} else {
				// TODO: find if the attacker can be captured
			}
		}
		return moves
	}

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		tos := m.SquareBlackTos(s, king)
		if len(tos) > 0 {
			moves = append(moves, peacemoves.FromTo{
				From: s,
				Tos:  tos,
			})
		}
	}

	return moves
}
