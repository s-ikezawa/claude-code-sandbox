package main

// NewFranc フランスフランを表すMoneyオブジェクトを作成する
func NewFranc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}
