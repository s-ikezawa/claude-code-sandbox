package main

import "testing"

/*
通貨システムのテスト項目:

TODOリスト:
[x] Money基底クラスのテスト
    [x] 金額を正しく保持できるか
    [x] 通貨タイプを正しく保持できるか
[ ] Dollarクラスのテスト
    [ ] Dollarオブジェクトが正しく作成できるか
    [ ] Dollar同士の掛け算が正しく動作するか
    [ ] Dollarオブジェクトの比較が正しく動作するか
[ ] Francクラスのテスト
    [ ] Francオブジェクトが正しく作成できるか
    [ ] Franc同士の掛け算が正しく動作するか
    [ ] Francオブジェクトの比較が正しく動作するか
[ ] 通貨間の換算テスト
    [ ] USD to CHFの換算が正しく動作するか
    [ ] 為替レート1USD:2CHFでの計算が正しいか
[ ] 異なる通貨の合計テスト
    [ ] 5USD + 10CHF = 10USDの計算が正しいか
*/

// Money基底クラスのテスト
func TestMoneyAmount(t *testing.T) {
	// 金額を正しく保持できるかテスト
	money := NewMoney(10, "USD")
	if money.Amount() != 10 {
		t.Errorf("Expected amount 10, got %d", money.Amount())
	}
}

func TestMoneyCurrency(t *testing.T) {
	// 通貨タイプを正しく保持できるかテスト
	money := NewMoney(10, "USD")
	if money.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", money.Currency())
	}
}