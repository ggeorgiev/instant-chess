package board

import "log"

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

func MustParseState(text string) State {
	state, err := ParseState(text)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return state
}

func (s State) Sprint() string {
	return s.Matrix.Sprint() + s.Rights.Sprint()
}
