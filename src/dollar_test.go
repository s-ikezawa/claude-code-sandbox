package main

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	tests := []struct {
		name             string
		createFunc       func(int) *Money
		amount           int
		expectedCurrency string
	}{
		{
			name: "Dollarを作成できる",
			createFunc: func(amount int) *Money {
				return &NewDollar(amount).Money
			},
			amount:           5,
			expectedCurrency: "USD",
		},
		{
			name: "Francを作成できる",
			createFunc: func(amount int) *Money {
				return &NewFranc(amount).Money
			},
			amount:           5,
			expectedCurrency: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := tt.createFunc(tt.amount)

			if money.Amount() != tt.amount {
				t.Errorf("Amount() = %d, want %d", money.Amount(), tt.amount)
			}

			if money.Currency() != tt.expectedCurrency {
				t.Errorf("Currency() = %s, want %s", money.Currency(), tt.expectedCurrency)
			}
		})
	}
}

func TestEquals(t *testing.T) {
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
		{
			name:   "同じ金額のDollarは等価である",
			money1: &NewDollar(5).Money,
			money2: &NewDollar(5).Money,
			want:   true,
		},
		{
			name:   "異なる金額のDollarは等価でない",
			money1: &NewDollar(5).Money,
			money2: &NewDollar(10).Money,
			want:   false,
		},
		{
			name:   "同じ金額のFrancは等価である",
			money1: &NewFranc(5).Money,
			money2: &NewFranc(5).Money,
			want:   true,
		},
		{
			name:   "異なる金額のFrancは等価でない",
			money1: &NewFranc(5).Money,
			money2: &NewFranc(10).Money,
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
}
