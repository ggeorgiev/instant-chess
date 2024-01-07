package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func whiteDirectInternalHelper() DirectsList {
	var directsList DirectsList

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var directs Directs
		for _, attackIndex := range FromKing[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.WhiteKing,
			})
		}
		for _, attackIndex := range FromKnight[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.WhiteKnight,
			})
		}
		for _, attackIndex := range FromWhitePawn[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.WhitePawn,
			})
		}

		directsList = append(directsList, directs)
	}

	return directsList
}

func blackDirectInternalHelper() DirectsList {
	var directsList DirectsList

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var directs Directs
		for _, attackIndex := range FromKing[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.BlackKing,
			})
		}
		for _, attackIndex := range FromKnight[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.BlackKnight,
			})
		}
		for _, attackIndex := range FromBlackPawn[s] {
			directs = append(directs, Direct{
				Index: attackIndex,
				Peace: peace.BlackPawn,
			})
		}

		directsList = append(directsList, directs)
	}

	return directsList
}
