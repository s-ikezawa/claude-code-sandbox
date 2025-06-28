package main

// Money 通貨全体を扱う基底クラス
type Money struct {
	amount   int
	currency string
}

// Amount 金額を返す
func (m *Money) Amount() int {
	return m.amount
}

// Currency 通貨を返す
func (m *Money) Currency() string {
	return m.currency
}

// Equals 通貨が等価かどうかを判定する
func (m *Money) Equals(other *Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}
