package move

import "github.com/ggeorgiev/instant-chess/src/square"

type Half struct {
	From square.Index
	Tos  square.Indexes
}

type Halfs []Half
