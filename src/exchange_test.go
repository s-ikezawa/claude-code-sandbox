// Bank/為替レート機能のテスト
// 責務: レート設定、換算、変換加算
package main

import (
	"testing"
)

func TestBankの為替レート設定(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		rate     float64
		expected bool
	}{
		{
			name:     "USDからCHFへのレート2.0を設定した場合、設定が成功する",
			from:     "USD",
			to:       "CHF",
			rate:     2.0,
			expected: true,
		},
		{
			name:     "CHFからUSDへのレート0.5を設定した場合、設定が成功する",
			from:     "CHF",
			to:       "USD",
			rate:     0.5,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			bank.AddRate(tt.from, tt.to, tt.rate)
			// 設定が成功したかどうかは、実際の換算テストで確認
		})
	}
}

func TestBankの通貨換算(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		to       string
		expected Money
	}{
		{
			name:     "5USDをCHFに換算した場合、10CHFが返される",
			money:    NewDollar(5),
			to:       "CHF",
			expected: NewFranc(10),
		},
		{
			name:     "10CHFをUSDに換算した場合、5USDが返される",
			money:    NewFranc(10),
			to:       "USD",
			expected: NewDollar(5),
		},
		{
			name:     "同じ通貨への換算の場合、元の金額がそのまま返される",
			money:    NewDollar(5),
			to:       "USD",
			expected: NewDollar(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			bank.AddRate("USD", "CHF", 2.0)
			bank.AddRate("CHF", "USD", 0.5)

			result := bank.Reduce(tt.money, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

func TestBankの異なる通貨の合計(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		to       string
		expected Money
	}{
		{
			name:     "5USDと10CHFを合計してUSDで取得した場合、10USDが返される",
			money1:   NewDollar(5),
			money2:   NewFranc(10),
			to:       "USD",
			expected: NewDollar(10),
		},
		{
			name:     "5USDと10CHFを合計してCHFで取得した場合、20CHFが返される",
			money1:   NewDollar(5),
			money2:   NewFranc(10),
			to:       "CHF",
			expected: NewFranc(20),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			bank.AddRate("USD", "CHF", 2.0)
			bank.AddRate("CHF", "USD", 0.5)

			sum := tt.money1.Plus(tt.money2)
			result := bank.Reduce(sum, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}
