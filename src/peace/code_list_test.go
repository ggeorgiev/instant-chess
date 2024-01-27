package peace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRights(t *testing.T) {
	codes := MustParse("")
	rightsList := codes.MoveRights()
	assert.Len(t, rightsList, 1)

	codes = MustParse("♚♔♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 3)

	codes = MustParse("♚♔♖♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 4)

	codes = MustParse("♚♔♜")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 3)

	codes = MustParse("♚♔♜♜")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 4)

	codes = MustParse("♚♔♜♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 9)

	codes = MustParse("♚♔♜♜♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 12)

	codes = MustParse("♚♔♜♖♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 12)

	codes = MustParse("♚♔♜♜♖♖")
	rightsList = codes.MoveRights()
	assert.Len(t, rightsList, 16)
}
