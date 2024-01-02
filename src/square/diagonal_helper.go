package square

func diagonalInternalHelper() []int8 {
	var diagonals []int8

	for s := ZeroIndex; s <= LastIndex; s++ {
		f := s.File()
		r := s.Rank()

		for f > 0 && r > 0 {
			f--
			r--
		}
		diagonals = append(diagonals, int8(NewIndex(f, r)))
	}
	return diagonals
}

func antiDiagonalInternalHelper() []int8 {
	var antiDiagonals []int8

	for s := ZeroIndex; s <= LastIndex; s++ {
		f := s.File()
		r := s.Rank()

		for f < 7 && r > 0 {
			f++
			r--
		}
		antiDiagonals = append(antiDiagonals, int8(NewIndex(f, r)))
	}
	return antiDiagonals
}
