package src

import (
	"testing"
)

func TestMoney_AccessorMethods(t *testing.T) {
	money := &Money{amount: 10, currency: "USD"}

	t.Run("金額を正しく取得できる", func(t *testing.T) {
		expected := 10
		actual := money.Amount()
		if actual != expected {
			t.Errorf("Amount() = %d, expected %d", actual, expected)
		}
	})

	t.Run("通貨を正しく取得できる", func(t *testing.T) {
		expected := "USD"
		actual := money.Currency()
		if actual != expected {
			t.Errorf("Currency() = %s, expected %s", actual, expected)
		}
	})
}

func TestMoney_Equals(t *testing.T) {
	testCases := []struct {
		name     string
		money1   *Money
		money2   *Money
		expected bool
	}{
		{"同じ金額と通貨の場合はtrueを返す", &Money{amount: 10, currency: "USD"}, &Money{amount: 10, currency: "USD"}, true},
		{"異なる金額の場合はfalseを返す", &Money{amount: 10, currency: "USD"}, &Money{amount: 20, currency: "USD"}, false},
		{"異なる通貨の場合はfalseを返す", &Money{amount: 10, currency: "USD"}, &Money{amount: 10, currency: "CHF"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.money1.Equals(tc.money2)
			if actual != tc.expected {
				t.Errorf("Equals() = %t, expected %t for %s", actual, tc.expected, tc.name)
			}
		})
	}
}

func TestMoney_Times(t *testing.T) {
	t.Run("指定した倍数で金額を乗算した新しいMoneyインスタンスを返す", func(t *testing.T) {
		money := &Money{amount: 10, currency: "USD"}
		result := money.Times(2)
		expected := &Money{amount: 20, currency: "USD"}
		if !result.Equals(expected) {
			t.Errorf("Times(2) = %+v, expected %+v", result, expected)
		}
	})
}

func TestMoney_Plus(t *testing.T) {
	t.Run("同じ通貨の場合は金額を加算した新しいMoneyインスタンスを返す", func(t *testing.T) {
		money1 := &Money{amount: 10, currency: "USD"}
		money2 := &Money{amount: 5, currency: "USD"}
		bank := NewBank()
		result := bank.Reduce(money1.Plus(money2), "USD")
		expected := &Money{amount: 15, currency: "USD"}
		if !result.Equals(expected) {
			t.Errorf("Plus() = %+v, expected %+v", result, expected)
		}
	})
}

// NewDollar/NewFrancファクトリ関数のテスト
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
