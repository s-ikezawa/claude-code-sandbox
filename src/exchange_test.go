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

		sum := usd.Plus(chf)
		result := bank.Reduce(sum, "USD")
		expected := NewDollar(10)

		if !result.Equals(expected) {
			t.Errorf("Currency exchange failed: got %+v, expected %+v", result, expected)
		}
	})

	t.Run("同じ通貨の場合は為替レート変換不要", func(t *testing.T) {
		usd1 := NewDollar(5)
		usd2 := NewDollar(10)

		bank := NewBank()
		sum := usd1.Plus(usd2)
		result := bank.Reduce(sum, "USD")
		expected := NewDollar(15)

		if !result.Equals(expected) {
			t.Errorf("Same currency addition failed: got %+v, expected %+v", result, expected)
		}
	})

	// 通貨変換テストケース
	conversionTests := []struct {
		name             string
		sourceAmount     int
		sourceCurrency   func(int) *Money
		targetCurrency   string
		rate             float64
		expectedAmount   int
		expectedCurrency func(int) *Money
	}{
		{"CHFからUSDへの変換", 10, NewFranc, "USD", 0.5, 5, NewDollar},
		{"USDからCHFへの変換", 5, NewDollar, "CHF", 2.0, 10, NewFranc},
	}

	for _, tc := range conversionTests {
		t.Run(tc.name, func(t *testing.T) {
			source := tc.sourceCurrency(tc.sourceAmount)
			bank := NewBank()
			bank.AddRate(source.Currency(), tc.targetCurrency, tc.rate)

			result := bank.Reduce(source, tc.targetCurrency)
			expected := tc.expectedCurrency(tc.expectedAmount)

			if !result.Equals(expected) {
				t.Errorf("%s failed: got %+v, expected %+v", tc.name, result, expected)
			}
		})
	}
}
