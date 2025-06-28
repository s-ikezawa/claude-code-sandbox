package main

import (
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("金額を取得できる", func(t *testing.T) {
		money := &Money{amount: 5}
		got := money.Amount()
		want := 5

		if got != want {
			t.Errorf("Amount() = %d, want %d", got, want)
		}
	})

	t.Run("通貨を取得できる", func(t *testing.T) {
		money := &Money{currency: "USD"}
		got := money.Currency()
		want := "USD"

		if got != want {
			t.Errorf("Currency() = %s, want %s", got, want)
		}
	})

	t.Run("等価性判定", func(t *testing.T) {
		tests := []struct {
			name   string
			money1 *Money
			money2 *Money
			want   bool
		}{
			{
				name:   "金額と通貨が等しい場合、等価である",
				money1: &Money{amount: 5, currency: "USD"},
				money2: &Money{amount: 5, currency: "USD"},
				want:   true,
			},
			{
				name:   "金額が異なる場合、等価でない",
				money1: &Money{amount: 5, currency: "USD"},
				money2: &Money{amount: 10, currency: "USD"},
				want:   false,
			},
			{
				name:   "通貨が異なる場合、等価でない",
				money1: &Money{amount: 5, currency: "USD"},
				money2: &Money{amount: 5, currency: "CHF"},
				want:   false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := tt.money1.Equals(tt.money2); got != tt.want {
					t.Errorf("Equals() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("通貨の加算", func(t *testing.T) {
		t.Run("同じ通貨同士の加算", func(t *testing.T) {
			money1 := &Money{amount: 5, currency: "USD"}
			money2 := &Money{amount: 10, currency: "USD"}
			result := money1.Add(money2)
			expected := &Money{amount: 15, currency: "USD"}

			if !result.Equals(expected) {
				t.Errorf("Add() = %v, want %v", result, expected)
			}
		})

		t.Run("異なる通貨間の加算", func(t *testing.T) {
			money1 := &Money{amount: 5, currency: "USD"}
			money2 := &Money{amount: 10, currency: "CHF"}
			result := money1.Add(money2)

			// 加算結果は加算対象の通貨になる
			if result.Currency() != "USD" {
				t.Errorf("Currency() = %s, want USD", result.Currency())
			}

			// 為替レート変換は別機能で実装するため、ここでは金額のみテスト
			if result.Amount() != 15 {
				t.Errorf("Amount() = %d, want 15", result.Amount())
			}
		})
	})
}
