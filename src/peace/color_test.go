package peace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorOponent(t *testing.T) {
	assert.Equal(t, WhiteColor.Oponent(), BlackColor)
	assert.Equal(t, BlackColor.Oponent(), WhiteColor)
}
