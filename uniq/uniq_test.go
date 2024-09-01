package uniq

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	testContent      = getTestFileContent("./test.txt")
	countriesContent = getTestFileContent("./countries.txt")
	bigContent       = getTestFileContent("./big.txt")
)

var content = bigContent

func getTestFileContent(fileName string) []byte {
	testContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return testContent
}

func length(s string) {
	// fmt.Println(s)
	fmt.Println(len(strings.Split(s, "\n")))
}

// func TestUniqV5(t *testing.T) {
// 	result := UniqV5(content)
// 	fmt.Println(result)
// 	if len(strings.Split(result, "\n")) != 1137034 {
// 		t.Errorf("got %v", len(strings.Split(result, "\n")))
// 	}
// }

var sink string

func BenchmarkUniq(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = Uniq(content)
	}
	sink = result
}

func BenchmarkUniqV2(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = UniqV2(content)
	}
	sink = result
}

func BenchmarkUniqV3(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = UniqV3(content)
	}
	sink = result
}

func BenchmarkUniqV4(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = UniqV4(content)
	}
	sink = result
}

func BenchmarkUniqV5(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = UniqV5(content)
	}
	sink = result
	// length(sink)
}
