package main 
import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " oneword ",
			expected: []string{"oneword"},
		},
		{
			input: " Charmander BULBA PIkaCHU",
			expected: []string{"charmander", "bulba", "pikachu"},
		},
		{
			input: " sPongebob\n to the \n rescue",
			expected: []string{"spongebob", "to", "the", "rescue"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("expected length: %v actual length: %v\n", len(c.expected), len(actual))
			continue
		}
		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if expectedWord != word {
				t.Errorf("expected: %v actual: %v\n", expectedWord, word)
			}
		}
	}
}