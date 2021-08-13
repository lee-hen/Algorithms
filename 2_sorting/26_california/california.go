package main

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"sort"
	"strings"
)

const CANDIDATE_ORDER = "RWQOJMVAHBSGZXNTCIEKUPDYFL"

func main() {
	candidates := []string{
		"LEONARD PADILLA",
		"LEO GALLAGHER",
		"LINGEL H. WINTERS",
		"LAWRENCE STEVEN STRAUSS",
		"WILLIAM \"BILL\" S. CHAMBERS",
		"WARREN FARRELL",
		"REVA RENEE RENZ",
		"RICHARD J. SIMMONS",
		"RICH GOSSE",
		"RALPH A. HERNANDEZ",
		"RANDALL D. SPRAGUE",
		"RONALD JASON PALMIERI",
		"RONALD J. FRIEDMAN",
		"ROBERT CULLENBINE",
		"ROBERT C. NEWMAN II",
		"ROBERT C. MANNHEIM",
		"ROBERT \"BUTCH\" DOLE",
	}

	order := CANDIDATE_ORDER

	sort.Slice(candidates, func(i, j int) bool {
		a, b := candidates[i], candidates[j]

		n := util.Min(len(a), len(b))

		for k := 0; k < n; k++ {
			aIdx := strings.Index(order, string(a[k]))
			bIdx := strings.Index(order, string(b[k]))

			if aIdx < bIdx {
				return true
			} else if aIdx > bIdx{
				return false
			}

		}
		return len(a) > len(b)
	})


	for i := 0; i < len(candidates); i++ {
		fmt.Println(candidates[i])
	}
}
