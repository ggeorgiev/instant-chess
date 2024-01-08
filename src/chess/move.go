package chess

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type ToAnswer struct {
	WhiteTo      square.Index
	BlackAnswers board.HalfMoves
}

func (ta ToAnswer) String() string {
	var ans []string
	for _, answer := range ta.BlackAnswers {
		var tos []string
		for _, to := range answer.Tos {
			tos = append(tos, to.String())
		}
		ans = append(ans, fmt.Sprintf("BlackFrom: %s, BlackTos: [%s]", answer.From, strings.Join(tos, ", ")))
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
