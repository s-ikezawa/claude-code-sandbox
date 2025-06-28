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

// 冗長なラッパーメソッドを削除
// 埋め込みによりMoney.Times()、Money.Equals()が直接利用可能
// GetMoneyメソッドも不要（埋め込みによる直接アクセス可能）
