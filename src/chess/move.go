package chess

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/square"
)

type Answer struct {
	BlackFrom square.Index
	BlackTos  square.Indexes
}

func (a Answer) String() string {
	var tos []string
	for _, t := range a.BlackTos {
		tos = append(tos, t.String())
	}
	return fmt.Sprintf("BlackFrom: %s, BlackTos: [%s]", a.BlackFrom, strings.Join(tos, ", "))
}

type ToAnswer struct {
	WhiteTo square.Index
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
	WhiteForm square.Index
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
	return strings.Join(mvs, "\n")
}

var KingMoves = []square.Indexes{
	{1, 8, 9},
	{0, 2, 8, 9, 10},
	{1, 3, 9, 10, 11},
	{2, 4, 10, 11, 12},
	{3, 5, 11, 12, 13},
	{4, 6, 12, 13, 14},
	{5, 7, 13, 14, 15},
	{6, 14, 15},
	{0, 1, 9, 16, 17},
	{0, 1, 2, 8, 10, 16, 17, 18},
	{1, 2, 3, 9, 11, 17, 18, 19},
	{2, 3, 4, 10, 12, 18, 19, 20},
	{3, 4, 5, 11, 13, 19, 20, 21},
	{4, 5, 6, 12, 14, 20, 21, 22},
	{5, 6, 7, 13, 15, 21, 22, 23},
	{6, 7, 14, 22, 23},
	{8, 9, 17, 24, 25},
	{8, 9, 10, 16, 18, 24, 25, 26},
	{9, 10, 11, 17, 19, 25, 26, 27},
	{10, 11, 12, 18, 20, 26, 27, 28},
	{11, 12, 13, 19, 21, 27, 28, 29},
	{12, 13, 14, 20, 22, 28, 29, 30},
	{13, 14, 15, 21, 23, 29, 30, 31},
	{14, 15, 22, 30, 31},
	{16, 17, 25, 32, 33},
	{16, 17, 18, 24, 26, 32, 33, 34},
	{17, 18, 19, 25, 27, 33, 34, 35},
	{18, 19, 20, 26, 28, 34, 35, 36},
	{19, 20, 21, 27, 29, 35, 36, 37},
	{20, 21, 22, 28, 30, 36, 37, 38},
	{21, 22, 23, 29, 31, 37, 38, 39},
	{22, 23, 30, 38, 39},
	{24, 25, 33, 40, 41},
	{24, 25, 26, 32, 34, 40, 41, 42},
	{25, 26, 27, 33, 35, 41, 42, 43},
	{26, 27, 28, 34, 36, 42, 43, 44},
	{27, 28, 29, 35, 37, 43, 44, 45},
	{28, 29, 30, 36, 38, 44, 45, 46},
	{29, 30, 31, 37, 39, 45, 46, 47},
	{30, 31, 38, 46, 47},
	{32, 33, 41, 48, 49},
	{32, 33, 34, 40, 42, 48, 49, 50},
	{33, 34, 35, 41, 43, 49, 50, 51},
	{34, 35, 36, 42, 44, 50, 51, 52},
	{35, 36, 37, 43, 45, 51, 52, 53},
	{36, 37, 38, 44, 46, 52, 53, 54},
	{37, 38, 39, 45, 47, 53, 54, 55},
	{38, 39, 46, 54, 55},
	{40, 41, 49, 56, 57},
	{40, 41, 42, 48, 50, 56, 57, 58},
	{41, 42, 43, 49, 51, 57, 58, 59},
	{42, 43, 44, 50, 52, 58, 59, 60},
	{43, 44, 45, 51, 53, 59, 60, 61},
	{44, 45, 46, 52, 54, 60, 61, 62},
	{45, 46, 47, 53, 55, 61, 62, 63},
	{46, 47, 54, 62, 63},
	{48, 49, 57},
	{48, 49, 50, 56, 58},
	{49, 50, 51, 57, 59},
	{50, 51, 52, 58, 60},
	{51, 52, 53, 59, 61},
	{52, 53, 54, 60, 62},
	{53, 54, 55, 61, 63},
	{54, 55, 62},
}

// King moves are simetrical
var AttackedFromKing = KingMoves

func kingMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		x := s.X()
		y := s.Y()

		if y > 0 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y-1))
			}
			squares = append(squares, square.NewIndex(x, y-1))
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y-1))
			}
		}

		if x > 0 {
			squares = append(squares, square.NewIndex(x-1, y))
		}

		if x < 7 {
			squares = append(squares, square.NewIndex(x+1, y))
		}

		if y < 7 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y+1))
			}
			squares = append(squares, square.NewIndex(x, y+1))
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}
