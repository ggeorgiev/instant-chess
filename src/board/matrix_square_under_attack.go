package board

import (
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) IsSquareUnderAttackFrom(s square.Index, from square.Index) bool {
	attackMask := peaceattacks.PeaceBitboardMasks(m[from])[from]
	return square.IndexMask[s]&attackMask != 0
}
