// Franc専用機能のテスト
// 責務: NewFranc関数の動作のみ
package main

import (
	"testing"
)

// TestNewFranc tests NewFranc function behavior
// 期待: 指定した金額でCHF通貨のMoneyオブジェクトが正しく作成される
func TestNewFranc(t *testing.T) {
	result := NewFranc(1)
	expected := Money{amount: 1, currency: "CHF"}
	if !result.Equals(expected) {
		t.Errorf("期待値 %+v, 実際 %+v", expected, result)
	}
}
