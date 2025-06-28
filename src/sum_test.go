package currency

import "testing"

// TestSumReduce tests Sum.Reduce method behavior
// 期待: Sumオブジェクトが指定通貨に正しく換算される
func TestSumReduce(t *testing.T) {
	tests := []struct {
		name     string
		sum      *Sum
		to       string
		expected *Money
	}{
		{
			name:     "5ドル + 5ドルを USD に換算すると 10ドル",
			sum:      &Sum{augend: NewDollar(5), addend: NewDollar(5)},
			to:       "USD",
			expected: NewDollar(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := NewBank()
			result := tt.sum.Reduce(bank, tt.to)
			if !result.Equals(tt.expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", tt.expected, result)
			}
		})
	}
}
