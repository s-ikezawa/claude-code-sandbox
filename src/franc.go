// Franc フランスフランを表すクラス
package main

// NewFranc 指定された金額のフランを作成する
func NewFranc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}
