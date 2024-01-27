package state

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
)

// Chess guid
type CGUID struct {
	Rights           move.Rights
	Peaces           peace.Codes
	PeacePermutation uint64
}
