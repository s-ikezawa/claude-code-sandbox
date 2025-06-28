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

// Times 掛け算メソッド
func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

// Equals 等価比較メソッド
func (f *Franc) Equals(other *Franc) bool {
	return f.amount == other.amount && f.currency == other.currency
}