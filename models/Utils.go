package models

import "sort"

func getSortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	// values := make([]int, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
		// values = append(values, v)
	}

	sort.Strings(keys)

	return keys
}
