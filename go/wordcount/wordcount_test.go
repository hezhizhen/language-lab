package wordcount

import (
	"maps"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "empty string",
			text: "",
			want: map[string]int{},
		},
		{
			name: "only separators",
			text: " ,.!? \t\n",
			want: map[string]int{},
		},
		{
			name: "single word",
			text: "word",
			want: map[string]int{"word": 1},
		},
		{
			name: "multiple distinct words",
			text: "one of each",
			want: map[string]int{"one": 1, "of": 1, "each": 1},
		},
		{
			name: "repeated words",
			text: "one fish two fish red fish blue fish",
			want: map[string]int{"one": 1, "fish": 4, "two": 1, "red": 1, "blue": 1},
		},
		{
			name: "case insensitive",
			text: "Go GO go gO",
			want: map[string]int{"go": 4},
		},
		{
			name: "punctuation and whitespace as separators",
			text: "car: carpet as java : javascript!!&@$%^&",
			want: map[string]int{"car": 1, "carpet": 1, "as": 1, "java": 1, "javascript": 1},
		},
		{
			name: "multiple consecutive separators",
			text: "  hello,,,   world\n\n\thello  ",
			want: map[string]int{"hello": 2, "world": 1},
		},
		{
			name: "digits are separators",
			text: "abc123def 456",
			want: map[string]int{"abc": 1, "def": 1},
		},
		{
			name: "apostrophe splits word",
			text: "don't",
			want: map[string]int{"don": 1, "t": 1},
		},
		{
			name: "non-ASCII letters are separators",
			text: "café 你好 naïve",
			want: map[string]int{"caf": 1, "na": 1, "ve": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.text)
			if !maps.Equal(got, tt.want) {
				t.Errorf("WordCount(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}
