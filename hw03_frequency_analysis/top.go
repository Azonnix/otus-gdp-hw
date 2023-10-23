package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(inputStr string) []string {
	if inputStr == "" {
		return []string{}
	}
	fmt.Println(inputStr)
	words := strings.Split(inputStr, " ")
	fmt.Println(words)
	countWordMap := make(map[string]int)

	for _, word := range words {
		fmt.Println(word)
		countWordMap[word]++
	}

	fmt.Println(countWordMap)

	counts := make([]int, len(countWordMap), len(countWordMap))
	itr := 0

	for _, val := range countWordMap {
		counts[itr] = val
		itr++
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	result := []string{}
	counts = counts[:10]

	for key, val := range countWordMap {
		if isInSlice(val, counts) {
			result = append(result, key)
		}
	}

	return result
}

func isInSlice(val int, sli []int) bool {
	for _, s := range sli {
		if val == s {
			return true
		}
	}

	return false
}
