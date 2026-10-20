package wordcount

import "strings"

func WordCount(text string) map[string]int {
	m := make(map[string]int)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z'))
	})
	for _, word := range words {
		m[strings.ToLower(word)]++
	}
	return m
}
