// Exchange 為替レート管理と通貨換算機能
package main

// Bank 為替レートを管理し、通貨換算を行う
type Bank struct {
	rates map[Pair]float64
}

// Pair 通貨ペアを表す構造体
type Pair struct {
	from string
	to   string
}

// NewBank 新しいBankインスタンスを作成する
func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]float64),
	}
}

// AddRate 為替レートを設定する
func (b *Bank) AddRate(from, to string, rate float64) {
	b.rates[Pair{from: from, to: to}] = rate
}

// Rate 指定した通貨ペアの為替レートを取得する
func (b *Bank) Rate(from, to string) float64 {
	if from == to {
		return 1.0
	}
	return b.rates[Pair{from: from, to: to}]
}

// Reduce 式を指定した通貨で評価する
func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(b, to)
}
