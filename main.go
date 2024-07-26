package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func calculate(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func toArabic(roman string) (int, error) {
	if value, exists := romanToArabic[roman]; exists {
		return value, nil
	}
	return 0, errors.New("invalid roman numeral")
}

func toRoman(arabic int) (string, error) {
	if arabic < 1 {
		return "", errors.New("resulting roman numeral is less than I")
	}

	roman := ""
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; arabic > 0; i++ {
		for arabic >= values[i] {
			arabic -= values[i]
			roman += symbols[i]
		}
	}

	return roman, nil
}

func isRoman(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var inputString string
	fmt.Println("input your expression:")
	inputString, _ = in.ReadString('\n')
	inputString = strings.TrimSpace(inputString)

	inputs := strings.Split(inputString, " ")
	if len(inputs) != 3 {
		fmt.Println("error: must contain two numbers and an operator")
		return
	}

	strA, operator, strB := inputs[0], inputs[1], inputs[2]

	var a, b, result int
	var err error
	isRomanOp := false

	if isRoman(strA) && isRoman(strB) {
		isRomanOp = true
		a, err = toArabic(strA)
		if err != nil {
			fmt.Fprintln(out, "error:", err)
			return
		}
		b, err = toArabic(strB)
		if err != nil {
			fmt.Fprintln(out, "error:", err)
			return
		}
	} else {
		a, err = strconv.Atoi(strA)
		if a > 10 || a < 1 {
			fmt.Fprintln(out, "numbers should be between 1 and 10")
			return
		}
		if err != nil {
			fmt.Fprintln(out, "error: incorrect format of number", strA)
			return
		}
		b, err = strconv.Atoi(strB)
		if b > 10 || b < 1 {
			fmt.Fprintln(out, "numbers should be between 1 and 10")
			return
		}
		if err != nil {
			fmt.Fprintln(out, "error: incorrect format of number", strB)
			return
		}
	}

	result, err = calculate(a, b, operator)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if isRomanOp {
		romResult, err := toRoman(result)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(romResult)
	} else {
		fmt.Println(result)
	}
}
