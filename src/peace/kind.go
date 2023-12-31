package peace

const (
	LinearMoverMask   = 0b0100
	DiagonalMoverMask = 0b1000
)

type Kind uint8
type Kinds []Kind

const (
	Pawn   Kind = 0b010000
	Bishop Kind = 0b000000 | DiagonalMoverMask
	Knight Kind = 0b100000
	Rook   Kind = 0b000000 | LinearMoverMask
	Queen  Kind = 0b000000 | LinearMoverMask | DiagonalMoverMask
	King   Kind = 0b110000
)

var AllKinds = Kinds{Pawn, Bishop, Knight, Rook, Queen, King}
