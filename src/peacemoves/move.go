package peacemoves

import "github.com/ggeorgiev/instant-chess/src/square"

type FromTo struct {
	From square.Index
	To   square.Indexes
}

type FromTos []FromTo
