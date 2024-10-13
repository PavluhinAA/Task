package calculator

import (
	"errors"
	"strings"
)

func romanArab(number string) (int, error) {
	romanMap := map[string]int{
		"M":    1000,
		"CM":   900,
		"D":    500,
		"CD":   400,
		"C":    100,
		"XC":   90,
		"L":    50,
		"XL":   40,
		"X":    10,
		"IX":   9,
		"VIII": 8,
		"VII":  7,
		"VI":   6,
		"V":    5,
		"IV":   4,
		"III":  3,
		"II":   2,
		"I":    1,
	}
	if number == "" {
		return 0, nil
	}
	var result int
	for key, value := range romanMap {
		if strings.Contains(number, key) {
			result = result + (value * strings.Count(number, key))
			number = strings.Trim(number, key)
		}
	}
	return result, nil
}
func arabRoman(input int) (string, error) {
	var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	arabMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		8:    "IIX",
		7:    "VII",
		6:    "VI",
		5:    "V",
		4:    "IV",
		3:    "III",
		2:    "II",
		1:    "I",
	}
	var (
		result string
		err    error
	)
	if input > 3999 {
		err = errors.New("римские числа не должны превышать MMMCMXCIX")
		return result, err
	}
	if input == 0 {
		return "", err
	}
	for i := 0; i < len(keys); i++ {
		if input >= keys[i] {
			result = result + arabMap[keys[i]]
			input -= keys[i]
			i = -1
		}
	}
	return result, err
}
