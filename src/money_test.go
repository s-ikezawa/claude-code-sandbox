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

	t.Run("金額と通貨が等しい場合、等価である", func(t *testing.T) {
		money1 := &Money{amount: 5, currency: "USD"}
		money2 := &Money{amount: 5, currency: "USD"}

		if !money1.Equals(money2) {
			t.Errorf("Equals() = false, want true")
		}
	})

	t.Run("等価判定のテスト", func(t *testing.T) {
		tests := []struct {
			name   string
			money1 *Money
			money2 *Money
			want   bool
		}{
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
}
