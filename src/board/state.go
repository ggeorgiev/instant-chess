package board

import (
	"log"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type State struct {
	Matrix *Matrix
	Rights *move.Rights
}

func ParseState(text string) (State, error) {
	rights, err := move.ParseRights(text)
	if err != nil {
		return State{}, err
	}
	matrix, err := ParseMatrix(text)
	if err != nil {
		return State{}, err
	}

	state := State{
		Matrix: &matrix,
		Rights: rights,
	}
	return state, nil
}

func MustParseState(text string) State {
	state, err := ParseState(text)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return state
}

func (s *State) Sprint() string {
	return s.Matrix.Sprint() + s.Rights.Sprint()
}

func (s *State) Invalid() (bool, square.Index) {
	invalid, offender := s.Matrix.Invalid()
	if invalid {
		return true, offender
	}

	whiteRookIsMisplaced :=
		(s.Rights.WhiteCastling.Kingside && s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] != peace.WhiteRook) ||
			(s.Rights.WhiteCastling.Queenside && s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] != peace.WhiteRook)

	blackRookIsMisplaced :=
		(s.Rights.BlackCastling.Kingside && s.Matrix[peaceplaces.BlackRookKingsideStartingPlace] != peace.BlackRook) ||
			(s.Rights.BlackCastling.Queenside && s.Matrix[peaceplaces.BlackRookQueensideStartingPlace] != peace.BlackRook)

	var offenders square.Indexes
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		figure := s.Matrix[i]
		if figure.IsNoFigure() {
			continue
		}
		if figure == peace.WhiteRook {
			if whiteRookIsMisplaced {
				if i == peaceplaces.WhiteRookKingsideStartingPlace && s.Rights.WhiteCastling.Kingside {
					continue
				}
				if i == peaceplaces.WhiteRookQueensideStartingPlace && s.Rights.WhiteCastling.Queenside {
					continue
				}
				return true, i
			}
		} else if figure == peace.BlackRook {
			if blackRookIsMisplaced {
				if i == peaceplaces.BlackRookKingsideStartingPlace && s.Rights.BlackCastling.Kingside {
					continue
				}
				if i == peaceplaces.BlackRookQueensideStartingPlace && s.Rights.BlackCastling.Queenside {
					continue
				}
				return true, i
			}
		} else if figure == peace.WhiteKing {
			if i != peaceplaces.WhiteKingStartingPlace &&
				(s.Rights.WhiteCastling.Kingside || s.Rights.WhiteCastling.Queenside) {
				offenders = append(offenders, i)
			}
		} else if figure == peace.BlackKing {
			if i != peaceplaces.BlackKingStartingPlace &&
				(s.Rights.BlackCastling.Kingside || s.Rights.BlackCastling.Queenside) {
				offenders = append(offenders, i)
			}
		}
	}
	return len(offenders) > 0, offenders.Max()
}
