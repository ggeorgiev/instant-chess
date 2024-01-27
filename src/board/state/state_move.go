package state

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/storage"
)

type Snapshot struct {
	Captured peace.Figure
	Rights   move.Rights
}

func (s *State) Moves() move.Fulls {
	var moves move.Fulls

	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		var answers move.Answers
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			blackFromTos := s.Matrix.BlackTos(blackKing)
			answers = append(answers, move.Answer{
				WhiteTo:      to,
				BlackAnswers: blackFromTos,
			})

			s.UndoWhite(snapshot, fromTos.From, to)
		}

		moves = append(moves, move.Full{
			WhiteForm: fromTos.From,
			Answers:   answers,
		})
	}

	return moves
}

func (s *State) M1() bool {
	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			mate := s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil

			s.UndoWhite(snapshot, fromTos.From, to)

			if mate {
				return true
			}
		}
	}

	return false
}

func (s *State) M1s() int {
	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	mates := 0

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			if s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil {
				mates++
			}

			s.UndoWhite(snapshot, fromTos.From, to)
		}
	}

	return mates
}

func (s *State) MateIn() (*storage.Data, error) {
	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, whiteFromTo := range whiteFromTos {
		for _, whiteTo := range whiteFromTo.Tos {
			whiteSnapshot := s.DoWhite(whiteFromTo.From, whiteTo)

			blackFromTos := s.Matrix.BlackTos(blackKing)

			for _, blackFromTo := range blackFromTos {
				for _, blackTo := range blackFromTo.Tos {
					blackSnapshot := s.DoBlack(blackFromTo.From, blackTo)

					figures, index := s.StorageLocation()

					readMap := storage.FetchReadMap(figures)
					if readMap == nil {
						// unknown
					}
					data := readMap.Get(index)
					if data == nil {
						// unkown
					}

					// TODO: get max

					s.UndoBlack(blackSnapshot, blackFromTo.From, blackTo)
				}
			}

			s.UndoWhite(whiteSnapshot, whiteFromTo.From, whiteTo)
		}
	}

	return nil, fmt.Errorf("it was not able to determine the mate moves at this time")
}
