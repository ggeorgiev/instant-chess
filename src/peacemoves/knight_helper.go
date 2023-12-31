package peacemoves

import "github.com/ggeorgiev/instant-chess/src/square"

func knightMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		x := s.X()
		y := s.Y()

		if y > 1 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y-2))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y-2))
			}
		}
		if y > 0 {
			if x > 1 {
				squares = append(squares, square.NewIndex(x-2, y-1))
			}
			if x < 6 {
				squares = append(squares, square.NewIndex(x+2, y-1))
			}
		}
		if y < 7 {
			if x > 1 {
				squares = append(squares, square.NewIndex(x-2, y+1))
			}
			if x < 6 {
				squares = append(squares, square.NewIndex(x+2, y+1))
			}
		}
		if y < 6 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y+2))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y+2))
			}
		}

		squaresList = append(squaresList, squares)
	}
	return squaresList
}
