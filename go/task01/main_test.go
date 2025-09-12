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

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		want     []int
	}{
		{
			name:     "normal case",
			input:    []int{1, 1, 2},
			expected: 2,
			want:     []int{1, 2},
		},
		{
			name:     "multiple duplicates",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
			want:     []int{0, 1, 2, 3, 4},
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
			want:     []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: 1,
			want:     []int{1},
		},
		{
			name:     "all same",
			input:    []int{1, 1, 1},
			expected: 1,
			want:     []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			got := removeDuplicates(input)
			if got != tt.expected {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.expected)
			}
			// 检查前 expected 个元素
			for i := 0; i < got; i++ {
				if input[i] != tt.want[i] {
					t.Errorf("array[%d] = %v, want %v", i, input[i], tt.want[i])
				}
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "normal case",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "carry case",
			input:    []int{1, 9, 9},
			expected: []int{2, 0, 0},
		},
		{
			name:     "all nines",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "single digit",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "zero",
			input:    []int{0},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			result := plusOne(input)
			if len(result) != len(tt.expected) {
				t.Errorf("plusOne() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := 0; i < len(tt.expected); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("plusOne()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "normal case",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "overlapping intervals",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "contained intervals",
			input:    [][]int{{1, 4}, {2, 3}},
			expected: [][]int{{1, 4}},
		},
		{
			name:     "empty input",
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			name:     "single interval",
			input:    [][]int{{1, 4}},
			expected: [][]int{{1, 4}},
		},
		{
			name:     "multiple overlaps",
			input:    [][]int{{1, 4}, {2, 5}, {3, 6}},
			expected: [][]int{{1, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("merge() length = %v, want %v", len(result), len(tt.expected))
				return
			}
			for i := 0; i < len(tt.expected); i++ {
				if len(result[i]) != 2 || len(tt.expected[i]) != 2 {
					t.Errorf("interval %d has wrong length", i)
					return
				}
				if result[i][0] != tt.expected[i][0] || result[i][1] != tt.expected[i][1] {
					t.Errorf("merge()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "normal case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "duplicate numbers",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "same numbers",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4},
			target:   -7,
			expected: []int{2, 3},
		},
		{
			name:     "no solution",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if tt.expected == nil {
				if result != nil {
					t.Errorf("twoSum() = %v, want nil", result)
				}
				return
			}
			if len(result) != 2 {
				t.Errorf("twoSum() length = %v, want 2", len(result))
				return
			}
			// 检查结果中的两个数之和是否等于目标值
			if tt.nums[result[0]]+tt.nums[result[1]] != tt.target {
				t.Errorf("twoSum() = %v, sum = %v, want %v", result, tt.nums[result[0]]+tt.nums[result[1]], tt.target)
			}
		})
	}
}
