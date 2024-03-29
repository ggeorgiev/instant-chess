package matrix

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestAttackBitboardMaskFromNonLinear(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   | ♘ |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   | ♙ |   |·5
··+---+---+---+---+---+---+---+---+··
4·| ♘ |   |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   | ♙ |·2
··+---+---+---+---+---+---+---+---+··
1·| ♚ |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`

	expected := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   | ● |   | ● |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   | ● |   |   |   | ● |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   | ● |   |   |   | ● |   | ● |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ● | ● | ● |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   | ● |   | ● | ● |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   | ● | ● | ● |   | ● |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ● |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`

	matrix, err := Parse(text)
	assert.NoError(t, err)

	mask := "\n" + square.SprintMask(matrix.AttackBitboardMaskFromWhite())

	assert.Equal(t, expected, mask, mask)
}

func TestAttackBitboardMaskFromDiagonals(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   | ♞ |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   | ♞ |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   | ♗ |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   | ♞ |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   | ♞ |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♚ |   |   | ♔ |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`

	expected := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   | ● |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   | ● |   | ● |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   | ● |   | ● |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   | ● |   |   |   | ● |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   | ● | ● | ● |   | ● |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   | ● |   | ● |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`

	matrix, err := Parse(text)
	assert.NoError(t, err)

	mask := "\n" + square.SprintMask(matrix.AttackBitboardMaskFromWhite())

	assert.Equal(t, expected, mask, mask)
}

func TestAttackBitboardMaskFromLinears(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   | ♞ |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ♞ | ♖ |   |   | ♞ |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   | ♞ |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♔ |   |   |   |   |   |   | ♚ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`

	expected := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   | ● |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   | ● |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ● |   | ● | ● | ● |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   | ● |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   | ● |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·| ● | ● |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   | ● |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`

	matrix, err := Parse(text)
	assert.NoError(t, err)

	mask := "\n" + square.SprintMask(matrix.AttackBitboardMaskFromWhite())
	assert.Equal(t, expected, mask, mask)
}

func TestAttackBitboardMaskFromLinearsAndDiagonals(t *testing.T) {
	text := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   | ♞ |   | ♞ |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ♞ | ♕ |   |   | ♚ |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   | ♞ |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   |   |   |   | ♞ |   |·2
··+---+---+---+---+---+---+---+---+··
1·| ♔ |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`

	expected := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   | ● |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   | ● |   |   |   | ● |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   | ● | ● | ● |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   | ● |   | ● | ● | ● |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   | ● | ● | ● |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   | ● |   | ● |   | ● |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·| ● | ● |   |   |   |   | ● |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   | ● |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`

	matrix, err := Parse(text)
	assert.NoError(t, err)

	mask := "\n" + square.SprintMask(matrix.AttackBitboardMaskFromWhite())

	assert.Equal(t, expected, mask, mask)
}
