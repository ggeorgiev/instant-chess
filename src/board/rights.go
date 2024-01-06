package board

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

type Rights struct {
	WhiteKingsideCastlingRight  bool
	WhiteQueensideCastlingRight bool
	BlackKingsideCastlingRight  bool
	BlackQueensideCastlingRight bool
	EnPassantRightFile          square.File
}

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
		rights.WhiteKingsideCastlingRight = true
	}
	if runes[WhiteQueensideCastlingRightRuneIndex] == 'w' {
		rights.WhiteQueensideCastlingRight = true
	}
	if runes[BlackKingsideCastlingRightRuneIndex] == 'b' {
		rights.BlackKingsideCastlingRight = true
	}
	if runes[BlackQueensideCastlingRightRuneIndex] == 'b' {
		rights.BlackQueensideCastlingRight = true
	}
	rights.EnPassantRightFile = square.FileFromRune(runes[EnPassantRightFileRuneIndex])

	return &rights, nil
}

func (r *Rights) Sprint() string {
	runes := []rune(rightsTemplate)
	if r.WhiteKingsideCastlingRight {
		runes[WhiteKingsideCastlingRightRuneIndex] = 'w'
	}
	if r.WhiteQueensideCastlingRight {
		runes[WhiteQueensideCastlingRightRuneIndex] = 'w'
	}
	if r.BlackKingsideCastlingRight {
		runes[BlackKingsideCastlingRightRuneIndex] = 'b'
	}
	if r.BlackQueensideCastlingRight {
		runes[BlackQueensideCastlingRightRuneIndex] = 'b'
	}
	runes[EnPassantRightFileRuneIndex] = r.EnPassantRightFile.Rune()
	return string(runes)
}
