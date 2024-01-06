package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Board [square.Number]peace.Figure

func (board Board) FindPeace(peace peace.Figure) square.Index {
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		if board[s] == peace {
			return s
		}

	}
	return square.InvalidIndex
}

func (board Board) Moves() Moves {
	var moves Moves

	wks := board.FindPeace(peace.WhiteKing)
	bks := board.FindPeace(peace.BlackKing)

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		tos := board.WhiteTos(s, wks)
		if len(tos) == 0 {
			continue
		}

		var toAnswers []ToAnswer
		for _, to := range tos {
			original := board[to]
			board[to] = board[s]
			board[s] = peace.NoFigure

			var answers []Answer

			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				if !board[s].IsBlack() {
					continue
				}
				blackTos := board.BlackTos(s, bks)
				if len(blackTos) > 0 {
					answers = append(answers, Answer{
						BlackFrom: s,
						BlackTos:  blackTos,
					})
				}
			}
			toAnswers = append(toAnswers, ToAnswer{
				WhiteTo: to,
				Answers: answers,
			})

			board[s] = board[to]
			board[to] = original
		}

		moves = append(moves, Move{WhiteForm: s, ToAnswers: toAnswers})
	}

	return moves
}

func (board Board) Print() {
	letters := "    a   b   c   d   e   f   g   h  "
	separator := "  +---+---+---+---+---+---+---+---+"

	fmt.Println(letters)
	fmt.Println(separator)

	for y := int8(8); y > 0; y-- {
		fmt.Printf("%d |", y)
		for x := int8(0); x < 8; x++ {
			fmt.Printf(" %s |", board[square.NewIndex(x, y-1)].Symbol())
		}
		fmt.Printf(" %d\n", y)
		fmt.Println(separator)
	}

	fmt.Println(letters)
}
