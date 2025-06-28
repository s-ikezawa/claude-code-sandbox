package main

import (
	"testing"
)

// TestNewFranc tests NewFranc constructor behavior
// 期待: 指定された金額で正しいCHF通貨のMoneyオブジェクトが作成される
func TestNewFranc(t *testing.T) {
	testCases := []struct {
		testName string
		amount   int
		expected Money
	}{
		{
			testName: "5フランのFrancオブジェクトが正しく作成される",
			amount:   5,
			expected: Money{amount: 5, currency: "CHF"},
		},
		{
			testName: "10フランのFrancオブジェクトが正しく作成される",
			amount:   10,
			expected: Money{amount: 10, currency: "CHF"},
		},
		{
			testName: "0フランのFrancオブジェクトが正しく作成される",
			amount:   0,
			expected: Money{amount: 0, currency: "CHF"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			result := NewFranc(tc.amount)
			if !result.Equals(tc.expected) {
				t.Errorf("期待値: %+v, 実際の値: %+v", tc.expected, result)
			}
		})
	}
}
