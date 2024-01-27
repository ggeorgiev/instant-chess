package alignment

import "github.com/ggeorgiev/instant-chess/src/square"

func IsSameRank(s1 square.Index, s2 square.Index) bool {
	return s1.Rank() == s2.Rank()
}

func IsSameFile(s1 square.Index, s2 square.Index) bool {
	return s1.File() == s2.File()
}

func IsSameDiagonal(s1 square.Index, s2 square.Index) bool {
	return s1.Diagonal() == s2.Diagonal()
}

func IsSameCounterDiagonal(s1 square.Index, s2 square.Index) bool {
	return s1.CounterDiagonal() == s2.CounterDiagonal()
}
