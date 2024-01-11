package board

import (
	"log"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type State struct {
	Matrix Matrix
	Rights *Rights
}

func ParseState(text string) (State, error) {
	rights, err := ParseRights(text)
	if err != nil {
		return State{}, err
	}
	matrix, err := ParseMatrix(text)
	if err != nil {
		return State{}, err
	}

	state := State{
		Matrix: matrix,
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

func (s State) Sprint() string {
	return s.Matrix.Sprint() + s.Rights.Sprint()
}

func (s State) Invalid() (bool, square.Index) {
	whiteKings := 0
	blackKings := 0
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		figure := s.Matrix[i]
		if figure.IsNoFigure() {
			continue
		}

		if figure == peace.WhiteKing {
			whiteKings++
			if whiteKings > 1 {
				return true, i
			}
		} else if figure == peace.BlackKing {
			blackKings++
			if blackKings > 1 {
				return true, i
			}

			attackers := s.Matrix.DirectAttackersOfSquareFromWhite(i)
			if len(attackers) > 0 {
				offender := attackers.Min()
				if offender < i {
					return true, i
				} else {
					return true, offender
				}
			}
			attackers = s.Matrix.BlockableAttackersOfSquareFromWhite(i)
			if len(attackers) > 0 {
				offender := attackers.Min()
				if offender > i {
					return true, offender
				}
				return true, square.InvalidIndex
			}
		}
	}
	return false, square.InvalidIndex
}

func (s State) M1() bool {
	return s.Matrix.M1()
}
