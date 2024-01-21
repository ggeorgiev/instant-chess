package batch

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
)

type Batch struct {
	Figures peace.Figures
	Rights  move.Rights
}

func Create(figures peace.Figures, rights move.Rights) *Batch {
	return &Batch{
		Figures: figures,
		Rights:  rights,
	}
}
