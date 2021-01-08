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

