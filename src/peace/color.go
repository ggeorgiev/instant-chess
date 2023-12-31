package peace

const (
	ColorMask = 0b11
	White     = 0b01
	Black     = 0b10
)

type Color uint8

const (
	WhiteColor Color = White
	BlackColor Color = Black
)

func (c Color) Oponent() Color {
	return c ^ ColorMask
}
