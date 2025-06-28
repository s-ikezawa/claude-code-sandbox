package main

import (
	"testing"
)

// TestBankExchange tests Bank.Exchange method behavior
// 期待: 為替レートに基づいて通貨が正しく変換される
func TestBankExchange(t *testing.T) {
	bank := NewBank()
	bank.AddRate("USD", "CHF", 2) // 1 USD = 2 CHF

	tests := []struct { //nolint:dupl
		name     string
		money    Money
		to       string
		expected Money
	}{
		{
			name:     "10 CHF を USD に変換すると 5 USD になる",
			money:    Money{amount: 10, currency: "CHF"},
			to:       "USD",
			expected: Money{amount: 5, currency: "USD"},
		},
		{
			name:     "同じ通貨の場合、変換されずそのまま返される",
			money:    Money{amount: 5, currency: "USD"},
			to:       "USD",
			expected: Money{amount: 5, currency: "USD"},
		},
		{
			name:     "5 USD を CHF に変換すると 10 CHF になる",
			money:    Money{amount: 5, currency: "USD"},
			to:       "CHF",
			expected: Money{amount: 10, currency: "CHF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bank.Exchange(tt.money, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestNewBank tests NewBank constructor behavior
// 期待: 新しいBankオブジェクトが作成される
func TestNewBank(t *testing.T) {
	bank := NewBank()
	if bank == nil {
		t.Error("expected Bank object, got nil")
	}
}
