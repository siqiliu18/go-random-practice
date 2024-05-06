package lc349decodestring

import "testing"

// TestDecodeString tests the decodeString function to ensure it returns the correct decoded string.
func TestDecodeString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "string with no digits",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "case 1",
			input:    "3[a]2[bc]",
			expected: "aaabcbc",
		},
		{
			name:     "case 2",
			input:    "3[ab2[cd]]",
			expected: "abcdcdabcdcdabcdcd",
		},
		{
			name:     "case 3",
			input:    "3[ab2[cd]e]",
			expected: "abcdcdeabcdcdeabcdcde",
		},
		{
			name:     "case 3",
			input:    "ok3[ab2[cd]e]fg",
			expected: "okabcdcdeabcdcdeabcdcdefg",
		},
		{
			name:     "case 4",
			input:    "100[leetcode]",
			expected: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := decodeString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, result)
			}
		})
	}
}
