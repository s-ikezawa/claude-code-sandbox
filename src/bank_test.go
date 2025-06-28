package currency

import "testing"

// TestBankExchangeRate tests Bank.AddRate method behavior
// 期待: 為替レートが正しく設定され、取得できる
func TestBankExchangeRate(t *testing.T) {
	tests := []struct {
		name         string
		from         string
		to           string
		rate         int
		expectedRate int
	}{
		{
			name:         "USD to CHF レート 1:2 が設定される",
			from:         "USD",
			to:           "CHF",
			rate:         2,
			expectedRate: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			bank.AddRate(tt.from, tt.to, tt.rate)
			actualRate := bank.Rate(tt.from, tt.to)
			if actualRate != tt.expectedRate {
				t.Errorf("期待値: %d, 実際の値: %d", tt.expectedRate, actualRate)
			}
		})
	}
}

// TestBankReduce tests Bank.Reduce method behavior
// 期待: 異なる通貨のMoneyオブジェクトが指定通貨に正しく換算される
func TestBankReduce(t *testing.T) {
	tests := []struct {
		name     string
		money    *Money
		to       string
		expected *Money
	}{
		{
			name:     "同じ通貨の場合、変換されずにそのまま返される",
			money:    NewDollar(5),
			to:       "USD",
			expected: NewDollar(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			result := bank.Reduce(tt.money, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", tt.expected, result)
			}
		})
	}
}
