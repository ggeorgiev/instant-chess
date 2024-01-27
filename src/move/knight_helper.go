package move

import "github.com/ggeorgiev/instant-chess/src/square"

func knightMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		file := s.File()
		rank := s.Rank()

		if rank > 1 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank-2))
			}
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank-2))
			}
		}
		if rank > 0 {
			if file > 1 {
				squares = append(squares, square.NewIndex(file-2, rank-1))
			}
			if file < 6 {
				squares = append(squares, square.NewIndex(file+2, rank-1))
			}
		}
		if rank < 7 {
			if file > 1 {
				squares = append(squares, square.NewIndex(file-2, rank+1))
			}
			if file < 6 {
				squares = append(squares, square.NewIndex(file+2, rank+1))
			}
		}
		if rank < 6 {
			if file > 0 {
				squares = append(squares, square.NewIndex(file-1, rank+2))
			}
			if file < 7 {
				squares = append(squares, square.NewIndex(file+1, rank+2))
			}
		}

		squaresList = append(squaresList, squares)
	}
	return squaresList
}
