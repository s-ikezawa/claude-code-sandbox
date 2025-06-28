package src

// Dollar はアメリカドルを表すクラス
type Dollar struct {
	*Money
}

// NewDollar は新しいDollarインスタンスを作成する
func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money: NewMoney(amount, "USD"),
	}
}

// Times は指定した倍数で金額を乗算した新しいDollarインスタンスを返す
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Plus は他のDollarと加算した新しいDollarインスタンスを返す
func (d *Dollar) Plus(other *Dollar) Expression {
	if other == nil {
		return d
	}
	return d.Money.Plus(other.Money)
}

// Equals は他のDollarインスタンスと等価かどうかを判定する
func (d *Dollar) Equals(other *Dollar) bool {
	if other == nil {
		return false
	}
	return d.Money.Equals(other.Money)
}