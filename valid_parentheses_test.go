package tdd_golang

import (
	"testing"
)

func TestShouldReturnFalseGivenParenthesesNotClosed(t *testing.T) {
	res := IsValidParentheses("(")
	if res != false {
		t.Errorf("expected to be false got %t", res)
	}
}
