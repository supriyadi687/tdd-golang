package foobarqix

import (
	"fmt"
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

	if shouldProcess(number) {
		result = processDivisibleRules(number, result)
		result = processDigits(input, result, false)
	} else {
		result = processDigits(input, result, true)
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


func processDigits(input string, result string, unprocessed bool) string {
	digits := strings.Split(input, "")
	for _, d := range digits {

		if d == "0" {
			result += "*"
		}

		if unprocessed && d != "0" {
			fmt.Println(d)
			result += d
			continue
		}

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

func isDivisible(input int, divisor int) bool {
	return input%divisor == 0
}

func shouldProcess(input int) bool {
	inputString := strconv.Itoa(input)
	process := isDivisible(input, FooNumber) ||
		  isDivisible(input, BarNumber) ||
		  isDivisible(input, QixNumber) ||
		strings.Contains(inputString, "3") ||
		strings.Contains(inputString, "5") ||
		strings.Contains(inputString, "7")
	return process
}