package chess

import (
	"fmt"
	"strings"
)

type Position struct {
	VerticalySymetric bool

	Result Result

	Matrix [8][8]Peace

	WhitePeaces int8
	BlackPeaces int8

	Valid bool
}

func CreatePosition(matrix [8][8]Peace) *Position {
	position := &Position{
		VerticalySymetric: true,

		Result: UnknownResult,

		Matrix: matrix,

		WhitePeaces: 0,
		BlackPeaces: 0,
	}

	whiteKings := 0
	blackKings := 0
	whitePawns := 0
	blackPawns := 0

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			peace := matrix[x][y]
			if peace.Type == King {
				if peace.Color == PeaceColorWhite {
					whiteKings++
				} else {
					blackKings++
				}
			}
			if peace.Type == Pawn {
				if peace.Color == PeaceColorWhite {
					whitePawns++
				} else {
					blackPawns++
				}
			}
		}
	}

	position.Valid = (whiteKings == 1 && blackKings == 1) &&
		(whitePawns <= 8 && blackPawns <= 8)

	return position
}

func runes(row string) []rune {
	runes := []rune(row)

	var result []rune
	for _, r := range runes {
		if r != 65038 {
			result = append(result, r)
		}
	}
	return result
}

func ParsePosition(text string) *Position {
	var matrix [8][8]Peace

	rows := strings.Split(text, "\n")

	row := 0
	if len(rows[row]) == 0 {
		row++
	}

	for y := 8; y > 0; y-- {
		row += 2
		runes := runes(rows[row])

		for x := 0; x < 8; x++ {
			matrix[x][y-1] = PeaceFromSymbol(runes[4+x*4])
		}

	}

	return CreatePosition(matrix)
}

func (p *Position) Print() {
	letters := "    a   b   c   d   e   f   g   h  "
	separator := "  +---+---+---+---+---+---+---+---+"

	fmt.Println(letters)
	fmt.Println(separator)

	for y := 8; y > 0; y-- {
		fmt.Printf("%d |", y)
		for x := 0; x < 8; x++ {
			fmt.Printf(" %s |", p.Matrix[x][y-1].Symbol())
		}
		fmt.Printf(" %d\n", y)
		fmt.Println(separator)
	}

	fmt.Println(letters)
}
