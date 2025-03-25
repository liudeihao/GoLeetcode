package main

import "fmt"

func findSubstring(s string, words []string) []int {
	nWords, wordLen := len(words), len(words[0])
	subLen := nWords * wordLen
	wordsNeeded := make(map[string]int, nWords)
	for _, word := range words {
		wordsNeeded[word]++
	}

	var ans []int
	for startIdx := range wordLen {
		wordsCount := make(map[string]int, nWords)
		start, end := startIdx, startIdx
		curLen := 0
		totalNeeded := 0
		for end+wordLen <= len(s) {
			if curLen == subLen {
				if totalNeeded == len(wordsNeeded) {
					ans = append(ans, start)
				}
				curLen -= wordLen
				leftS := s[start : start+wordLen]
				if wordsCount[leftS] == wordsNeeded[leftS] {
					totalNeeded--
				}
				wordsCount[leftS]--
				start += wordLen
			}

			rightS := s[end : end+wordLen]
			if _, ok := wordsNeeded[rightS]; !ok {
				start, end = end+wordLen, end+wordLen
				curLen = 0
				totalNeeded = 0
				wordsCount = make(map[string]int, nWords)
				continue
			}

			wordsCount[rightS]++
			if wordsCount[rightS] == wordsNeeded[rightS] {
				totalNeeded++
			}
			end += wordLen
			curLen += wordLen
			fmt.Println(s[start:end])
		}
		if curLen == subLen && totalNeeded == len(wordsNeeded) {
			ans = append(ans, start)
		}
	}
	return ans
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
}
