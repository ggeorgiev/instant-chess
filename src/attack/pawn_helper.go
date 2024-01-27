package attack

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func whitePawnSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		file := s.File()
		rank := s.Rank()

		if rank != 7 && rank > 0 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank+1))
			}
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func blackPawnSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		file := s.File()
		rank := s.Rank()

		if rank != 0 && rank < 7 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank-1))
			}
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank-1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func whitePawnBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromWhitePawn

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func blackPawnBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromBlackPawn

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
