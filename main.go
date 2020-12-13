package main

import (
	"errors"
	"math"
)

func GetText() string {
	return "Hello world"
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

type Form interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

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

var (
	ErrorInsufficientFunds = errors.New("Insufficient funds")
)

func (w *Wallet) Withdraw(amount float64) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}
