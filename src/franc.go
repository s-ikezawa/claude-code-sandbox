package main

// Franc フランス フランを表すクラス
type Franc struct {
	*Money
}

// NewFranc Francのコンストラクタ
func NewFranc(amount int) *Franc {
	return &Franc{
		Money: NewMoney(amount, "CHF"),
	}
}

// Times 掛け算メソッド（Francを返すラッパー）
func (f *Franc) Times(multiplier int) *Franc {
	result := f.Money.Times(multiplier)
	return &Franc{Money: result}
}

// Equals 等価比較メソッド（Francとして比較）
func (f *Franc) Equals(other *Franc) bool {
	return f.Money.Equals(other.Money)
}

// GetMoney Moneyインスタンスを取得
func (f *Franc) GetMoney() *Money {
	return f.Money
}