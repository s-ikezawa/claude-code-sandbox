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

// Plus 他のMoneyオブジェクトとの合計を行う（Expression型を返す）
func (m Money) Plus(addend Money) Expression {
	return Sum{augend: m, addend: addend}
}

// Reduce 式を指定した通貨で評価する（Expression interface実装）
func (m Money) Reduce(bank *Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return Money{
		amount:   int(float64(m.amount) * rate),
		currency: to,
	}
}

// Expression 式を表すインターフェース
type Expression interface {
	Reduce(bank *Bank, to string) Money
}

// Sum 加算式を表す構造体
type Sum struct {
	augend Money
	addend Money
}

// Reduce 加算式を指定した通貨で評価する
func (s Sum) Reduce(bank *Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return Money{amount: amount, currency: to}
}
