package chess

import "fmt"

type PeaceColor bool

const (
	PeaceColorWhite PeaceColor = true
	PeaceColorBlack PeaceColor = false
)

type PeaceType uint

const (
	Empty = iota
	Pawn
	Bishop
	Knight
	Rook
	Queen
	King
)

type Peace struct {
	Color PeaceColor
	Type  PeaceType
}

func (p Peace) Symbol() string {
	switch p.Color {
	case PeaceColorWhite:
		switch p.Type {
		case Empty:
			return ` `
		case Pawn:
			return `♙`
		case Bishop:
			return `♗`
		case Knight:
			return `♘`
		case Rook:
			return `♖`
		case Queen:
			return `♕`
		case King:
			return `♔`
		}
	case PeaceColorBlack:
		switch p.Type {
		case Empty:
			return ` `
		case Pawn:
			return `♟︎`
		case Bishop:
			return `♝`
		case Knight:
			return `♞`
		case Rook:
			return `♜`
		case Queen:
			return `♛`
		case King:
			return `♚`
		}
	}
	panic("impossible")
}

func PeaceFromSymbol(symbol rune) Peace {
	switch symbol {
	case 32:
		return Peace{Type: Empty}
	case 9820:
		return Peace{Type: Rook, Color: PeaceColorBlack}
	case 9822:
		return Peace{Type: Knight, Color: PeaceColorBlack}
	case 9821:
		return Peace{Type: Bishop, Color: PeaceColorBlack}
	case 9819:
		return Peace{Type: Queen, Color: PeaceColorBlack}
	case 9818:
		return Peace{Type: King, Color: PeaceColorBlack}
	case 9823:
		return Peace{Type: Pawn, Color: PeaceColorBlack}
	case 9817:
		return Peace{Type: Pawn, Color: PeaceColorWhite}
	case 9814:
		return Peace{Type: Rook, Color: PeaceColorWhite}
	case 9816:
		return Peace{Type: Knight, Color: PeaceColorWhite}
	case 9815:
		return Peace{Type: Bishop, Color: PeaceColorWhite}
	case 9813:
		return Peace{Type: Queen, Color: PeaceColorWhite}
	case 9812:
		return Peace{Type: King, Color: PeaceColorWhite}
	}
	panic(fmt.Sprintf("imposible: %d", symbol))
}
