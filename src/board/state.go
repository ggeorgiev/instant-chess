package board

import (
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

type State struct {
	Peaces [square.Number]peace.Figure
}

func ParseState(text string) State {
	var state State

	rows := strings.Split(text, "\n")

	row := 0
	if len(rows[row]) == 0 {
		row++
	}

	for y := int8(8); y > 0; y-- {
		row += 2
		runes := util.Runes(rows[row])

		for x := int8(0); x < 8; x++ {
			state.Peaces[square.NewIndex(x, y-1)] = peace.FromSymbol(runes[4+x*4])
		}

	}
	return state
}
