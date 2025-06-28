//nolint:dupl
package main

import (
	"testing"
)

// TestNewFranc tests NewFranc constructor behavior //nolint:dupl
// 期待: 指定された金額で正しいCHF通貨のMoneyオブジェクトが作成される
func TestNewFranc(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected Money
	}{
		{
			name:     "5フランのMoneyオブジェクトが作成される",
			amount:   5,
			expected: Money{amount: 5, currency: "CHF"},
		},
		{
			name:     "10フランのMoneyオブジェクトが作成される",
			amount:   10,
			expected: Money{amount: 10, currency: "CHF"},
		},
		{
			name:     "0フランのMoneyオブジェクトが作成される",
			amount:   0,
			expected: Money{amount: 0, currency: "CHF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFranc(tt.amount)
			if !result.Equals(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
