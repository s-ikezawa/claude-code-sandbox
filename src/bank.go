package src

// Bank は為替レートを管理し、通貨変換を行うクラス
type Bank struct {
	rates map[string]map[string]float64
}

// NewBank は新しいBankインスタンスを作成する
func NewBank() *Bank {
	return &Bank{
		rates: make(map[string]map[string]float64),
	}
}

// AddRate は為替レートを追加する
func (b *Bank) AddRate(from, to string, rate float64) {
	if b.rates[from] == nil {
		b.rates[from] = make(map[string]float64)
	}
	b.rates[from][to] = rate
}

// Reduce は指定した通貨に変換したMoneyを返す
func (b *Bank) Reduce(source Expression, toCurrency string) *Money {
	if source == nil {
		return nil
	}

	return source.Reduce(b, toCurrency)
}

// getRate は為替レートを取得する
func (b *Bank) getRate(from, to string) float64 {
	if from == to {
		return 1.0
	}

	if b.rates[from] != nil {
		if rate, exists := b.rates[from][to]; exists {
			return rate
		}
	}

	return 0
}

// ConvertToMoney はMoneyを返すヘルパー関数（後方互換性のため残存）
func ConvertToMoney(money interface{}) *Money {
	switch m := money.(type) {
	case *Money:
		return m
	default:
		return nil
	}
}
