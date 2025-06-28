package main

// Money 通貨全体を扱う基底クラス
type Money struct {
	amount   int
	currency string
}

// NewMoney Moneyのコンストラクタ
func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

// Amount 金額を取得
func (m *Money) Amount() int {
	return m.amount
}

// Currency 通貨タイプを取得
func (m *Money) Currency() string {
	return m.currency
}