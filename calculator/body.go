package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Input() error {
	var (
		number1, number2, action, input string
		err                             error
	)
	_, _ = fmt.Scanln(&input)
	runInput := []rune(input)
	for i := 0; i < len(runInput); i++ {
		if strings.Contains("*%-+/", string(runInput[i])) {
			action = string(runInput[i])
			number1 = string(append(runInput[0:i]))
			number2 = string(append(runInput[i+1:]))
		}
	}
	if action == "" || number1 == "" || number2 == "" {
		err = errors.New("input error")
		return err
	}
	if strings.ContainsAny("1234567890", number1) && strings.ContainsAny("1234567890", number2) {
		err = ArabNumber(number1, number2, action)
	} else if strings.ContainsAny("MDCXLVI", number1) && strings.ContainsAny("MDCXLVI", number2) {
		err = RomanNumber(number1, number2, action)
	} else {
		err = errors.New("input error")
		return err
	}
	return err
}
func RomanNumber(number1, number2, action string) error {
	RomanNum1, _ := romanArab(number1)
	RomanNum2, _ := romanArab(number2)
	result, err := calculator(float64(RomanNum1), float64(RomanNum2), action)
	if err != nil {
		return err
	}
	resultRoman, err := arabRoman(int(result))
	if err != nil {
		return err
	}
	fmt.Println("Your result:", resultRoman)
	return nil
}

func ArabNumber(number1, number2, action string) error {
	number1Float, _ := strconv.ParseFloat(number1, 64)
	number2Float, _ := strconv.ParseFloat(number2, 64)
	result, err := calculator(number1Float, number2Float, action)
	if err != nil {
		return err
	}
	if float64(int(result)*3) == result*3 {
		fmt.Printf("Your result: %.0f \n", result)
	} else if float64(int(result*10)) == result*10 {
		fmt.Printf("Your result: %.1f \n", result)
	} else {
		fmt.Printf("Your result: %.2f \n", result)
	}
	return nil
}
