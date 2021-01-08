package foobarqix

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDivisibleBy3ShouldReturnFoo(t *testing.T) {
	result := Compute("6")
	assert.Equal(t, result, "Foo")
}
