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
[x] Francクラスのテスト
    [x] Francオブジェクトが正しく作成できるか
    [x] Franc同士の掛け算が正しく動作するか
    [x] Francオブジェクトの比較が正しく動作するか
[x] 通貨間の換算テスト
    [x] USD to CHFの換算が正しく動作するか
    [x] 為替レート1USD:2CHFでの計算が正しいか
[x] 異なる通貨の合計テスト
    [x] 5USD + 10CHF = 10USDの計算が正しいか

リファクタリングTODOリスト:
[x] 冗長なラッパーメソッドの削除テスト
    [x] Dollar.Times()削除後でもMoney.Times()が正常動作するか
    [x] Franc.Times()削除後でもMoney.Times()が正常動作するか
    [x] Dollar.Equals()削除後でもMoney.Equals()が正常動作するか
    [x] Franc.Equals()削除後でもMoney.Equals()が正常動作するか
[x] 不要なGetMoneyメソッドの削除テスト
    [x] Dollar.GetMoney()削除後でも埋め込みによる直接アクセスが可能か
    [x] Franc.GetMoney()削除後でも埋め込みによる直接アクセスが可能か
[x] 既存機能の保持確認
    [x] リファクタリング後でも全ての既存テストが通るか
    [x] 型固有の戻り値が不要な場面では基底クラスメソッドを直接使用するか
[x] テストコードの重複統合
    [x] 作成テストをテーブル駆動テストに統合できるか
    [x] 掛け算テストをテーブル駆動テストに統合できるか
    [x] 比較テストをテーブル駆動テストに統合できるか
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

// 作成テストのテーブル駆動テスト統合
func TestCurrencyCreation(t *testing.T) {
	tests := []struct {
		name           string
		createFunc     func(int) interface{ Amount() int; Currency() string }
		amount         int
		expectedCurrency string
	}{
		{"Dollar作成テスト", func(a int) interface{ Amount() int; Currency() string } { return NewDollar(a) }, 5, "USD"},
		{"Franc作成テスト", func(a int) interface{ Amount() int; Currency() string } { return NewFranc(a) }, 5, "CHF"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currency := tt.createFunc(tt.amount)
			if currency.Amount() != tt.amount {
				t.Errorf("Expected amount %d, got %d", tt.amount, currency.Amount())
			}
			if currency.Currency() != tt.expectedCurrency {
				t.Errorf("Expected currency %s, got %s", tt.expectedCurrency, currency.Currency())
			}
		})
	}
}

// 掛け算テストのテーブル駆動テスト統合
func TestCurrencyMultiply(t *testing.T) {
	tests := []struct {
		name           string
		currencyType   string
		amount         int
		multiplier     int
		expectedAmount int
		expectedCurrency string
	}{
		{"Dollar掛け算テスト", "dollar", 5, 2, 10, "USD"},
		{"Franc掛け算テスト", "franc", 5, 2, 10, "CHF"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Money
			if tt.currencyType == "dollar" {
				dollar := NewDollar(tt.amount)
				result = dollar.Money.Times(tt.multiplier)
			} else if tt.currencyType == "franc" {
				franc := NewFranc(tt.amount)
				result = franc.Money.Times(tt.multiplier)
			}
			
			if result.Amount() != tt.expectedAmount {
				t.Errorf("Expected amount %d, got %d", tt.expectedAmount, result.Amount())
			}
			if result.Currency() != tt.expectedCurrency {
				t.Errorf("Expected currency %s, got %s", tt.expectedCurrency, result.Currency())
			}
		})
	}
}

// 比較テストのテーブル駆動テスト統合
func TestCurrencyEquals(t *testing.T) {
	tests := []struct {
		name           string
		currencyType   string
		amount1        int
		amount2        int
		expectedEqual  bool
	}{
		{"Dollar等価比較（同じ値）", "dollar", 5, 5, true},
		{"Dollar等価比較（異なる値）", "dollar", 5, 10, false},
		{"Franc等価比較（同じ値）", "franc", 5, 5, true},
		{"Franc等価比較（異なる値）", "franc", 5, 10, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c1, c2 *Money
			if tt.currencyType == "dollar" {
				c1 = NewDollar(tt.amount1).Money
				c2 = NewDollar(tt.amount2).Money
			} else if tt.currencyType == "franc" {
				c1 = NewFranc(tt.amount1).Money
				c2 = NewFranc(tt.amount2).Money
			}
			
			result := c1.Equals(c2)
			if result != tt.expectedEqual {
				t.Errorf("Expected %v, got %v", tt.expectedEqual, result)
			}
		})
	}
}


