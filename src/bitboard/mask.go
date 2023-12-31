package bitboard

import "fmt"

type Mask uint64
type Masks []Mask

const Empty = Mask(0)

func (mask Mask) String() string {
	return fmt.Sprintf("0b%064b", mask)
}

func (masks Masks) String() string {
	result := ""
	for _, mask := range masks {
		result += fmt.Sprintf("%s,\n", mask.String())
	}
	return result
}
