package chess

import (
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peacemoves"
)

type Moves []peacemoves.Full

func (ms Moves) String() string {
	var mvs []string
	for _, m := range ms {
		mvs = append(mvs, m.String())
	}
	return strings.Join(mvs, "\n")
}
