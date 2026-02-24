package hashing

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	results := [][]string{}
	grouped := make(map[string][]string, len(strs))

	for _, v := range strs {
		s := []byte(v)
		slices.Sort(s)
		sorted := string(s)
		grouped[sorted] = append(grouped[sorted], v)
	}

	for _, v := range grouped {
		results = append(results, v)
	}

	return results
}
