package main

import "testing"

func TestIsNice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "should return true for 'ugknbfddgicrmopn' input",
			input:    "ugknbfddgicrmopn",
			expected: true,
		},
		{
			name:     "should return true for 'aaa' input",
			input:    "aaa",
			expected: true,
		},
		{
			name:     "should return false for 'jchzalrnumimnmhp' input",
			input:    "jchzalrnumimnmhp",
			expected: false,
		},
		{
			name:     "should return false for 'haegwjzuvuyypxyu' input",
			input:    "haegwjzuvuyypxyu",
			expected: false,
		},
		{
			name:     "should return false for 'dvszwmarrgswjxmb' input",
			input:    "dvszwmarrgswjxmb",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNice(tt.input)
			assertEqualBool(t, tt.expected, got)
		})
	}
}

func TestIsNice2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "should return true for 'qjhvhtzxzqqjkmpb' input",
			input:    "qjhvhtzxzqqjkmpb",
			expected: true,
		},
		{
			name:     "should return true for 'xxyxx' input",
			input:    "xxyxx",
			expected: true,
		},
		{
			name:     "should return false for 'uurcxstgmygtbstg' input",
			input:    "uurcxstgmygtbstg",
			expected: false,
		},
		{
			name:     "should return false for 'ieodomkazucvgmuy' input",
			input:    "ieodomkazucvgmuy",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNice(tt.input)
			assertEqualBool(t, tt.expected, got)
		})
	}
}

func assertEqualBool(t testing.TB, want, got bool) {
	t.Helper()

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
