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

// Times 掛け算メソッド（Dollarを返すラッパー）
func (d *Dollar) Times(multiplier int) *Dollar {
	result := d.Money.Times(multiplier)
	return &Dollar{Money: result}
}

// Equals 等価比較メソッド（Dollarとして比較）
func (d *Dollar) Equals(other *Dollar) bool {
	return d.Money.Equals(other.Money)
}

// GetMoney Moneyインスタンスを取得
func (d *Dollar) GetMoney() *Money {
	return d.Money
}