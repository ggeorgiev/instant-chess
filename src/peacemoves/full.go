package peacemoves

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/square"
)

type Answer struct {
	WhiteTo      square.Index
	BlackAnswers Halfs
}
type Answers []Answer

type Full struct {
	WhiteForm square.Index
	Answers   Answers
}

type Fulls []Full

func (a Answer) String() string {
	var ans []string
	for _, answer := range a.BlackAnswers {
		var tos []string
		for _, to := range answer.Tos {
			tos = append(tos, to.String())
		}
		ans = append(ans, fmt.Sprintf("BlackFrom: %s, BlackTos: [%s]", answer.From, strings.Join(tos, ", ")))
	}
	return fmt.Sprintf("  Answers: %s\n", strings.Join(ans, ", "))
}

func (m Full) String() string {
	var tas []string
	for _, ta := range m.Answers {
		tas = append(tas, fmt.Sprintf("White: %s:%s\n%s", m.WhiteForm, ta.WhiteTo, ta.String()))
	}
	return strings.Join(tas, "")
}

func (fs Fulls) String() string {
	var mvs []string
	for _, m := range fs {
		mvs = append(mvs, m.String())
	}
	return strings.Join(mvs, "\n")
}
