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

type Rights uint8
type RightsList []Rights

const (
	EnPassantFileActiveMask    = uint8(0b00001000)
	EnPassantFileValueMask     = uint8(0b00000111)
	EnPassantFileMask          = EnPassantFileActiveMask | EnPassantFileValueMask
	WhiteKingsideCastlingMask  = uint8(0b00010000)
	WhiteQueensideCastlingMask = uint8(0b00100000)
	WhiteBothCastlingMask      = WhiteQueensideCastlingMask | WhiteKingsideCastlingMask
	BlackKingsideCastlingMask  = uint8(0b01000000)
	BlackQueensideCastlingMask = uint8(0b10000000)
	BlackBothCastlingMask      = BlackQueensideCastlingMask | BlackKingsideCastlingMask
)

const (
	NoRights                     = Rights(0)
	WhiteKingsideCastlingRights  = Rights(WhiteKingsideCastlingMask)
	WhiteQueensideCastlingRights = Rights(WhiteQueensideCastlingMask)
	WhiteBothCastlingRights      = WhiteKingsideCastlingRights | WhiteQueensideCastlingRights
	BlackKingsideCastlingRights  = Rights(BlackKingsideCastlingMask)
	BlackQueensideCastlingRights = Rights(BlackQueensideCastlingMask)
	BlackBothCastlingRights      = BlackKingsideCastlingRights | BlackQueensideCastlingRights
	AllCastlingRights            = WhiteBothCastlingRights | BlackBothCastlingRights
)

var (
	NoEnPassantFile        = NoRights.SetEnPassantFile(square.InvalidFile)
	AllEnPassantFileRights = []Rights{
		NoEnPassantFile,
		NoRights.SetEnPassantFile(square.FileA),
		NoRights.SetEnPassantFile(square.FileB),
		NoRights.SetEnPassantFile(square.FileC),
		NoRights.SetEnPassantFile(square.FileD),
		NoRights.SetEnPassantFile(square.FileE),
		NoRights.SetEnPassantFile(square.FileF),
		NoRights.SetEnPassantFile(square.FileG),
		NoRights.SetEnPassantFile(square.FileH),
	}
)

func (r Rights) Index() uint8 {
	index := uint8(r) ^ EnPassantFileActiveMask
	return index<<4 | index>>4
}

func RightsFromIndex(index uint8) Rights {
	index = index<<4 | index>>4
	return Rights(index ^ EnPassantFileMask)
}

func (r Rights) SetWhiteKingsideCastling() Rights {
	return Rights(uint8(r) | WhiteKingsideCastlingMask)
}

func (r Rights) ResetWhiteKingsideCastling() Rights {
	return Rights(uint8(r) & ^WhiteKingsideCastlingMask)
}

func (r Rights) IsWhiteKingsideCastling() bool {
	return uint8(r)&WhiteKingsideCastlingMask != 0
}

func (r Rights) SetWhiteQueensideCastling() Rights {
	return Rights(uint8(r) | WhiteQueensideCastlingMask)
}

func (r Rights) ResetWhiteQueensideCastling() Rights {
	return Rights(uint8(r) & ^WhiteQueensideCastlingMask)
}

func (r Rights) IsWhiteQueensideCastling() bool {
	return uint8(r)&WhiteQueensideCastlingMask != 0
}

func (r Rights) SetWhiteBothCastling() Rights {
	return Rights(uint8(r) | WhiteBothCastlingMask)
}

func (r Rights) ResetWhiteBothCastling() Rights {
	return Rights(uint8(r) & ^WhiteBothCastlingMask)
}

func (r Rights) IsWhiteAnyCastling() bool {
	return uint8(r)&(WhiteBothCastlingMask) != 0
}

func (r Rights) SetBlackKingsideCastling() Rights {
	return Rights(uint8(r) | BlackKingsideCastlingMask)
}

func (r Rights) ResetBlackKingsideCastling() Rights {
	return Rights(uint8(r) & ^BlackKingsideCastlingMask)
}

func (r Rights) IsBlackKingsideCastling() bool {
	return uint8(r)&BlackKingsideCastlingMask != 0
}

func (r Rights) SetBlackQueensideCastling() Rights {
	return Rights(uint8(r) | BlackQueensideCastlingMask)
}

func (r Rights) ResetBlackQueensideCastling() Rights {
	return Rights(uint8(r) & ^BlackQueensideCastlingMask)
}

