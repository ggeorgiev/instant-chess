package move

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRights(t *testing.T) {
	texts := []string{
		"· O-O: --, O-O-O: --, En Passant: - ·\n",
		"· O-O: w-, O-O-O: --, En Passant: - ·\n",
		"· O-O: -b, O-O-O: --, En Passant: - ·\n",
		"· O-O: --, O-O-O: w-, En Passant: - ·\n",
		"· O-O: --, O-O-O: -b, En Passant: - ·\n",
		"· O-O: --, O-O-O: --, En Passant: A ·\n",
		"· O-O: --, O-O-O: --, En Passant: H ·\n",
	}

	for _, text := range texts {
		rights, err := ParseRights("blah" + text + "blah")
		assert.NoError(t, err, text)

		result := rights.Sprint()
		assert.Equal(t, result, text, text)
	}
}

func TestGenerateRightsList(t *testing.T) {
	rightsList := GenerateRightsList(0, 0, 0, 0)
	assert.Equal(t, 1, len(rightsList))

	rightsList = GenerateRightsList(1, 0, 0, 0)
	assert.Equal(t, 3, len(rightsList))

	rightsList = GenerateRightsList(2, 0, 0, 0)
	assert.Equal(t, 4, len(rightsList))

	rightsList = GenerateRightsList(1, 1, 0, 0)
	assert.Equal(t, 9, len(rightsList))

	rightsList = GenerateRightsList(2, 1, 0, 0)
	assert.Equal(t, 12, len(rightsList))

	rightsList = GenerateRightsList(2, 2, 0, 0)
	assert.Equal(t, 16, len(rightsList))

	rightsList = GenerateRightsList(2, 2, 1, 1)
	assert.Equal(t, 144, len(rightsList))

	for _, rights := range rightsList {
		index := rights.Index()
		assert.Less(t, index, uint8(144))
	}
}

func TestWhiteKingsideCastling(t *testing.T) {
	rights := NoRights
	assert.False(t, rights.IsWhiteKingsideCastling())
	rights = rights.SetWhiteKingsideCastling()
	assert.True(t, rights.IsWhiteKingsideCastling())
}

func TestWhiteQueensideCastling(t *testing.T) {
	rights := NoRights
	assert.False(t, rights.IsWhiteQueensideCastling())
	rights = rights.SetWhiteQueensideCastling()
	assert.True(t, rights.IsWhiteQueensideCastling())
}

func TestBlackKingsideCastling(t *testing.T) {
	rights := NoRights
	assert.False(t, rights.IsBlackKingsideCastling())
	rights = rights.SetBlackKingsideCastling()
	assert.True(t, rights.IsBlackKingsideCastling())
}

func TestBlackQueensideCastling(t *testing.T) {
	rights := NoRights
	assert.False(t, rights.IsBlackQueensideCastling())
	rights = rights.SetBlackQueensideCastling()
	assert.True(t, rights.IsBlackQueensideCastling())
}
