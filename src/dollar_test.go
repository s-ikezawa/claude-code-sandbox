// Dollar専用機能のテスト
// 責務: NewDollar関数の動作のみ
package main

import (
	"testing"
)

// TestNewDollar tests NewDollar function behavior
// 期待: 指定した金額でUSD通貨のMoneyオブジェクトが正しく作成される
func TestNewDollar(t *testing.T) {
	result := NewDollar(1)
	expected := Money{amount: 1, currency: "USD"}
	if !result.Equals(expected) {
		t.Errorf("期待値 %+v, 実際 %+v", expected, result)
	}
}
