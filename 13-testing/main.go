package main

import "errors"

func main() {}

func Divide(dividend, divisor float64) (float64, error) {
	var result float64

	// cannot divide by 0
	if divisor == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = dividend / divisor

	return result, nil
}

func Add(num1, num2 float64) (float64, error) {
	result := num1 + num2
	return result, nil
}
