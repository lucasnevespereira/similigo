package similarity

import (
	"math"
	"strings"
)

/*func ContainmentSimilarity(text1, text2 string) float64 {
	keywordTokens := strings.Fields(strings.ToLower(text1))
	titleTokens := strings.Fields(strings.ToLower(text2))

	matchCount := 0.0
	for _, kToken := range keywordTokens {
		for _, tToken := range titleTokens {
			if strings.Contains(kToken, tToken) || strings.Contains(tToken, kToken) {
				matchCount++
			}
		}
	}
	return matchCount
}*/

func ContainmentSimilarity(text1, text2 string) float64 {
	set1 := make(map[string]bool)
	set2 := make(map[string]bool)

	for _, token := range strings.Fields(text1) {
		set1[token] = true
	}
	for _, token := range strings.Fields(text2) {
		set2[token] = true
	}

	intersection := 0
	for token := range set1 {
		if set2[token] {
			intersection++
		}
	}

	minLength := math.Min(float64(len(set1)), float64(len(set2)))
	if minLength == 0 {
		return 0
	}
	return float64(intersection) / minLength
}
