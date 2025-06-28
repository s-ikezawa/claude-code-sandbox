package src

import (
	"testing"
)

func TestCurrencyCreation(t *testing.T) {
	testCases := []struct {
		name     string
		creator  func(int) *Money
		amount   int
		currency string
	}{
		{"NewDollar creates USD currency", NewDollar, 10, "USD"},
		{"NewFranc creates CHF currency", NewFranc, 10, "CHF"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			money := tc.creator(tc.amount)
			if money.Amount() != tc.amount {
				t.Errorf("%s(%d).Amount() = %d, expected %d", tc.name, tc.amount, money.Amount(), tc.amount)
			}
			if money.Currency() != tc.currency {
				t.Errorf("%s(%d).Currency() = %s, expected %s", tc.name, tc.amount, money.Currency(), tc.currency)
			}
		})
	}
}

func TestCurrencyTimes(t *testing.T) {
	testCases := []struct {
		name     string
		creator  func(int) *Money
		currency string
	}{
		{"Dollar multiplication", NewDollar, "USD"},
		{"Franc multiplication", NewFranc, "CHF"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			money := tc.creator(5)
			result := money.Times(2)
			expected := tc.creator(10)
			if !result.Equals(expected) {
				t.Errorf("%s.Times(2) = %+v, expected %+v", tc.name, result, expected)
			}
		})
	}
}

func TestCurrencyPlus(t *testing.T) {
	testCases := []struct {
		name     string
		creator  func(int) *Money
		currency string
	}{
		{"Dollar addition", NewDollar, "USD"},
		{"Franc addition", NewFranc, "CHF"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			money1 := tc.creator(5)
			money2 := tc.creator(10)
			bank := NewBank()
			result := bank.Reduce(money1.Plus(money2), tc.currency)
			expected := tc.creator(15)
			if !result.Equals(expected) {
				t.Errorf("%s.Plus() = %+v, expected %+v", tc.name, result, expected)
			}
		})
	}
}

func TestCurrencyEquals(t *testing.T) {
	testCases := []struct {
		name     string
		creator  func(int) *Money
		currency string
	}{
		{"Dollar equality", NewDollar, "USD"},
		{"Franc equality", NewFranc, "CHF"},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" - equal amounts", func(t *testing.T) {
			money1 := tc.creator(5)
			money2 := tc.creator(5)
			if !money1.Equals(money2) {
				t.Errorf("Equals() = false, expected true for equal %s", tc.currency)
			}
		})

		t.Run(tc.name+" - different amounts", func(t *testing.T) {
			money1 := tc.creator(5)
			money2 := tc.creator(10)
			if money1.Equals(money2) {
				t.Errorf("Equals() = true, expected false for different %s amounts", tc.currency)
			}
		})
	}
}
