package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "test case 1",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "test case 2",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "test case 3",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("singleNumber() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "positive palindrome",
			num:      121,
			expected: true,
		},
		{
			name:     "negative number",
			num:      -121,
			expected: false,
		},
		{
			name:     "single digit",
			num:      5,
			expected: true,
		},
		{
			name:     "non-palindrome",
			num:      123,
			expected: false,
		},
		{
			name:     "zero",
			num:      0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.num)
			if result != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.num, result, tt.expected)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid simple",
			input:    "()",
			expected: true,
		},
		{
			name:     "valid multiple",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "valid nested",
			input:    "{[]}",
			expected: true,
		},
		{
			name:     "invalid simple",
			input:    "(]",
			expected: false,
		},
		{
			name:     "invalid order",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single bracket",
			input:    "(",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.input)
			if result != tt.expected {
				t.Errorf("isValid(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
} 

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"common prefix", []string{"flower", "flow", "flight"}, "fl"},
		{"no common prefix", []string{"dog", "racecar", "car"}, ""},
		{"all same", []string{"test", "test", "test"}, "test"},
		{"single string", []string{"alone"}, "alone"},
		{"empty array", []string{}, ""},
		{"prefix is empty string", []string{"", "abc"}, ""},
		{"partial match", []string{"abc", "ab", "a"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