// 通貨間の換算テスト
func TestExchangeRate(t *testing.T) {
	// USD to CHFの換算が正しく動作するかテスト
	bank := NewBank()
	bank.AddRate("USD", "CHF", 2)
	
	usd := NewDollar(1)
	chf := bank.Exchange(usd.Money, "CHF")
	
	if chf.Amount() != 2 {
		t.Errorf("Expected amount 2, got %d", chf.Amount())
	}
	if chf.Currency() != "CHF" {
		t.Errorf("Expected currency CHF, got %s", chf.Currency())
	}
}

func TestCurrencyCalculation(t *testing.T) {
	// 為替レート1USD:2CHFでの計算が正しいかテスト
	bank := NewBank()
	bank.AddRate("USD", "CHF", 2)
	bank.AddRate("CHF", "USD", 0.5)
	
	dollar := NewDollar(5)
	franc := NewFranc(10)
	
	// 5USD * 2 = 10USD (5USD + 10CHF = 10USD)
	result := bank.Add(dollar.Money, franc.Money, "USD")
	
	if result.Amount() != 10 {
		t.Errorf("Expected amount 10, got %d", result.Amount())
	}
	if result.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", result.Currency())
	}
}

// 異なる通貨の合計テスト
func TestCurrencySum(t *testing.T) {
	// 5USD + 10CHF = 10USDの計算が正しいかテスト
	bank := NewBank()
	bank.AddRate("CHF", "USD", 0.5)
	
	dollar := NewDollar(5)
	franc := NewFranc(10)
	
	result := bank.Add(dollar.Money, franc.Money, "USD")
	
	if result.Amount() != 10 {
		t.Errorf("Expected amount 10, got %d", result.Amount())
	}
	if result.Currency() != "USD" {
		t.Errorf("Expected currency USD, got %s", result.Currency())
	}
}

// リファクタリング後の構造をテストする新しいテスト

// Dollar.Times()削除後でもMoney.Times()が正常動作するかテスト
func TestDollarTimesRefactored(t *testing.T) {
	dollar := NewDollar(5)
	result := dollar.Money.Times(2)
	expected := NewMoney(10, "USD")
	
	if !result.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Franc.Times()削除後でもMoney.Times()が正常動作するかテスト
func TestFrancTimesRefactored(t *testing.T) {
	franc := NewFranc(5)
	result := franc.Money.Times(2)
	expected := NewMoney(10, "CHF")
	
	if !result.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// Dollar.Equals()削除後でもMoney.Equals()が正常動作するかテスト
func TestDollarEqualsRefactored(t *testing.T) {
	dollar1 := NewDollar(5)
	dollar2 := NewDollar(5)
	dollar3 := NewDollar(10)
	
	if !dollar1.Money.Equals(dollar2.Money) {
		t.Error("Expected dollar1.Money to equal dollar2.Money")
	}
	if dollar1.Money.Equals(dollar3.Money) {
		t.Error("Expected dollar1.Money not to equal dollar3.Money")
	}
}

// Franc.Equals()削除後でもMoney.Equals()が正常動作するかテスト
func TestFrancEqualsRefactored(t *testing.T) {
	franc1 := NewFranc(5)
	franc2 := NewFranc(5)
	franc3 := NewFranc(10)
	
	if !franc1.Money.Equals(franc2.Money) {
		t.Error("Expected franc1.Money to equal franc2.Money")
	}
	if franc1.Money.Equals(franc3.Money) {
		t.Error("Expected franc1.Money not to equal franc3.Money")
	}
}

// 埋め込みによる直接アクセステスト
func TestEmbeddedDirectAccess(t *testing.T) {
	dollar := NewDollar(100)
	franc := NewFranc(200)
	
	// 埋め込みにより直接Moneyのメソッドにアクセス可能
	if dollar.Amount() != 100 {
		t.Errorf("Expected dollar amount 100, got %d", dollar.Amount())
	}
	if dollar.Currency() != "USD" {
		t.Errorf("Expected dollar currency USD, got %s", dollar.Currency())
	}
	
	if franc.Amount() != 200 {
		t.Errorf("Expected franc amount 200, got %d", franc.Amount())
	}
	if franc.Currency() != "CHF" {
		t.Errorf("Expected franc currency CHF, got %s", franc.Currency())
	}
}