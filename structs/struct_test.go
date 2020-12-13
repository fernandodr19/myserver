package structs

import "testing"

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

		if err != ErrorInsufficientFunds {
			t.Errorf("Test failed, test case should receive an Insufficient funds error")
		}

		result := wallet.Balance()
		expected := 10.0

		if result != expected {
			t.Errorf("Result %.2f; Expected %.2f", result, expected)
		}
	})
}
