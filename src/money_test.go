package main

import (
	"testing"
)

// TestEquality tests Money.Equals method behavior
// 期待: 同じ値と通貨のオブジェクト同士は等しく、異なる値や通貨の場合は等しくないと判定される
func TestEquality(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		expected bool
	}{
		{
			name:     "同じ値と通貨のMoney同士の場合、等しいと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "USD"},
			expected: true,
		},
		{
			name:     "異なる値のMoney同士の場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 10, currency: "USD"},
			expected: false,
		},
		{
			name:     "異なる通貨のMoney同士の場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "CHF"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Equals(tt.money2)
			if result != tt.expected {
				t.Errorf("期待値: %v, 実際の値: %v", tt.expected, result)
			}
		})
	}
}

// TestMultiplication tests Money.Multiply method behavior
// 期待: 金額に指定された値を掛けた新しいMoneyオブジェクトが返される
func TestMultiplication(t *testing.T) {
	tests := []struct {
		name       string
		money      Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5USDを2倍すると10USDになる",
			money:      Money{amount: 5, currency: "USD"},
			multiplier: 2,
			expected:   Money{amount: 10, currency: "USD"},
		},
		{
			name:       "10CHFを3倍すると30CHFになる",
			money:      Money{amount: 10, currency: "CHF"},
			multiplier: 3,
			expected:   Money{amount: 30, currency: "CHF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Multiply(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", tt.expected, result)
			}
		})
	}
}

// TestCurrency tests Money.GetCurrency method behavior
// 期待: Moneyオブジェクトの通貨が正しく取得される
func TestCurrency(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		expected string
	}{
		{
			name:     "USDの通貨が正しく取得される",
			money:    Money{amount: 5, currency: "USD"},
			expected: "USD",
		},
		{
			name:     "CHFの通貨が正しく取得される",
			money:    Money{amount: 10, currency: "CHF"},
			expected: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.GetCurrency()
			if result != tt.expected {
				t.Errorf("期待値: %s, 実際の値: %s", tt.expected, result)
			}
		})
	}
}
