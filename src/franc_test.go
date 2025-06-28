// Franc専用機能のテスト
// 責務: Franc作成、Franc固有動作
package main

import (
	"testing"
)

func TestFrancの作成(t *testing.T) {
	tests := []struct {
		name     string
		amount   int
		expected Money
	}{
		{
			name:     "5フランを作成した場合、金額5でCHF通貨のMoneyが返される",
			amount:   5,
			expected: Money{amount: 5, currency: "CHF"},
		},
		{
			name:     "10フランを作成した場合、金額10でCHF通貨のMoneyが返される",
			amount:   10,
			expected: Money{amount: 10, currency: "CHF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewFranc(tt.amount)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestFrancの乗算(t *testing.T) {
	tests := []struct {
		name       string
		franc      Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5フランを2倍した場合、10フランが返される",
			franc:      NewFranc(5),
			multiplier: 2,
			expected:   NewFranc(10),
		},
		{
			name:       "3フランを4倍した場合、12フランが返される",
			franc:      NewFranc(3),
			multiplier: 4,
			expected:   NewFranc(12),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.franc.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestFrancの等価性判定(t *testing.T) {
	tests := []struct {
		name     string
		franc1   Money
		franc2   Money
		expected bool
	}{
		{
			name:     "同じ金額のフラン同士の場合、等しいと判定される",
			franc1:   NewFranc(5),
			franc2:   NewFranc(5),
			expected: true,
		},
		{
			name:     "異なる金額のフラン同士の場合、等しくないと判定される",
			franc1:   NewFranc(5),
			franc2:   NewFranc(10),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.franc1.Equals(tt.franc2)
			if result != tt.expected {
				t.Errorf("期待値 %v, 実際 %v", tt.expected, result)
			}
		})
	}
}
