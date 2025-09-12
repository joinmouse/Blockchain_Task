package main

import "testing"

func TestAddTen(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "positive number",
			input:    5,
			expected: 15,
		},
		{
			name:     "zero",
			input:    0,
			expected: 10,
		},
		{
			name:     "negative number",
			input:    -5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num := tt.input
			addTen(&num)
			if num != tt.expected {
				t.Errorf("addTen() = %v, want %v", num, tt.expected)
			}
		})
	}
}

func TestDoubleSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "negative numbers",
			input:    []int{-1, -2, -3},
			expected: []int{-2, -4, -6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.input
			doubleSlice(&nums)
			if len(nums) != len(tt.expected) {
				t.Errorf("doubleSlice() length = %v, want %v", len(nums), len(tt.expected))
				return
			}
			for i := range nums {
				if nums[i] != tt.expected[i] {
					t.Errorf("doubleSlice()[%d] = %v, want %v", i, nums[i], tt.expected[i])
				}
			}
		})
	}
} 
