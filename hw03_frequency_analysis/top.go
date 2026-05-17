package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func normalizeToken(token string) string {
	if token == "" {
		return ""
	}

	normalized := strings.ToLower(token)
	normalized = strings.TrimFunc(normalized, func(r rune) bool {
		return unicode.IsPunct(r) && r != '-'
	})

	if normalized == "" {
		return ""
	}

	if isHyphenOnly(normalized) {
		if len([]rune(normalized)) == 1 {
			return ""
		}

		return normalized
	}

	return normalized
}

func isHyphenOnly(token string) bool {
	for _, r := range token {
		if r != '-' {
			return false
		}
	}

	return token != ""
}

func Top10(text string) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	freq := make(map[string]int, len(words))
	for _, word := range words {
		normalized := normalizeToken(word)
		if normalized == "" {
			continue
		}

		freq[normalized]++
	}

	if len(freq) == 0 {
		return []string{}
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
