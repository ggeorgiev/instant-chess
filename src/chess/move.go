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
