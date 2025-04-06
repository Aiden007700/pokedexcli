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
			input:    "hi",
			expected: []string{"hi"},
		},
		{
			input:    "hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hEllo  World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  helloworld  ",
			expected: []string{"helloworld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(c.input) > 0 && len(actual) == 0 {
			t.Errorf("actual [%v] is empty for c.input[%v]", actual, c.input)
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("word [%v] does not match expectedWord [%v]", word, expectedWord)
				t.Fail()
			}
		}
	}
}
