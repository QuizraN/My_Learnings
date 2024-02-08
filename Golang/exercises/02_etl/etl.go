package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	for key, values := range in {
		for _, value := range values {
			result[strings.ToLower(value)] = key
		}
	}
	return result
}
