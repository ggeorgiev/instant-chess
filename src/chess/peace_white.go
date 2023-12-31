package chess

func (p Peace) IsWhite() bool {
	return p&White == White
}

func (p Peace) IsEmptyOrWhite() bool {
	return uint8(p)&Black == 0
}

func (p Peace) IsWhiteLinearMover() bool {
	return uint8(p)&(LinearMoverMask|ColorMask) == LinearMoverMask|White
}
