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

func (board Board) FindWhiteKing() (int, int) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if board[x][y].Color == PeaceColorWhite && board[x][y].Type == King {
				return x, y
			}
		}
	}
	return -1, -1
}

func (board Board) FindBlackKing() (int, int) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if board[x][y].Color == PeaceColorBlack && board[x][y].Type == King {
				return x, y
			}
		}
	}
	return -1, -1
}

func (board Board) kingTos(x int, y int) []Square {
	var tos []Square
	color := board[x][y].Color
	oponentColor := color.Oponent()

	original := board[x][y]
	board[x][y] = EmptyPeace

	check := func(square Square) {
		peace := board[square.X][square.Y]
		if (peace.Type == Empty || peace.Color == oponentColor) && !board.SquareUnderAttack(square.X, square.Y, oponentColor) {
			tos = append(tos, square)
		}
	}

	if y < 7 {
		if x > 0 {
			check(Square{X: x - 1, Y: y + 1})
		}
		check(Square{X: x, Y: y + 1})
		if x < 7 {
			check(Square{X: x + 1, Y: y + 1})
		}
	}
	if x > 0 {
		check(Square{X: x - 1, Y: y})
	}
	if x < 7 {
		check(Square{X: x + 1, Y: y})
	}

	if y > 0 {
		if x > 0 {
			check(Square{X: x - 1, Y: y - 1})
		}
		check(Square{X: x, Y: y - 1})
		if x < 7 {
			check(Square{X: x + 1, Y: y - 1})
		}
	}

	board[x][y] = original
	return tos
}

func (board Board) rookTos(x int, y int, kX int, kY int) []Square {
	var tos []Square
	color := board[x][y].Color
	oponentColor := color.Oponent()

	check := func(square Square) bool {
		peace := board[square.X][square.Y]
		if peace.Type == Empty || peace.Color == oponentColor {
			original := board[square.X][square.Y]
			board[square.X][square.Y] = board[x][y]
			board[x][y] = EmptyPeace

			if !board.SquareUnderAttack(kX, kY, oponentColor) {
				tos = append(tos, square)
			}

			board[x][y] = board[square.X][square.Y]
			board[square.X][square.Y] = original
		}
		return peace.Type == Empty
	}

	square := Square{}

	square.Y = y
	for square.X = x - 1; square.X >= 0; square.X-- {
		if !check(square) {
			break
		}
	}
	for square.X = x + 1; square.X < 8; square.X++ {
		if !check(square) {
			break
		}
	}

	square.X = x
	for square.Y = y - 1; square.Y >= 0; square.Y-- {
		if !check(square) {
			break
		}
	}
	for square.Y = y + 1; square.Y < 8; square.Y++ {
		if !check(square) {
			break
		}
	}
	return tos
}

func (board Board) Tos(x int, y int, bkX int, bkY int) []Square {
	if board[x][y].Type == King {
		return board.kingTos(x, y)
	}
	if board[x][y].Type == Rook {
		return board.rookTos(x, y, bkX, bkY)
	}

	return nil
}

func (board Board) Move(x int, y int) Move {
	wkX, wkY := board.FindWhiteKing()

	tos := board.Tos(x, y, wkX, wkY)

	bkX, bkY := board.FindBlackKing()

	var toAnswers []ToAnswer
	for _, to := range tos {
		original := board[to.X][to.Y]
		board[to.X][to.Y] = board[x][y]
		board[x][y] = EmptyPeace

		var answers []Answer

		for x := 0; x < 8; x++ {
			for y := 0; y < 8; y++ {
				peace := board[x][y]
				if peace.Type != Empty && peace.Color == PeaceColorBlack {
					blackTos := board.Tos(x, y, bkX, bkY)
					if len(blackTos) > 0 {
						answers = append(answers, Answer{
							BlackFrom: Square{X: x, Y: y},
							BlackTos:  blackTos,
						})
					}
				}
			}
		}
		toAnswers = append(toAnswers, ToAnswer{
			WhiteTo: to,
			Answers: answers,
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
