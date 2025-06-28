package main

// Bank 為替レートを管理し、通貨換算を行うクラス
type Bank struct {
	rates map[string]map[string]float64
}

// NewBank Bankのコンストラクタ
func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]map[string]float64),
	}
}

// AddRate 為替レートを追加
func (b *Bank) AddRate(from, to string, rate float64) {
	if b.rates[from] == nil {
		b.rates[from] = make(map[string]float64)
	}
	b.rates[from][to] = rate
}

// Exchange 通貨換算
func (b *Bank) Exchange(money *Money, toCurrency string) *Money {
	if money.Currency() == toCurrency {
		return money
	}

	rate := b.rates[money.Currency()][toCurrency]
	newAmount := int(float64(money.Amount()) * rate)

	return NewMoney(newAmount, toCurrency)
}

// Add 異なる通貨の合計（targetCurrencyに換算して合計）
func (b *Bank) Add(money1, money2 *Money, targetCurrency string) *Money {
	converted1 := b.Exchange(money1, targetCurrency)
	converted2 := b.Exchange(money2, targetCurrency)

	totalAmount := converted1.Amount() + converted2.Amount()
	return NewMoney(totalAmount, targetCurrency)
}
