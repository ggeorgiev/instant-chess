package chess

func (p Peace) IsBlack() bool {
	return p&Black == Black
}

func (p Peace) IsEmptyOrBlack() bool {
	return uint8(p)&White == 0
}

func (p Peace) IsBlackLinearMover() bool {
	return uint8(p)&(LinearMoverMask|ColorMask) == LinearMoverMask|Black
}

func (p Peace) IsBlackDiagonalMover() bool {
	return uint8(p)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|Black
}
