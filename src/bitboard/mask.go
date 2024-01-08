package bitboard

import (
	"fmt"
)

type Mask uint64
type Masks []Mask
type MasksList []Masks

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

func (masksList MasksList) String() string {
	result := ""
	for _, masks := range masksList {
		result += fmt.Sprintf("{\n%s},\n", masks.String())
	}
	return result
}
