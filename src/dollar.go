package main

// NewDollar アメリカドルを表すMoneyオブジェクトを作成する
func NewDollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}
