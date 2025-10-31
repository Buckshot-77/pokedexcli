package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "!!aNd tHen   ,teEth     ",
			expected: []string{"!!and", "then", ",teeth"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if actualLength != expectedLength {
			t.Errorf("expected length is %d and slice has length %d", expectedLength, actualLength)
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word '%s' does not match expected '%s'", word, expectedWord)
				t.Fail()
			}
		}
	}
}
