package currency

import "fmt"

// Bank 為替レート管理と通貨換算を行う構造体
type Bank struct {
	rates map[string]int
}

// NewBank 新しいBankインスタンスを作成
func NewBank() *Bank {
	return &Bank{rates: make(map[string]int)}
}

// AddRate 為替レートを設定
func (b *Bank) AddRate(from, to string, rate int) {
	key := fmt.Sprintf("%s_%s", from, to)
	b.rates[key] = rate
}

// Rate 為替レートを取得
func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	key := fmt.Sprintf("%s_%s", from, to)
	return b.rates[key]
}

// Reduce Moneyオブジェクトを指定通貨に換算
func (b *Bank) Reduce(money *Money, to string) *Money {
	if money.currency == to {
		return money
	}
	rate := b.Rate(money.currency, to)
	return &Money{amount: money.amount / rate, currency: to}
}
