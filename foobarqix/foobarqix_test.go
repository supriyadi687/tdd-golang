package foobarqix

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNotDivisibleBy35And7ShouldReturnNumber(t *testing.T) {
	result := Compute(1)
	assert.Equal(t, result, "1")
}
