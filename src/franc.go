package main

// Franc フランスフランを表すクラス
type Franc struct {
	Money
}

// NewFranc Francを生成する
func NewFranc(amount int) *Franc {
	return &Franc{
		Money: Money{
			amount:   amount,
			currency: "CHF",
		},
	}
}
