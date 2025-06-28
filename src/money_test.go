// Money基底クラスの機能テスト
// 責務: 基本操作、等価性判定、共通機能
package main

import (
	"testing"
)

func TestMoneyの等価性判定(t *testing.T) {
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
			name:     "金額が異なる場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 10, currency: "USD"},
			expected: false,
		},
		{
			name:     "通貨が異なる場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "CHF"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Equals(tt.money2)
			if result != tt.expected {
				t.Errorf("期待値 %v, 実際 %v", tt.expected, result)
			}
		})
	}
}

func TestMoneyの乗算(t *testing.T) {
	tests := []struct {
		name       string
		money      Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5USDを2倍した場合、10USDが返される",
			money:      Money{amount: 5, currency: "USD"},
			multiplier: 2,
			expected:   Money{amount: 10, currency: "USD"},
		},
		{
			name:       "10CHFを3倍した場合、30CHFが返される",
			money:      Money{amount: 10, currency: "CHF"},
			multiplier: 3,
			expected:   Money{amount: 30, currency: "CHF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestMoneyの通貨取得(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		expected string
	}{
		{
			name:     "USD通貨のMoneyの場合、USDが返される",
			money:    Money{amount: 5, currency: "USD"},
			expected: "USD",
		},
		{
			name:     "CHF通貨のMoneyの場合、CHFが返される",
			money:    Money{amount: 10, currency: "CHF"},
			expected: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Currency()
			if result != tt.expected {
				t.Errorf("期待値 %s, 実際 %s", tt.expected, result)
			}
		})
	}
}
