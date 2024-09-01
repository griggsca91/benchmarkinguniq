package uniq

import (
	"hash/fnv"
	"strings"
	"unsafe"
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
	cache := make(map[string]int)

	var position int
	var startPosition int
	for i, b := range content {
		if b != '\n' {
			continue
		}
		word := content[startPosition:i]
		l := *(*string)(unsafe.Pointer(&word))
		if _, ok := cache[l]; !ok {
			cache[l] = position
			position++
		}
		startPosition = i + 1
	}

	output := make([]string, len(cache))
	for key, value := range cache {
		output[value] = key
	}

	return strings.Join(output, "\n")
}

// hash could be incorrect, but we have a wide range of values so i'm assuming it's all good
func HashFnv1a(data []byte) uint32 {
	h := fnv.New32a()
	_, _ = h.Write(data)
	return h.Sum32()
}

func UniqV5(content []byte) string {
	// average size of a word according to google is 4.7 characters long, and assuming all ascii, 4.7 * 8 = 37.6 which is ~= 40
	// 25 is faster, but uses more memory
	// 40 uses less memory but small trade off for perf, which is ok
	// 100 uses less memory even more, but speed gains start to go down noticiable, maybe worth making this and env variable if we care that much
	cache := make(map[uint32]struct{}, len(content)/40)

	var startPosition int

	str := make([]byte, 0, len(content)/10)

	for i, b := range content {
		if b != '\n' {
			continue
		}
		lineBytes := content[startPosition:i]
		var hash uint32 = 2166136261
		for _, c := range lineBytes {
			hash ^= uint32(c)
			hash *= 16777619
		}

		if _, ok := cache[hash]; !ok {
			cache[hash] = struct{}{}
			str = append(str, lineBytes...)
			str = append(str, '\n')
		}
		startPosition = i + 1
	}

	return unsafe.String(unsafe.SliceData(str), len(str))
}
