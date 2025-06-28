package main

import (
	"testing"
)

// TestMoneyEquals tests Money.Equals method behavior
// 期待: 同じ金額と通貨の場合は等しく、異なる場合は等しくないと判定される
func TestMoneyEquals(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		expected bool
	}{
		{
			name:     "同じ金額と通貨の場合、等しいと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "USD"},
			expected: true,
		},
		{
			name:     "異なる金額の場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 10, currency: "USD"},
			expected: false,
		},
		{
			name:     "異なる通貨の場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "CHF"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Equals(tt.money2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestMoneyTimes tests Money.Times method behavior
// 期待: 金額に指定した倍数を掛けた新しいMoneyオブジェクトが返される
func TestMoneyTimes(t *testing.T) {
	tests := []struct {
		name       string
		money      Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5 USD x 2 = 10 USD",
			money:      Money{amount: 5, currency: "USD"},
			multiplier: 2,
			expected:   Money{amount: 10, currency: "USD"},
		},
		{
			name:       "10 CHF x 3 = 30 CHF",
			money:      Money{amount: 10, currency: "CHF"},
			multiplier: 3,
			expected:   Money{amount: 30, currency: "CHF"},
		},
		{
			name:       "5 USD x 0 = 0 USD",
			money:      Money{amount: 5, currency: "USD"},
			multiplier: 0,
			expected:   Money{amount: 0, currency: "USD"},
		},
	}

	for _, tt := range tests { //nolint:dupl
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestMoneyCurrency tests Money.Currency method behavior
// 期待: Moneyオブジェクトの通貨が正しく返される
func TestMoneyCurrency(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		expected string
	}{
		{
			name:     "USD通貨のMoneyオブジェクトの場合、USDが返される",
			money:    Money{amount: 5, currency: "USD"},
			expected: "USD",
		},
		{
			name:     "CHF通貨のMoneyオブジェクトの場合、CHFが返される",
			money:    Money{amount: 10, currency: "CHF"},
			expected: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Currency()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestMoneyPlus tests Money.Plus method behavior for same currency
// 期待: 同じ通貨同士の場合、金額が加算された新しいMoneyオブジェクトが返される
func TestMoneyPlus(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		expected Money
	}{
		{
			name:     "5 USD + 5 USD = 10 USD",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "USD"},
			expected: Money{amount: 10, currency: "USD"},
		},
		{
			name:     "10 CHF + 5 CHF = 15 CHF",
			money1:   Money{amount: 10, currency: "CHF"},
			money2:   Money{amount: 5, currency: "CHF"},
			expected: Money{amount: 15, currency: "CHF"},
		},
		{
			name:     "0 USD + 5 USD = 5 USD",
			money1:   Money{amount: 0, currency: "USD"},
			money2:   Money{amount: 5, currency: "USD"},
			expected: Money{amount: 5, currency: "USD"},
		},
	}

	for _, tt := range tests { //nolint:dupl
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Plus(tt.money2)
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
