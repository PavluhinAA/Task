package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var arabMap = []struct {
	decVal int
	symVal string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{8, "IIX"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}
var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"},
	{500, "D"},
	{100, "C"},
	{50, "L"},
	{10, "X"},
	{5, "V"},
	{1, "I"},
}

func romanArab(input string) (int, error) {
	var (
		err         error
		resultRoman int
		lastIndex   = 10000
	)
	if input == "" {
		return 0, err
	}
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if strings.Contains("MDCLXVI", string(runes[i])) {
			continue
		} else {
			err = errors.New("неправильный ввод римских цифр")
			break
		}
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	input = string(runes)
	for _, roman := range romanMap {
		if strings.Contains(input, roman.symbol) {
			if strings.Index(input, roman.symbol) >= lastIndex {
				resultRoman = resultRoman - (roman.decVal * strings.Count(input, roman.symbol))
				input = strings.Trim(input, roman.symbol)
			} else {
				lastIndex = strings.Index(input, roman.symbol)
				resultRoman = resultRoman + roman.decVal*strings.Count(input, roman.symbol)
				input = strings.Trim(input, roman.symbol)
			}
		}
	}
	return resultRoman, err
}
func arabRoman(input int) (error, string) {
	var (
		result string
		err    error
	)
	if input > 1200 {
		err = errors.New("Римские числа не должны превышать MMMCMXCIX")
	}
	if input == 0 {
		return err, ""
	}
	for _, decNum := range arabMap {
		for i := 0; i < len(strconv.Itoa(input)); i++ {
			if input >= decNum.decVal {
				result = result + decNum.symVal
				input -= decNum.decVal
			}
		}
	}
	return err, result
}
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
func calculator(number1, number2 float64, input string) (float64, error) {
	var (
		result float64
		err    error
	)
	switch input {
	case `*`:
		result = multiplication(number1, number2)
	case "/":
		if number2 == 0 {
			err = errors.New("деление на ноль невозможно")
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
		err = errors.New("неизвестный знак действия")
	}
	return result, err
}
func RomanNumber() {
	var (
		number1, number2, action string
		result                   float64
	)
	fmt.Println("Введите запрос\nПример ввода: XX * X\nПри нецелом результате он будет округлён")
	for {
		_, err := fmt.Scanln(&number1, &action, &number2)
		if err != nil {
			fmt.Println("проверьте правильность ввода")
			fmt.Println(err)
			err = nil
			continue
		}
		RomanNum1, err := romanArab(number1)
		RomanNum2, err := romanArab(number2)
		if err != nil {
			fmt.Println(err)
			err = nil
			continue
		}
		result, err = calculator(float64(RomanNum1), float64(RomanNum2), action)
		if err != nil {
			fmt.Println(err)
			err = nil
			continue
		}
		err, resultRoman := arabRoman(int(result))
		if err != nil {
			fmt.Println(err)
			err = nil
			continue
		}
		fmt.Println("Ваш результат:", resultRoman)
	}
}
func ArabNumber() {
	var (
		input                    string
		number1, number2, result float64
		err                      error
	)
	fmt.Println("Введите запрос\nПример ввода: 5 * 2")
	for {
		_, err = fmt.Scanf("%g %s %g", &number1, &input, &number2)
		if err != nil {
			fmt.Println("проверьте правильность ввода")
			err = nil
			continue
		}
		result, err = calculator(number1, number2, input)
		if err != nil {
			fmt.Println(err)
			err = nil
			continue
		}
		if float64(int(result)*3) == result*3 {
			fmt.Printf("Ваш результат: %.0f \n", result)
		} else if float64(int(result*10)) == result*10 {
			fmt.Printf("Ваш результат: %.1f \n", result)
		} else {
			fmt.Printf("Ваш результат: %.2f \n", result)
		}
	}
}
func main() {
	var (
		input string
	)
	fmt.Println("Чтобы выбрать цифры для работы введите\nДля арабских- A\nДля римских-R")
	for {
		fmt.Scan(&input)
		if input == "A" {
			ArabNumber()
		} else if input == "R" {
			RomanNumber()
		} else {
			fmt.Println("проверьте правильность ввода")
			continue
		}
	}
}

// Как тебе мой велосипед?
