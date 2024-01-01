package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func riseLeftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		i := x
		j := y

		for i > 0 && j < 7 {
			i--
			j++
			squares = append(squares, square.NewIndex(i, j))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func riseRightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		i := x
		j := y

		for i < 7 && j < 7 {
			i++
			j++
			squares = append(squares, square.NewIndex(i, j))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallLeftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		i := x
		j := y

		for i > 0 && j > 0 {
			i--
			j--
			squares = append(squares, square.NewIndex(i, j))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallRightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		i := x
		j := y

		for i < 7 && j > 0 {
			i++
			j--
			squares = append(squares, square.NewIndex(i, j))
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

func riseBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for i := range RiseLeftBitboardMasks {
		masks = append(masks, RiseLeftBitboardMasks[i]|RiseRightBitboardMasks[i])
	}

	return masks
}

func fallBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for i := range FallLeftBitboardMasks {
		masks = append(masks, FallLeftBitboardMasks[i]|FallRightBitboardMasks[i])
	}

	return masks
}
