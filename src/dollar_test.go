//nolint:dupl
package main

import (
	"testing"
)

// TestNewDollar tests NewDollar constructor behavior //nolint:dupl
// 期待: 指定された金額で正しいUSD通貨のMoneyオブジェクトが作成される
func TestNewDollar(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected Money
	}{
		{
			name:     "5ドルのMoneyオブジェクトが作成される",
			amount:   5,
			expected: Money{amount: 5, currency: "USD"},
		},
		{
			name:     "10ドルのMoneyオブジェクトが作成される",
			amount:   10,
			expected: Money{amount: 10, currency: "USD"},
		},
		{
			name:     "0ドルのMoneyオブジェクトが作成される",
			amount:   0,
			expected: Money{amount: 0, currency: "USD"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewDollar(tt.amount)
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
