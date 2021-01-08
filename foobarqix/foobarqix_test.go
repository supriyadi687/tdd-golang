package foobarqix

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDivisibleBy3ShouldReturnFoo(t *testing.T) {
	result := Compute("6")
	assert.Equal(t, result, "Foo")
	result = Compute("9")
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

func TestDivisibleBy3AndContains3ShouldReturnFooFoo(t *testing.T) {
	result := Compute("3")
	assert.Equal(t, result, "FooFoo")
}

func TestDivisibleBy3AndContains5ShouldReturnBarBar(t *testing.T) {
	result := Compute("5")
	assert.Equal(t, result, "BarBar")
}

func TestDivisibleBy3AndContains7ShouldReturnBarBar(t *testing.T) {
	result := Compute("7")
	assert.Equal(t, result, "QixQix")
}

func TestGivenInputIs13ShouldReturnFoo(t *testing.T) {
	result := Compute("13")
	assert.Equal(t, result, "Foo")
}


