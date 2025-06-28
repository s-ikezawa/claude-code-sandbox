// Dollar アメリカドルを表すクラス
package main

// NewDollar 指定された金額のドルを作成する
func NewDollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}
