// Money 通貨全体を扱う基底クラス
package main

// Money 金額と通貨を表す基底構造体
type Money struct {
	amount   int
	currency string
}

// Equals 他のMoneyオブジェクトとの等価性を判定する
func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// Times 金額を指定した倍数で乗算する
func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// Currency 通貨コードを取得する
func (m Money) Currency() string {
	return m.currency
}
