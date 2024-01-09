package board

import (
	"log"

	"github.com/ggeorgiev/instant-chess/src/peace"
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

func (s State) M1() bool {
	moves := s.Matrix.Moves()
	for _, move := range moves {
		for _, toAnswer := range move.Answers {
			if len(toAnswer.BlackAnswers) == 0 {
				original := s.Matrix[toAnswer.WhiteTo]
				s.Matrix[toAnswer.WhiteTo] = s.Matrix[move.WhiteForm]
				s.Matrix[move.WhiteForm] = peace.NoFigure

				bk := s.Matrix.FindSinglePeace(peace.BlackKing)
				mate := s.Matrix.IsSquareUnderAttackFromWhite(bk)

				s.Matrix[move.WhiteForm] = s.Matrix[toAnswer.WhiteTo]
				s.Matrix[toAnswer.WhiteTo] = original
				if mate {
					return true
				}
			}
		}
	}
	return false
}
