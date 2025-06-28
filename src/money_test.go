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
}
