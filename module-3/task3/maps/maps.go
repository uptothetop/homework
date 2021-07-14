package maps

import (
	"sort"
	"strings"
)

func SortedValues(m map[int]string) string {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	vals := make([]string, 0)
	for _, v := range keys {
		vals = append(vals, m[v])
	}

	return strings.Join(vals[:], "")
}
