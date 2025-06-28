package main

import "testing"

/*
通貨システムのテスト項目:

TODOリスト:
[x] Money基底クラスのテスト
    [x] 金額を正しく保持できるか
    [x] 通貨タイプを正しく保持できるか
[x] Dollarクラスのテスト
    [x] Dollarオブジェクトが正しく作成できるか
    [x] Dollar同士の掛け算が正しく動作するか
    [x] Dollarオブジェクトの比較が正しく動作するか
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

// Dollarクラスのテスト
func TestDollarCreation(t *testing.T) {
	// Dollarオブジェクトが正しく作成できるかテスト
	dollar := NewDollar(5)
	if dollar.Amount() != 5 {
		t.Errorf("Expected amount 5, got %d", dollar.Amount())
	}
	if dollar.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", dollar.Currency())
	}
}

func TestDollarMultiply(t *testing.T) {
	// Dollar同士の掛け算が正しく動作するかテスト
	dollar := NewDollar(5)
	result := dollar.Times(2)
	if result.Amount() != 10 {
		t.Errorf("Expected amount 10, got %d", result.Amount())
	}
	if result.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", result.Currency())
	}
}

func TestDollarEquals(t *testing.T) {
	// Dollarオブジェクトの比較が正しく動作するかテスト
	dollar1 := NewDollar(5)
	dollar2 := NewDollar(5)
	dollar3 := NewDollar(10)
	
	if !dollar1.Equals(dollar2) {
		t.Error("Expected dollar1 to equal dollar2")
	}
	if dollar1.Equals(dollar3) {
		t.Error("Expected dollar1 not to equal dollar3")
	}
}