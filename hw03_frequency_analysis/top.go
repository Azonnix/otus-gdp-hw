package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(inputStr string) []string {
	if inputStr == "" {
		return []string{}
	}

	words := strings.Fields(inputStr)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}

	counts := make([]int, len(wordCountMap), len(wordCountMap))
	itr := 0

	for _, val := range wordCountMap {
		counts[itr] = val
		itr++
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	result := []string{}
	countWordsResult := 10
	counts = counts[:countWordsResult]
	countWordsMap := make(map[int][]string)

	for _, c := range counts {
		if _, ok := countWordsMap[c]; ok == true {
			continue
		}
		countWordsMap[c] = getMapKeys(c, wordCountMap)
		sort.Strings(countWordsMap[c])
		result = append(result, countWordsMap[c]...)
	}

	return result
}

func getMapKeys(val int, inputMap map[string]int) []string {
	result := []string{}

	for k, v := range inputMap {
		if v == val {
			result = append(result, k)
		}
	}

	return result
}
