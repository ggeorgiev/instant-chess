package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func riseSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		for i := y + 1; i < 8; i++ {
			squares = append(squares, square.NewIndex(x, i))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func rightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		for i := x + 1; i < 8; i++ {
			squares = append(squares, square.NewIndex(i, y))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		for i := y - 1; i >= 0; i-- {
			squares = append(squares, square.NewIndex(x, i))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func leftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		for i := x - 1; i >= 0; i-- {
			squares = append(squares, square.NewIndex(i, y))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func riseBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := RiseSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func rightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := RightSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func fallBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FallSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func leftBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := LeftSquareIndexes

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func linearsRiseRightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for i := range RiseBitboardMasks {
		masks = append(masks, RiseBitboardMasks[i]|RightBitboardMasks[i])
	}

	return masks
}

func linearsFallLeftBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for i := range FallBitboardMasks {
		masks = append(masks, FallBitboardMasks[i]|LeftBitboardMasks[i])
	}

	return masks
}
