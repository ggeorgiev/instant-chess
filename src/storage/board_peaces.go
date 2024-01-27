package storage

import (
	"strings"

	"github.com/ggeorgiev/instant-chess/src/peace"
)

type BoardPeaces [32]peace.Figure

func (bf BoardPeaces) String() string {
	var sb strings.Builder
	for _, f := range bf {
		sb.WriteString(f.Symbol())
	}
	return sb.String()
}
