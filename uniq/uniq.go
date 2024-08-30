package uniq

import (
	"strings"
)

func Uniq(content []byte) string {
	cache := make(map[string]bool)

	var output []string

	for _, l := range strings.Split(string(content), "\n") {
		if len(l) == 0 {
			continue
		}
		if _, ok := cache[l]; !ok {
			output = append(output, l)
			cache[l] = true
		}
	}

	return strings.Join(output, "\n")
}

func UniqV2(content []byte) string {
	cache := make(map[string]bool)

	var output []string

	var word []byte
	for _, b := range content {
		if b != '\n' {
			word = append(word, b)
			continue
		}
		l := string(word)
		if _, ok := cache[l]; !ok {
			output = append(output, l)
			cache[l] = true
		}
		word = word[0:0]
	}

	return strings.Join(output, "\n")
}

func UniqV3(content []byte) string {
	cache := make(map[string]int)

	var word []byte
	var position int
	for _, b := range content {
		if b != '\n' {
			word = append(word, b)
			continue
		}
		l := string(word)
		if _, ok := cache[l]; !ok {
			cache[l] = position
			position++
		}
		word = word[0:0]
	}

	output := make([]string, len(cache))
	for key, value := range cache {
		output[value] = key
	}

	return strings.Join(output, "\n")
}

func UniqV4(content []byte) string {
	cache := &Node{}

	var word []byte
	var position int
	for _, b := range content {
		if b != '\n' {
			word = append(word, b)
			continue
		}
		if ok := cache.Has(word); !ok {
			cache.Insert(word, position)
			position++
		}
		word = word[0:0]
	}

	output := make([]string, position)
	for value := range cache.Nodes() {
		output[value.Position] = string(value.Value)
	}

	return strings.Join(output, "\n")
}
