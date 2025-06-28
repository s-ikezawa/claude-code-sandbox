package main

import (
	"testing"
)

// TestNewDollar tests NewDollar constructor behavior
// 期待: 指定された金額で正しいUSD通貨のMoneyオブジェクトが作成される
func TestNewDollar(t *testing.T) {
	t.Run("5ドルのDollarオブジェクトが正しく作成される", func(t *testing.T) {
		result := NewDollar(5)
		expected := Money{amount: 5, currency: "USD"}
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})

	t.Run("10ドルのDollarオブジェクトが正しく作成される", func(t *testing.T) {
		result := NewDollar(10)
		expected := Money{amount: 10, currency: "USD"}
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})

	t.Run("0ドルのDollarオブジェクトが正しく作成される", func(t *testing.T) {
		result := NewDollar(0)
		expected := Money{amount: 0, currency: "USD"}
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})
}
