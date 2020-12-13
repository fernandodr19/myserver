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

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, form Form, expected float64) {
		result := form.Area()
		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		verifyArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		verifyArea(t, circle, 314.1592653589793)
	})
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10.0)

		result := wallet.Balance()
		expected := 10.0

		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	})

	t.Run("withdraw_ok", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10.0)
		err := wallet.Withdraw(5.0)

		if err != nil {
			t.Errorf("Not expected error: %s", err.Error())
		}

		result := wallet.Balance()
		expected := 5.0

		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	})

	t.Run("withdraw_fail", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10.0)
		err := wallet.Withdraw(15.0)

		if err == nil {
			t.Errorf("Test failed, test case should receive an Insufficient funds error")
		}

		result := wallet.Balance()
		expected := 10.0

		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	})
}
