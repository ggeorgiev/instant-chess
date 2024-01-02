package peacemoves

import "github.com/ggeorgiev/instant-chess/src/square"

func kingMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		file := s.File()
		rank := s.Rank()

		if rank > 0 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank-1))
			}
			squares = append(squares, square.NewIndex(file, rank-1))
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank-1))
			}
		}

		if file > 0 {
			squares = append(squares, square.NewIndex(file-1, rank))
		}

		if file < 7 {
			squares = append(squares, square.NewIndex(file+1, rank))
		}

		if rank < 7 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank+1))
			}
			squares = append(squares, square.NewIndex(file, rank+1))
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}
