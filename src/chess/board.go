package chess

import (
	"fmt"
	"strings"
)

type Board [8][8]Peace

func ParseBoard(text string) Board {
	var board Board

	rows := strings.Split(text, "\n")

	row := 0
	if len(rows[row]) == 0 {
		row++
	}

	for y := 8; y > 0; y-- {
		row += 2
		runes := runes(rows[row])

		for x := 0; x < 8; x++ {
			board[x][y-1] = PeaceFromSymbol(runes[4+x*4])
		}

	}
	return board
}

func (board Board) SquareUnderAttack(x int, y int, fromColor PeaceColor) bool {
	if x > 0 {
		if y > 0 {
			peace := board[x-1][y-1]
			if peace.Color == fromColor && peace.Type == King {
				return true
			}
		}
		peace := board[x-1][y]
		if peace.Color == fromColor && peace.Type == King {
			return true
		}
		if y < 7 {
			peace := board[x-1][y+1]
			if peace.Color == fromColor && peace.Type == King {
				return true
			}
		}
	}
	if y > 0 {
		peace := board[x][y-1]
		if peace.Color == fromColor && peace.Type == King {
			return true
		}
	}
	if y < 7 {
		peace := board[x][y+1]
		if peace.Color == fromColor && peace.Type == King {
			return true
		}
	}
	if x < 7 {
		if y > 0 {
			peace := board[x+1][y-1]
			if peace.Color == fromColor && peace.Type == King {
				return true
			}
		}
		peace := board[x+1][y]
		if peace.Color == fromColor && peace.Type == King {
			return true
		}
		if y < 7 {
			peace := board[x+1][y+1]
			if peace.Color == fromColor && peace.Type == King {
				return true
			}
		}
	}

	for i := x - 1; i >= 0; i-- {
		peace := board[i][y]
		if peace.Color == fromColor && (peace.Type == Rook || peace.Type == Queen) {
			return true
		}
		if peace.Type != Empty {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		peace := board[i][y]
		if peace.Color == fromColor && (peace.Type == Rook || peace.Type == Queen) {
			return true
		}
		if peace.Type != Empty {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		peace := board[x][i]
		if peace.Color == fromColor && (peace.Type == Rook || peace.Type == Queen) {
			return true
		}
		if peace.Type != Empty {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		peace := board[x][i]
		if peace.Color == fromColor && (peace.Type == Rook || peace.Type == Queen) {
			return true
		}
		if peace.Type != Empty {
			break
		}
	}
	return false
}

func (board Board) kingTos(x int, y int) []Square {
	var tos []Square
	color := board[x][y].Color
	oponentColor := color.Oponent()

	original := board[x][y]
	board[x][y] = EmptyPeace

	check := func(square Square) {
		peace := board[square.X][square.Y]
		if (peace.Type == Empty || peace.Color == color) && !board.SquareUnderAttack(square.X, square.Y, oponentColor) {
			tos = append(tos, square)
		}
	}

	if x > 0 {
		if y > 0 {
			check(Square{X: x - 1, Y: y - 1})
		}
		check(Square{X: x - 1, Y: y})
		if y < 7 {
			check(Square{X: x - 1, Y: y + 1})
		}
	}
	if y > 0 {
		check(Square{X: x, Y: y - 1})
	}
	if y < 7 {
		check(Square{X: x, Y: y + 1})
	}
	if x < 7 {
		if y > 0 {
			check(Square{X: x + 1, Y: y - 1})
		}
		check(Square{X: x + 1, Y: y})
		if y < 7 {
			check(Square{X: x + 1, Y: y + 1})
		}
	}

	board[x][y] = original
	return tos
}

func (board Board) Tos(x int, y int) []Square {
	if board[x][y].Type == King {
		return board.kingTos(x, y)
	}

	return nil
}

func (board Board) Move(x int, y int) Move {
	tos := board.Tos(x, y)

	var toAnswers []ToAnswer
	for _, to := range tos {
		original := board[to.X][to.Y]
		board[to.X][to.Y] = board[x][y]
		board[x][y] = EmptyPeace

		for y := 8; y > 0; y-- {
			for x := 0; x < 8; x++ {
			}
		}

		toAnswers = append(toAnswers, ToAnswer{
			WhiteTo: to,
		})

		board[x][y] = board[to.X][to.Y]
		board[to.X][to.Y] = original
	}

	return Move{
		WhiteForm: Square{
			X: x,
			Y: y,
		},
		ToAnswers: toAnswers,
	}
}

func (board Board) Print() {
	letters := "    a   b   c   d   e   f   g   h  "
	separator := "  +---+---+---+---+---+---+---+---+"

	fmt.Println(letters)
	fmt.Println(separator)

	for y := 8; y > 0; y-- {
		fmt.Printf("%d |", y)
		for x := 0; x < 8; x++ {
			fmt.Printf(" %s |", board[x][y-1].Symbol())
		}
		fmt.Printf(" %d\n", y)
		fmt.Println(separator)
	}

	fmt.Println(letters)
}
