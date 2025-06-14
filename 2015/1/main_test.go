package main

import "testing"

func TestGetFloor(t *testing.T) {
	t.Run("it should be 0 with (()) and ()()", func(t *testing.T) {
		want := 0
		input1 := "()()"
		input2 := "(())"

		result1 := GetFloor([]byte(input1))
		result2 := GetFloor([]byte(input2))

		assertResult(t, want, result1)
		assertResult(t, want, result2)
	})

	t.Run("it should be 3 with ((( and (()(()(", func(t *testing.T) {
		want := 3
		input1 := "((("
		input2 := "(()(()("

		result1 := GetFloor([]byte(input1))
		result2 := GetFloor([]byte(input2))

		assertResult(t, want, result1)
		assertResult(t, want, result2)
	})

	t.Run("it should not count another type of characters", func(t *testing.T) {
		want := 3
		input := "(kdfjdf90(("

		result := GetFloor([]byte(input))

		assertResult(t, want, result)
	})

}

func TestGetPosition(t *testing.T) {
	t.Run("it should be 1 with )", func(t *testing.T) {
		want := 1
		input := ")"

		result := GetPosition([]byte(input))
		assertResult(t, want, result)
	})

	t.Run("it should be 5 with ()())", func(t *testing.T) {
		want := 5
		input := "()())"

		result := GetPosition([]byte(input))
		assertResult(t, want, result)
	})

	t.Run("it should not count another type of characters", func(t *testing.T) {
		want := 1
		input := "fjdf38994)"

		result := GetPosition([]byte(input))
		assertResult(t, want, result)
	})
}

func assertResult(t testing.TB, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
