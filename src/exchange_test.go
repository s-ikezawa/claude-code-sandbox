package main

import (
	"testing"
)

func TestBank(t *testing.T) {
	t.Run("為替レートを設定できる", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2.0)

		rate := bank.GetRate("CHF", "USD")
		if rate != 2.0 {
			t.Errorf("GetRate() = %v, want 2.0", rate)
		}
	})

	t.Run("為替レートによる換算", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2.0) // 1 CHF = 2 USD

		franc := NewFranc(5)
		result := bank.Convert(&franc.Money, "USD")
		expected := &Money{amount: 10, currency: "USD"}

		if !result.Equals(expected) {
			t.Errorf("Convert() = %v, want %v", result, expected)
		}
	})

	t.Run("為替レート換算による通貨加算", func(t *testing.T) {
		tests := []struct {
			name           string
			rate           float64
			usdAmount      int
			chfAmount      int
			expectedAmount int
			description    string
		}{
			{
				name:           "通常の換算 (1 CHF = 2 USD)",
				rate:           2.0,
				usdAmount:      5,
				chfAmount:      10,
				expectedAmount: 25,
				description:    "5 USD + (10 CHF * 2) = 25 USD",
			},
			{
				name:           "仕様書の例 (USD 1 : CHF 2)",
				rate:           0.5,
				usdAmount:      5,
				chfAmount:      10,
				expectedAmount: 10,
				description:    "5 USD + (10 CHF * 0.5) = 10 USD",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				bank := NewBank()
				bank.AddRate("CHF", "USD", tt.rate)

				usd := NewDollar(tt.usdAmount)
				chf := NewFranc(tt.chfAmount)

				result := bank.AddWithConversion(&usd.Money, &chf.Money)
				expected := &Money{amount: tt.expectedAmount, currency: "USD"}

				if !result.Equals(expected) {
					t.Errorf("AddWithConversion() = %v, want %v (%s)", result, expected, tt.description)
				}
			})
		}
	})
}
