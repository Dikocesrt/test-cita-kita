package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type TestCase struct {
    ID        int      `json:"id"`
    RootWords []string `json:"root_words"`
    Sentence  string   `json:"sentence"`
}

type TestData struct {
    TestCases []TestCase `json:"test_cases"`
}

func replaceWithRootWords(rootWords []string, sentence string) string {
	temp := make([]string, len(rootWords))

	copy(temp, rootWords)
	sort.Slice(temp, func(i, j int) bool {
		return len(temp[i]) < len(temp[j])
	})

	listWords := strings.Split(sentence, " ")
	for idx, word := range listWords {
		replaced := false
		for _, root := range temp {
			if strings.HasPrefix(word, root) {
				listWords[idx] = root
				replaced = true
				break
			}
		}
		if !replaced {
			listWords[idx] = word
		}
	}
	return strings.Join(listWords, " ")
}

func loadTestData(filename string) (*TestData, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var testData TestData
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&testData)
    if err != nil {
        return nil, err
    }

    return &testData, nil
}

func main(){
	// Soal 1
	rootWords1 := []string{"cat", "bat", "rat"}
	sentence1 := "the cattle was rattled by the battery"
	result1 := replaceWithRootWords(rootWords1, sentence1)
	fmt.Println("Soal 1 Output:")
	fmt.Println(result1)
	fmt.Println()

	// Soal 2
	rootWords2 := []string{"dog", "car", "bike"}
	sentence2 := "the dogs were barking near the cars and bikers"
	result2 := replaceWithRootWords(rootWords2, sentence2)
	fmt.Println("Soal 2 Output:")
	fmt.Println(result2)

	testData, err := loadTestData(filepath.Join("soal_a", "test.json"))
    if err != nil {
        fmt.Printf("Error loading test data: %v\n", err)
        return
    }

	// Additional test cases from JSON
	fmt.Println(strings.Repeat("-", 50))
    for _, testCase := range testData.TestCases {
        fmt.Printf("\nTest Case %d:\n", testCase.ID)
        fmt.Printf("Root Words: %v\n", testCase.RootWords)
        fmt.Printf("Original Sentence: %s\n", testCase.Sentence)
        
        result := replaceWithRootWords(testCase.RootWords, testCase.Sentence)
        fmt.Printf("Result: %s\n", result)
        fmt.Println(strings.Repeat("-", 50))
    }
}