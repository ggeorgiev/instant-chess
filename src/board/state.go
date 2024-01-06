package board

type State struct {
	Matrix Matrix
	Rights *Rights
}

func ParseState(text string) (State, error) {
	rights, err := ParseRights(text)
	if err != nil {
		return State{}, err
	}
	matrix, err := ParseMatrix(text)
	if err != nil {
		return State{}, err
	}

	state := State{
		Matrix: matrix,
		Rights: rights,
	}
	return state, nil
}

func (s State) Sprint() string {
	return s.Matrix.Sprint() + s.Rights.Sprint()
}
