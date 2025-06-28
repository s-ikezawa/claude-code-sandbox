package src

import (
	"testing"
)

func TestMoney_Amount(t *testing.T) {
	t.Run("金額を正しく取得できる", func(t *testing.T) {
		money := &Money{amount: 10, currency: "USD"}
		expected := 10
		actual := money.Amount()
		if actual != expected {
			t.Errorf("Amount() = %d, expected %d", actual, expected)
		}
	})
}

func TestMoney_Currency(t *testing.T) {
	t.Run("通貨を正しく取得できる", func(t *testing.T) {
		money := &Money{amount: 10, currency: "USD"}
		expected := "USD"
		actual := money.Currency()
		if actual != expected {
			t.Errorf("Currency() = %s, expected %s", actual, expected)
		}
	})
}

func TestMoney_Equals(t *testing.T) {
	t.Run("同じ金額と通貨の場合はtrueを返す", func(t *testing.T) {
		money1 := &Money{amount: 10, currency: "USD"}
		money2 := &Money{amount: 10, currency: "USD"}
		if !money1.Equals(money2) {
			t.Errorf("Equals() = false, expected true for equal money")
		}
	})

	t.Run("異なる金額の場合はfalseを返す", func(t *testing.T) {
		money1 := &Money{amount: 10, currency: "USD"}
		money2 := &Money{amount: 20, currency: "USD"}
		if money1.Equals(money2) {
			t.Errorf("Equals() = true, expected false for different amounts")
		}
	})

	t.Run("異なる通貨の場合はfalseを返す", func(t *testing.T) {
		money1 := &Money{amount: 10, currency: "USD"}
		money2 := &Money{amount: 10, currency: "CHF"}
		if money1.Equals(money2) {
			t.Errorf("Equals() = true, expected false for different currencies")
		}
	})
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