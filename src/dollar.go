package main

// NewDollar creates a new Money object with USD currency
func NewDollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}
