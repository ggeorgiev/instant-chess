package state

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
)

// Chess guid
type CGUID struct {
	Rights            move.Rights
	Figures           peace.Figures
	FigurePermutation uint64
}
