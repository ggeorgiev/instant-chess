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
		(s.Rights.WhiteKingsideCastlingRight && s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] != peace.WhiteRook) ||
			(s.Rights.WhiteQueensideCastlingRight && s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] != peace.WhiteRook)

	blackRookIsMisplaced :=
		(s.Rights.BlackKingsideCastlingRight && s.Matrix[peaceplaces.BlackRookKingsideStartingPlace] != peace.BlackRook) ||
			(s.Rights.BlackQueensideCastlingRight && s.Matrix[peaceplaces.BlackRookQueensideStartingPlace] != peace.BlackRook)

	var offenders square.Indexes
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		figure := s.Matrix[i]
		if figure.IsNoFigure() {
			continue
		}
		if figure == peace.WhiteRook {
			if whiteRookIsMisplaced {
				if i == peaceplaces.WhiteRookKingsideStartingPlace && s.Rights.WhiteKingsideCastlingRight {
					continue
				}
				if i == peaceplaces.WhiteRookQueensideStartingPlace && s.Rights.WhiteQueensideCastlingRight {
					continue
				}
				return true, i
			}
		} else if figure == peace.BlackRook {
			if blackRookIsMisplaced {
				if i == peaceplaces.BlackRookKingsideStartingPlace && s.Rights.BlackKingsideCastlingRight {
					continue
				}
				if i == peaceplaces.BlackRookQueensideStartingPlace && s.Rights.BlackQueensideCastlingRight {
					continue
				}
				return true, i
			}
		} else if figure == peace.WhiteKing {
			if i != peaceplaces.WhiteKingStartingPlace &&
				(s.Rights.WhiteKingsideCastlingRight || s.Rights.WhiteQueensideCastlingRight) {
				offenders = append(offenders, i)
			}
		} else if figure == peace.BlackKing {
			if i != peaceplaces.BlackKingStartingPlace &&
				(s.Rights.BlackKingsideCastlingRight || s.Rights.BlackQueensideCastlingRight) {
				offenders = append(offenders, i)
			}
		}
	}
	return len(offenders) > 0, offenders.Max()
}
