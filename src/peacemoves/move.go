package peacemoves

import "github.com/ggeorgiev/instant-chess/src/square"

type FromTo struct {
	From square.Index
	Tos  square.Indexes
}

type FromTos []FromTo
