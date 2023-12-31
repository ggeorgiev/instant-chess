package chess

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
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

var KnightMoves = []square.Indexes{
	{10, 17},
	{11, 16, 18},
	{8, 12, 17, 19},
	{9, 13, 18, 20},
	{10, 14, 19, 21},
	{11, 15, 20, 22},
	{12, 21, 23},
	{13, 22},
	{2, 18, 25},
	{3, 19, 24, 26},
	{0, 4, 16, 20, 25, 27},
	{1, 5, 17, 21, 26, 28},
	{2, 6, 18, 22, 27, 29},
	{3, 7, 19, 23, 28, 30},
	{4, 20, 29, 31},
	{5, 21, 30},
	{1, 10, 26, 33},
	{0, 2, 11, 27, 32, 34},
	{1, 3, 8, 12, 24, 28, 33, 35},
	{2, 4, 9, 13, 25, 29, 34, 36},
	{3, 5, 10, 14, 26, 30, 35, 37},
	{4, 6, 11, 15, 27, 31, 36, 38},
	{5, 7, 12, 28, 37, 39},
	{6, 13, 29, 38},
	{9, 18, 34, 41},
	{8, 10, 19, 35, 40, 42},
	{9, 11, 16, 20, 32, 36, 41, 43},
	{10, 12, 17, 21, 33, 37, 42, 44},
	{11, 13, 18, 22, 34, 38, 43, 45},
	{12, 14, 19, 23, 35, 39, 44, 46},
	{13, 15, 20, 36, 45, 47},
	{14, 21, 37, 46},
	{17, 26, 42, 49},
	{16, 18, 27, 43, 48, 50},
	{17, 19, 24, 28, 40, 44, 49, 51},
	{18, 20, 25, 29, 41, 45, 50, 52},
	{19, 21, 26, 30, 42, 46, 51, 53},
	{20, 22, 27, 31, 43, 47, 52, 54},
	{21, 23, 28, 44, 53, 55},
	{22, 29, 45, 54},
	{25, 34, 50, 57},
	{24, 26, 35, 51, 56, 58},
	{25, 27, 32, 36, 48, 52, 57, 59},
	{26, 28, 33, 37, 49, 53, 58, 60},
	{27, 29, 34, 38, 50, 54, 59, 61},
	{28, 30, 35, 39, 51, 55, 60, 62},
	{29, 31, 36, 52, 61, 63},
	{30, 37, 53, 62},
	{33, 42, 58},
	{32, 34, 43, 59},
	{33, 35, 40, 44, 56, 60},
	{34, 36, 41, 45, 57, 61},
	{35, 37, 42, 46, 58, 62},
	{36, 38, 43, 47, 59, 63},
	{37, 39, 44, 60},
	{38, 45, 61},
	{41, 50},
	{40, 42, 51},
	{41, 43, 48, 52},
	{42, 44, 49, 53},
	{43, 45, 50, 54},
	{44, 46, 51, 55},
	{45, 47, 52},
	{46, 53},
}

