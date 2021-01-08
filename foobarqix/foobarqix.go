package foobarqix

import (
	"strconv"
	"strings"
)

func Compute(input string) string {
	number, err := strconv.Atoi(input)
	if err != nil {
		return "unable to parse input"
	}

	result := ""

	if number % 3 == 0 {
		result += "Foo"
	}

	if number % 5 == 0 {
		result += "Bar"
	}

	if number % 7 == 0 {
		result += "Qix"
	}

	digits := strings.Split(input, "")
	for _, d := range digits {
		if d == "3" {
			result += "Foo"
		}

		if d == "5" {
			result += "Bar"
		}

		if d == "7" {
			result += "Qix"
		}
	}

	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
