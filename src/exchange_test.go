package src

import (
	"testing"
)

func TestCurrencyExchange(t *testing.T) {
	t.Run("USDとCHFの為替レート変換（1USD=2CHF）", func(t *testing.T) {
		// 5 USD + 10 CHF = 10 USD (when 1 USD = 2 CHF)
		usd := NewDollar(5)
		chf := NewFranc(10)
		
		bank := NewBank()
		bank.AddRate("CHF", "USD", 0.5) // 1 CHF = 0.5 USD (2 CHF = 1 USD)
		
		sum := usd.Money.Plus(chf.Money)
		result := bank.Reduce(sum, "USD")
		expected := NewDollar(10)
		
		if !result.Equals(expected.Money) {
			t.Errorf("Currency exchange failed: got %+v, expected %+v", result, expected.Money)
		}
	})

	t.Run("同じ通貨の場合は為替レート変換不要", func(t *testing.T) {
		usd1 := NewDollar(5)
		usd2 := NewDollar(10)
		
		bank := NewBank()
		sum := usd1.Money.Plus(usd2.Money)
		result := bank.Reduce(sum, "USD")
		expected := NewDollar(15)
		
		if !result.Equals(expected.Money) {
			t.Errorf("Same currency addition failed: got %+v, expected %+v", result, expected.Money)
		}
	})

	t.Run("CHFからUSDへの変換", func(t *testing.T) {
		chf := NewFranc(10)
		
		bank := NewBank()
		bank.AddRate("CHF", "USD", 0.5) // 1 CHF = 0.5 USD
		
		result := bank.Reduce(chf.Money, "USD")
		expected := NewDollar(5)
		
		if !result.Equals(expected.Money) {
			t.Errorf("CHF to USD conversion failed: got %+v, expected %+v", result, expected.Money)
		}
	})

	t.Run("USDからCHFへの変換", func(t *testing.T) {
		usd := NewDollar(5)
		
		bank := NewBank()
		bank.AddRate("USD", "CHF", 2.0) // 1 USD = 2 CHF
		
		result := bank.Reduce(usd.Money, "CHF")
		expected := NewFranc(10)
		
		if !result.Equals(expected.Money) {
			t.Errorf("USD to CHF conversion failed: got %+v, expected %+v", result, expected.Money)
		}
	})
}