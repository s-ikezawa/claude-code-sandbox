package main

// Money 通貨全体を扱う基底クラス
type Money struct {
	amount   int
	currency string
}

// Amount 金額を返す
func (m *Money) Amount() int {
	return m.amount
}

// Currency 通貨を返す
func (m *Money) Currency() string {
	return m.currency
}

// Equals 通貨が等価かどうかを判定する
func (m *Money) Equals(other *Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

// Add 通貨を加算する
func (m *Money) Add(other *Money) *Money {
	// 同じ通貨の場合は単純加算
	if m.currency == other.currency {
		return &Money{
			amount:   m.amount + other.amount,
			currency: m.currency,
		}
	}

	// 異なる通貨の場合は基準通貨（呼び出し元の通貨）で結果を返す
	// 為替レート変換は将来実装予定
	return &Money{
		amount:   m.amount + other.amount,
		currency: m.currency,
	}
}
