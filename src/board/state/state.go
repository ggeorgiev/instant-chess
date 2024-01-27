package state

import (
	"log"

	"github.com/ggeorgiev/instant-chess/src/board/matrix"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/storage"
)

type State struct {
	Matrix *matrix.Matrix
	Rights move.Rights
}

func Parse(text string) (State, error) {
	rights, err := move.ParseRights(text)
	if err != nil {
		return State{}, err
	}
	matrix, err := matrix.Parse(text)
	if err != nil {
		return State{}, err
	}

	state := State{
		Matrix: &matrix,
		Rights: rights,
	}
	return state, nil
}

func MustParse(text string) State {
	state, err := Parse(text)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return state
}

func (s *State) Sprint() string {
	return s.Matrix.Sprint() + s.Rights.Sprint()
}

func (s *State) Invalid() (bool, square.Index) {
	invalid, offender := s.Matrix.Invalid()
	if invalid {
		return true, offender
	}

	whiteRookIsMisplaced :=
		(s.Rights.IsWhiteKingsideCastling() && s.Matrix[place.WhiteRookKingsideStarting] != peace.WhiteRook) ||
			(s.Rights.IsWhiteQueensideCastling() && s.Matrix[place.WhiteRookQueensideStarting] != peace.WhiteRook)

	blackRookIsMisplaced :=
		(s.Rights.IsBlackKingsideCastling() && s.Matrix[place.BlackRookKingsideStarting] != peace.BlackRook) ||
			(s.Rights.IsBlackQueensideCastling() && s.Matrix[place.BlackRookQueensideStarting] != peace.BlackRook)

	var offenders square.Indexes
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		pc := s.Matrix[i]
		if pc.IsNull() {
			continue
		}
		if pc == peace.WhiteRook {
			if whiteRookIsMisplaced {
				if i == place.WhiteRookKingsideStarting && s.Rights.IsWhiteKingsideCastling() {
					continue
				}
				if i == place.WhiteRookQueensideStarting && s.Rights.IsWhiteQueensideCastling() {
					continue
				}
				return true, i
			}
		} else if pc == peace.BlackRook {
			if blackRookIsMisplaced {
				if i == place.BlackRookKingsideStarting && s.Rights.IsBlackKingsideCastling() {
					continue
				}
				if i == place.BlackRookQueensideStarting && s.Rights.IsBlackQueensideCastling() {
					continue
				}
				return true, i
			}
		} else if pc == peace.WhiteKing {
			if i != place.WhiteKingStarting &&
				s.Rights.IsWhiteAnyCastling() {
				offenders = append(offenders, i)
			}
		} else if pc == peace.BlackKing {
			if i != place.BlackKingStarting &&
				s.Rights.IsBlackAnyCastling() {
				offenders = append(offenders, i)
			}
		}
	}
	return len(offenders) > 0, offenders.Max()
}

func (s *State) StorageLocation() (storage.BoardPeaces, storage.BoardStateIndex) {
	var peaces storage.BoardPeaces

	n := 0
	for sq := square.ZeroIndex; sq <= square.LastIndex; sq++ {
		if s.Matrix[sq] != peace.Null {
			peaces[n] = s.Matrix[sq]
			n++
		}
	}
	bitmask := s.Matrix.Mask()
	index := math.BitsetToIndex(uint64(n), uint64(bitmask))

	return peaces, storage.BoardStateIndex(index)
}
