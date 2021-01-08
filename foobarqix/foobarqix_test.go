package foobarqix

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDivisibleBy3ShouldReturnFoo(t *testing.T) {
	result := Compute("6")
	assert.Equal(t, result, "Foo")
}

func TestDivisibleBy5ShouldReturnBar(t *testing.T) {
	result := Compute("10")
	assert.Equal(t, result, "Bar")
}

func TestDivisibleBy7ShouldReturnQix(t *testing.T) {
	result := Compute("14")
	assert.Equal(t, result, "Qix")
}

func TestNotDivisibleBy35And7ShouldReturnNumber(t *testing.T) {
	result := Compute("1")
	assert.Equal(t, result, "1")
	result = Compute("2")
	assert.Equal(t, result, "2")
	result = Compute("4")
	assert.Equal(t, result, "4")
	result = Compute("8")
	assert.Equal(t, result, "8")
}


