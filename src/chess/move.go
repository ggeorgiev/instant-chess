package chess

import (
	"fmt"
	"strings"
)

type Answer struct {
	BlackFrom Square
	BlackTos  []Square
}

func (a Answer) String() string {
	var tos []string
	for _, t := range a.BlackTos {
		tos = append(tos, t.String())
	}
	return fmt.Sprintf("BlackFrom: %s, BlackTos: [%s]", a.BlackFrom, strings.Join(tos, ", "))
}

type ToAnswer struct {
	WhiteTo Square
	Answers []Answer
}

func (ta ToAnswer) String() string {
	var ans []string
	for _, a := range ta.Answers {
		ans = append(ans, a.String())
	}
	return fmt.Sprintf("  Answers: %s\n", strings.Join(ans, ", "))
}

type Move struct {
	WhiteForm Square
	ToAnswers []ToAnswer
}

func (m Move) String() string {
	var tas []string
	for _, ta := range m.ToAnswers {
		tas = append(tas, fmt.Sprintf("White: %s:%s\n%s", m.WhiteForm, ta.WhiteTo, ta.String()))
	}
	return strings.Join(tas, "")
}

type Moves []Move

func (ms Moves) String() string {
	var mvs []string
	for _, m := range ms {
		mvs = append(mvs, m.String())
	}
	return fmt.Sprintf("%s", strings.Join(mvs, "\n"))
}

var KingMoves = [][]Square{
	{8, 1, 9},
	{0, 8, 9, 2, 10},
	{1, 9, 10, 3, 11},
	{2, 10, 11, 4, 12},
	{3, 11, 12, 5, 13},
	{4, 12, 13, 6, 14},
	{5, 13, 14, 7, 15},
	{6, 14, 15},
	{0, 16, 1, 9, 17},
	{0, 8, 16, 1, 17, 2, 10, 18},
	{1, 9, 17, 2, 18, 3, 11, 19},
	{2, 10, 18, 3, 19, 4, 12, 20},
	{3, 11, 19, 4, 20, 5, 13, 21},
	{4, 12, 20, 5, 21, 6, 14, 22},
	{5, 13, 21, 6, 22, 7, 15, 23},
	{6, 14, 22, 7, 23},
	{8, 24, 9, 17, 25},
	{8, 16, 24, 9, 25, 10, 18, 26},
	{9, 17, 25, 10, 26, 11, 19, 27},
	{10, 18, 26, 11, 27, 12, 20, 28},
	{11, 19, 27, 12, 28, 13, 21, 29},
	{12, 20, 28, 13, 29, 14, 22, 30},
	{13, 21, 29, 14, 30, 15, 23, 31},
	{14, 22, 30, 15, 31},
	{16, 32, 17, 25, 33},
	{16, 24, 32, 17, 33, 18, 26, 34},
	{17, 25, 33, 18, 34, 19, 27, 35},
	{18, 26, 34, 19, 35, 20, 28, 36},
	{19, 27, 35, 20, 36, 21, 29, 37},
	{20, 28, 36, 21, 37, 22, 30, 38},
	{21, 29, 37, 22, 38, 23, 31, 39},
	{22, 30, 38, 23, 39},
	{24, 40, 25, 33, 41},
	{24, 32, 40, 25, 41, 26, 34, 42},
	{25, 33, 41, 26, 42, 27, 35, 43},
	{26, 34, 42, 27, 43, 28, 36, 44},
	{27, 35, 43, 28, 44, 29, 37, 45},
	{28, 36, 44, 29, 45, 30, 38, 46},
	{29, 37, 45, 30, 46, 31, 39, 47},
	{30, 38, 46, 31, 47},
	{32, 48, 33, 41, 49},
	{32, 40, 48, 33, 49, 34, 42, 50},
	{33, 41, 49, 34, 50, 35, 43, 51},
	{34, 42, 50, 35, 51, 36, 44, 52},
	{35, 43, 51, 36, 52, 37, 45, 53},
	{36, 44, 52, 37, 53, 38, 46, 54},
	{37, 45, 53, 38, 54, 39, 47, 55},
	{38, 46, 54, 39, 55},
	{40, 56, 41, 49, 57},
	{40, 48, 56, 41, 57, 42, 50, 58},
	{41, 49, 57, 42, 58, 43, 51, 59},
	{42, 50, 58, 43, 59, 44, 52, 60},
	{43, 51, 59, 44, 60, 45, 53, 61},
	{44, 52, 60, 45, 61, 46, 54, 62},
	{45, 53, 61, 46, 62, 47, 55, 63},
	{46, 54, 62, 47, 63},
	{48, 49, 57},
	{48, 56, 49, 50, 58},
	{49, 57, 50, 51, 59},
	{50, 58, 51, 52, 60},
	{51, 59, 52, 53, 61},
	{52, 60, 53, 54, 62},
	{53, 61, 54, 55, 63},
	{54, 62, 55},
}

// King moves are simetrical
var AttackedFromKing = KingMoves

func kingMovesInternalHelper() [][]Square {
	var squaresList [][]Square

	for s := Square(0); s < 64; s++ {
		var squares []Square

		x := s.X()
		y := s.Y()

		if x > 0 {
			if y > 0 {
				squares = append(squares, NewSquare(x-1, y-1))
			}
			squares = append(squares, NewSquare(x-1, y))
			if y < 7 {
				squares = append(squares, NewSquare(x-1, y+1))
			}
		}
		if y > 0 {
			squares = append(squares, NewSquare(x, y-1))
		}
		if y < 7 {
			squares = append(squares, NewSquare(x, y+1))
		}
		if x < 7 {
			if y > 0 {
				squares = append(squares, NewSquare(x+1, y-1))
			}
			squares = append(squares, NewSquare(x+1, y))
			if y < 7 {
				squares = append(squares, NewSquare(x+1, y+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}