func (r Rights) IsBlackQueensideCastling() bool {
	return uint8(r)&BlackQueensideCastlingMask != 0
}

func (r Rights) SetBlackBothCastling() Rights {
	return Rights(uint8(r) | BlackBothCastlingMask)
}

func (r Rights) ResetBlackBothCastling() Rights {
	return Rights(uint8(r) & ^BlackBothCastlingMask)
}

func (r Rights) IsBlackAnyCastling() bool {
	return uint8(r)&(BlackBothCastlingMask) != 0
}

func (r Rights) SetEnPassantFile(file square.File) Rights {
	if file == square.InvalidFile {
		return Rights(uint8(r) & ^EnPassantFileMask)
	}
	return Rights((uint8(r) & ^EnPassantFileMask) | uint8(file) | EnPassantFileActiveMask)
}

func (r Rights) EnPassantFile() square.File {
	if uint8(r)&EnPassantFileActiveMask == 0 {
		return square.InvalidFile
	}
	return square.File(uint8(r) & EnPassantFileValueMask)
}

func ParseRights(text string) (Rights, error) {
	match := rightsPattern.FindAllString(text, -1)
	if len(match) == 0 {
		return NoRights, fmt.Errorf("failed to find the rights description")
	}
	if len(match) > 1 {
		return NoRights, fmt.Errorf("found to many rights descriptions")
	}

	runes := util.Runes(match[0])

	rights := NoRights

	if runes[WhiteKingsideCastlingRightRuneIndex] == 'w' {
		rights = rights.SetWhiteKingsideCastling()
	}
	if runes[WhiteQueensideCastlingRightRuneIndex] == 'w' {
		rights = rights.SetWhiteQueensideCastling()
	}
	if runes[BlackKingsideCastlingRightRuneIndex] == 'b' {
		rights = rights.SetBlackKingsideCastling()
	}
	if runes[BlackQueensideCastlingRightRuneIndex] == 'b' {
		rights = rights.SetBlackQueensideCastling()
	}
	rights = rights.SetEnPassantFile(square.FileFromRune(runes[EnPassantRightFileRuneIndex]))
	return rights, nil
}

func GenerateRightsList(whiteRooks int, blackRooks int, whitePawns int, blackPawns int) RightsList {
	whiteCastling := []Rights{NoRights}
	if whiteRooks >= 1 {
		whiteCastling = append(whiteCastling, WhiteKingsideCastlingRights, WhiteQueensideCastlingRights)
		if whiteRooks >= 2 {
			whiteCastling = append(whiteCastling, WhiteBothCastlingRights)
		}
	}

	blackCastling := []Rights{NoRights}
	if blackRooks >= 1 {
		blackCastling = append(blackCastling, BlackKingsideCastlingRights, BlackQueensideCastlingRights)
		if blackRooks >= 2 {
			blackCastling = append(blackCastling, BlackBothCastlingRights)
		}
	}

	enPassantFileRightsList := AllEnPassantFileRights
	if whitePawns == 0 || blackPawns == 0 {
		enPassantFileRightsList = []Rights{NoEnPassantFile}
	}

	return CombineRights(whiteCastling, blackCastling, enPassantFileRightsList)
}

func CombineRights(whiteCastling []Rights, blackCastling []Rights, possibleEnPassant []Rights) RightsList {
	var rightsList RightsList
	for _, wc := range whiteCastling {
		for _, bc := range blackCastling {
			for _, ep := range possibleEnPassant {
				rightsList = append(rightsList, Rights(uint8(wc)|uint8(bc)|uint8(ep)))
			}
		}
	}
	return rightsList
}

func (r Rights) Sprint() string {
	runes := []rune(rightsTemplate)
	if r.IsWhiteKingsideCastling() {
		runes[WhiteKingsideCastlingRightRuneIndex] = 'w'
	}
	if r.IsWhiteQueensideCastling() {
		runes[WhiteQueensideCastlingRightRuneIndex] = 'w'
	}
	if r.IsBlackKingsideCastling() {
		runes[BlackKingsideCastlingRightRuneIndex] = 'b'
	}
	if r.IsBlackQueensideCastling() {
		runes[BlackQueensideCastlingRightRuneIndex] = 'b'
	}
	runes[EnPassantRightFileRuneIndex] = r.EnPassantFile().Rune()
	return string(runes)
}
