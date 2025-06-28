package main

import (
	"testing"
)

// Franc専用のテストファイル

func TestFranc(t *testing.T) {
	t.Run("Francを作成できる", func(t *testing.T) {
		franc := NewFranc(5)

		if franc.Amount() != 5 {
			t.Errorf("Amount() = %d, want 5", franc.Amount())
		}

		if franc.Currency() != "CHF" {
			t.Errorf("Currency() = %s, want CHF", franc.Currency())
		}
	})

	t.Run("FrancはMoneyである", func(t *testing.T) {
		franc := NewFranc(10)
		money := &franc.Money

		if money.Amount() != 10 {
			t.Errorf("Amount() = %d, want 10", money.Amount())
		}
	})

	t.Run("同じ金額のFrancは等価である", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(5)

		if !franc1.Equals(&franc2.Money) {
			t.Errorf("Equals() = false, want true")
		}
	})

	t.Run("異なる金額のFrancは等価でない", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(10)

		if franc1.Equals(&franc2.Money) {
			t.Errorf("Equals() = true, want false")
		}
	})

	t.Run("Francの加算", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(10)
		result := franc1.Add(&franc2.Money)
		expected := NewFranc(15)

		if !result.Equals(&expected.Money) {
			t.Errorf("Add() = %v, want %v", result, &expected.Money)
		}
	})
}
