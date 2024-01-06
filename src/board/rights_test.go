package board

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
		"· O-O: --, O-O-O: -b, En Passant: A ·\n",
		"· O-O: --, O-O-O: -b, En Passant: H ·\n",
	}

	for _, text := range texts {
		rights, err := ParseRights("blah" + text + "blah")
		assert.NoError(t, err, text)

		result := rights.Sprint()
		assert.Equal(t, text, result, text)
	}
}
