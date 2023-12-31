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
	king := Combine(fromColor, King)

	attackedFromKing := AttackedFromKing[s]

	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == king {
			return true
		}
	}

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		peace := board[NewSquare(i, y)]
		if peace.IsLinearMoverFrom(fromColor) {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		peace := board[NewSquare(i, y)]
		if peace.IsLinearMoverFrom(fromColor) {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		peace := board[NewSquare(x, i)]
		if peace.IsLinearMoverFrom(fromColor) {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		peace := board[NewSquare(x, i)]
		if peace.IsLinearMoverFrom(fromColor) {
			return true
		}
		if peace != Empty {
			break
		}
	}
	return false
}

func (board Board) FindPeace(peace Peace) Square {
	for s := Square(0); s < 64; s++ {
		if board[s] == peace {
			return s
		}

	}
	return Square(-1)
}

func (board Board) kingTos(s Square) []Square {
	var tos []Square
	color := board[s].Color()
	oponentColor := color.Oponent()

	original := board[s]
	board[s] = Empty

	kingMoves := KingMoves[s]

	for _, square := range kingMoves {
		peace := board[square]
		if peace.IsEmptyOrNot(color) && !board.SquareUnderAttack(square, oponentColor) {
			tos = append(tos, square)
		}
	}

	board[s] = original
	return tos
}

func (board Board) rookTos(s Square, ks Square) []Square {
	var tos []Square
	color := board[s].Color()
	oponentColor := color.Oponent()

	check := func(square Square) bool {
		peace := board[square]
		if peace.IsEmptyOrNot(color) {
			original := board[square]
			board[square] = board[s]
			board[s] = Empty

			if !board.SquareUnderAttack(ks, oponentColor) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
		return peace.IsEmpty()
	}

	x := s.X()
	y := s.Y()

	for i := x; i > 0; {
		i--
		if !check(NewSquare(i, y)) {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		if !check(NewSquare(i, y)) {
			break
		}
	}

	for i := y; i > 0; {
		i--
		if !check(NewSquare(x, i)) {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		if !check(NewSquare(x, i)) {
			break
		}
	}
	return tos
}

func (board Board) Tos(s Square, bks Square) []Square {
	if board[s].IsKing() {
		return board.kingTos(s)
	}
	if board[s].IsRook() {
		return board.rookTos(s, bks)
	}

	return nil
}

func (board Board) Move(s Square) Move {
	wks := board.FindPeace(WhiteKing)
	tos := board.Tos(s, wks)

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
				blackTos := board.Tos(square, bks)
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
