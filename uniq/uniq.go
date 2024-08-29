package uniq

import "strings"

func Uniq(content []byte) string {
	cache := make(map[string]bool)

	var output []string

	for _, l := range strings.Split(string(content), "\n") {
		if _, ok := cache[l]; !ok {
			output = append(output, l)
			cache[l] = true
		}
	}

	return strings.Join(output, "\n")
}
