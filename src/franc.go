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

// 冗長なラッパーメソッドを削除
// 埋め込みによりMoney.Times()、Money.Equals()が直接利用可能
// GetMoneyメソッドも不要（埋め込みによる直接アクセス可能）