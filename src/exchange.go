package main

// Bank 為替レート管理と通貨換算を行うクラス
type Bank struct {
	rates map[string]map[string]float64
}

// NewBank Bankを生成する
func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]map[string]float64),
	}
}

// AddRate 為替レートを設定する
func (b *Bank) AddRate(from, to string, rate float64) {
	if b.rates[from] == nil {
		b.rates[from] = make(map[string]float64)
	}
	b.rates[from][to] = rate
}

// GetRate 為替レートを取得する
func (b *Bank) GetRate(from, to string) float64 {
	if b.rates[from] == nil {
		return 1.0 // 同一通貨の場合は1
	}
	rate, exists := b.rates[from][to]
	if !exists {
		return 1.0 // レートが設定されていない場合は1
	}
	return rate
}

// Convert 通貨を換算する
func (b *Bank) Convert(money *Money, toCurrency string) *Money {
	if money.Currency() == toCurrency {
		return money // 同一通貨の場合はそのまま返す
	}

	rate := b.GetRate(money.Currency(), toCurrency)
	return &Money{
		amount:   int(float64(money.Amount()) * rate),
		currency: toCurrency,
	}
}

// AddWithConversion 為替レート換算による通貨加算
func (b *Bank) AddWithConversion(money1, money2 *Money) *Money {
	// money1の通貨を基準通貨とする
	converted := b.Convert(money2, money1.Currency())
	return money1.Add(converted)
}
