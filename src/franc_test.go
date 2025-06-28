package src

import (
	"testing"
)

func TestNewFranc(t *testing.T) {
	t.Run("指定した金額でFrancインスタンスを作成できる", func(t *testing.T) {
		franc := NewFranc(10)
		expected := 10
		actual := franc.Amount()
		if actual != expected {
			t.Errorf("NewFranc(10).Amount() = %d, expected %d", actual, expected)
		}
	})

	t.Run("通貨がCHFに設定される", func(t *testing.T) {
		franc := NewFranc(10)
		expected := "CHF"
		actual := franc.Currency()
		if actual != expected {
			t.Errorf("NewFranc(10).Currency() = %s, expected %s", actual, expected)
		}
	})
}

func TestFranc_Times(t *testing.T) {
	t.Run("指定した倍数で金額を乗算した新しいFrancインスタンスを返す", func(t *testing.T) {
		franc := NewFranc(5)
		result := franc.Times(2)
		expected := NewFranc(10)
		if !result.Equals(expected) {
			t.Errorf("Franc.Times(2) = %+v, expected %+v", result, expected)
		}
	})
}

func TestFranc_Plus(t *testing.T) {
	t.Run("他のFrancと加算した新しいFrancインスタンスを返す", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(10)
		bank := NewBank()
		result := bank.Reduce(franc1.Plus(franc2), "CHF")
		expected := NewFranc(15)
		if !result.Equals(expected.Money) {
			t.Errorf("Franc.Plus() = %+v, expected %+v", result, expected.Money)
		}
	})
}

func TestFranc_Equals(t *testing.T) {
	t.Run("同じ金額のFrancの場合はtrueを返す", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(5)
		if !franc1.Equals(franc2) {
			t.Errorf("Equals() = false, expected true for equal francs")
		}
	})

	t.Run("異なる金額のFrancの場合はfalseを返す", func(t *testing.T) {
		franc1 := NewFranc(5)
		franc2 := NewFranc(10)
		if franc1.Equals(franc2) {
			t.Errorf("Equals() = true, expected false for different amounts")
		}
	})
}