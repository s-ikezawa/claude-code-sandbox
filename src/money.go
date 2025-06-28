package src

// Expression は計算式を表すインターフェース
type Expression interface {
	Reduce(bank *Bank, to string) *Money
}

// Money は通貨全体を扱う基底クラス
type Money struct {
	amount   int
	currency string
}

// NewMoney は新しいMoneyインスタンスを作成する
func NewMoney(amount int, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

// NewDollar は新しいUSDのMoneyインスタンスを作成する
func NewDollar(amount int) *Money {
	return NewMoney(amount, "USD")
}

// NewFranc は新しいCHFのMoneyインスタンスを作成する
func NewFranc(amount int) *Money {
	return NewMoney(amount, "CHF")
}

// Amount は金額を返す
func (m *Money) Amount() int {
	return m.amount
}

// Currency は通貨を返す
func (m *Money) Currency() string {
	return m.currency
}

// Equals は他のMoneyインスタンスと等価かどうかを判定する
func (m *Money) Equals(other *Money) bool {
	if other == nil {
		return false
	}
	return m.amount == other.amount && m.currency == other.currency
}

// Times は指定した倍数で金額を乗算した新しいMoneyインスタンスを返す
func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus は同じ通貨の場合は金額を加算した新しいMoneyインスタンスを返す
// 異なる通貨の場合は加算式を表すSumを返す
func (m *Money) Plus(other *Money) Expression {
	if other == nil {
		return m
	}
	if m.currency == other.currency {
		return &Money{
			amount:   m.amount + other.amount,
			currency: m.currency,
		}
	}
	return &Sum{augend: m, addend: other}
}

// Reduce はBankを使って指定した通貨に変換する
func (m *Money) Reduce(bank *Bank, to string) *Money {
	if m.currency == to {
		return m
	}

	rate := bank.getRate(m.currency, to)
	if rate == 0 {
		return m
	}

	convertedAmount := int(float64(m.amount) * rate)
	return NewMoney(convertedAmount, to)
}

// Sum は2つのExpressionの和を表すクラス
type Sum struct {
	augend Expression
	addend Expression
}

// Reduce はSumを評価してMoneyに変換する
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	augendMoney := s.augend.Reduce(bank, to)
	addendMoney := s.addend.Reduce(bank, to)

	if augendMoney == nil {
		return addendMoney
	}
	if addendMoney == nil {
		return augendMoney
	}

	return &Money{
		amount:   augendMoney.amount + addendMoney.amount,
		currency: to,
	}
}
