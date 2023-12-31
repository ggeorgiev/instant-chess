package peace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKindUniqness(t *testing.T) {
	kinds := make(map[Kind]struct{})
	for _, kind := range AllKinds {
		kinds[kind] = struct{}{}
	}
	assert.Equal(t, len(kinds), len(AllKinds))
}

func TestKindNoConflictWithColor(t *testing.T) {
	for _, kind := range AllKinds {
		assert.Equal(t, 0, int(kind&ColorMask))
	}
}
