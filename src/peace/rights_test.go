package peace

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestRights(t *testing.T) {
	rf := rightsInternalHelper()

	sb := &strings.Builder{}
	for _, codes := range rf {
		var indexes []int
		for s := range codes {
			indexes = append(indexes, int(s))
		}
		sort.Ints(indexes)

		sb.WriteString("{")

		for _, i := range indexes {
			sb.WriteString(fmt.Sprintf("%02d: ", i))
			switch codes[square.Index(i)] {
			case wk:
				sb.WriteString("wk")
			case wr:
				sb.WriteString("wr")
			case wp:
				sb.WriteString("wp")
			case bk:
				sb.WriteString("bk")
			case br:
				sb.WriteString("br")
			case bp:
				sb.WriteString("bp")
			}
			sb.WriteString(", ")
		}
		sb.WriteString("},\n")
	}

	assert.Equal(t, rightsMapList, rf, sb.String())
}
