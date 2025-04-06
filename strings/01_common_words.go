package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Given a string paragraph and a string array of the banned words banned, return the most
frequent word that is not banned. It is guaranteed there is at least one word that is not banned, and that the answer is unique.

The words in paragraph are case-insensitive and the answer should be returned in lowercase
https://leetcode.com/problems/most-common-word/description/
*/

func mostCommonWord(paragraph string, banned []string) string {
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	bannedWords := make(map[string]int)
	for _, word := range banned {
		bannedWords[strings.ToLower(word)]++
	}
	m := make(map[string]int)
	maxFreq := 0
	mcw := ""
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		if _, ok := bannedWords[word]; !ok {
			if _, ok := m[word]; ok {
				m[word]++
			} else {
				m[word] = 1
			}
			if m[word] > maxFreq {
				maxFreq = m[word]
				mcw = word
			}
		}
	}
	return mcw
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))

	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))

}
