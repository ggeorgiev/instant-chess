package peacemoves

import "github.com/ggeorgiev/instant-chess/src/square"

func kingMovesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes

		x := s.X()
		y := s.Y()

		if y > 0 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y-1))
			}
			squares = append(squares, square.NewIndex(x, y-1))
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y-1))
			}
		}

		if x > 0 {
			squares = append(squares, square.NewIndex(x-1, y))
		}

		if x < 7 {
			squares = append(squares, square.NewIndex(x+1, y))
		}

		if y < 7 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y+1))
			}
			squares = append(squares, square.NewIndex(x, y+1))
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}
