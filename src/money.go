package main

// Money 通貨全体を扱う基底クラス
type Money struct {
	amount   int
	currency string
}

// Equals 他のMoneyオブジェクトと値と通貨が等しいかを判定する
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// Multiply 金額に指定された値を掛けた新しいMoneyオブジェクトを返す
func (m Money) Multiply(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// GetCurrency 通貨を取得する
func (m Money) GetCurrency() string {
	return m.currency
}
