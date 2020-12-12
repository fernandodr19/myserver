package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello", func(t *testing.T) {
		result := GetText()
		expected := "Hello world"

		if result != expected {
			t.Errorf("Result: '%s'; Expected '%s'", result, expected)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("say hello", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if expected != result {
			t.Errorf("Result %d; Expected %d, given %v", result, expected, numbers)
		}
	})

	t.Run("say hello", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if expected != result {
			t.Errorf("Result %d; Expected %d, given %v", result, expected, numbers)
		}
	})
}
