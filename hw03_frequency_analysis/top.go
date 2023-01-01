package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsM := map[string]int{}

	allWords := strings.Fields(text)
	allWords2 := []string{}

	f := func(c rune) bool {
		return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' || c == '"'
	}

	for _, w := range allWords {
		s := strings.FieldsFunc(w, f)
		allWords2 = append(allWords2, s...)
	}

	if len(allWords2) == 0 {
		return nil
	}

	for _, s := range allWords2 {
		if s != "-" && s != "" {
			wordsM[strings.ToLower(s)]++
		}
	}

	if len(wordsM) == 0 {
		return nil
	}

	words := make([]string, 0, len(wordsM))

	for key := range wordsM {
		words = append(words, key)
	}

	sort.Slice(words, func(i, j int) bool {
		if wordsM[words[i]] == wordsM[words[j]] {
			return words[i] < words[j]
		}
		return wordsM[words[i]] > wordsM[words[j]]
	})

	countOut := 10
	if len(words) < countOut {
		countOut = len(words)
	}

	return words[:countOut]
}
