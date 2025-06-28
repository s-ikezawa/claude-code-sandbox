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
