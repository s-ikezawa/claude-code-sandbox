// Dollar専用機能のテスト
// 責務: Dollar作成、Dollar固有動作
package main

import (
	"testing"
)

func TestDollarの作成(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected Money
	}{
		{
			name:     "5ドルを作成した場合、金額5でUSD通貨のMoneyが返される",
			amount:   5,
			expected: Money{amount: 5, currency: "USD"},
		},
		{
			name:     "10ドルを作成した場合、金額10でUSD通貨のMoneyが返される",
			amount:   10,
			expected: Money{amount: 10, currency: "USD"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewDollar(tt.amount)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestDollarの乗算(t *testing.T) {
	tests := []struct {
		name       string
		dollar     Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5ドルを2倍した場合、10ドルが返される",
			dollar:     NewDollar(5),
			multiplier: 2,
			expected:   NewDollar(10),
		},
		{
			name:       "3ドルを4倍した場合、12ドルが返される",
			dollar:     NewDollar(3),
			multiplier: 4,
			expected:   NewDollar(12),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dollar.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestDollarの等価性判定(t *testing.T) {
	tests := []struct {
		name     string
		dollar1  Money
		dollar2  Money
		expected bool
	}{
		{
			name:     "同じ金額のドル同士の場合、等しいと判定される",
			dollar1:  NewDollar(5),
			dollar2:  NewDollar(5),
			expected: true,
		},
		{
			name:     "異なる金額のドル同士の場合、等しくないと判定される",
			dollar1:  NewDollar(5),
			dollar2:  NewDollar(10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dollar1.Equals(tt.dollar2)
			if result != tt.expected {
				t.Errorf("期待値 %v, 実際 %v", tt.expected, result)
			}
		})
	}
}
