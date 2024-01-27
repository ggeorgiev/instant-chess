package peace

import (
	"strings"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

type Codes []Code

// Provide the sort interface for Codes
func (codes Codes) Len() int           { return len(codes) }
func (codes Codes) Less(i, j int) bool { return codes[i] < codes[j] }
func (codes Codes) Swap(i, j int)      { codes[i], codes[j] = codes[j], codes[i] }

func Parse(text string) (Codes, error) {
	runes := util.Runes(text)
	position := make([]int, len(runes))
	peaces := make(Codes, len(runes))
	for i, symbol := range runes {
		position[i] = 0
		var err error
		peaces[i], err = FromSymbol(symbol)
		if err != nil {
			return nil, err
		}
	}
	return peaces, nil
}

func MustParse(text string) Codes {
	codes, err := Parse(text)
	if err != nil {
		panic(err.Error())
	}
	return codes
}

func (fs Codes) MoveRights() move.RightsList {
	whiteKing := false
	blackKing := false
	whiteRooks := 0
	blackRooks := 0
	whitePawn := false
	blackPawn := false

	for _, f := range fs {
		switch f {
		case WhiteKing:
			whiteKing = true
		case BlackKing:
			blackKing = true
		case WhiteRook:
			whiteRooks++
		case BlackRook:
			blackRooks++
		case WhitePawn:
			whitePawn = true
		case BlackPawn:
			blackPawn = true
		}
	}

	whiteCasting := move.RightsList{move.NoRights}
	if whiteKing {
		if whiteRooks >= 1 {
			whiteCasting = append(whiteCasting, move.WhiteKingsideCastlingRights, move.WhiteQueensideCastlingRights)
		}
		if whiteRooks >= 2 {
			whiteCasting = append(whiteCasting, move.WhiteBothCastlingRights)
		}
	}

	blackCasting := move.RightsList{move.NoRights}
	if blackKing {
		if blackRooks >= 1 {
			blackCasting = append(blackCasting, move.BlackKingsideCastlingRights, move.BlackQueensideCastlingRights)
		}
		if blackRooks >= 2 {
			blackCasting = append(blackCasting, move.BlackBothCastlingRights)
		}
	}

	enPassantFiles := move.RightsList{move.NoEnPassantFile}
	if whitePawn && blackPawn {
		for f := square.ZeroFile; f <= square.LastFile; f++ {
			enPassantFiles = append(enPassantFiles, move.NoRights.SetEnPassantFile(f))
		}
	}

	return move.CombineRights(whiteCasting, blackCasting, enPassantFiles)
}

func (fs Codes) Copy() Codes {
	peaces := make(Codes, len(fs))
	copy(peaces, fs)
	return peaces
}

func (fs Codes) String() string {
	var sb strings.Builder
	for _, f := range fs {
		sb.WriteString(f.Symbol())
	}
	return sb.String()
}

func (fs Codes) RemoveOne(f Code) Codes {
	for i, pc := range fs {
		if pc == f {
			newSize := len(fs) - 1
			fs[i] = fs[newSize]
			return fs[:newSize]
		}
	}
	panic("Removing pc that does not exists")
}
