package board

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

const (
	letters   = "····a···b···c···d···e···f···g···h····\n"
	separator = "··+---+---+---+---+---+---+---+---+··\n"
)

var runesExpectedLength = len(util.Runes(separator)) - 1

type Matrix [square.Number]peace.Figure

func ParseMatrix(text string) (Matrix, error) {
	var matrix Matrix

	rows := strings.Split(text, "\n")

	row := 0
	// Allowing starting with an empty line
	if len(rows[row]) == 0 {
		row++
	}

	for r := square.LastRank; r >= square.ZeroRank; r-- {
		row += 2
		runes := util.Runes(rows[row])
		if len(runes) != runesExpectedLength {
			return matrix, fmt.Errorf("on rank %d unexpected length %d instead %d", r+1, len(runes), runesExpectedLength)
		}

		for f := square.ZeroFile; f <= square.LastFile; f++ {
			matrix[square.NewIndex(f, r)] = peace.FromSymbol(runes[4+f*4])
		}
	}
	return matrix, nil
}

func (m Matrix) Sprint() string {
	var sb strings.Builder
	sb.Grow(len(letters)*2 + len(separator)*9 + 8*(9*4+1))

	sb.WriteString(letters)
	sb.WriteString(separator)

	for r := square.LastRank; r >= square.ZeroRank; r-- {
		sb.WriteString(fmt.Sprintf("%d·|", r+1))
		for f := square.ZeroFile; f <= square.LastFile; f++ {
			sb.WriteString(fmt.Sprintf(" %s |", m[square.NewIndex(f, r)].Symbol()))
		}
		sb.WriteString(fmt.Sprintf("·%d\n", r+1))
		sb.WriteString(separator)
	}

	sb.WriteString(letters)
	return sb.String()
}

func (m Matrix) FindSinglePeace(peace peace.Figure) square.Index {
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		if m[s] == peace {
			return s
		}
	}
	return square.InvalidIndex
}
