package chess

import (
	"fmt"
	"log"

	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Position struct {
	HorizontalySymetric bool
	VerticalySymetric   bool

	Result Result

	BoardState board.State

	WhitePeaces int8
	BlackPeaces int8

	Valid bool

	Moves Moves
}

func CreatePosition(boardState board.State) *Position {
	position := &Position{
		HorizontalySymetric: true,
		VerticalySymetric:   true,

		Result: UnknownResult,

		BoardState: boardState,

		WhitePeaces: 0,
		BlackPeaces: 0,

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

		if peace.IsWhite() {
			position.WhitePeaces++
		} else {
			position.BlackPeaces++
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

	position.Moves = Board(boardState.Matrix).Moves()

	position.Valid = position.Valid &&
		(whiteKings == 1 && blackKings == 1) &&
		(whitePawns <= 8 && blackPawns <= 8)

	return position
}

func ParsePosition(text string) *Position {
	board, err := board.ParseState(text)
	if err != nil {
		log.Panicf("%s", err.Error())
	}
	return CreatePosition(board)
}

func (p *Position) Print() {
	fmt.Print(p.BoardState.Sprint())
}

func (p *Position) M1() bool {
	for _, move := range p.Moves {
		for _, toAnswer := range move.ToAnswers {
			if len(toAnswer.Answers) == 0 {
				brd := Board(p.BoardState.Matrix)
				original := brd[toAnswer.WhiteTo]
				brd[toAnswer.WhiteTo] = brd[move.WhiteForm]
				brd[move.WhiteForm] = peace.NoFigure

				bk := board.Matrix(brd).FindSinglePeace(peace.BlackKing)
				mate := board.Matrix(brd).IsSquareUnderAttackFromWhite(bk)

				brd[move.WhiteForm] = brd[toAnswer.WhiteTo]
				brd[toAnswer.WhiteTo] = original
				if mate {
					return true
				}
			}
		}
	}
	return false
}
