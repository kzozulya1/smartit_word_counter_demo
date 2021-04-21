package analyzer

import (
	"sort"
)

// rankMapStringInt sorts map[string]int by values, return indexes
func rankMapStringIntDesc(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value

	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
