// Money基底クラスの機能テスト
// 責務: 基本操作、等価性判定、共通機能
package main

import (
	"testing"
)

// TestEquality tests Money.Equals method behavior
// 期待: 同じ金額と通貨の場合は等しく、異なる場合は等しくないと判定される
func TestEquality(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		expected bool
	}{
		{
			name:     "同じ金額と通貨の場合、等しいと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "USD"},
			expected: true,
		},
		{
			name:     "金額が異なる場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 10, currency: "USD"},
			expected: false,
		},
		{
			name:     "通貨が異なる場合、等しくないと判定される",
			money1:   Money{amount: 5, currency: "USD"},
			money2:   Money{amount: 5, currency: "CHF"},
			expected: false,
		},
		{
			name:     "同じ金額のDollar同士の場合、等しいと判定される",
			money1:   NewDollar(5),
			money2:   NewDollar(5),
			expected: true,
		},
		{
			name:     "異なる金額のDollar同士の場合、等しくないと判定される",
			money1:   NewDollar(5),
			money2:   NewDollar(10),
			expected: false,
		},
		{
			name:     "同じ金額のFranc同士の場合、等しいと判定される",
			money1:   NewFranc(5),
			money2:   NewFranc(5),
			expected: true,
		},
		{
			name:     "異なる金額のFranc同士の場合、等しくないと判定される",
			money1:   NewFranc(5),
			money2:   NewFranc(10),
			expected: false,
		},
		{
			name:     "DollarとFrancで同じ金額でも、等しくないと判定される",
			money1:   NewDollar(5),
			money2:   NewFranc(5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Equals(tt.money2)
			if result != tt.expected {
				t.Errorf("期待値 %v, 実際 %v", tt.expected, result)
			}
		})
	}
}

// TestMultiplication tests Money.Times method behavior
// 期待: 金額を指定した倍数で乗算し、同じ通貨で新しいMoneyオブジェクトが返される
func TestMultiplication(t *testing.T) {
	tests := []struct {
		name       string
		money      Money
		multiplier int
		expected   Money
	}{
		{
			name:       "5USDを2倍した場合、10USDが返される",
			money:      Money{amount: 5, currency: "USD"},
			multiplier: 2,
			expected:   Money{amount: 10, currency: "USD"},
		},
		{
			name:       "10CHFを3倍した場合、30CHFが返される",
			money:      Money{amount: 10, currency: "CHF"},
			multiplier: 3,
			expected:   Money{amount: 30, currency: "CHF"},
		},
		{
			name:       "NewDollarで作成したMoneyを2倍した場合、正しく動作する",
			money:      NewDollar(5),
			multiplier: 2,
			expected:   NewDollar(10),
		},
		{
			name:       "NewFrancで作成したMoneyを4倍した場合、正しく動作する",
			money:      NewFranc(3),
			multiplier: 4,
			expected:   NewFranc(12),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

// TestCurrency tests Money.Currency method behavior
// 期待: Moneyオブジェクトの通貨コードが正しく取得される
func TestCurrency(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		expected string
	}{
		{
			name:     "USD通貨のMoneyの場合、USDが返される",
			money:    Money{amount: 5, currency: "USD"},
			expected: "USD",
		},
		{
			name:     "CHF通貨のMoneyの場合、CHFが返される",
			money:    Money{amount: 10, currency: "CHF"},
			expected: "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Currency()
			if result != tt.expected {
				t.Errorf("期待値 %s, 実際 %s", tt.expected, result)
			}
		})
	}
}

// TestAddition tests Money.Plus method behavior
// 期待: 2つのMoneyオブジェクトを加算してSumオブジェクトが返される
func TestAddition(t *testing.T) {
	tests := []struct {
		name     string
		money1   Money
		money2   Money
		expected Expression
	}{
		{
			name:     "5USDと3USDを加算した場合、Sumが返される",
			money1:   NewDollar(5),
			money2:   NewDollar(3),
			expected: Sum{augend: NewDollar(5), addend: NewDollar(3)},
		},
		{
			name:     "5USDと10CHFを加算した場合、Sumが返される",
			money1:   NewDollar(5),
			money2:   NewFranc(10),
			expected: Sum{augend: NewDollar(5), addend: NewFranc(10)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Plus(tt.money2)
			// Sumの内容確認（構造体の比較）
			if resultSum, ok := result.(Sum); ok {
				expectedSum := tt.expected.(Sum)
				if !resultSum.augend.Equals(expectedSum.augend) || !resultSum.addend.Equals(expectedSum.addend) {
					t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
				}
			} else {
				t.Errorf("Sumが返されませんでした: %+v", result)
			}
		})
	}
}

// TestReduce tests Money.Reduce method behavior
// 期待: 指定した通貨に換算されたMoneyオブジェクトが返される
func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		money    Money
		bank     *Bank
		to       string
		expected Money
	}{
		{
			name:     "同じ通貨への変換の場合、元の金額がそのまま返される",
			money:    NewDollar(1),
			bank:     NewBank(),
			to:       "USD",
			expected: NewDollar(1),
		},
		{
			name:  "2:1のレートでUSDからCHFに変換した場合、正しく換算される",
			money: NewDollar(2),
			bank: func() *Bank {
				bank := NewBank()
				bank.AddRate("USD", "CHF", 2)
				return bank
			}(),
			to:       "CHF",
			expected: NewFranc(4),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Reduce(tt.bank, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}

// TestConstructors tests NewDollar and NewFranc functions
// 期待: 指定した金額で正しい通貨のMoneyオブジェクトが作成される
func TestConstructors(t *testing.T) {
	tests := []struct {
		name        string
		constructor func(int) Money
		amount      int
		currency    string
	}{
		{
			name:        "NewDollarで5ドルを作成した場合、正しいUSD通貨のMoneyが返される",
			constructor: NewDollar,
			amount:      5,
			currency:    "USD",
		},
		{
			name:        "NewFrancで10フランを作成した場合、正しいCHF通貨のMoneyが返される",
			constructor: NewFranc,
			amount:      10,
			currency:    "CHF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.constructor(tt.amount)
			expected := Money{amount: tt.amount, currency: tt.currency}
			if !result.Equals(expected) {
				t.Errorf("期待値 %+v, 実際 %+v", expected, result)
			}
		})
	}
}

// TestSpecificationExample tests the main specification example
// 期待: 5USD + 10CHF = 10USD（USD 1 : CHF 2のレート）が正しく計算される
func TestSpecificationExample(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (*Bank, Money, Money)
		expected Money
	}{
		{
			name: "5USD+10CHF=10USD（仕様書の例：USD1:CHF2のレート）",
			setup: func() (*Bank, Money, Money) {
				bank := NewBank()
				bank.AddRate("USD", "CHF", 2.0)
				bank.AddRate("CHF", "USD", 0.5)
				return bank, NewDollar(5), NewFranc(10)
			},
			expected: NewDollar(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank, money1, money2 := tt.setup()
			sum := money1.Plus(money2)
			result := bank.Reduce(sum, "USD")
			if !result.Equals(tt.expected) {
				t.Errorf("期待値 %+v, 実際 %+v", tt.expected, result)
			}
		})
	}
}
