package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) IsSquareUnderAttackFrom(s square.Index, from square.Index) bool {
	attackMask := attack.PeaceBitboardMasks(m[from])[from]
	return square.IndexMask[s]&attackMask != 0
}
