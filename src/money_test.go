package currency

import "testing"

// TestNewDollar tests NewDollar constructor behavior
// 期待: 指定された金額で正しいUSD通貨のMoneyオブジェクトが作成される
//
//nolint:dupl
func TestNewDollar(t *testing.T) {
	tests := []struct {
		name   string
		amount int
	}{
		{
			name:   "5ドルのMoneyオブジェクトが作成される",
			amount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewDollar(tt.amount)
			if money.amount != tt.amount {
				t.Errorf("期待値: %d, 実際の値: %d", tt.amount, money.amount)
			}
			if money.currency != "USD" {
				t.Errorf("期待される通貨: USD, 実際の通貨: %s", money.currency)
			}
		})
	}
}

// TestNewFranc tests NewFranc constructor behavior
// 期待: 指定された金額で正しいCHF通貨のMoneyオブジェクトが作成される
//
//nolint:dupl
func TestNewFranc(t *testing.T) {
	tests := []struct {
		name   string
		amount int
	}{
		{
			name:   "10フランのMoneyオブジェクトが作成される",
			amount: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			money := NewFranc(tt.amount)
			if money.amount != tt.amount {
				t.Errorf("期待値: %d, 実際の値: %d", tt.amount, money.amount)
			}
			if money.currency != "CHF" {
				t.Errorf("期待される通貨: CHF, 実際の通貨: %s", money.currency)
			}
		})
	}
}

// TestEquality tests Money.Equals method behavior
// 期待: 同じ値と通貨のオブジェクト同士は等しく、異なる場合は等しくないと判定される
func TestEquality(t *testing.T) {
	tests := []struct {
		name     string
		money1   *Money
		money2   *Money
		expected bool
	}{
		{
			name:     "同じ金額と通貨の場合、等しいと判定される",
			money1:   NewDollar(5),
			money2:   NewDollar(5),
			expected: true,
		},
		{
			name:     "金額が異なる場合、等しくないと判定される",
			money1:   NewDollar(5),
			money2:   NewDollar(6),
			expected: false,
		},
		{
			name:     "通貨が異なる場合、等しくないと判定される",
			money1:   NewDollar(5),
			money2:   NewFranc(5),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money1.Equals(tt.money2)
			if result != tt.expected {
				t.Errorf("期待値: %v, 実際の値: %v", tt.expected, result)
			}
		})
	}
}

// TestMultiplication tests Money.Times method behavior
// 期待: 指定された倍率で金額が正しく計算され、同じ通貨が維持される
func TestMultiplication(t *testing.T) {
	tests := []struct {
		name       string
		money      *Money
		multiplier int
		expected   *Money
	}{
		{
			name:       "5ドル × 2 = 10ドル",
			money:      NewDollar(5),
			multiplier: 2,
			expected:   NewDollar(10),
		},
		{
			name:       "10フラン × 3 = 30フラン",
			money:      NewFranc(10),
			multiplier: 3,
			expected:   NewFranc(30),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.money.Times(tt.multiplier)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", tt.expected, result)
			}
		})
	}
}
