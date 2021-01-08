package foobarqix

import "strconv"

func Compute(input string) string {
	number, err := strconv.Atoi(input)
	if err != nil {
		return "unable to parse input"
	}
	if number % 3 == 0 {
		return "Foo"
	}

	if number % 5 == 0 {
		return "Bar"
	}

	return strconv.Itoa(number)
}
