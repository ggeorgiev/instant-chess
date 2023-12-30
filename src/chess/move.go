package chess

import (
	"fmt"
	"strings"
)

type Square struct {
	X int
	Y int
}

// String converts a Square's coordinates to chess board notation.
func (s Square) String() string {
	// Convert the X coordinate to a letter (A-H) and the Y coordinate to a number (1-8).
	// Assuming X: 1-8 maps to A-H and Y: 1-8 is the same.
	// Adjust the ranges if your coordinates are 0-indexed or differently ranged.
	letters := "ABCDEFGH"
	if s.X >= 1 && s.X <= 8 && s.Y >= 1 && s.Y <= 8 {
		return fmt.Sprintf("%c%d", letters[s.X-1], s.Y)
	}
	// Return an error message if out of bounds
	return "Invalid Square"
}

type Answer struct {
	BlackFrom Square
	BlackTos  []Square
}

func (a Answer) String() string {
	var tos []string
	for _, t := range a.BlackTos {
		tos = append(tos, t.String())
	}
	return fmt.Sprintf("BlackFrom: %s, BlackTos: [%s]", a.BlackFrom, strings.Join(tos, ", "))
}

type ToAnswer struct {
	WhiteTo Square
	Answers []Answer
}

func (ta ToAnswer) String() string {
	var ans []string
	for _, a := range ta.Answers {
		ans = append(ans, a.String())
	}
	return fmt.Sprintf("  Answers: %s\n", strings.Join(ans, ", "))
}

type Move struct {
	WhiteForm Square
	ToAnswers []ToAnswer
}

func (m Move) String() string {
	var tas []string
	for _, ta := range m.ToAnswers {
		tas = append(tas, fmt.Sprintf("White: %s:%s\n%s", m.WhiteForm, ta.WhiteTo, ta.String()))
	}
	return strings.Join(tas, "")
}

type Moves []Move

func (ms Moves) String() string {
	var mvs []string
	for _, m := range ms {
		mvs = append(mvs, m.String())
	}
	return fmt.Sprintf("%s", strings.Join(mvs, "\n"))
}
