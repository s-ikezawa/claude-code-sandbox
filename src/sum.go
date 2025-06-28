package currency

// Sum 二つのMoneyオブジェクトの合計を表す構造体
type Sum struct {
	augend *Money // 被加数
	addend *Money // 加数
}

// Reduce Sumオブジェクトを指定通貨に換算
func (s *Sum) Reduce(bank *Bank, to string) *Money {
	augendReduced := bank.Reduce(s.augend, to)
	addendReduced := bank.Reduce(s.addend, to)
	return &Money{
		amount:   augendReduced.amount + addendReduced.amount,
		currency: to,
	}
}
