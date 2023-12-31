package chess

import (
	"fmt"
	"strings"
)

type Board [64]Peace

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
			board[NewSquare(x, y-1)] = PeaceFromSymbol(runes[4+x*4])
		}

	}
	return board
}

func (board Board) SquareUnderAttack(s Square, fromColor PeaceColor) bool {
	if fromColor == White {
		return board.SquareUnderAttackWhite(s)
	}
	return board.SquareUnderAttackBlack(s)
}

func (board Board) FindPeace(peace Peace) Square {
	for s := Square(0); s < 64; s++ {
		if board[s] == peace {
			return s
		}

	}
	return Square(-1)
}

func (board Board) Move(s Square) Move {
	wks := board.FindPeace(WhiteKing)
	tos := board.WhiteTos(s, wks)

	bks := board.FindPeace(BlackKing)

	var toAnswers []ToAnswer
	for _, to := range tos {
		original := board[to]
		board[to] = board[s]
		board[s] = Empty

		var answers []Answer

		for square := Square(0); square < 64; square++ {
			peace := board[square]
			if peace.IsBlack() {
				blackTos := board.BlackTos(square, bks)
				if len(blackTos) > 0 {
					answers = append(answers, Answer{
						BlackFrom: square,
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
			fmt.Printf(" %s |", board[NewSquare(x, y-1)].Symbol())
		}
		fmt.Printf(" %d\n", y)
		fmt.Println(separator)
	}

	fmt.Println(letters)
}
