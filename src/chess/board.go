package chess

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Board [square.Number]peace.Figure

func ParseBoard(text string) Board {
	var board Board

	rows := strings.Split(text, "\n")

	row := 0
	if len(rows[row]) == 0 {
		row++
	}

	for y := int8(8); y > 0; y-- {
		row += 2
		runes := runes(rows[row])

		for x := int8(0); x < 8; x++ {
			board[square.NewIndex(x, y-1)] = peace.FromSymbol(runes[4+x*4])
		}

	}
	return board
}

func (board Board) FindPeace(peace peace.Figure) square.Index {
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		if board[s] == peace {
			return s
		}

	}
	return square.InvalidIndex
}

func (board Board) Move(s square.Index) Move {
	wks := board.FindPeace(peace.WhiteKing)
	tos := board.WhiteTos(s, wks)

	bks := board.FindPeace(peace.BlackKing)

	var toAnswers []ToAnswer
	for _, to := range tos {
		original := board[to]
		board[to] = board[s]
		board[s] = peace.NoFigure

		var answers []Answer

		for s := square.ZeroIndex; s <= square.LastIndex; s++ {
			peace := board[s]
			if peace.IsBlack() {
				blackTos := board.BlackTos(s, bks)
				if len(blackTos) > 0 {
					answers = append(answers, Answer{
						BlackFrom: s,
						BlackTos:  blackTos,
					})
				}
			}
		}
		toAnswers = append(toAnswers, ToAnswer{
			WhiteTo: to,
			Answers: answers,
		})

		board[s] = board[to]
		board[to] = original
	}

	return Move{
		WhiteForm: s,
		ToAnswers: toAnswers,
	}
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
