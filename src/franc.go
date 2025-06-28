package src

// Franc はフランスフランを表すクラス
type Franc struct {
	*Money
}

// NewFranc は新しいFrancインスタンスを作成する
func NewFranc(amount int) *Franc {
	return &Franc{
		Money: NewMoney(amount, "CHF"),
	}
}

// Times は指定した倍数で金額を乗算した新しいFrancインスタンスを返す
func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

// Plus は他のFrancと加算した新しいFrancインスタンスを返す
func (f *Franc) Plus(other *Franc) Expression {
	if other == nil {
		return f
	}
	return f.Money.Plus(other.Money)
}

// Equals は他のFrancインスタンスと等価かどうかを判定する
func (f *Franc) Equals(other *Franc) bool {
	if other == nil {
		return false
	}
	return f.Money.Equals(other.Money)
}