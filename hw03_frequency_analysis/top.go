package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	freq := make(map[string]int, len(words))
	for _, word := range words {
		// Base homework contract counts exact tokens as-is: no normalization.
		freq[word]++
	}

	result := make([]string, 0, len(freq))
	for word := range freq {
		result = append(result, word)
	}

	sort.Slice(result, func(i, j int) bool {
		if freq[result[i]] == freq[result[j]] {
			return result[i] < result[j]
		}

		return freq[result[i]] > freq[result[j]]
	})

	if len(result) > 10 {
		result = result[:10]
	}

	return result
}
