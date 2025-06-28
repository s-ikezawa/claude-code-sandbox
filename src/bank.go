package main

// Bank 通貨換算機能を提供するクラス
type Bank struct {
	rates map[string]float64
}

// NewBank 新しいBankインスタンスを作成する
func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]float64),
	}
}

// AddRate 為替レートを設定する
func (b *Bank) AddRate(from, to string, rate float64) {
	key := from + "-" + to
	b.rates[key] = rate
}

// GetRate 為替レートを取得する
func (b *Bank) GetRate(from, to string) float64 {
	if from == to {
		return 1.0
	}
	key := from + "-" + to
	if rate, exists := b.rates[key]; exists {
		return rate
	}
	return 1.0
}

// Reduce Moneyオブジェクトを指定通貨に換算する
func (b *Bank) Reduce(money Money, toCurrency string) Money {
	rate := b.GetRate(money.GetCurrency(), toCurrency)
	convertedAmount := int(float64(money.amount) * rate)
	return Money{
		amount:   convertedAmount,
		currency: toCurrency,
	}
}

// Sum 複数のMoneyオブジェクトを合計し、指定通貨に換算する
func (b *Bank) Sum(money1, money2 Money, toCurrency string) Money {
	reduced1 := b.Reduce(money1, toCurrency)
	reduced2 := b.Reduce(money2, toCurrency)
	return Money{
		amount:   reduced1.amount + reduced2.amount,
		currency: toCurrency,
	}
}
