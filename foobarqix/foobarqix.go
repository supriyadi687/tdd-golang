package foobarqix

import (
	"strconv"
	"strings"
)

const FOO = "Foo"
const BAR = "Bar"
const QIX = "Qix"
const FooNumber = 3
const BarNumber = 5
const QixNumber = 7

func Compute(input string) string {
	result := ""

	number, err := strconv.Atoi(input)
	if err != nil {
		return "unable to parse input"
	}

	result = processDivisibleRules(number, result)
	result = processDigits(input, result)

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}

func processDivisibleRules(number int, result string) string {
	if isDivisible(number, FooNumber) {
		result += FOO
	}

	if isDivisible(number, BarNumber) {
		result += BAR
	}

	if isDivisible(number, QixNumber) {
		result += QIX
	}
	return result
}


func processDigits(input string, result string) string {
	digits := strings.Split(input, "")
	for _, d := range digits {
		if d == strconv.Itoa(FooNumber) {
			result += FOO
		}

		if d == strconv.Itoa(BarNumber) {
			result += BAR
		}

		if d == strconv.Itoa(QixNumber) {
			result += QIX
		}
	}
	return result
}

func isDivisible(number int, divisor int) bool {
	return number%divisor == 0
}
