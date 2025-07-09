package main

import (
	"fmt"
	"sort"
	"strings"
)

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
}