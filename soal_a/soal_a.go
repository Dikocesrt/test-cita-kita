package main

import (
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

}