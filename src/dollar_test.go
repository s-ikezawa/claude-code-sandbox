// Dollar専用機能のテスト
// 責務: NewDollar関数の動作のみ
package main

import (
	"testing"
)

func TestNewDollar関数(t *testing.T) {
	result := NewDollar(1)
	expected := Money{amount: 1, currency: "USD"}
	if !result.Equals(expected) {
		t.Errorf("期待値 %+v, 実際 %+v", expected, result)
	}
}
