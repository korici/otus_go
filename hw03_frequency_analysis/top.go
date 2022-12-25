package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type TypeWord struct {
	Word  string
	Count int
}

func Top10(text string) []string {
	wordsM := map[string]int{}

	allWords := strings.Fields(text)
	allWords2 := []string{}

	for _, w := range allWords {
		s := strings.Split(w, ".")
		allWords2 = append(allWords2, s...)
	}

	allWords = nil

	for _, w := range allWords2 {
		s := strings.Split(w, ",")
		allWords = append(allWords, s...)
	}

	allWords2 = nil

	for _, w := range allWords {
		s := strings.Split(w, "!")
		allWords2 = append(allWords2, s...)
	}

	allWords = nil

	for _, w := range allWords2 {
		s := strings.Split(w, "?")
		allWords = append(allWords, s...)
	}

	allWords2 = nil

	for _, w := range allWords {
		s := strings.Split(w, ":")
		allWords2 = append(allWords2, s...)
	}

	allWords = nil

	for _, w := range allWords2 {
		s := strings.Split(w, ";")
		allWords = append(allWords, s...)
	}

	allWords2 = nil

	for _, w := range allWords {
		s := strings.Split(w, "\"")
		allWords2 = append(allWords2, s...)
	}

	if len(allWords) == 0 {
		return nil
	}

	for _, s := range allWords {
		if s != "-" && s != "" {
			i := wordsM[strings.ToLower(s)]
			wordsM[strings.ToLower(s)] = i + 1
		}
	}

	if len(wordsM) == 0 {
		return nil
	}

	wordForSort := []TypeWord{}
	for key, value := range wordsM {
		w := TypeWord{
			Word:  key,
			Count: value,
		}
		wordForSort = append(wordForSort, w)
		fmt.Println(w.Word, w.Count)
	}

	sort.Slice(wordForSort, func(i, j int) bool {
		if wordForSort[i].Count == wordForSort[j].Count {
			return wordForSort[i].Word < wordForSort[j].Word
		}
		return wordForSort[i].Count > wordForSort[j].Count
	})

	words := []string{}
	l := 10
	if len(wordForSort) < l {
		l = len(wordForSort)
	}

	for i := 0; i < l; i++ {
		words = append(words, wordForSort[i].Word)
	}

	return words
}
