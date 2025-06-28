package main

// NewFranc creates a new Money object with CHF currency
func NewFranc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}
