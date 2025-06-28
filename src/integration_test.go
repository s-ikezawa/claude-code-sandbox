package currency

import "testing"

// TestCurrencyIntegration tests multi-currency addition and exchange rate conversion
// 期待: 異なる通貨の加算時に為替レートに基づいて正しく換算された結果が得られる
func TestCurrencyIntegration(t *testing.T) {
	tests := []struct {
		name           string
		usdAmount      int
		chfAmount      int
		exchangeRate   int // USD 1 : CHF exchangeRate
		expectedResult int // USD単位での期待結果
	}{
		{
			name:           "5 USD + 10 CHF = 10 USD (レート1:2)",
			usdAmount:      5,
			chfAmount:      10,
			exchangeRate:   2,
			expectedResult: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// USD Moneyオブジェクトの作成
			usd := NewDollar(tt.usdAmount)
			// CHF Moneyオブジェクトの作成
			chf := NewFranc(tt.chfAmount)
			// 銀行と為替レートの設定
			bank := NewBank()
			bank.AddRate("CHF", "USD", tt.exchangeRate)
			// 加算処理の実行
			sum := usd.Plus(chf)
			result := sum.Reduce(bank, "USD")
			// 結果の検証
			expected := NewDollar(tt.expectedResult)
			if !result.Equals(expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", expected, result)
			}
		})
	}
}
