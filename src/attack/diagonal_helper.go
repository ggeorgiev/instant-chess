package attack

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func riseLeftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		f := s.File()
		r := s.Rank()

		for f > 0 && r < 7 {
			f--
			r++
			squares = append(squares, square.NewIndex(f, r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func riseRightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		f := s.File()
		r := s.Rank()

		for f < 7 && r < 7 {
			f++
			r++
			squares = append(squares, square.NewIndex(f, r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallLeftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		f := s.File()
		r := s.Rank()

		for f > 0 && r > 0 {
			f--
			r--
			squares = append(squares, square.NewIndex(f, r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallRightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		f := s.File()
		r := s.Rank()

		for f < 7 && r > 0 {
			f++
			r--
			squares = append(squares, square.NewIndex(f, r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func riseLeftBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := RiseLeftSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func riseRightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := RiseRightSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func fallLeftBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FallLeftSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func fallRightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FallRightSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func diagonalsRiseBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for f := range RiseLeftBitboardMasks {
		masks = append(masks, RiseLeftBitboardMasks[f]|RiseRightBitboardMasks[f])
	}

	return masks
}

func diagonalsFallBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for f := range FallLeftBitboardMasks {
		masks = append(masks, FallLeftBitboardMasks[f]|FallRightBitboardMasks[f])
	}

	return masks
}

func diagonalsBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for f := range FallLeftBitboardMasks {
		masks = append(masks, FallLeftBitboardMasks[f]|FallRightBitboardMasks[f]|RiseLeftBitboardMasks[f]|RiseRightBitboardMasks[f])
	}

	return masks
}
