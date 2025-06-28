package main

import (
	"testing"
)

func TestDollar(t *testing.T) {
	t.Run("Dollarを作成できる", func(t *testing.T) {
		dollar := NewDollar(5)

		if dollar.Amount() != 5 {
			t.Errorf("Amount() = %d, want 5", dollar.Amount())
		}

		if dollar.Currency() != "USD" {
			t.Errorf("Currency() = %s, want USD", dollar.Currency())
		}
	})

	t.Run("DollarはMoneyである", func(t *testing.T) {
		dollar := NewDollar(10)
		money := &dollar.Money

		if money.Amount() != 10 {
			t.Errorf("Amount() = %d, want 10", money.Amount())
		}
	})

	t.Run("同じ金額のDollarは等価である", func(t *testing.T) {
		dollar1 := NewDollar(5)
		dollar2 := NewDollar(5)

		if !dollar1.Equals(&dollar2.Money) {
			t.Errorf("Equals() = false, want true")
		}
	})

	t.Run("異なる金額のDollarは等価でない", func(t *testing.T) {
		dollar1 := NewDollar(5)
		dollar2 := NewDollar(10)

		if dollar1.Equals(&dollar2.Money) {
			t.Errorf("Equals() = true, want false")
		}
	})
}
