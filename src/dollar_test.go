package src

import (
	"testing"
)

func TestNewDollar(t *testing.T) {
	t.Run("指定した金額でDollarインスタンスを作成できる", func(t *testing.T) {
		dollar := NewDollar(10)
		expected := 10
		actual := dollar.Amount()
		if actual != expected {
			t.Errorf("NewDollar(10).Amount() = %d, expected %d", actual, expected)
		}
	})

	t.Run("通貨がUSDに設定される", func(t *testing.T) {
		dollar := NewDollar(10)
		expected := "USD"
		actual := dollar.Currency()
		if actual != expected {
			t.Errorf("NewDollar(10).Currency() = %s, expected %s", actual, expected)
		}
	})
}

func TestDollar_Times(t *testing.T) {
	t.Run("指定した倍数で金額を乗算した新しいDollarインスタンスを返す", func(t *testing.T) {
		dollar := NewDollar(5)
		result := dollar.Times(2)
		expected := NewDollar(10)
		if !result.Equals(expected) {
			t.Errorf("Dollar.Times(2) = %+v, expected %+v", result, expected)
		}
	})
}

func TestDollar_Plus(t *testing.T) {
	t.Run("他のDollarと加算した新しいDollarインスタンスを返す", func(t *testing.T) {
		dollar1 := NewDollar(5)
		dollar2 := NewDollar(10)
		bank := NewBank()
		result := bank.Reduce(dollar1.Plus(dollar2), "USD")
		expected := NewDollar(15)
		if !result.Equals(expected.Money) {
			t.Errorf("Dollar.Plus() = %+v, expected %+v", result, expected.Money)
		}
	})
}

func TestDollar_Equals(t *testing.T) {
	t.Run("同じ金額のDollarの場合はtrueを返す", func(t *testing.T) {
		dollar1 := NewDollar(5)
		dollar2 := NewDollar(5)
		if !dollar1.Equals(dollar2) {
			t.Errorf("Equals() = false, expected true for equal dollars")
		}
	})

	t.Run("異なる金額のDollarの場合はfalseを返す", func(t *testing.T) {
		dollar1 := NewDollar(5)
		dollar2 := NewDollar(10)
		if dollar1.Equals(dollar2) {
			t.Errorf("Equals() = true, expected false for different amounts")
		}
	})
}