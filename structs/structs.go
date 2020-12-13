package structs

import "errors"

var (
	ErrorInsufficientFunds = errors.New("Insufficient funds")
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amount float64) {
	// Must receive a pointer in order to
	// edit the original object
	w.balance += amount
}

func (w *Wallet) Withdraw(amount float64) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}
