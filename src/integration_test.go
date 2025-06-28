package main

import (
	"testing"
)

// TestCurrencyAdditionWithExchange tests the complete currency addition with exchange functionality
// 期待: 仕様書の例「5 USD + 10 CHF = 10 USD」が正しく動作する
func TestCurrencyAdditionWithExchange(t *testing.T) {
	bank := NewBank()
	bank.AddRate("USD", "CHF", 2) // 1 USD = 2 CHF

	fiveDollars := NewDollar(5)
	tenFrancs := NewFranc(10)

	// 10 CHF を USD に変換してから加算
	tenFrancsInUSD := bank.Exchange(tenFrancs, "USD")
	result := fiveDollars.Plus(tenFrancsInUSD)

	expected := NewDollar(10)
	if !result.Equals(expected) {
		t.Errorf("5 USD + 10 CHF should equal 10 USD, got %v", result)
	}
}

// TestCurrencyMultiplicationExample tests currency multiplication as shown in specification
// 期待: 仕様書の例「5 USD x 2 = 10 USD」が正しく動作する
func TestCurrencyMultiplicationExample(t *testing.T) {
	fiveDollars := NewDollar(5)
	result := fiveDollars.Times(2)
	expected := NewDollar(10)

	if !result.Equals(expected) {
		t.Errorf("5 USD x 2 should equal 10 USD, got %v", result)
	}
}
