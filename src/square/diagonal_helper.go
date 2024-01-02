package square

func diagonalsInternalHelper() []Diagonal {
	var diagonals []Diagonal

	for s := ZeroIndex; s <= LastIndex; s++ {
		f := s.File()
		r := s.Rank()

		for f > 0 && r > 0 {
			f--
			r--
		}
		diagonals = append(diagonals, Diagonal(NewIndex(f, r)))
	}
	return diagonals
}

func counterDiagonalsInternalHelper() []CounterDiagonal {
	var counterDiagonals []CounterDiagonal

	for s := ZeroIndex; s <= LastIndex; s++ {
		f := s.File()
		r := s.Rank()

		for f < 7 && r > 0 {
			f++
			r--
		}
		counterDiagonals = append(counterDiagonals, CounterDiagonal(NewIndex(f, r)))
	}
	return counterDiagonals
}
