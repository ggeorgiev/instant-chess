package move

import (
	"fmt"
	"regexp"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

const (
	rightsTemplate = "· O-O: --, O-O-O: --, En Passant: - ·\n"

	WhiteKingsideCastlingRightRuneIndex  = 7
	WhiteQueensideCastlingRightRuneIndex = 18
	BlackKingsideCastlingRightRuneIndex  = 8
	BlackQueensideCastlingRightRuneIndex = 19
	EnPassantRightFileRuneIndex          = 34
)

var (
	rightsPattern = regexp.MustCompile("· O-O: [-w][-b], O-O-O: [-w][-b], En Passant: [-A-H] ·")
)

type CastlingRights struct {
	Kingside  bool
	Queenside bool
}

type CastlingRightsList []CastlingRights

type Rights struct {
	WhiteCastling CastlingRights
	BlackCastling CastlingRights
	EnPassantFile square.File
}

type RightsList []Rights

var (
	NoCastling        = CastlingRights{false, false}
	KingsideCastling  = CastlingRights{true, false}
	QueensideCastling = CastlingRights{false, true}
	BothCastlings     = CastlingRights{true, true}
)

func ParseRights(text string) (*Rights, error) {
	match := rightsPattern.FindAllString(text, -1)
	if len(match) == 0 {
		return nil, fmt.Errorf("failed to find the rights description")
	}
	if len(match) > 1 {
		return nil, fmt.Errorf("found to many rights descriptions")
	}

	runes := util.Runes(match[0])

	rights := Rights{}

	if runes[WhiteKingsideCastlingRightRuneIndex] == 'w' {
		rights.WhiteCastling.Kingside = true
	}
	if runes[WhiteQueensideCastlingRightRuneIndex] == 'w' {
		rights.WhiteCastling.Queenside = true
	}
	if runes[BlackKingsideCastlingRightRuneIndex] == 'b' {
		rights.BlackCastling.Kingside = true
	}
	if runes[BlackQueensideCastlingRightRuneIndex] == 'b' {
		rights.BlackCastling.Queenside = true
	}
	rights.EnPassantFile = square.FileFromRune(runes[EnPassantRightFileRuneIndex])

	return &rights, nil
}

func GenerateRightsList(whiteRooks int, blackRooks int, whitePawns int, blackPawns int) RightsList {
	whiteCastling := []CastlingRights{{false, false}}
	if whiteRooks >= 1 {
		whiteCastling = append(whiteCastling, CastlingRights{true, false})
		whiteCastling = append(whiteCastling, CastlingRights{false, true})
	}
	if whiteRooks >= 2 {
		whiteCastling = append(whiteCastling, CastlingRights{true, true})
	}

	blackCastling := []CastlingRights{{false, false}}
	if blackRooks >= 1 {
		blackCastling = append(blackCastling, CastlingRights{true, false})
		blackCastling = append(blackCastling, CastlingRights{false, true})
	}
	if blackRooks >= 2 {
		blackCastling = append(blackCastling, CastlingRights{true, true})
	}

	possibleEnPassant := square.Files{square.InvalidFile}
	if whitePawns*blackPawns > 0 {
		possibleEnPassant = square.AllFiles
	}

	return CombineRights(whiteCastling, blackCastling, possibleEnPassant)
}

func CombineRights(whiteCastling []CastlingRights, blackCastling []CastlingRights, possibleEnPassant square.Files) RightsList {
	var rightsList RightsList
	for _, wc := range whiteCastling {
		for _, bc := range blackCastling {
			for _, ep := range possibleEnPassant {
				rightsList = append(rightsList, Rights{
					WhiteCastling: wc,
					BlackCastling: bc,
					EnPassantFile: ep,
				})
			}
		}
	}
	return rightsList
}

func (r *Rights) Sprint() string {
	runes := []rune(rightsTemplate)
	if r.WhiteCastling.Kingside {
		runes[WhiteKingsideCastlingRightRuneIndex] = 'w'
	}
	if r.WhiteCastling.Queenside {
		runes[WhiteQueensideCastlingRightRuneIndex] = 'w'
	}
	if r.BlackCastling.Kingside {
		runes[BlackKingsideCastlingRightRuneIndex] = 'b'
	}
	if r.BlackCastling.Queenside {
		runes[BlackQueensideCastlingRightRuneIndex] = 'b'
	}
	runes[EnPassantRightFileRuneIndex] = r.EnPassantFile.Rune()
	return string(runes)
}
