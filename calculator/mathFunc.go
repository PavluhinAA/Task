package calculator

import (
	"errors"
)

func division(number1, number2 float64) float64 {
	res := number1 / number2
	return res
}
func multiplication(number1, number2 float64) float64 {
	res := number1 * number2
	return res
}
func subtraction(number1, number2 float64) float64 {
	res := number1 - number2
	return res
}
func addition(number1, number2 float64) float64 {
	res := number1 + number2
	return res
}
func find_perc(number1, number2 float64) float64 {
	res := number1 / 100 * number2
	return res
}
func calculator(number1, number2 float64, action string) (float64, error) {
	var (
		result float64
		err    error
	)
	switch action {
	case `*`:
		result = multiplication(number1, number2)
	case "/":
		if number2 == 0 {
			err = errors.New("division by zero is not possible")
			return result, err
		} else {
			result = division(number1, number2)
		}
	case "+":
		result = addition(number1, number2)
	case "-":
		result = subtraction(number1, number2)
	case "%":
		result = find_perc(number1, number2)
	default:
		err = errors.New("unknown action sign")
	}
	return result, err
}
