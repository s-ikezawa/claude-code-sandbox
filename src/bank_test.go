package main

import (
	"testing"
)

// TestBankReduce tests Bank.Reduce method behavior
// 期待: Bankが為替レートを使用してMoneyオブジェクトを指定通貨に換算する
func TestBankReduce(t *testing.T) {
	t.Run("USD通貨のMoneyをUSDに換算すると同じ値が返される", func(t *testing.T) {
		bank := NewBank()
		money := NewDollar(5)
		result := bank.Reduce(money, "USD")
		expected := NewDollar(5)
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})

	t.Run("CHF通貨のMoneyをUSDに換算すると正しい値が返される", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 0.5) // CHF 1 = USD 0.5 (CHF 2 = USD 1)
		money := NewFranc(10)
		result := bank.Reduce(money, "USD")
		expected := NewDollar(5)
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})
}

// TestBankSum tests Bank.Sum method behavior
// 期待: Bankが複数のMoneyオブジェクトを合計し、指定通貨に換算する
func TestBankSum(t *testing.T) {
	t.Run("5USD+10CHFをUSDで合計すると10USDになる", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 0.5) // CHF 2 = USD 1
		usd := NewDollar(5)
		chf := NewFranc(10)
		result := bank.Sum(usd, chf, "USD")
		expected := NewDollar(10)
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})

	t.Run("同じ通貨同士の合計は正しく計算される", func(t *testing.T) {
		bank := NewBank()
		usd1 := NewDollar(5)
		usd2 := NewDollar(3)
		result := bank.Sum(usd1, usd2, "USD")
		expected := NewDollar(8)
		if !result.Equals(expected) {
			t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
		}
	})
}

// TestBankRateManagement tests Bank rate management behavior
// 期待: Bankが為替レートを正しく管理する
func TestBankRateManagement(t *testing.T) {
	t.Run("為替レートが正しく設定・取得される", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 0.5)
		rate := bank.GetRate("CHF", "USD")
		expected := 0.5
		if rate != expected {
			t.Errorf("期待値: %f, 実際の値: %f", expected, rate)
		}
	})

	t.Run("未設定の為替レートは1.0が返される", func(t *testing.T) {
		bank := NewBank()
		rate := bank.GetRate("USD", "USD")
		expected := 1.0
		if rate != expected {
			t.Errorf("期待値: %f, 実際の値: %f", expected, rate)
		}
	})
}
