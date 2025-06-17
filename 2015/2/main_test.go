package main

import (
	"reflect"
	"testing"
)

func TestGetPresentWrapper(t *testing.T) {
	tests := []struct {
		name     string
		present  Present
		expected PresentWrapper
	}{
		{
			name: "should return correct wrapper for 2x3x4 present",
			present: Present{
				length: 2,
				width:  3,
				height: 4,
			},
			expected: PresentWrapper{
				length:     2,
				width:      3,
				height:     4,
				squareFeet: 52,
			},
		},
		{
			name: "should return correct wrapper for 1x1x10 present",
			present: Present{
				length: 1,
				width:  1,
				height: 10,
			},
			expected: PresentWrapper{
				length:     1,
				width:      1,
				height:     10,
				squareFeet: 42,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPresentWrapper(tt.present)
			assertPresentWrapper(t, tt.expected, got)
		})
	}
}

func TestGetExtraPaper(t *testing.T) {
	tests := []struct {
		name     string
		present  Present
		expected int
	}{
		{
			name: "should return 6 for 2x3x4 present",
			present: Present{
				length: 2,
				width:  3,
				height: 4,
			},
			expected: 6,
		},
		{
			name: "should return 1 for 1x1x10 present",
			present: Present{
				length: 1,
				width:  1,
				height: 10,
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := PresentWrapper{
				length: tt.present.length,
				width:  tt.present.width,
				height: tt.present.height,
			}
			got := wrapper.getExtraPaper()
			assertEqual(t, tt.expected, got)
		})
	}
}

func TestGetWrapRibbon(t *testing.T) {
	tests := []struct {
		name     string
		present  Present
		expected int
	}{
		{
			name: "should return 10 for 2x3x4 present",
			present: Present{
				length: 2,
				width:  3,
				height: 4,
			},
			expected: 10,
		},
		{
			name: "should return 4 for 1x1x10 present",
			present: Present{
				length: 1,
				width:  1,
				height: 10,
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWrapRibbon(tt.present)
			assertEqual(t, tt.expected, got)
		})
	}
}

func TestGetWrapBow(t *testing.T) {
	tests := []struct {
		name     string
		present  Present
		expected int
	}{
		{
			name: "should return 24 for 2x3x4 present",
			present: Present{
				length: 2,
				width:  3,
				height: 4,
			},
			expected: 24,
		},
		{
			name: "should return 10 for 1x1x10 present",
			present: Present{
				length: 1,
				width:  1,
				height: 10,
			},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWrapBow(tt.present)
			assertEqual(t, tt.expected, got)
		})
	}
}

func TestGetTotalRibbon(t *testing.T) {
	tests := []struct {
		name     string
		present  Present
		expected int
	}{
		{
			name: "should return 34 for 2x3x4 present",
			present: Present{
				length: 2,
				width:  3,
				height: 4,
			},
			expected: 34,
		},
		{
			name: "should return 14 for 1x1x10 present",
			present: Present{
				length: 1,
				width:  10,
				height: 1,
			},
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTotalRibbon(tt.present)
			assertEqual(t, tt.expected, got)
		})
	}
}

func assertPresentWrapper(t testing.TB, expected, got PresentWrapper) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
}

func assertEqual(t testing.TB, expected, got int) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
