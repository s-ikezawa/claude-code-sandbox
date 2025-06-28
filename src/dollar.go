package main

// Dollar アメリカドルを表すクラス
type Dollar struct {
	Money
}

// NewDollar Dollarを生成する
func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money: Money{
			amount:   amount,
			currency: "USD",
		},
	}
}
