package storage

import (
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peace"
)

type BoardFigures [32]peace.Figure

func (bf BoardFigures) String() string {
	var sb strings.Builder
	for _, f := range bf {
		sb.WriteString(f.Symbol())
	}
	return sb.String()
}
