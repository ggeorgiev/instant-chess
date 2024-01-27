package attack

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func riseSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		for r := s.Rank() + 1; r <= square.LastRank; r++ {
			squares = append(squares, square.NewIndex(s.File(), r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func rightSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		for f := s.File() + 1; f <= square.LastFile; f++ {
			squares = append(squares, square.NewIndex(f, s.Rank()))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func fallSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		for r := s.Rank() - 1; r >= square.ZeroRank; r-- {
			squares = append(squares, square.NewIndex(s.File(), r))
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func leftSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		for f := s.File() - 1; f >= square.ZeroFile; f-- {
			squares = append(squares, square.NewIndex(f, s.Rank()))
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

	for f := range RiseBitboardMasks {
		masks = append(masks, RiseBitboardMasks[f]|RightBitboardMasks[f])
	}

	return masks
}

func linearsFallLeftBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for f := range FallBitboardMasks {
		masks = append(masks, FallBitboardMasks[f]|LeftBitboardMasks[f])
	}

	return masks
}

func linearsBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	for f := range FallBitboardMasks {
		masks = append(masks, FallBitboardMasks[f]|LeftBitboardMasks[f]|RiseBitboardMasks[f]|RightBitboardMasks[f])
	}

	return masks
}
