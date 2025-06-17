package main

import "testing"

func TestStartDelivering(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "should return 2 for > input",
			input:    ">",
			expected: 2,
		},
		{
			name:     "should return 4 for ^>v< input",
			input:    "^>v<",
			expected: 4,
		},
		{
			name:     "should return 2 for ^v^v^v^v^v input",
			input:    "^v^v^v^v^v",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(startDelivering([]byte(tt.input)))
			assertEqual(t, tt.expected, got)
		})
	}
}

func TestStartDeliveringWithRobot(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "should return 3 for ^v input",
			input:    "^v",
			expected: 3,
		},
		{
			name:     "should return 3 for ^>v< input",
			input:    "^>v<",
			expected: 3,
		},
		{
			name:     "should return 11 for ^v^v^v^v^v input",
			input:    "^v^v^v^v^v",
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(startDeliveringWithRobot([]byte(tt.input)))
			assertEqual(t, tt.expected, got)
		})
	}
}

func assertEqual(t testing.TB, expected, got int) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
