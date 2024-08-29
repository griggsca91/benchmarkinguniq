package uniq

import (
	"os"
	"testing"
)

var (
	testContent      = getTestFileContent("./test.txt")
	countriesContent = getTestFileContent("./countries.txt")
	bigContent       = getTestFileContent("./big.txt")
)

func getTestFileContent(fileName string) []byte {
	testContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return testContent
}

func TestUniq(t *testing.T) {
	expectedContent := `line1
line2
line3
line4
`
	if content := Uniq(testContent); content != expectedContent {
		t.Errorf("Got %v but expected %v", content, expectedContent)
	}
}

var sink string

func BenchmarkUniq(t *testing.B) {
	var result string
	for i := 0; i < t.N; i++ {
		result = Uniq(bigContent)
	}
	sink = result
}
