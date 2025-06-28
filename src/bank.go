package main

// Bank handles currency exchange rates and conversions
type Bank struct {
	rates map[string]map[string]int
}

// NewBank creates a new Bank object
func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]map[string]int),
	}
}

// AddRate adds an exchange rate from one currency to another
func (b *Bank) AddRate(from, to string, rate int) {
	if b.rates[from] == nil {
		b.rates[from] = make(map[string]int)
	}
	b.rates[from][to] = rate
}

// Exchange converts money from one currency to another using stored rates
func (b *Bank) Exchange(money Money, to string) Money {
	if money.currency == to {
		return money
	}

	rate, exists := b.rates[money.currency][to]
	if !exists {
		// Try reverse rate
		reverseRate, reverseExists := b.rates[to][money.currency]
		if !reverseExists {
			panic("exchange rate not found")
		}
		// Convert using reverse rate (divide by reverse rate)
		return Money{
			amount:   money.amount / reverseRate,
			currency: to,
		}
	}

	// Convert using direct rate (multiply by rate)
	return Money{
		amount:   money.amount * rate,
		currency: to,
	}
}
