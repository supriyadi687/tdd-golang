package tdd_golang_test

import (
	"github.com/supriyadi687/tdd-golang"
	"testing"
)

func TestShouldReturn0GivenEmptyString(t *testing.T) {
	res, _ := tdd_golang.Add("")
	if res != 0 {
		t.Errorf("expected to be 0 got %d", res)
	}
}
