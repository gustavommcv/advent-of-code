package main

import "testing"

func TestWhatFloor(t *testing.T) {
	t.Run("it should be 0 with (()) and ()()", func(t *testing.T) {
		want := 0
		input1 := "()()"
		input2 := "(())"

		result1 := WhatFloor([]byte(input1))
		result2 := WhatFloor([]byte(input2))

		assertResult(t, want, result1)
		assertResult(t, want, result2)
	})

	t.Run("it should be 3 with ((( and (()(()(", func(t *testing.T) {
		want := 3
		input1 := "((("
		input2 := "(()(()("

		result1 := WhatFloor([]byte(input1))
		result2 := WhatFloor([]byte(input2))

		assertResult(t, want, result1)
		assertResult(t, want, result2)
	})
}

func assertResult(t testing.TB, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
