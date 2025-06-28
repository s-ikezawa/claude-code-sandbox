package main

// Money represents a currency amount
type Money struct {
	amount   int
	currency string
}

// Equals checks if two Money objects are equal
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// Times returns a new Money object with the amount multiplied by the given multiplier
func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// Currency returns the currency of the Money object
func (m Money) Currency() string {
	return m.currency
}