var KnightAttackBitboardMasks = bitboard.Masks{
	0b0000000000000000000000000000000000000000000000100000010000000000,
	0b0000000000000000000000000000000000000000000001010000100000000000,
	0b0000000000000000000000000000000000000000000010100001000100000000,
	0b0000000000000000000000000000000000000000000101000010001000000000,
	0b0000000000000000000000000000000000000000001010000100010000000000,
	0b0000000000000000000000000000000000000000010100001000100000000000,
	0b0000000000000000000000000000000000000000101000000001000000000000,
	0b0000000000000000000000000000000000000000010000000010000000000000,
	0b0000000000000000000000000000000000000010000001000000000000000100,
	0b0000000000000000000000000000000000000101000010000000000000001000,
	0b0000000000000000000000000000000000001010000100010000000000010001,
	0b0000000000000000000000000000000000010100001000100000000000100010,
	0b0000000000000000000000000000000000101000010001000000000001000100,
	0b0000000000000000000000000000000001010000100010000000000010001000,
	0b0000000000000000000000000000000010100000000100000000000000010000,
	0b0000000000000000000000000000000001000000001000000000000000100000,
	0b0000000000000000000000000000001000000100000000000000010000000010,
	0b0000000000000000000000000000010100001000000000000000100000000101,
	0b0000000000000000000000000000101000010001000000000001000100001010,
	0b0000000000000000000000000001010000100010000000000010001000010100,
	0b0000000000000000000000000010100001000100000000000100010000101000,
	0b0000000000000000000000000101000010001000000000001000100001010000,
	0b0000000000000000000000001010000000010000000000000001000010100000,
	0b0000000000000000000000000100000000100000000000000010000001000000,
	0b0000000000000000000000100000010000000000000001000000001000000000,
	0b0000000000000000000001010000100000000000000010000000010100000000,
	0b0000000000000000000010100001000100000000000100010000101000000000,
	0b0000000000000000000101000010001000000000001000100001010000000000,
	0b0000000000000000001010000100010000000000010001000010100000000000,
	0b0000000000000000010100001000100000000000100010000101000000000000,
	0b0000000000000000101000000001000000000000000100001010000000000000,
	0b0000000000000000010000000010000000000000001000000100000000000000,
	0b0000000000000010000001000000000000000100000000100000000000000000,
	0b0000000000000101000010000000000000001000000001010000000000000000,
	0b0000000000001010000100010000000000010001000010100000000000000000,
	0b0000000000010100001000100000000000100010000101000000000000000000,
	0b0000000000101000010001000000000001000100001010000000000000000000,
	0b0000000001010000100010000000000010001000010100000000000000000000,
	0b0000000010100000000100000000000000010000101000000000000000000000,
	0b0000000001000000001000000000000000100000010000000000000000000000,
	0b0000001000000100000000000000010000000010000000000000000000000000,
	0b0000010100001000000000000000100000000101000000000000000000000000,
	0b0000101000010001000000000001000100001010000000000000000000000000,
	0b0001010000100010000000000010001000010100000000000000000000000000,
	0b0010100001000100000000000100010000101000000000000000000000000000,
	0b0101000010001000000000001000100001010000000000000000000000000000,
	0b1010000000010000000000000001000010100000000000000000000000000000,
	0b0100000000100000000000000010000001000000000000000000000000000000,
	0b0000010000000000000001000000001000000000000000000000000000000000,
	0b0000100000000000000010000000010100000000000000000000000000000000,
	0b0001000100000000000100010000101000000000000000000000000000000000,
	0b0010001000000000001000100001010000000000000000000000000000000000,
	0b0100010000000000010001000010100000000000000000000000000000000000,
	0b1000100000000000100010000101000000000000000000000000000000000000,
	0b0001000000000000000100001010000000000000000000000000000000000000,
	0b0010000000000000001000000100000000000000000000000000000000000000,
	0b0000000000000100000000100000000000000000000000000000000000000000,
	0b0000000000001000000001010000000000000000000000000000000000000000,
	0b0000000000010001000010100000000000000000000000000000000000000000,
	0b0000000000100010000101000000000000000000000000000000000000000000,
	0b0000000001000100001010000000000000000000000000000000000000000000,
	0b0000000010001000010100000000000000000000000000000000000000000000,
	0b0000000000010000101000000000000000000000000000000000000000000000,
	0b0000000000100000010000000000000000000000000000000000000000000000,
}

// Knight moves are simetrical
var AttackedFromKnight = KnightMoves

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

func knightMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		x := s.X()
		y := s.Y()

		if y > 1 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y-2))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y-2))
			}
		}
		if y > 0 {
			if x > 1 {
				squares = append(squares, square.NewIndex(x-2, y-1))
			}
			if x < 6 {
				squares = append(squares, square.NewIndex(x+2, y-1))
			}
		}
		if y < 7 {
			if x > 1 {
				squares = append(squares, square.NewIndex(x-2, y+1))
			}
			if x < 6 {
				squares = append(squares, square.NewIndex(x+2, y+1))
			}
		}
		if y < 6 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y+2))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y+2))
			}
		}

		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func knightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := knightMovesInternalHelper()

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
