package chess

import (
	"fmt"
	"log"

	"github.com/ggeorgiev/instant-chess/src/board/state"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Position struct {
	HorizontalySymetric bool
	VerticalySymetric   bool

	BoardState state.State

	Valid bool
}

func CreatePosition(boardState state.State) *Position {
	position := &Position{
		HorizontalySymetric: true,
		VerticalySymetric:   true,

		BoardState: boardState,

		Valid: true,
	}

	whiteKings := 0
	blackKings := 0
	whitePawns := 0
	blackPawns := 0

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		peace := boardState.Matrix[s]
		if peace.IsNoFigure() {
			continue
		}

		if peace.IsKing() {
			if peace.IsWhite() {
				whiteKings++
			} else {
				if boardState.Matrix.IsSquareUnderAttackFromWhite(s) {
					position.Valid = false
				}

				blackKings++
			}
		}
	}

	position.Valid = position.Valid &&
		(whiteKings == 1 && blackKings == 1) &&
		(whitePawns <= 8 && blackPawns <= 8)

	if !position.Valid {
		return position
	}

	return position
}

func ParsePosition(text string) *Position {
	board, err := state.Parse(text)
	if err != nil {
		log.Panicf("%s", err.Error())
	}
	return CreatePosition(board)
}

func (p *Position) Print() {
	fmt.Print(p.BoardState.Sprint())
}
