package currency

// Money 金額と通貨を表す構造体
type Money struct {
	amount   int
	currency string
}

// NewDollar USD通貨のMoneyオブジェクトを作成
func NewDollar(amount int) *Money {
	return &Money{amount: amount, currency: "USD"}
}

// NewFranc CHF通貨のMoneyオブジェクトを作成
func NewFranc(amount int) *Money {
	return &Money{amount: amount, currency: "CHF"}
}

// Equals 他のMoneyオブジェクトとの等価性を判定
func (m *Money) Equals(other *Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// Times 金額を指定された倍率で計算し、新しいMoneyオブジェクトを返す
func (m *Money) Times(multiplier int) *Money {
	return &Money{amount: m.amount * multiplier, currency: m.currency}
}
