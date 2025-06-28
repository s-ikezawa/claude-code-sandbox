package main

// Dollar アメリカ ドル を表すクラス
type Dollar struct {
	*Money
}

// NewDollar Dollarのコンストラクタ
func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money: NewMoney(amount, "USD"),
	}
}

// Times 掛け算メソッド
func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals 等価比較メソッド
func (d *Dollar) Equals(other *Dollar) bool {
	return d.amount == other.amount && d.currency == other.currency
}