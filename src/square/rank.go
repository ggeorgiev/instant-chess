package square

type Rank int8
type Ranks []Rank

const (
	InvalidRank = Rank(-1)
	ZeroRank    = Rank(0)
	LastRank    = Rank(7)
)

var ValidRanks = Ranks{ZeroRank, Rank(1), Rank(2), Rank(3), Rank(4), Rank(5), Rank(6), LastRank}
var AllRanks = Ranks{InvalidRank, ZeroRank, Rank(1), Rank(2), Rank(3), Rank(4), Rank(5), Rank(6), LastRank}

func RankFromRune(r rune) Rank {
	if r == '-' {
		return InvalidRank
	}

	return Rank(r - '1')
}

func (r Rank) Rune() rune {
	if r == InvalidRank {
		return '-'
	}
	return '1' + rune(r)
}
