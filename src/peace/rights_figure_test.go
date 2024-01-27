package peace

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestRightsFigures(t *testing.T) {
	rf := rightsFigureInternalHelper()

	sb := &strings.Builder{}
	for _, figures := range rf {
		var indexes []int
		for s := range figures {
			indexes = append(indexes, int(s))
		}
		sort.Ints(indexes)

		sb.WriteString("{")

		for _, i := range indexes {
			sb.WriteString(fmt.Sprintf("%02d: ", i))
			switch figures[square.Index(i)] {
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

	assert.Equal(t, rightsFigures, rf, sb.String())
}
